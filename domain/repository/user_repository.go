package repository

import (
	"mygram/domain/entity"

	"gorm.io/gorm"
)

// IUserRepository
type IUserRepository interface {
	Create(entity.User) (entity.User, error)
	Update(entity.User) (entity.User, error)
	Delete(entity.User) (entity.User, error)
	GetByUsername(string) (entity.User, error)
	GetById(uint) (entity.User, error)
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

func (ur *UserRepository) Create(newUser entity.User) (entity.User, error) {
	tx := ur.db.Create(&newUser)
	return newUser, tx.Error
}
func (ur *UserRepository) GetByUsername(username string) (entity.User, error) {
	var user entity.User
	tx := ur.db.First(&user, "username = ?", username)
	return user, tx.Error
}

func (ur *UserRepository) GetById(id uint) (entity.User, error) {
	var user entity.User
	tx := ur.db.First(&user, id)
	return user, tx.Error
}

func (ur *UserRepository) Update(user entity.User) (entity.User, error) {
	tx := ur.db.Save(&user)
	return user, tx.Error
}

func (ur *UserRepository) Delete(user entity.User) (entity.User, error) {
	tx := ur.db.Delete(&user)
	return user, tx.Error
}
