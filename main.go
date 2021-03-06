package main

import (
	"context"
	"errors"
	"log"

	"github.com/Mersock/golang-mongo-testing/dbiface"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//User struct
type User struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
}

func inserData(collection dbiface.CollectionAPIs, user User) (*mongo.InsertOneResult, error) {
	if user.FirstName != "knz" {
		return nil, errors.New("name not correct")
	}
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return res, err
	}
	return res, nil
}

func deleteOne(collection dbiface.CollectionAPIs, user User) (*mongo.DeleteResult, error) {
	res, err := collection.DeleteOne(context.Background(), user)
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
	res1, err := inserData(col, User{"knz", "phumthawan"})
	log.Println(res1, err)
	res2, err := deleteOne(col, User{"knz", "phumthawan"})
	log.Println(res2, err)
}
