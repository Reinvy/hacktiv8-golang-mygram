package model

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaURL string
	UserID         uint
}
