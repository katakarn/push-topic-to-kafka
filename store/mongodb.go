package store

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(ctx context.Context) *mongo.Client {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		AuthSource:    viper.GetString("MONGO_DB_NAME"),
		Username:      viper.GetString("MONGO_DB_USER"),
		Password:      viper.GetString("MONGO_DB_PASSWORD"),
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("MONGO_DB_URI")).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB")
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}
