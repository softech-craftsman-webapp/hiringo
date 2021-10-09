package model

import (
	"time"

	"gorm.io/gorm"
)

type Rating struct {
	gorm.Model

	ID                     string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	PointsFromRecruiter    int            `gorm:"type:int" json:"points_from_recruiter"`
	PointsFromProfessional int            `gorm:"type:int" json:"points_from_professional"`
	CreatedAt              time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt              time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Handled by Auth Service
	ProfessionalID string `gorm:"type:uuid" json:"professional_id"`
	RecruiterID    string `gorm:"type:uuid" json:"recruiter_id"`

	// From Contract Table
	ContractId string `gorm:"type:uuid" json:"contract_id"`
}
