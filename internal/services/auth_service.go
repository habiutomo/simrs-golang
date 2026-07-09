package services

import (
	"errors"

	"simrs-golang/internal/database"
	"simrs-golang/internal/middleware"
	"simrs-golang/internal/models"
	"simrs-golang/internal/repositories"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

type LoginResponse struct {
	Token    string      `json:"token"`
	User     models.User `json:"user"`
}

func (s *AuthService) Login(username, password string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if !database.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid username or password")
	}

	if !user.IsActive {
		return nil, errors.New("account is disabled")
	}

	token, err := middleware.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *AuthService) GetProfile(userID uint) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
