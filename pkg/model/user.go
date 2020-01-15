package model

import (
	"time"

	"github.com/mustafa-korkmaz/goapitemplate/pkg/enum"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User represents user model
type User struct {
	ID          *primitive.ObjectID  `bson:"_id"`
	Username    string               `bson:"username"`
	Email       string               `bson:"email"`
	FullName    string               `bson:"full_name"`
	Password    string               `bson:"password"`
	AccessLevel enum.AccessLevelType `bson:"access_level"`
	Status      enum.UserStatusType  `bson:"status"`
	CreatedAt   time.Time            `bson:"created_at"`
	//..
	//..
}

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID          string
	Username    string
	Email       string
	AccessLevel enum.AccessLevelType
}
