package repositories

import (
	"mikrotikapp/internal/domain/models"

	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

var (
	ErrDuplicateRouter = errors.New("duplicate router")
	ErrRouterNotFound  = errors.New("router not found")
)

type CoreRouterRepository interface {
	Create(router *models.CoreRouter) error
	FindByTenant(tenantID uuid.UUID) ([]models.CoreRouter, error)
	FindByID(tenantID, id uuid.UUID) (*models.CoreRouter, error)
	Delete(tenantID, id uuid.UUID) error
}

type coreRouterRepository struct {
	db *gorm.DB
}

func NewCoreRouterRepository(db *gorm.DB) CoreRouterRepository {
	return &coreRouterRepository{db: db}
}

func (r *coreRouterRepository) Create(router *models.CoreRouter) error {
	err := r.db.Create(router).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateRouter
		}
		return err
	}
	return nil
}

func (r *coreRouterRepository) FindByTenant(tenantID uuid.UUID) ([]models.CoreRouter, error) {
	var routers []models.CoreRouter
	err := r.db.Where("tenant_id = ?", tenantID).Find(&routers).Error
	return routers, err
}

func (r *coreRouterRepository) FindByID(tenantID, id uuid.UUID) (*models.CoreRouter, error) {
	var router models.CoreRouter
	err := r.db.
		Where("id = ? AND tenant_id = ?", id, tenantID).
		First(&router).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRouterNotFound
		}
		return nil, err
	}

	return &router, nil
}

func (r *coreRouterRepository) Delete(tenantID, id uuid.UUID) error {
	res := r.db.
		Where("id = ? AND tenant_id = ?", id, tenantID).
		Delete(&models.CoreRouter{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrRouterNotFound
	}
	return nil
}