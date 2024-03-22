package service

import (
	"mygram/domain/dto"
	"mygram/domain/entity"
	"mygram/domain/repository"
	"mygram/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentService struct {
	commentRepository repository.ICommentRepository
}

func NewCommentService(db *gorm.DB) *commentService {
	return &commentService{
		commentRepository: repository.NewCommentRepository(db),
	}
}

func (cc *commentService) Create(c *gin.Context) {
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

func (cc *commentService) GetAll(c *gin.Context) {
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