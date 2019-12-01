package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

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
