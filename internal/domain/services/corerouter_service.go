package services

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"mikrotikapp/internal/domain/repositories"
	"mikrotikapp/pkg/utils"

	"github.com/google/uuid"
)

var (
	ErrInvalidTenant       = errors.New("invalid tenant")
	ErrMissingFields       = errors.New("missing required fields")
	ErrRouterAlreadyExists = errors.New("router already exists")
	ErrRouterNotFound      = errors.New("router not found")
)

type CoreRouterService struct {
	repo repositories.CoreRouterRepository
}

func NewCoreRouterService(repo repositories.CoreRouterRepository) *CoreRouterService {
	return &CoreRouterService{repo: repo}
}

func (s *CoreRouterService) Create(
	tenantID uuid.UUID,
	name, host, username, password string,
	apiPort int,
) error {

	if tenantID == uuid.Nil {
		return ErrInvalidTenant
	}
	if name == "" || host == "" || username == "" || password == "" {
		return ErrMissingFields
	}
	if apiPort == 0 {
		apiPort = 8728
	}

	encPwd, err := utils.Encrypt(password)
	if err != nil {
		return err
	}

	router := &models.CoreRouter{
		TenantID:          tenantID,
		Name:              name,
		Host:              host,
		ApiPort:           apiPort,
		Username:          username,
		PasswordEncrypted: encPwd,
	}

	err = s.repo.Create(router)
	if err != nil {
		if errors.Is(err, repositories.ErrDuplicateRouter) {
			return ErrRouterAlreadyExists
		}
		return err
	}

	return nil
}

func (s *CoreRouterService) List(tenantID uuid.UUID) ([]models.CoreRouter, error) {
	if tenantID == uuid.Nil {
		return nil, ErrInvalidTenant
	}
	return s.repo.FindByTenant(tenantID)
}

func (s *CoreRouterService) Delete(tenantID, id uuid.UUID) error {
	if tenantID == uuid.Nil {
		return ErrInvalidTenant
	}
	if id == uuid.Nil {
		return ErrRouterNotFound
	}

	err := s.repo.Delete(tenantID, id)
	if err != nil {
		if errors.Is(err, repositories.ErrRouterNotFound) {
			return ErrRouterNotFound
		}
		return err
	}
	return nil
}