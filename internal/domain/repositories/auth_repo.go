package repositories

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"mikrotikapp/pkg/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db       *gorm.DB
	userRepo *UserRepository
}

func NewAuthRepository(db *gorm.DB, userRepo *UserRepository) *AuthRepository {
	return &AuthRepository{db: db, userRepo: userRepo}
}

func (s *AuthRepository) Register(username, password string, tenantID *uuid.UUID, role string) (*models.User, error) {
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
