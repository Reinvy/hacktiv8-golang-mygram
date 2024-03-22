package controller

import (
	"mygram/model/dto"
	"mygram/model/entity"
	"mygram/model/repository"
	"mygram/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserController(db *gorm.DB) *userController {
	return &userController{
		userRepository: repository.NewUserRepository(db),
	}
}

func (uc *userController) Login(c *gin.Context) {
	var userLogin dto.UserLogin
	err := c.ShouldBindJSON(&userLogin)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	user, err := uc.userRepository.GetByUsername(userLogin.Username)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if !util.HashMatched([]byte(user.Password), []byte(userLogin.Password)) {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Wrong password",
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	token, err := util.GenerateJWTToken(true, user.ID)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Login success",
		Data: gin.H{
			"token": token,
		},
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}

func (uc *userController) Register(c *gin.Context) {
	var userInput dto.UserRegister
	var newUser entity.User

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: strings.Split(err.Error(), "Error:")[1],
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	hashedPassword, err := util.Hash([]byte(userInput.Password))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	userInput.Password = string(hashedPassword)
	newUser = userInput.ToEntity()

	user, err := uc.userRepository.Create(newUser)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: strings.Split(err.Error(), "ERROR: ")[1],
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var userResponse dto.UserResponse
	userResponse.FromEntity(user)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "User created successfully",
		Data:    userResponse,
	}
	c.AbortWithStatusJSON(http.StatusCreated, r)

}

func (uc *userController) Update(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsId := c.Param("id")
	userInput := dto.UserUpdate{}

	id, _ := strconv.ParseUint(paramsId, 10, 64)

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	user, err := uc.userRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	claims, _ := util.GetJWTClaims(strings.Split(token, " ")[1])
	userID, err := util.GetSubFromClaims(claims)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if uint(id) != userID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	user.Username = userInput.Username
	user.Email = userInput.Email

	updatedUser, err := uc.userRepository.Update(user)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var userResponse dto.UserResponse
	userResponse.FromEntity(updatedUser)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "User updated successfully",
		Data:    userResponse,
	}
	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (uc *userController) Delete(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsId := c.Param("id")
	id, _ := strconv.ParseUint(paramsId, 10, 64)

	claims, _ := util.GetJWTClaims(strings.Split(token, " ")[1])
	userID, err := util.GetSubFromClaims(claims)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	if uint(id) != userID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	user, err := uc.userRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusNotFound, r)
		return
	}

	_, err = uc.userRepository.Delete(user)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, r)
		return
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "User deleted successfully",
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}
