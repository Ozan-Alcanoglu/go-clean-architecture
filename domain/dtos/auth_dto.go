package dtos

type LoginDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterResponse struct {
	AccessToken string `json:"access_token"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
