package repository

import (
	"mygram/domain/entity"

	"gorm.io/gorm"
)

type IPhotoRepository interface {
	Create(entity.Photo) (entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	GetById(uint) (entity.Photo, error)
	Update(entity.Photo) (entity.Photo, error)
	Delete(entity.Photo) (entity.Photo, error)
}

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) Create(newPhoto entity.Photo) (entity.Photo, error) {
	tx := pr.db.Preload("User").Create(&newPhoto)
	return newPhoto, tx.Error
}

func (pr *PhotoRepository) GetAll() ([]entity.Photo, error) {
	var photos []entity.Photo
	tx := pr.db.Preload("User").Find(&photos)
	return photos, tx.Error
}

func (pr *PhotoRepository) GetById(id uint) (entity.Photo, error) {
	var photo entity.Photo
	tx := pr.db.Preload("User").First(&photo, id)
	return photo, tx.Error
}

func (pr *PhotoRepository) Update(photo entity.Photo) (entity.Photo, error) {
	tx := pr.db.Save(&photo)
	return photo, tx.Error
}

func (pr *PhotoRepository) Delete(photo entity.Photo) (entity.Photo, error) {
	tx := pr.db.Delete(&photo)
	return photo, tx.Error
}
