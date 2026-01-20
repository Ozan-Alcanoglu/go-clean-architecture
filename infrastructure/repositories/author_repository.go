package repositories

import (
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) interfaces.AuthorRepository {
	return &authorRepository{db: db}
}

func (ar *authorRepository) CreateAuthor(author *models.Author) error {

	return ar.db.Create(author).Error

}

func (ar *authorRepository) UpdateAuthor(author *models.Author) (*models.Author, error) {

	if err := ar.db.Save(&author).Error; err != nil {
		return nil, err
	}

	var updated models.Author

	if err := ar.db.Where("id = ?", author.Id).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil

}

func (ar *authorRepository) GetById(id uuid.UUID) (*models.Author, error) {

	var author models.Author

	if err := ar.db.Where("id = ?", id).First(&author).Error; err != nil {
		return nil, err
	}

	return &author, nil

}

func (ar *authorRepository) DeleteAuthor(id uuid.UUID) error {

	err := ar.db.Where("id= ?", id).Delete(&models.Author{}).Error

	if err != nil {
		return err
	}

	return nil

}
