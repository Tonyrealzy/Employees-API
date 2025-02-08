package repository

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://tonyrealzy:AxQQubTrrCOUzwtj@tonyrealzy.alipl.mongodb.net/?retryWrites=true&w=majority&appName=Tonyrealzy"))
	if err != nil {
		log.Fatal("an error occurred while connecting to mongodb", err)
	}
	log.Println("connected to mongodb successfully")
	err = mongoTestClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}
	log.Println("ping successful")
	return mongoTestClient
}
