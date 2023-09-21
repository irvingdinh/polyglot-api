package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `bson:"_id"`

	Username string `bson:"username"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`

	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
