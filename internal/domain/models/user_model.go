package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey" json:"id"`
	TenantID  *uuid.UUID `gorm:"type:uuid" json:"tenant_id"`
	Username  string     `gorm:"uniqueIndex;not null" json:"username"`
	Password  string     `gorm:"column:password_hash;not null" json:"password"` // pakai password_hash
	Role      string     `gorm:"not null;default:TENANT_ADMIN" json:"role"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil{
		u.ID = uuid.New()
	}

	return
}
