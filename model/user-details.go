package model

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	gorm.Model

	ID        string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string         `gorm:"type:varchar(255);unique_index" json:"email"`
	Telephone string         `gorm:"type:varchar(255);unique_index" json:"telephone"`
	Bio       string         `gorm:"type:text;null" json:"bio"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Handled by Auth Service
	UserID string `gorm:"type:uuid;not null" json:"user_id"`

	// Location from table
	LocationID string `gorm:"type:uuid;not null" json:"location_id"`
}
