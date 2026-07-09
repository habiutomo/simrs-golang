package handlers

import (
	"net/http"

	"simrs-golang/internal/dto"
	"simrs-golang/internal/services"
	"simrs-golang/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request", err.Error())
		return
	}

	result, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	response.OK(c, "Login successful", result)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, _ := c.Get("user_id")

	user, err := h.authService.GetProfile(userID.(uint))
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}

	response.OK(c, "Profile retrieved", user)
}
