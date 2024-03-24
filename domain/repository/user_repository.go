package repository

import (
	"mygram/domain/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(entity.User) (entity.User, error)
	Update(entity.User) (entity.User, error)
	Delete(entity.User) (entity.User, error)
	GetByUsername(string) (entity.User, error)
	GetByEmail(string) (entity.User, error)
	GetById(uint) (entity.User, error)
}

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

func (ur *UserRepository) GetByEmail(email string) (entity.User, error) {
	var user entity.User
	tx := ur.db.First(&user, "email = ?", email)
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
	var photos []entity.Photo

	tx := ur.db.Where("user_id = ?", user.ID).Delete(&entity.Comment{})
	if tx.Error != nil {
		return user, tx.Error
	}

	tx = ur.db.Where("user_id = ?", user.ID).Find(&photos)
	if tx.Error != nil {
		return user, tx.Error
	}
	for _, photo := range photos {
		tx = ur.db.Where("photo_id = ?", photo.ID).Delete(&entity.Comment{})
		if tx.Error != nil {
			return user, tx.Error
		}
	}

	tx = ur.db.Where("photo_id = ?", user.ID).Delete(&entity.Comment{})
	if tx.Error != nil {
		return user, tx.Error
	}

	tx = ur.db.Where("user_id = ?", user.ID).Delete(&entity.Photo{})
	if tx.Error != nil {
		return user, tx.Error
	}

	tx = ur.db.Where("user_id = ?", user.ID).Delete(&entity.SocialMedia{})
	if tx.Error != nil {
		return user, tx.Error
	}

	tx = ur.db.Delete(&user)
	return user, tx.Error
}
