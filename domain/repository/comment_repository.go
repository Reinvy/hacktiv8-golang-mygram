package repository

import (
	"mygram/domain/entity"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(newComment entity.Comment) (entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	GetById(id uint) (entity.Comment, error)
	Update(comment entity.Comment) (entity.Comment, error)
	Delete(comment entity.Comment) (entity.Comment, error)
	DeleteByID(id uint) (entity.Comment, error)
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) Create(newComment entity.Comment) (entity.Comment, error) {
	tx := cr.db.Create(&newComment)
	return newComment, tx.Error
}

func (cr *CommentRepository) GetAll() ([]entity.Comment, error) {
	var comments []entity.Comment
	tx := cr.db.Preload("Photo").Preload("User").Find(&comments)
	return comments, tx.Error
}

func (cr *CommentRepository) GetById(id uint) (entity.Comment, error) {
	var comment entity.Comment
	tx := cr.db.Preload("Photo").Preload("User").First(&comment, id)
	return comment, tx.Error
}

func (cr *CommentRepository) Update(comment entity.Comment) (entity.Comment, error) {
	tx := cr.db.Save(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) Delete(comment entity.Comment) (entity.Comment, error) {
	tx := cr.db.Delete(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) DeleteByID(id uint) (entity.Comment, error) {
	var comment entity.Comment
	tx := cr.db.Where("id = ?", id).Delete(&comment)
	return comment, tx.Error
}
