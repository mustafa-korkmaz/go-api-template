package repository

import (
	"context"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/model"
	mongodb "github.com/mustafa-korkmaz/goapitemplate/pkg/mongodb/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollectionName = "users"

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

//GetUserByEmail returns user by email. returns nil if user not exists
func (u *User) GetUserByEmail(email string) *model.User {

	filter := bson.D{primitive.E{Key: "email", Value: email}}

	var res = u.FindOnebyDocument(filter)

	if res == nil {
		//record not found
		return nil
	}

	var user = new(model.User)

	res.Decode(user)

	return user
}

//GetUserByUsername returns user by username. returns nil if user not exists
func (u *User) GetUserByUsername(email string) *model.User {

	var user = new(model.User)

	//var tag = model.GetBsonTag(user, "username")
	filter := bson.D{primitive.E{Key: "username", Value: email}}

	var res = u.FindOnebyDocument(filter)

	if res == nil {
		//record not found
		return nil
	}

	res.Decode(user)

	return user
}

//Register creates new user
func (u *User) Register(userModel *model.User) error {
	return nil
}

//New creates a new olive repository object
func New(c *mongo.Client, dbName string) *User {
	var user = User{}
	user.CollectionName = userCollectionName
	user.DBName = dbName
	user.SetClient(c)

	return &user
}
