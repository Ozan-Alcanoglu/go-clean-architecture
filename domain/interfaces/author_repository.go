package interfaces

import (
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type AuthorRepository interface {
	CreateAuthor(author *models.Author) error
	UpdateAuthor(author *models.Author) (*models.Author, error)
	GetById(id uuid.UUID) (*models.Author, error)
	DeleteAuthor(id uuid.UUID) error
}
