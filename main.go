package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User struct
type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

func inserData(collection *mongo.Collection, user User) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:123456@mongo/?authSource=admin"))
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := inserData(col, User{"knz", "phumthawan"})
	log.Println(res, err)
}
