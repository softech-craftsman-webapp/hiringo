package model

import (
	"time"

	"gorm.io/gorm"
)

type Job struct {
	gorm.Model

	ID                  string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name                string         `gorm:"index;type:text;not null" json:"name"`
	Description         string         `gorm:"type:text;not null" json:"description"`
	Image               string         `gorm:"type:text;null" json:"image"`
	UserID              string         `gorm:"type:uuid;not null" json:"recruiter_id"`
	IsPremium           bool           `gorm:"default:false" json:"is_premium"`
	IsEquipmentRequired bool           `gorm:"default:false" json:"is_equipment_required"`
	ValidUntil          time.Time      `gorm:"type:timestamp;" json:"valid_until"`
	CreatedAt           time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// From Category Table
	CategoryID string `gorm:"type:uuid;not null" json:"category_id"`

	// From Transaction Table
	TransactionID string `gorm:"type:uuid;not null" json:"transaction_id"`

	// Contracts
	Contracts        []Contract `gorm:"foreignKey:JobID" json:"contracts"`
	IsContractSigned bool       `gorm:"default:false" json:"is_contract_signed"`

	// Location
	Latitude  float64 `gorm:"type:float;not null" json:"latitude"`
	Longitude float64 `gorm:"type:float;not null" json:"longitude"`
}
