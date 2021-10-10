package view

import "time"

type JobView struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"user_id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Image               string    `json:"image"`
	ValidUntil          time.Time `json:"valid_until"`
	IsEquipmentRequired bool      `json:"is_equipment_required"`
}

type JobEmptyView struct {
	ID string `json:"id"`
}
