package models

import (
	"time"

	"github.com/google/uuid"
)

type CoreRouter struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TenantID          uuid.UUID `gorm:"type:uuid;not null;index:idx_tenant_host,unique"`
	Name              string    `gorm:"size:100;not null"`
	Host              string    `gorm:"type:inet;not null;index:idx_tenant_host,unique"`
	ApiPort           int       `gorm:"not null;default:8728"`
	Username          string    `gorm:"size:50;not null"`
	PasswordEncrypted string    `gorm:"type:text;not null"`
	CreatedAt         time.Time
}

func (CoreRouter) TableName() string {
	return "core_router"
}