package repository

import "go.mongodb.org/mongo-driver/mongo"

const DB = "polyglot"

type Repository interface {
	AccessTokenRepository() AccessTokenRepository
	UserRepository() UserRepository
}

func New(
	mongoClient *mongo.Client,
) Repository {
	return &repositoryImpl{
		accessTokenRepository: NewAccessTokenRepository(mongoClient),
		userRepository:        NewUserRepository(mongoClient),
	}
}

type repositoryImpl struct {
	accessTokenRepository AccessTokenRepository
	userRepository        UserRepository
}

func (i *repositoryImpl) AccessTokenRepository() AccessTokenRepository {
	return i.accessTokenRepository
}

func (i *repositoryImpl) UserRepository() UserRepository {
	return i.userRepository
}
