package security

import (
	"errors"
	"gobackend/domain/interfaces"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type jwtService struct {
	secretKey []byte
	issuer    string
}

func NewJWTService() interfaces.TokenService {
	return &jwtService{
		secretKey: []byte(os.Getenv("JWT_SECRET")),
		issuer:    "gobackend",
	}
}

func (s *jwtService) GenerateAccessToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID.String(),
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
		"iss":     s.issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *jwtService) GenerateRefreshToken() (string, error) {
	return uuid.New().String(), nil
}

func (s *jwtService) ValidateAccessToken(tokenString string) (*uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDStr, ok := claims["user_id"].(string)
		if !ok {
			return nil, errors.New("invalid token claims")
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return nil, err
		}

		return &userID, nil
	}

	return nil, errors.New("invalid token")
}
