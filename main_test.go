package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockCollection struct {
}

func (m *mockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	c := &mongo.InsertOneResult{}
	return c, nil
}

func (m *mockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	c := &mongo.DeleteResult{}
	return c, nil
}

func TestInsertData(t *testing.T) {
	mockCol := &mockCollection{}
	res1, err := inserData(mockCol, User{"knz", "phumthawan"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res1)
	res2, err := inserData(mockCol, User{"bbb", "ccc"})
	assert.NotNil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res2)
}

func testDeleteData(t *testing.T) {
	mockCol := &mockCollection{}
	res1, err := inserData(mockCol, User{"knz", "phumthawan"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.InsertOneResult{}, res1)
	res2, err := deleteOne(mockCol, User{"knz", "phumthawan"})
	assert.Nil(t, err)
	assert.IsType(t, &mongo.DeleteResult{}, res2)
}
