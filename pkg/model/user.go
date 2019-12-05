package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//User represents user model
type User struct {
	ID          *primitive.ObjectID `bson:"_id"`
	Username    string              `bson:"username"`
	Email       string              `bson:"email"`
	AccessLevel string              `bson:"access_level"`
	//..
	//..
}
