package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"os"
	"time"
)

var DB *mongo.Client

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		fmt.Println("DATABASE_URL not set in environment")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	fmt.Println(err)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("Error pinging MongoDB:", err)
	}

	DB = client

	fmt.Println("Connected to MongoDB!")
}
