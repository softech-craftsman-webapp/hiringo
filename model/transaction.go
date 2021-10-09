package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model

	ID        string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID    string         `gorm:"type:uuid;not null" json:"user_id"`
	Amount    float64        `gorm:"type:decimal(10,2);not null" json:"amount"`
	Currency  string         `gorm:"default:EUR;type:varchar(3);not null" json:"currency"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
