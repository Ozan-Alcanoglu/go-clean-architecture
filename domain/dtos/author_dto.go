package dtos

import "github.com/google/uuid"

type CreateAuthorDto struct {
	Name string `json:"name" binding:"required"`
	Bio  string `json:"bio" binding:"required"`
}

type UpdateAuthorDto struct {
	Name *string `json:"name"`
	Bio  *string `json:"bio"`
}

type AuthorResponseDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Bio  string    `json:"bio"`
}
