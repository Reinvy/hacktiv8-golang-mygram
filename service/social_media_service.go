package service

import (
	"mygram/domain/dto"
	"mygram/domain/entity"
	"mygram/domain/repository"
	"mygram/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type socialMediaService struct {
	socialMediaRepository repository.ISocialMediaRepository
}

func NewSocialMediaService(db *gorm.DB) *socialMediaService {
	return &socialMediaService{
		socialMediaRepository: repository.NewSocialMediaRepository(db),
	}
}

func (sms *socialMediaService) GetAll(c *gin.Context) {
	socialMedias, err := sms.socialMediaRepository.GetAll()
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var socialMediaResponses []dto.SocialMediaResponse
	for _, socialMedia := range socialMedias {
		var socialMediaResponse dto.SocialMediaResponse
		socialMediaResponse.FromEntity(socialMedia)
		socialMediaResponses = append(socialMediaResponses, socialMediaResponse)
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Get all social media successfully",
		Data:    socialMediaResponses,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (sms *socialMediaService) Create(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var socialMediaRequest dto.SocialMediaPost
	var newSocialMedia entity.SocialMedia

	err := c.ShouldBindJSON(&socialMediaRequest)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	claims, _ := util.GetJWTClaims(token)
	userID, _ := util.GetSubFromClaims(claims)

	newSocialMedia = socialMediaRequest.ToEntity()
	newSocialMedia.UserID = uint(userID)

	socialMedia, err := sms.socialMediaRepository.Create(newSocialMedia)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var socialMediaResponse dto.SocialMediaResponse
	socialMediaResponse.FromEntity(socialMedia)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Create social media successfully",
		Data:    socialMediaResponse,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (sms *socialMediaService) Update(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)

	var socialMediaRequest dto.SocialMediaUpdate
	err := c.ShouldBindJSON(&socialMediaRequest)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	claims, _ := util.GetJWTClaims(token)
	userID, _ := util.GetSubFromClaims(claims)

	socialMedia, err := sms.socialMediaRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if uint(userID) != socialMedia.UserID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	socialMedia.Name = socialMediaRequest.Name
	socialMedia.SocialMediaURL = socialMediaRequest.SocialMediaURL

	updatedSocialMedia, err := sms.socialMediaRepository.Update(socialMedia)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var socialMediaResponse dto.SocialMediaResponse
	socialMediaResponse.FromEntity(updatedSocialMedia)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Update social media successfully",
		Data:    socialMediaResponse,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (sms *socialMediaService) Delete(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)

	claims, _ := util.GetJWTClaims(token)
	userID, _ := util.GetSubFromClaims(claims)

	socialMedia, err := sms.socialMediaRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if uint(userID) != socialMedia.UserID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	socialMedia, err = sms.socialMediaRepository.DeleteByID(uint(id))
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
		Message: "Delete social media successfully",
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (sms *socialMediaService) GetSocialMediaByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)

	socialMedia, err := sms.socialMediaRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var socialMediaResponse dto.SocialMediaResponse
	socialMediaResponse.FromEntity(socialMedia)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Get social media successfully",
		Data:    socialMediaResponse,
	}

	c.AbortWithStatusJSON(http.StatusOK, r)
}
