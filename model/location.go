package model

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	gorm.Model

	ID            string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID        string         `gorm:"type:uuid;not null" json:"user_id"`
	Address       string         `gorm:"index;type:varchar(255);not null" json:"address"`
	Longitude     float64        `gorm:"type:float;not null" json:"longitude"`
	Latitude      float64        `gorm:"type:float;not null" json:"latitude"`
	Street        string         `gorm:"type:varchar(255)" json:"street"`
	HouseNumber   string         `gorm:"type:varchar(255)" json:"house_number"`
	Suburb        string         `gorm:"type:varchar(255)" json:"suburb"`
	Postcode      string         `gorm:"type:varchar(255)" json:"postcode"`
	State         string         `gorm:"type:varchar(255)" json:"state"`
	StateCode     string         `gorm:"type:varchar(255)" json:"state_code"`
	StateDistrict string         `gorm:"type:varchar(255)" json:"state_district"`
	County        string         `gorm:"type:varchar(255)" json:"county"`
	Country       string         `gorm:"type:varchar(255)" json:"country"`
	CountryCode   string         `gorm:"type:varchar(255)" json:"country_code"`
	City          string         `gorm:"type:varchar(255)" json:"city"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
