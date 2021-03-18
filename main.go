package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Mersock/golang-mongo-testing/dbiface"
	"go.mongodb.org/mongo-driver/bson"
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

func findData(collection dbiface.CollectionAPIs) ([]User, error) {
	var users []User
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Printf("find error :%v\n", err)
		return users, err
	}
	fmt.Printf("cursor :%+v\n", cur.Current)
	err = cur.All(context.Background(), &users)
	if err != nil {
		return users, err
	}
	return users, nil
}

func main() {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:123456@mongo/?authSource=admin"))
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res1, err := inserData(col, User{"knz", "phumthawan"})
	if err != nil {
		fmt.Printf("insert failures: %+v\n", err)
	}
	fmt.Println("insert done ", res1)
	res2, err := findData(col)
	if err != nil {
		fmt.Printf("find failures: %+v\n", err)
	}
	fmt.Println("find done ", res2)
	res3, err := deleteOne(col, User{"knz", "phumthawan"})
	if err != nil {
		fmt.Printf("delete failures: %+v\n", err)
	}
	fmt.Println("delete done", res3)

}
