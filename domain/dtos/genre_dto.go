package dtos

import "github.com/google/uuid"

type CreateGenreDto struct {
	Name string `json:"name" binding:"required"`
}

type UpdateGenreDto struct {
	Name *string `json:"name" binding:"required"`
}

type GenreResponseDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
