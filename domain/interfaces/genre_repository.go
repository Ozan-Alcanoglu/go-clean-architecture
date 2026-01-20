package interfaces

import (
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type GenreRepository interface {
	CreateGenre(genre *models.Genre) error
	GetById(id uuid.UUID) (*models.Genre, error)
	DeleteGenreById(id uuid.UUID) error
	UpdateGenre(genre *models.Genre) (*models.Genre, error)
}
