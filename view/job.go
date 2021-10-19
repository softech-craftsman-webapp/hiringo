package view

import (
	"hiringo/model"
	"time"
)

type JobView struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"user_id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Image               string    `json:"image"`
	ValidUntil          time.Time `json:"valid_until"`
	IsEquipmentRequired bool      `json:"is_equipment_required"`
	CategoryID          string    `json:"category_id"`
	LocationID          string    `json:"location_id"`
	TransactionID       string    `json:"transaction_id"`
}

type JobEmptyView struct {
	ID string `json:"id"`
}

func JobModelToView(job model.Job) JobView {
	return JobView{
		ID:                  job.ID,
		UserID:              job.UserID,
		Name:                job.Name,
		Description:         job.Description,
		Image:               job.Image,
		ValidUntil:          job.ValidUntil,
		IsEquipmentRequired: job.IsEquipmentRequired,
		CategoryID:          job.CategoryID,
		LocationID:          job.LocationID,
		TransactionID:       job.TransactionID,
	}
}
