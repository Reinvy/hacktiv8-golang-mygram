package dto

import (
	"mygram/domain/entity"
	"time"
)

type CommentDTO interface {
	FromEntity(entity.Comment)
	ToEntity() entity.Comment
}

type CommentResponse struct {
	ID        uint          `gorm:"primaryKey" json:"id"`
	Message   string        `json:"message"`
	UserID    uint          `json:"user_id"`
	PhotoID   uint          `json:"photo_id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	User      UserResponse  `json:"user"`
	Photo     PhotoResponse `json:"photo"`
}

type CommentRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoID uint   `json:"photo_id" binding:"required"`
}

func (cr *CommentResponse) FromEntity(comment entity.Comment) {
	cr.ID = comment.ID
	cr.Message = comment.Message
	cr.UserID = comment.UserID
	cr.PhotoID = comment.PhotoID
	cr.CreatedAt = comment.CreatedAt
	cr.UpdatedAt = comment.UpdatedAt
	cr.User.FromEntity(comment.User)
	cr.Photo.FromEntity(comment.Photo)
}

func (cr *CommentResponse) ToEntity() entity.Comment {
	return entity.Comment{
		ID:        cr.ID,
		Message:   cr.Message,
		UserID:    cr.UserID,
		PhotoID:   cr.PhotoID,
		CreatedAt: cr.CreatedAt,
		UpdatedAt: cr.UpdatedAt,
		Photo:     cr.Photo.ToEntity(),
	}
}

func (cr *CommentRequest) ToEntity() entity.Comment {
	return entity.Comment{
		Message: cr.Message,
		PhotoID: cr.PhotoID,
	}
}
