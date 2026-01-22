package services

import (
	"errors"
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"gobackend/domain/models"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	userRepo         interfaces.UserRepository
	tokenService     interfaces.TokenService
	refreshTokenRepo interfaces.RefreshTokenRepository
}

func NewAuthService(userRepo interfaces.UserRepository, tokenService interfaces.TokenService, refreshTokenRepo interfaces.RefreshTokenRepository) interfaces.AuthService {
	return &authService{
		userRepo:         userRepo,
		tokenService:     tokenService,
		refreshTokenRepo: refreshTokenRepo,
	}
}

func (s *authService) Register(registerDto *dtos.RegisterDto) (*dtos.RegisterResponse, string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	user := &models.User{
		Name:     registerDto.Name,
		Email:    registerDto.Email,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, "", err
	}

	accessToken, err := s.tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return nil, "", err
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	if err := s.refreshTokenRepo.Create(refreshTokenModel); err != nil {
		return nil, "", err
	}

	return &dtos.RegisterResponse{AccessToken: accessToken}, refreshToken, nil
}

func (s *authService) Login(loginDto *dtos.LoginDto) (*dtos.LoginResponse, string, error) {
	user, err := s.userRepo.GetByEmail(loginDto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("invalid credentials")
		}
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	accessToken, err := s.tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return nil, "", err
	}

	refreshTokenModel := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	}

	if err := s.refreshTokenRepo.Create(refreshTokenModel); err != nil {
		return nil, "", err
	}

	return &dtos.LoginResponse{AccessToken: accessToken}, refreshToken, nil
}

func (s *authService) RenewAccessToken(refreshToken string) (*dtos.RefreshTokenResponse, error) {
	storedToken, err := s.refreshTokenRepo.GetByToken(refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if storedToken.Revoked || storedToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("refresh token expired or revoked")
	}

	accessToken, err := s.tokenService.GenerateAccessToken(storedToken.UserID)
	if err != nil {
		return nil, err
	}

	return &dtos.RefreshTokenResponse{AccessToken: accessToken}, nil
}
