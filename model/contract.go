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
	BookingTime              time.Time      `gorm:"type:timestamp;not null" json:"booking_time"`
	SignedByRecruiterTime    *time.Time     `gorm:"type:timestamp;" json:"signed_by_recruiter_time"`
	SignedByProfessionalTime *time.Time     `gorm:"type:timestamp;" json:"signed_by_professional_time"`
	CreatedAt                time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Handled by Auth Service
	RecruiterId    string `gorm:"type:uuid;not null" json:"recruiter_id"`
	ProfessionalId string `gorm:"type:uuid;not null" json:"professional_id"`

	// From Transaction Table
	TransactionId string `gorm:"type:uuid;not null" json:"transaction_id"`
}
