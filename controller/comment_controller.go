package controller

import (
	"mygram/model/dto"
	"mygram/model/entity"
	"mygram/model/repository"
	"mygram/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentController struct {
	commentRepository repository.ICommentRepository
}

func NewCommentController(db *gorm.DB) *commentController {
	return &commentController{
		commentRepository: repository.NewCommentRepository(db),
	}
}

func (cc *commentController) Create(c *gin.Context) {
	var commentRequest dto.CommentRequest
	var newComment entity.Comment

	err := c.ShouldBindJSON(&commentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	claims, _ := util.GetJWTClaims(strings.Split(c.GetHeader("Authorization"), " ")[1])
	userID, _ := util.GetSubFromClaims(claims)

	newComment = commentRequest.ToEntity()
	newComment.UserID = uint(userID)

	comment, err := cc.commentRepository.Create(newComment)
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
		Message: "Comment created successfully",
		Data: gin.H{
			"id":         comment.ID,
			"message":    comment.Message,
			"photo_id":   comment.PhotoID,
			"user_id":    comment.UserID,
			"created_at": comment.CreatedAt},
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}

func (cc *commentController) GetAll(c *gin.Context) {
	comments, err := cc.commentRepository.GetAll()
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
		Message: "Get all comments successfully",
		Data:    comments,
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}
