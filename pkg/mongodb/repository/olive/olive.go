package repository

import (
	"context"

	mongodb "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const oliveCollectionName = "olives"

// Repository represents olive collection operations interface
type Repository interface {
	mongodb.Repository
	GetOlivesCount() (int64, error)
}

// Olive respresents the struct for mongo db operations for olive collection
type Olive struct {
	mongodb.MongoBase
}

//GetOlivesCount returns the document count for olive collection
func (o *Olive) GetOlivesCount() (int64, error) {
	collection := o.GetCollection()

	var docCount, err = collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	return docCount, nil
}

//New creates a new olive repository object
func New(c *mongo.Client, dbName string) *Olive {
	var olive = Olive{}
	olive.CollectionName = oliveCollectionName
	olive.DBName = dbName
	olive.SetClient(c)

	return &olive
}
