package repositories

import (
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type genreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) interfaces.GenreRepository {

	return &genreRepository{db: db}

}

func (ur *genreRepository) CreateGenre(genre *models.Genre) error {

	return ur.db.Create(genre).Error

}

func (ur *genreRepository) GetById(id uuid.UUID) (*models.Genre, error) {

	var genre models.Genre

	err := ur.db.Where("id= ?", id).First(&genre).Error

	if err != nil {
		return nil, err
	}

	return &genre, nil

}

func (ur *genreRepository) DeleteGenreById(id uuid.UUID) error {
	err := ur.db.Where("id = ?", id).Delete(&models.Genre{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *genreRepository) UpdateGenre(genre *models.Genre) (*models.Genre, error) {
	if err := ur.db.Save(genre).Error; err != nil {
		return nil, err
	}

	var updated models.Genre
	if err := ur.db.Where("id = ?", genre.ID).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}
