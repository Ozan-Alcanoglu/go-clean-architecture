package dtos

import "github.com/google/uuid"

type CreateBookDTO struct {
	Title         string    `json:"title" binding:"required"`
	AuthorID      uuid.UUID `json:"author_id" binding:"required"`
	PublishedYear int       `json:"published_year" binding:"required"`
	Stock         int       `json:"stock" binding:"required"`
}

type UpdateBookDTO struct {
	Title         *string    `json:"title"`
	AuthorID      *uuid.UUID `json:"author_id"`
	PublishedYear *int       `json:"published_year"`
	Stock         *int       `json:"stock"`
}

type BookResponseDTO struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	AuthorID      uuid.UUID `json:"author_id"`
	PublishedYear int       `json:"published_year"`
	Stock         int       `json:"stock"`
}
