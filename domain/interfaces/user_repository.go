package interfaces

import (
	"gobackend/domain/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetById(id uuid.UUID) (*models.User, error)
	DeleteUserById(id uuid.UUID) error
	UpdateUser(user *models.User) (*models.User, error)
}
