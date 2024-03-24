package dto

import (
	"mygram/domain/entity"
	"time"
)

type UserDTO interface {
	FromEntity(entity.User)
	ToEntity() entity.User
}

type UserResponse struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6"`
}

type UserRegister struct {
	Username string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6"`
	Age      int    `binding:"required,min=8"`
}

type UserUpdate struct {
	Username string `binding:"required"`
	Email    string `binding:"required,email"`
}

func (ur *UserRegister) ToEntity() entity.User {
	return entity.User{
		Username: ur.Username,
		Email:    ur.Email,
		Password: ur.Password,
		Age:      ur.Age,
	}
}

func (u *UserResponse) FromEntity(user entity.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Email = user.Email
	u.Age = user.Age
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt
}
