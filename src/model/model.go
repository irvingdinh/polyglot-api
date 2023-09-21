package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	ID primitive.ObjectID `bson:"_id"`

	CreatedAt primitive.DateTime `bson:"created_at"`
	UpdatedAt primitive.DateTime `bson:"updated_at"`
}
