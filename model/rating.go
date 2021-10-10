package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model

	ID            string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Points        int            `gorm:"type:int;not null" json:"points"`
	SubmittedByID string         `gorm:"type:uuid;not null" json:"submitted_by_id"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Handled by Auth Service
	UserID string `gorm:"type:uuid" json:"user_id"`

	// From Contract Table
	ContractID string `gorm:"type:uuid" json:"contract_id"`
}
