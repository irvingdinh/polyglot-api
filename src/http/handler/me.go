package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/irvingdinh/polyglot-api/src/repository"
	"github.com/irvingdinh/polyglot-api/src/service"
	"net/http"
)

type MeHandler interface {
	Me(c *gin.Context)
}

func NewMeHandler(
	repositoryObj repository.Repository,
) MeHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &meHandlerImpl{
		repository: repositoryObj,

		validate: validate,
	}
}

type meHandlerImpl struct {
	repository repository.Repository

	validate *validator.Validate
}

func (i *meHandlerImpl) Me(c *gin.Context) {
	ctx := c.Request.Context()

	userId, err := service.AccessToken{}.GinToUserID(ctx, i.repository.AccessTokenRepository(), c)
	if err != nil {
		if errors.Is(err, service.ErrUnauthenticated) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user, err := i.repository.UserRepository().FindOneByID(ctx, userId.Hex())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":       user.ID.Hex(),
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
	})
}
