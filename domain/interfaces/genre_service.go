package interfaces

import (
	"gobackend/domain/dtos"

	"github.com/google/uuid"
)

type GenreService interface {
	CreateGenre(createGenreDto *dtos.CreateGenreDto) (*dtos.GenreResponseDto, error)
	GetGenreById(id uuid.UUID) (*dtos.GenreResponseDto, error)
	UpdateGenre(id uuid.UUID, updateGenreDto *dtos.UpdateGenreDto) (*dtos.GenreResponseDto, error)
	DeleteGenre(id uuid.UUID) error
}
