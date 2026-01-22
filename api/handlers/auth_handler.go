package handlers

import (
	"gobackend/domain/dtos"
	"gobackend/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service interfaces.AuthService
}

func NewAuthHandler(service interfaces.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var registerDto dtos.RegisterDto
	if err := c.ShouldBindJSON(&registerDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, refreshToken, err := h.service.Register(&registerDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var loginDto dtos.LoginDto
	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, refreshToken, err := h.service.Login(&loginDto)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "localhost", false, true)
	c.JSON(http.StatusOK, response)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token required"})
		return
	}

	response, err := h.service.RenewAccessToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
