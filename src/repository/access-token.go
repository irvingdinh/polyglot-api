package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/irvingdinh/polyglot-api/src/model"
	"github.com/irvingdinh/polyglot-api/src/utils"
)

type AccessTokenRepository interface {
	FindOneByID(ctx context.Context, idAsString string) (*model.AccessToken, error)
	InsertOne(ctx context.Context, userId primitive.ObjectID) (*model.AccessToken, string, error)
	DeleteOneByID(ctx context.Context, idAsString string) error
}

func NewAccessTokenRepository(
	mongoClient *mongo.Client,
) AccessTokenRepository {
	return &accessTokenRepositoryImpl{
		mongoClient: mongoClient,
	}
}

type accessTokenRepositoryImpl struct {
	mongoClient *mongo.Client
}

func (i *accessTokenRepositoryImpl) FindOneByID(ctx context.Context, idAsString string) (*model.AccessToken, error) {
	col := i.mongoClient.Database(DB).Collection("access_tokens")

	id, err := primitive.ObjectIDFromHex(idAsString)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", id}}

	var result model.AccessToken

	err = col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (i *accessTokenRepositoryImpl) InsertOne(ctx context.Context, userId primitive.ObjectID) (*model.AccessToken, string, error) {
	col := i.mongoClient.Database(DB).Collection("access_tokens")

	secret := utils.String{}.Random(32)

	secretAsBytes, err := bcrypt.GenerateFromPassword([]byte(secret), 14)
	if err != nil {
		return nil, "", err
	}

	accessToken := model.AccessToken{
		ID:        primitive.NewObjectID(),
		UserID:    userId,
		Secret:    string(secretAsBytes),
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		ExpiredAt: primitive.NewDateTimeFromTime(time.Now().AddDate(0, 1, 0)),
	}

	_, err = col.InsertOne(ctx, accessToken)
	if err != nil {
		return nil, "", err
	}

	return &accessToken, secret, nil
}

func (i *accessTokenRepositoryImpl) DeleteOneByID(ctx context.Context, idAsString string) error {
	col := i.mongoClient.Database(DB).Collection("access_tokens")

	id, err := primitive.ObjectIDFromHex(idAsString)
	if err != nil {
		return err
	}

	_, err = col.DeleteOne(ctx, bson.D{{"_id", id}})

	return err
}
