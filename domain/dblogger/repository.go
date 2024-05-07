package dblogger

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertLog(mongoDb *mongo.Database) func(ctx context.Context, data DbLoggerData) error {
	return func(ctx context.Context, data DbLoggerData) error {
		collection := mongoDb.Collection("dblogger")

		_, err := collection.InsertOne(ctx, data)
		if err != nil {
			return err
		}

		return nil
	}
}
