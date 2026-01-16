package services

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"mikrotikapp/internal/domain/repositories"

	"github.com/google/uuid"
)

type TenantService struct {
	repo *repositories.TenantRepository
}

func NewTenantService(repo *repositories.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) Create(name string, role string) (*models.Tenant, error) {
	if role != "SUPERADMIN" {
		return nil, errors.New("halo")
	}

	if s.repo.ExistsByName(name) {
		return nil, errors.New("tenant already exists")
	}

	tenant := &models.Tenant{
		ID:   uuid.New().String(),
		Name: name,
	}

	if err := s.repo.Create(tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) List(role string) ([]models.Tenant, error) {
	if role != "SUPERADMIN" {
		return nil, errors.New("forbidden")
	}

	return s.repo.List()
}

func (s *TenantService) Delete(id uuid.UUID, role string) error {
	if role != "SUPERADMIN" {
		return errors.New("forbidden")
	}

	return s.repo.Delete(id.String())
}

