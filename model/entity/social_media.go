package entity

import "gorm.io/gorm"

type SocialMedia struct {
	gorm.Model
	Name           string
	SocialMediaURL string
	UserID         uint `foreignKey:"UserID"`
}
