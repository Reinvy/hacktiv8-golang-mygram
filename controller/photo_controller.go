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

type photoController struct {
	photoRepository repository.IPhotoRepository
}

func NewPhotoController(db *gorm.DB) *photoController {
	return &photoController{
		photoRepository: repository.NewPhotoRepository(db),
	}
}

func (pc *photoController) Create(c *gin.Context) {
	var photoRequest dto.PhotoRequest
	var newPhoto entity.Photo

	err := c.ShouldBindJSON(&photoRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	claims, _ := util.GetJWTClaims(strings.Split(c.GetHeader("Authorization"), " ")[1])
	userID, _ := util.GetSubFromClaims(claims)

	newPhoto = photoRequest.ToEntity()
	newPhoto.UserID = uint(userID)

	photo, err := pc.photoRepository.Create(newPhoto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Photo created successfully",
		Data: gin.H{
			"id":         photo.ID,
			"title":      photo.Title,
			"caption":    photo.Caption,
			"photo_url":  photo.PhotoUrl,
			"user_id":    photo.UserID,
			"created_at": photo.CreatedAt,
		},
	}

	c.AbortWithStatusJSON(http.StatusCreated, r)

}

func (pc *photoController) GetAll(c *gin.Context) {
	photos, err := pc.photoRepository.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	var photoResponses []dto.PhotoResponse
	for _, photo := range photos {
		var photoResponse dto.PhotoResponse
		photoResponse.FromEntity(photo)
		photoResponses = append(photoResponses, photoResponse)
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Photo retrieved successfully",
		Data:    photoResponses,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (pc *photoController) GetPhotoByID(c *gin.Context) {
	StringId := c.Param("id")

	id, err := strconv.ParseUint(StringId, 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.Response{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	photo, err := pc.photoRepository.GetById(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, dto.Response{
			Status:  "Error",
			Message: err.Error(),
		})
		return
	}

	var photoResponse dto.PhotoResponse
	photoResponse.FromEntity(photo)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Photo retrieved successfully",
		Data:    photoResponse,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (pc *photoController) Update(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsId := c.Param("id")
	var photoRequest dto.PhotoRequest

	id, _ := strconv.ParseUint(paramsId, 10, 64)
	err := c.ShouldBindJSON(&photoRequest)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	photo, err := pc.photoRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusNotFound, r)
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
	if photo.UserID != userID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	photo.Title = photoRequest.Title
	photo.Caption = photoRequest.Caption
	photo.PhotoUrl = photoRequest.PhotoUrl

	updatedPhoto, err := pc.photoRepository.Update(photo)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Photo updated successfully",
		Data: gin.H{
			"id":         updatedPhoto.ID,
			"title":      updatedPhoto.Title,
			"caption":    updatedPhoto.Caption,
			"photo_url":  updatedPhoto.PhotoUrl,
			"user_id":    updatedPhoto.UserID,
			"updated_at": updatedPhoto.UpdatedAt,
		},
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (pc *photoController) Delete(c *gin.Context) {
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

	photo, err := pc.photoRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusNotFound, r)
		return
	}

	if photo.UserID != userID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	_, err = pc.photoRepository.Delete(photo)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Photo deleted successfully",
	})

}
