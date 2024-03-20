package repository

import (
	"mygram/model"

	"gorm.io/gorm"
)

// IUserRepository
type IUserRepository interface {
	Create(model.User) (model.User, error)
	GetByUsername(string) (model.User, error)
}

// userRepository
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(newUser model.User) (model.User, error) {

	tx := ur.db.Create(&newUser)

	return newUser, tx.Error

}
func (ur *UserRepository) GetByUsername(username string) (model.User, error) {
	var user model.User
	tx := ur.db.First(&user, "username = ?", username)
	return user, tx.Error

}
