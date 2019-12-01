package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoBaseRepository represents mongo db base operations interface
type MongoBaseRepository interface {
	FindOne(objectID string) *mongo.SingleResult
	// Insert(entity string)
	// Update(entity string)
	// Delete(id string)
}

// MongoBase respresents the struct for mongo db base operations
type MongoBase struct {
	client         *mongo.Client
	DBName         string
	CollectionName string
}

//FindOne gets the result by ObjectId
func (repository *MongoBase) FindOne(objectID string) *mongo.SingleResult {

	doc := bson.D{primitive.E{Key: "_id", Value: objectID}}
	collection := repository.client.Database(repository.DBName).Collection(repository.CollectionName)

	return collection.FindOne(context.TODO(), doc)
}
