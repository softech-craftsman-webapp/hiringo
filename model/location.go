package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model

	ID         string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name       string         `gorm:"index;type:varchar(64);not null" json:"name"`
	Langtitude float64        `gorm:"type:float;not null" json:"langtitude"`
	Latitude   float64        `gorm:"type:float;not null" json:"latitude"`
	CreatedAt  time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
