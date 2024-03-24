package dto

import (
	"mygram/domain/entity"
	"time"
)

type SocialMediaResponse struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	Name           string       `json:"name"`
	SocialMediaURL string       `json:"social_media_url"`
	UserID         uint         `json:"user_id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	User           UserResponse `json:"user"`
}

func (smr *SocialMediaResponse) FromEntity(socialMedia entity.SocialMedia) {
	smr.ID = socialMedia.ID
	smr.Name = socialMedia.Name
	smr.SocialMediaURL = socialMedia.SocialMediaURL
	smr.UserID = socialMedia.UserID
	smr.CreatedAt = socialMedia.CreatedAt
	smr.UpdatedAt = socialMedia.UpdatedAt
	smr.User.FromEntity(socialMedia.User)
}

func (smr *SocialMediaResponse) ToEntity() entity.SocialMedia {
	return entity.SocialMedia{
		ID:             smr.ID,
		Name:           smr.Name,
		SocialMediaURL: smr.SocialMediaURL,
		UserID:         smr.UserID,
		CreatedAt:      smr.CreatedAt,
		UpdatedAt:      smr.UpdatedAt,
	}
}

type SocialMediaPost struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (smp *SocialMediaPost) ToEntity() entity.SocialMedia {
	return entity.SocialMedia{
		Name:           smp.Name,
		SocialMediaURL: smp.SocialMediaURL,
	}
}

type SocialMediaUpdate struct {
	Name           string `json:"name"`
	SocialMediaURL string `json:"social_media_url"`
}

func (smu *SocialMediaUpdate) ToEntity() entity.SocialMedia {
	return entity.SocialMedia{
		Name:           smu.Name,
		SocialMediaURL: smu.SocialMediaURL,
	}
}
