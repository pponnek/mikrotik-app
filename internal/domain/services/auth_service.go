package services

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/pkg/utils"

	"github.com/google/uuid"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthServcie(userRepo *repositories.UserRepository) *AuthService {
	return  &AuthService{userRepo: userRepo}
}

func (s *AuthService) Login(username, password string)(string, string, string, error) {
	user, err := s.userRepo.FindByUsername(username)

	if err != nil {
		return "","","", errors.New("Username tidak ditemukan")
	}

	if !utils.CheckPasswordHash(user.Password, password) {
		return "", "", "", errors.New("Password salah")
	}

	token, _ := utils.GenerateToken(
		user.ID.String(),
		utils.UUIDPtrToString(user.TenantID),
		user.Role,
	)

	return user.ID.String(), utils.UUIDPtrToString(user.TenantID), token, nil
}

func (s *AuthService) Register(username, password string, tenantID *uuid.UUID, role string) (*models.User, error) {
	if role == "" {
		role = "TENANT_ADMIN"
	}

	if role == "SUPER_ADMIN" {
		tenantID = nil
	}

	if _, err := s.userRepo.FindByUsername(username); err == nil {
		return nil, errors.New("username already exists")
	}

	hashed, _ := utils.HashPassword(password)
	user := &models.User{
		ID:       uuid.New(),
		Username: username,
		Password: hashed,
		Role:     role,
		TenantID: tenantID,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Profile(userID uuid.UUID) (*models.User, error) {
	return s.userRepo.FindByID(userID)
}
