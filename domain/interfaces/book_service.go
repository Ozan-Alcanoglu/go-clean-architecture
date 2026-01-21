package interfaces

import (
	"gobackend/domain/dtos"

	"github.com/google/uuid"
)

type BookService interface {
	CreateBook(createBookDto *dtos.CreateBookDTO) (*dtos.BookResponseDTO, error)
	GetBookById(id uuid.UUID) (*dtos.BookResponseDTO, error)
	UpdateBook(id uuid.UUID, updateBookDto *dtos.UpdateBookDTO) (*dtos.BookResponseDTO, error)
	DeleteBook(id uuid.UUID) error
}
