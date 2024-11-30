package service

import (
	"errors"
	"loyalty_central/internal/dto"
	"loyalty_central/internal/repository"
	"loyalty_central/internal/utils"
	"os"
	"time"
)

// AuthService provides authentication services
type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	adminRepo repository.AdminRepository
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(adminRepo repository.AdminRepository) AuthService {
	return &authService{adminRepo}
}

// Login authenticates a user and returns a JWT token if successful
func (s *authService) Login(username, password string) (string, error) {
	admin, err := s.adminRepo.GetAdminByUsername(username)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	if !utils.CheckPasswordHash(password, admin.Password) {
		return "", errors.New("invalid username or password")
	}

	token, err := utils.GenerateJWTToken(dto.AdminDTO{
		ID:       admin.ID,
		Username: admin.Username,
		Name:     admin.Name,
		Role:     admin.Role,
		Phone:    admin.PhoneNumber,
	}, os.Getenv("JWT_SECRET"), time.Hour*24)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
