package interfaces

import (
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type BookRepository interface {
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) (*models.Book, error)
	DeleteBook(id uuid.UUID) error
	GetById(id uuid.UUID) (*models.Book, error)
}
