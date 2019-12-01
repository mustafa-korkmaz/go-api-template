package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const oliveCollectionName = "olives"

// OliveRepository represents olive collection operations interface
type OliveRepository interface {
	MongoBaseRepository
	GetOlivesCount() (int64, error)
}

// Olive respresents the struct for mongo db operations for olive collection
type Olive struct {
	MongoBase
}

//GetOlivesCount returns the document count for olive collection
func (repository *Olive) GetOlivesCount() (int64, error) {
	collection := repository.client.Database(repository.DBName).Collection(repository.CollectionName)

	var docCount, err = collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	return docCount, nil
}

//New creates a new olive repository object
func New(c *mongo.Client, dbName string) OliveRepository {
	var repository = Olive{}
	repository.CollectionName = oliveCollectionName
	repository.DBName = dbName
	repository.MongoBase.client = c

	return &repository
}
