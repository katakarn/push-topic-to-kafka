package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"testKafka/client"
	"testKafka/config"
	"testKafka/domain/health"
	"testKafka/domain/registration"
	ginrouter "testKafka/gin-router"
	"testKafka/middleware"
	"testKafka/producer"
	"testKafka/store"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sony/gobreaker"
)

func shouldBeSwitchedToOpen(counts gobreaker.Counts) bool {
	failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
	return counts.Requests >= 3 && failureRatio >= 0.6
}

var cb client.CircuitBreakerProxy

func init() {
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = shouldBeSwitchedToOpen
	st.OnStateChange = func(_ string, from gobreaker.State, to gobreaker.State) {
		log.Println("state changed from", from.String(), "to", to.String())
	}
	cb.Gb = gobreaker.NewCircuitBreaker(st)
}

func main() {
	conf := config.LoadConfig(".")
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	zapLog, err := zapConfig.Build()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer zapLog.Sync()
	cb.Logger = *zapLog

	producer := producer.NewProducer(zapLog, &kafka.ConfigMap{
		"bootstrap.servers": conf.KafkaServer,
	})

	if err := producer.Init(); err != nil {
		zapLog.Error("Error initializing Kafka producer", zap.Error(err))
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Init MongoDB
	mongoClient := store.InitMongoDB(ctx)
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	_ = mongoClient.Database(viper.GetString("MONGO_DB_NAME"))

	// Init PostgresDB
	pgDB := config.GetConnection(&conf)
	defer pgDB.Close()

	// Init RegDB
	regDB := store.GetRegDB(&conf)
	if regDB != nil {
		defer regDB.Close()
	}

	//Init redisDB
	redisDB := store.InitRedis(&conf)
	defer redisDB.Close()

	cf := cors.DefaultConfig()
	cf.AllowOrigins = []string{"*"}
	cf.AllowHeaders = []string{"*"}
	cf.AllowHeaders = append(
		cf.AllowHeaders,
		"Authorization",
		"X-API-KEY",
	)

	cf.AllowCredentials = true

	sr := ginrouter.SetupRouter()
	r := sr.Group("/api/registration")
	r.GET("/pushKafka", registration.PushKafkaHandler(producer))
	r.Use(cors.New(cf))
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/api/registration/health"}}))
	r.Use(gin.Recovery())

	r.GET("/health", health.HealthCheckHandler())
	r.Use(middleware.UserMobileAuthorizes())
	r.Use(middleware.LoggingMiddleware())

	server := &http.Server{
		Addr:           ":3000",
		Handler:        sr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("listen :3000", err)
		}
	}()

	<-ctx.Done()
	stop()
	log.Fatalf("force shutdown")

	timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(timeout); err != nil {
		log.Fatal(err)
	}

}
