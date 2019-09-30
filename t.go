package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	collection := client.Database("QAREALTIME").Collection("account")

	ctx, _ = context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(
		ctx,
		bson.D{
			{"broker_name", "QUANTAXIS"},
		})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		// do something with result....
		log.Println(result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
