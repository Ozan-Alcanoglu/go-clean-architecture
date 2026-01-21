package interfaces

import (
	"gobackend/domain/dtos"

	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(createUserDto *dtos.CreateUserDto) (*dtos.UserResponseDto, error)
	GetUserById(id uuid.UUID) (*dtos.UserResponseDto, error)
	UpdateUser(id uuid.UUID, updateUserDto *dtos.UpdateUserDto) (*dtos.UserResponseDto, error)
	DeleteUser(id uuid.UUID) error
}
