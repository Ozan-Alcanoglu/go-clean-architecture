package interfaces

import "gobackend/domain/dtos"

type AuthService interface {
	Register(registerDto *dtos.RegisterDto) (*dtos.RegisterResponse, string, error)
	Login(loginDto *dtos.LoginDto) (*dtos.LoginResponse, string, error)
	RenewAccessToken(refreshToken string) (*dtos.RefreshTokenResponse, error)
}
