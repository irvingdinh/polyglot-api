package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccessToken struct {
	ID primitive.ObjectID `bson:"_id"`

	UserID primitive.ObjectID `bson:"user_id"`

	Secret string `bson:"secret"`

	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
	ExpiredAt primitive.DateTime `bson:"expired_at"`
}
