package repository

import (
	"mygram/domain/entity"

	"gorm.io/gorm"
)

type ISocialMediaRepository interface {
	Create(newSocialMedia entity.SocialMedia) (entity.SocialMedia, error)
	GetAll() ([]entity.SocialMedia, error)
	GetById(id uint) (entity.SocialMedia, error)
	DeleteByID(id uint) (entity.SocialMedia, error)
	Update(socialMedia entity.SocialMedia) (entity.SocialMedia, error)
}
type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (s *SocialMediaRepository) Create(newSocialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	tx := s.db.Create(&newSocialMedia)
	return newSocialMedia, tx.Error
}

func (s *SocialMediaRepository) GetAll() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia
	tx := s.db.Preload("User").Find(&socialMedias)
	return socialMedias, tx.Error
}

func (s *SocialMediaRepository) GetById(id uint) (entity.SocialMedia, error) {
	var socialMedia entity.SocialMedia
	tx := s.db.Preload("User").First(&socialMedia, id)
	return socialMedia, tx.Error
}

func (s *SocialMediaRepository) DeleteByID(id uint) (entity.SocialMedia, error) {
	var socialMedia entity.SocialMedia
	tx := s.db.Where("id = ?", id).Delete(&socialMedia)
	return socialMedia, tx.Error
}

func (s *SocialMediaRepository) Update(socialMedia entity.SocialMedia) (entity.SocialMedia, error) {
	tx := s.db.Save(&socialMedia)
	return socialMedia, tx.Error
}
