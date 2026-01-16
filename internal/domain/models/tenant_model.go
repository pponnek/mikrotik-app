package models

import (
	"time"
)

type Tenant struct {
	ID        string `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string `gorm:"column:name;type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Tenant) TableName() string {
	return "tenants"
}