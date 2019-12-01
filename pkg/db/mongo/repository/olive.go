package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// OliveService represents olive collection operations interface
type OliveService interface {
	GetOlivesCount() (int64, error)
}

// OliveRepository respresents the struct for mongo db base operations
type OliveRepository struct {
	MongoBaseRepository
	ekmek string
}

//GetOlivesCount returns the document count for olive collection
func (repository *OliveRepository) GetOlivesCount() (int64, error) {
	collection := repository.client.Database(repository.DBName).Collection(repository.CollectionName)

	var docCount, err = collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	return docCount, nil
}
