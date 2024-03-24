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

type commentService struct {
	commentRepository repository.ICommentRepository
}

func NewCommentService(db *gorm.DB) *commentService {
	return &commentService{
		commentRepository: repository.NewCommentRepository(db),
	}
}

func (cs *commentService) Create(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var commentRequest dto.CommentPost
	var newComment entity.Comment

	err := c.ShouldBindJSON(&commentRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}
	claims, _ := util.GetJWTClaims(token)
	userID, _ := util.GetSubFromClaims(claims)

	newComment = commentRequest.ToEntity()
	newComment.UserID = uint(userID)

	comment, err := cs.commentRepository.Create(newComment)
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

func (cs *commentService) GetAll(c *gin.Context) {
	comments, err := cs.commentRepository.GetAll()
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var commentResponses []dto.CommentResponse
	for _, comment := range comments {
		var commentResponse dto.CommentResponse
		commentResponse.FromEntity(comment)
		commentResponses = append(commentResponses, commentResponse)
	}

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Get all comments successfully",
		Data:    commentResponses,
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}

func (cs *commentService) GetCommentByID(c *gin.Context) {
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)
	comment, err := cs.commentRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	var commentResponse dto.CommentResponse
	commentResponse.FromEntity(comment)
	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Get comment successfully",
		Data:    commentResponse,
	}
	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (cs *commentService) Update(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)

	var commentRequest dto.CommentPut
	err := c.ShouldBindJSON(&commentRequest)
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

	comment, err := cs.commentRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if uint(userID) != comment.UserID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	comment.Message = commentRequest.Message

	updatedComment, err := cs.commentRepository.Update(comment)
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	commentResponse := dto.CommentResponse{}
	commentResponse.FromEntity(updatedComment)

	var r dto.Response = dto.Response{
		Status:  "Success",
		Message: "Update comment successfully",
		Data:    commentResponse,
	}
	c.AbortWithStatusJSON(http.StatusOK, r)
}

func (cs *commentService) Delete(c *gin.Context) {
	token := c.GetHeader("Authorization")
	paramsID := c.Param("id")
	id, _ := strconv.ParseUint(paramsID, 10, 64)

	claims, _ := util.GetJWTClaims(token)
	userID, _ := util.GetSubFromClaims(claims)

	comment, err := cs.commentRepository.GetById(uint(id))
	if err != nil {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, r)
		return
	}

	if uint(userID) != comment.UserID {
		var r dto.Response = dto.Response{
			Status:  "Error",
			Message: "Unauthorized",
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	comment, err = cs.commentRepository.DeleteByID(uint(id))
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
		Message: "Delete comment successfully",
	}
	c.AbortWithStatusJSON(http.StatusOK, r)

}
