package main

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertData(t *testing.T) {
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:123456@mongo/?authSource=admin"))
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	db := c.Database("tronics")
	col := db.Collection("products")
	res, err := inserData(col, User{"knz", "phumthawan"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res)
}
