package model

import (
	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User represents user model
type User struct {
	ID          *primitive.ObjectID `bson:"_id"`
	Username    string              `bson:"username"`
	Email       string              `bson:"email"`
	FullName    string              `bson:"full_name"`
	AccessLevel byte                `bson:"access_level"`
	Status      enum.UserStatusType `bson:"status"`
	//..
	//..
}
