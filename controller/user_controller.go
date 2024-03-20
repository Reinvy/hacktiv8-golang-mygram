package controller

import (
	"mygram/helper"
	"mygram/model"
	"mygram/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		userRepository: userRepository,
	}
}

func (uc *userController) Login(c *gin.Context) {

	var requestedUser model.User
	err := c.ShouldBindJSON(&requestedUser)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	user, err := uc.userRepository.GetByUsername(requestedUser.Username)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	if !helper.HashMatched([]byte(user.Password), []byte(requestedUser.Password)) {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	token, err := helper.GenerateJWTToken(true, user.Username)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	c.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"token": token,
	}, ""))
}

func (uc *userController) Register(c *gin.Context) {
	var userInput model.UserInput
	var newUser model.User

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	hashedPassword, err := helper.Hash([]byte(userInput.Password))
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	newUser.Username = userInput.Username
	newUser.Email = userInput.Email
	newUser.Age = userInput.Age
	newUser.Password = string(hashedPassword)

	user, err := uc.userRepository.Create(newUser)
	if err != nil {
		var r model.Response = model.Response{
			Success: false,
			Error:   err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	c.JSON(http.StatusCreated, helper.CreateResponse(true, gin.H{
		"id": user.ID,
	}, ""))

}
