package repository

import (
	"context"

	mongodb "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const oliveOilsCollectionName = "oliveOils"

// Repository represents olive collection operations interface
type Repository interface {
	mongodb.Repository
	GetOliveOilAmount() (float32, error)
}

// OliveOil respresents the struct for mongo db operations for oliveOils collection
type OliveOil struct {
	mongodb.MongoBase
}

//GetOliveOilAmount returns the total amount of oils
func (repository *OliveOil) GetOliveOilAmount() (float32, error) {
	collection := repository.GetCollection()

	//todo some stuff for calculating total amount of oliveOils

	var docCount, err = collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	return float32(docCount), nil
}

//New creates a new olive repository object
func New(c *mongo.Client, dbName string) *OliveOil {
	var repository = OliveOil{}
	repository.CollectionName = oliveOilsCollectionName
	repository.DBName = dbName
	repository.MongoBase.SetClient(c)

	return &repository
}
