package interfaces

import (
	"gobackend/domain/models"
)

type RefreshTokenRepository interface {
	Create(token *models.RefreshToken) error
	GetByToken(token string) (*models.RefreshToken, error)
	Revoke(token string) error
	DeleteExpired() error
}
