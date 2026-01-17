package repositories

import (
	"errors"
	"mikrotikapp/internal/domain/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrDuplicateCustomerRouter = errors.New("customer router already exists")
	ErrCustomerRouterNotFound  = errors.New("customer router not found")
)

type CustomerRouterRepository interface {
	Create(router *models.CustomerRouter) error
	FindByTenant(tenantID uuid.UUID) ([]models.CustomerRouter, error)
	Delete(tenantID, id uuid.UUID) error
	SetBlocked(tenantID uuid.UUID, id uuid.UUID, blocked bool) error

}

type customerRouterRepository struct {
	db *gorm.DB
}

func NewCustomerRouterRepository(db *gorm.DB) CustomerRouterRepository {
	return &customerRouterRepository{db: db}
}

func (r *customerRouterRepository) Create(router *models.CustomerRouter) error {
	err := r.db.Create(router).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ErrDuplicateCustomerRouter
		}
		return err
	}
	return nil
}

func (r *customerRouterRepository) FindByTenant(tenantID uuid.UUID) ([]models.CustomerRouter, error) {
	var routers []models.CustomerRouter
	err := r.db.
		Where("tenant_id = ?", tenantID).
		Find(&routers).Error
	return routers, err
}

func (r *customerRouterRepository) Delete(tenantID, id uuid.UUID) error {
	res := r.db.
		Where("id = ? AND tenant_id = ?", id, tenantID).
		Delete(&models.CustomerRouter{})

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrCustomerRouterNotFound
	}
	return nil
}

func (r *customerRouterRepository) SetBlocked(
	tenantID uuid.UUID,
	id uuid.UUID,
	blocked bool,
) error {

	res := r.db.Model(&models.CustomerRouter{}).
		Where("id = ? AND tenant_id = ?", id, tenantID).
		Update("is_blocked", blocked)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return ErrCustomerRouterNotFound
	}
	return nil
}
