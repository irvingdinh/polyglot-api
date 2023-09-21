package service

import (
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"

	"github.com/irvingdinh/polyglot-api/src/repository"
)

type AccessToken struct{}

func (i AccessToken) GinToUserID(
	ctx context.Context,
	accessTokenRepository repository.AccessTokenRepository,
	c *gin.Context,
) (*primitive.ObjectID, error) {
	plainTextToken, _ := c.Cookie("POLYGLOT")
	if plainTextToken == "" {
		plainTextToken = c.GetHeader("Authorization")
	}
	if plainTextToken == "" {
		return nil, ErrUnauthenticated
	}

	fragments := strings.Split(plainTextToken, "-")
	if len(fragments) != 2 {
		return nil, ErrUnauthenticated
	}

	accessToken, err := accessTokenRepository.FindOneByID(ctx, fragments[0])
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrUnauthenticated
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accessToken.Secret), []byte(fragments[1]))
	if err != nil {
		return nil, ErrUnauthenticated
	}

	return &accessToken.UserID, nil
}
