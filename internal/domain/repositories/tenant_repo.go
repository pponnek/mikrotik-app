package repositories

import (
	"errors"
	"mikrotikapp/internal/domain/models"

	"gorm.io/gorm"
)

type TenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(tenant *models.Tenant) error {
	return r.db.Create(tenant).Error
}

func (r *TenantRepository) List() ([]models.Tenant, error) {
	var tenant []models.Tenant

	if err := r.db.Order("created_at desc").Find(&tenant).Error; err != nil {
		return nil, err
	}
	return tenant, nil
}

func (r *TenantRepository) Delete(id string) error {
	result := r.db.Delete(&models.Tenant{}, "id = ?", id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("tenant not found")
	}

	return nil
}


func (r *TenantRepository) ExistsByName(name string) bool {
	var count int64
	r.db.Model(&models.Tenant{}).
		Where("name = ?", name).Count(&count)

	return count > 0
}
