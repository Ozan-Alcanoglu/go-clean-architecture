package interfaces

import (
	"gobackend/domain/dtos"

	"github.com/google/uuid"
)

type AuthorService interface {
	CreateAuthor(createAuthorDto *dtos.CreateAuthorDto) (*dtos.AuthorResponseDto, error)
	GetAuthorById(id uuid.UUID) (*dtos.AuthorResponseDto, error)
	UpdateAuthor(id uuid.UUID, updateAuthorDto *dtos.UpdateAuthorDto) (*dtos.AuthorResponseDto, error)
	DeleteAuthor(id uuid.UUID) error
}
