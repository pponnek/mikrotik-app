package services

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"mikrotikapp/internal/domain/repositories"

	"github.com/google/uuid"
)

var (
	ErrInvalidCustomerRouter  = errors.New("invalid customer router")
	ErrCustomerRouterExists   = errors.New("customer router already exists")
	ErrCustomerRouterNotFound = errors.New("customer router not found")
)

type CustomerRouterService struct {
	repo     repositories.CustomerRouterRepository
	coreRepo repositories.CoreRouterRepository
}

func NewCustomerRouterService(
	repo repositories.CustomerRouterRepository,
	coreRepo repositories.CoreRouterRepository,
) *CustomerRouterService {
	return &CustomerRouterService{
		repo:     repo,
		coreRepo: coreRepo,
	}
}

func (s *CustomerRouterService) Create(
	tenantID uuid.UUID,
	coreRouterID uuid.UUID,
	name string,
	staticIP string,
) error {

	if tenantID == uuid.Nil || coreRouterID == uuid.Nil {
		return ErrInvalidCustomerRouter
	}
	if name == "" || staticIP == "" {
		return ErrInvalidCustomerRouter
	}

	// pastikan core router milik tenant
	_, err := s.coreRepo.FindByID(tenantID, coreRouterID)
	if err != nil {
		return err
	}

	router := &models.CustomerRouter{
		TenantID:     tenantID,
		CoreRouterID: coreRouterID,
		Name:         name,
		StaticIP:     staticIP,
	}

	err = s.repo.Create(router)
	if err != nil {
		if errors.Is(err, repositories.ErrDuplicateCustomerRouter) {
			return ErrCustomerRouterExists
		}
		return err
	}
	return nil
}

func (s *CustomerRouterService) List(tenantID uuid.UUID) ([]models.CustomerRouter, error) {
	return s.repo.FindByTenant(tenantID)
}

func (s *CustomerRouterService) Delete(tenantID, id uuid.UUID) error {
	err := s.repo.Delete(tenantID, id)
	if err != nil {
		if errors.Is(err, repositories.ErrCustomerRouterNotFound) {
			return ErrCustomerRouterNotFound
		}
		return err
	}
	return nil
}

func (s *CustomerRouterService) Block(
	tenantID uuid.UUID,
	id uuid.UUID,
) error {

	if tenantID == uuid.Nil || id == uuid.Nil {
		return ErrInvalidCustomerRouter
	}

	err := s.repo.SetBlocked(tenantID, id, true)
	if err != nil {
		if errors.Is(err, repositories.ErrCustomerRouterNotFound) {
			return ErrCustomerRouterNotFound
		}
		return err
	}
	return nil
}

func (s *CustomerRouterService) Unblock(
	tenantID uuid.UUID,
	id uuid.UUID,
) error {

	if tenantID == uuid.Nil || id == uuid.Nil {
		return ErrInvalidCustomerRouter
	}

	err := s.repo.SetBlocked(tenantID, id, false)
	if err != nil {
		if errors.Is(err, repositories.ErrCustomerRouterNotFound) {
			return ErrCustomerRouterNotFound
		}
		return err
	}
	return nil
}

