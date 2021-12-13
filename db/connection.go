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
func ConnectToDatabase() int64 {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	uriString := fmt.Sprintf("mongodb+srv://kaancetinkayasf:%s@cluster0.tiask.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", config.Password)

	// Connection to MongoDB Atlas
	clientOptions := options.Client().
		ApplyURI(uriString)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Db and collection declerations
	quotesDatabase := client.Database("test")
	quotesCollection := quotesDatabase.Collection("quotes")

	DocumentCount, err := quotesCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	return DocumentCount

	// var restaurants bson.M
	// err = quotesCollection.FindOne(ctx, bson.M{}).Decode(&restaurants)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(restaurants)

}
