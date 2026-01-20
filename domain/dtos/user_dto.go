package dtos

import "github.com/google/uuid"

type CreateUserDto struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type UpdateUserDto struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
}

type UserResponseDto struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}
