package entity

import (
	"time"
)

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Message   string    `json:"message"`
	UserID    uint      `foreignKey:"UserID"`
	PhotoID   uint      `json:"photo_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	User      User      `gorm:"foreignKey:UserID"`
	Photo     Photo     `gorm:"foreignKey:PhotoID"`
}
