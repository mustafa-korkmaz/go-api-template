package repository

import (
	"context"

	mongodb "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollectionName = "users"

// Repository represents olive collection operations interface
type Repository interface {
	mongodb.Repository
	GetOlivesCount() (int64, error)
}

// User respresents the struct for mongo db operations for olive collection
type User struct {
	mongodb.MongoBase
}

//GetOlivesCount returns the document count for olive collection
func (u *User) GetOlivesCount() (int64, error) {
	collection := u.GetCollection()

	var docCount, err = collection.CountDocuments(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	return docCount, nil
}

//New creates a new olive repository object
func New(c *mongo.Client, dbName string) *User {
	var user = User{}
	user.CollectionName = userCollectionName
	user.DBName = dbName
	user.SetClient(c)

	return &user
}
