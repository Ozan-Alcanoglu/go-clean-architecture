package interfaces

import "github.com/google/uuid"

type TokenService interface {
	GenerateAccessToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken() (string, error)
	ValidateAccessToken(tokenString string) (*uuid.UUID, error)
}
