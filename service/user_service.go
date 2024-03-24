package service

import (
	"mygram/domain/dto"
	"mygram/domain/entity"
	"mygram/domain/repository"
	"mygram/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userService struct {
	userRepository repository.IUserRepository
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{
		userRepository: repository.NewUserRepository(db),
	}
}

func (uc *userService) Login(c *gin.Context) {
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

	user, err := uc.userRepository.GetByEmail(userLogin.Email)
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

func (uc *userService) Register(c *gin.Context) {
	var userInput dto.UserRegister
	var newUser entity.User

	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
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

func (uc *userService) Update(c *gin.Context) {
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

	claims, _ := util.GetJWTClaims(token)
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

func (uc *userService) Delete(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsId := c.Param("id")
	id, _ := strconv.ParseUint(paramsId, 10, 64)

	claims, _ := util.GetJWTClaims(token)
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
