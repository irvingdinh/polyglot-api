package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/irvingdinh/polyglot-api/src/model"
	"github.com/irvingdinh/polyglot-api/src/repository"
)

type UserHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func NewUserHandler(
	repositoryObj repository.Repository,
) UserHandler {
	validate := validator.New(validator.WithRequiredStructEnabled())

	return &userHandlerImpl{
		repository: repositoryObj,

		validate: validate,
	}
}

type userHandlerImpl struct {
	repository repository.Repository

	validate *validator.Validate
}

func (i *userHandlerImpl) Register(c *gin.Context) {
	ctx := c.Request.Context()

	type RequestBody struct {
		Password string `validate:"required,gte=6"`
		Name     string `validate:"required,gte=6,lte=64"`
		Email    string `validate:"required,email,lte=256"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := i.validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	results, err := i.repository.UserRepository().FindAll(ctx, bson.D{
		{"email", body.Email},
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(results) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is taken",
		})

		return
	}

	passwordAsBytes, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	user := model.User{
		Username: body.Email,
		Password: string(passwordAsBytes),
		Name:     body.Name,
		Email:    body.Email,
	}

	result, err := i.repository.UserRepository().InsertOne(ctx, user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"_id": result.InsertedID})
}

func (i *userHandlerImpl) Login(c *gin.Context) {
	ctx := c.Request.Context()

	type RequestBody struct {
		Email    string `validate:"required,email,lte=256"`
		Password string `validate:"required,gte=6"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := i.validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	results, err := i.repository.UserRepository().FindAll(ctx, bson.D{
		{"email", body.Email},
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if len(results) != 1 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := results[0]

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, secret, err := i.repository.AccessTokenRepository().InsertOne(ctx, user.ID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	plainText := fmt.Sprintf("%s-%s", accessToken.ID.Hex(), secret)

	c.SetCookie("POLYGLOT", plainText, 60*60*24*30, "/", c.Request.Host, false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": plainText,
	})
}
