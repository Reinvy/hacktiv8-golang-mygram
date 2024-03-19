package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Message string
	UserID  uint
	PhotoID uint
}
