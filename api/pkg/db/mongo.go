package db

import (
	"context"
	"fmt"
	"forum/api/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoClient() (*mongo.Database, error) {
	fmt.Println("Connection to db")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// ignore loading error if could not connect to database error will be raised
	dbHost, _ := config.Config("DB_HOST", "string")
	dbPort, _ := config.Config("DB_PORT", "string")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	dbName, err := config.Config("DB_NAME", "string")
	if err != nil {
		return nil, err
	}

	db_client := client.Database(dbName.(string))
	fmt.Println("Successfully connected to database")
	return db_client, nil
}
