package model

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model

	ID                       string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	StartTime                time.Time      `gorm:"type:timestamp;not null" json:"start_time"`
	EndTime                  time.Time      `gorm:"type:timestamp;not null" json:"end_time"`
	SignedByRecruiterTime    time.Time      `gorm:"type:timestamp;" json:"signed_by_recruiter_time"`
	SignedByProfessionalTime time.Time      `gorm:"type:timestamp;" json:"signed_by_professional_time"`
	CreatedAt                time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Handled by Auth Service
	RecruiterID    string `gorm:"type:uuid;not null" json:"recruiter_id"`
	ProfessionalID string `gorm:"type:uuid;not null" json:"professional_id"`

	// From Job Table
	JobID string `gorm:"type:uuid;not null" json:"job_id"`

	// From Rating Table
	Ratings []Rating `gorm:"foreignKey:ContractID" json:"rating"`
}
