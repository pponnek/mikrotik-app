package models

import (
	"time"

	"github.com/google/uuid"
)

type CustomerRouter struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	TenantID     uuid.UUID `gorm:"type:uuid;not null;index:uniq_customer_router,unique"`
	CoreRouterID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name         string    `gorm:"size:100;not null"`
	StaticIP     string    `gorm:"type:inet;not null;index:uniq_customer_router,unique"`
	IsBlocked    bool      `gorm:"not null;default:false"`
	CreatedAt    time.Time
}

func (CustomerRouter) TableName() string {
	return "customer_router"
}