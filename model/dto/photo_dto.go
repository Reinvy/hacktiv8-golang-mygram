package dto

import (
	"mygram/model/entity"
	"time"
)

type PhotoDTO interface {
	FromEntity(entity.Photo)
	ToEntity() entity.Photo
}

type PhotoResponse struct {
	ID        uint         `gorm:"primaryKey" json:"id"`
	Title     string       `json:"title"`
	Caption   string       `json:"caption"`
	PhotoUrl  string       `json:"photo_url"`
	UserID    uint         `json:"user_id"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func (pr *PhotoResponse) FromEntity(photo entity.Photo) {
	pr.ID = photo.ID
	pr.Title = photo.Title
	pr.Caption = photo.Caption
	pr.PhotoUrl = photo.PhotoUrl
	pr.UserID = photo.UserID
	pr.User.FromEntity(photo.User)
	pr.CreatedAt = photo.CreatedAt
	pr.UpdatedAt = photo.UpdatedAt
}

func (pr *PhotoResponse) ToEntity() entity.Photo {
	return entity.Photo{
		ID:        pr.ID,
		Title:     pr.Title,
		Caption:   pr.Caption,
		PhotoUrl:  pr.PhotoUrl,
		UserID:    pr.UserID,
		CreatedAt: pr.CreatedAt,
		UpdatedAt: pr.UpdatedAt,
	}
}

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

func (pr *PhotoRequest) ToEntity() entity.Photo {
	return entity.Photo{
		Title:    pr.Title,
		Caption:  pr.Caption,
		PhotoUrl: pr.PhotoUrl,
	}
}
