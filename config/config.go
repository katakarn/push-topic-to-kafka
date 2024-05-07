package config

import (
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	KafkaServer string `mapstructure:"KAFKA_SERVER"`
	Topic       string `mapstructure:"TOPIC"`

	PgDbURL      string `mapstructure:"DATABASE_URL"` // DATABASE_URL will be used in preference if it exists
	PgDbHost     string `mapstructure:"POSTGRES_HOST"`
	PgDbPort     string `mapstructure:"POSTGRES_PORT"`
	PgDbUser     string `mapstructure:"POSTGRES_USER"`
	PgDbPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PgDbName     string `mapstructure:"POSTGRES_DB"`

	DbServer   string `mapstructure:"DB_SERVER"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbService  string `mapstructure:"DB_SERVICE"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`

	RedisClusterServer string `mapstructure:"REDIS_CLUSTER_SERVER"`

	RegistrationServiceBaseURL string `mapstructure:"REGISTRATION_SERVICE_BASE_URL"`
	RegistrationServiceKey     string `mapstructure:"REGISTRATION_SERVICE_KEY"`
}

func LoadConfig(path string) (config Config) {
	viper.SetDefault("KAFKA_SERVER", "localhost:8097,localhost:8098,localhost:8099")
	viper.SetDefault("MONGO_DB_URI", "mongodb://localhost:27017/")
	viper.SetDefault("MONGO_DB_NAME", "newsdb")
	viper.SetDefault("MONGO_DB_USER", "root")
	viper.SetDefault("MONGO_DB_PASSWORD", "password")

	viper.SetDefault("POSTGRES_HOST", "localhost")
	viper.SetDefault("POSTGRES_PORT", "5492")
	viper.SetDefault("POSTGRES_USER", "postgres")
	viper.SetDefault("POSTGRES_PASSWORD", "mysecretpassword")
	viper.SetDefault("POSTGRES_DB", "identity")

	viper.SetDefault("DB_SERVER", "172.27.2.5")
	viper.SetDefault("DB_PORT", "1521")
	viper.SetDefault("DB_SERVICE", "SUREG")
	viper.SetDefault("DB_USER", "REGSMARTAPP")
	viper.SetDefault("DB_PASSWORD", "App#651226@su")

	viper.SetDefault("REGISTRATION_SERVICE_BASE_URL", "https://smartplus-reg-dev.su.ac.th/SmartPlus")
	viper.SetDefault("REGISTRATION_SERVICE_KEY", "b2S7xei8ni7Q6UVkArNxBaGzmOlStH6X")

	viper.SetDefault("REDIS_CLUSTER_SERVER", "localhost:6397, localhost:6398, localhost:6399")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	AppConfig = config
	return AppConfig
}

func GetConnection(c *Config) *pg.DB {
	// if DATABASE_URL is valid, we will use its constituent values in preference
	validConfig, err := validPostgresURL(c.PgDbURL)
	if err == nil {
		c = validConfig
	}
	db := pg.Connect(&pg.Options{
		Addr:     c.PgDbHost + ":" + c.PgDbPort,
		User:     c.PgDbUser,
		Password: c.PgDbPassword,
		Database: c.PgDbName,
		PoolSize: 150,
	})
	return db
}

func validPostgresURL(URL string) (*Config, error) {
	if URL == "" || strings.TrimSpace(URL) == "" {
		return nil, errors.New("database url is blank")
	}

	validURL, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	c.PgDbURL = URL
	c.PgDbHost = validURL.Host
	c.PgDbName = validURL.Path
	c.PgDbPort = validURL.Port()
	c.PgDbUser = validURL.User.Username()
	c.PgDbPassword, _ = validURL.User.Password()
	return c, nil
}
