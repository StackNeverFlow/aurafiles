package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	url string = "mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&ssl=false"
)

//Connect used to connect to the mongodb database
func Connect(database string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongodb database!")
	return client.Database(database)
}
