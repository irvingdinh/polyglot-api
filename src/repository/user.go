package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/irvingdinh/polyglot-api/src/model"
)

type UserRepository interface {
	FindAll(ctx context.Context, filter bson.D) ([]model.User, error)
	FindOneByID(ctx context.Context, idAsString string) (*model.User, error)
	InsertOne(ctx context.Context, user model.User) (*mongo.InsertOneResult, error)
}

func NewUserRepository(
	mongoClient *mongo.Client,
) UserRepository {
	return &userRepositoryImpl{
		mongoClient: mongoClient,
	}
}

type userRepositoryImpl struct {
	mongoClient *mongo.Client
}

func (i *userRepositoryImpl) FindAll(ctx context.Context, filter bson.D) ([]model.User, error) {
	col := i.mongoClient.Database(DB).Collection("users")

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var results []model.User
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (i *userRepositoryImpl) FindOneByID(ctx context.Context, idAsString string) (*model.User, error) {
	col := i.mongoClient.Database(DB).Collection("users")

	id, err := primitive.ObjectIDFromHex(idAsString)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var result model.User

	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (i *userRepositoryImpl) InsertOne(ctx context.Context, user model.User) (*mongo.InsertOneResult, error) {
	col := i.mongoClient.Database(DB).Collection("users")

	user.ID = primitive.NewObjectID()
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}
