package model

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model

	ID          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedByID string         `gorm:"type:uuid" json:"created_by"`
	Name        string         `gorm:"index;type:varchar(64);not null;unique" json:"name"`
	Description string         `gorm:"null" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// From Job Table
	Jobs []Job `gorm:"foreignKey:CategoryID" json:"jobs"`
}
