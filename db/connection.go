package db

import (
	"context"
	util "daily-quote/utils"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connects to the MongoDB atlas cluster and returns the total amount of documents in collection
func ConnectToDatabase() []Quote {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	uriString := fmt.Sprintf("mongodb+srv://kaancetinkayasf:%s@cluster0.tiask.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", config.Password)
	clientOptions := options.Client().
		ApplyURI(uriString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	quotesDatabase := client.Database("test")
	quotesCollection := quotesDatabase.Collection("quotes")

	var quote []Quote

	aggr := bson.D{{"$sample", bson.D{{"size", 1}}}}

	cursor, err := quotesCollection.Aggregate(ctx, mongo.Pipeline{aggr})

	if err != nil {
		fmt.Println(err)
	}
	cursor.All(ctx, &quote)
	return quote
}
