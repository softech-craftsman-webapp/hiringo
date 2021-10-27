package view

import (
	helper "hiringo/helper"
	model "hiringo/model"

	"time"

	paginator "github.com/dmitryburov/gorm-paginator"
)

type JobView struct {
	ID                  string    `json:"id"`
	UserID              string    `json:"user_id"`
	Name                string    `json:"name"`
	Description         string    `json:"description"`
	Image               string    `json:"image"`
	ValidUntil          time.Time `json:"valid_until"`
	IsPremium           bool      `json:"is_premium"`
	IsEquipmentRequired bool      `json:"is_equipment_required"`
	CategoryID          string    `json:"category_id"`
	TransactionID       string    `json:"transaction_id"`
	Latitude            float64   `json:"latitude"`
	Longitude           float64   `json:"longitude"`
	Distance            float64   `json:"distance"`
	IsContractSigned    bool      `json:"is_contract_signed"`
}

type JobEmptyView struct {
	ID string `json:"id"`
}

type JobPagination struct {
	Items      []model.Job           `json:"items"`
	Pagination *paginator.Pagination `json:"pagination"`
}

type JobPaginationView struct {
	Items      []JobView             `json:"items"`
	Pagination *paginator.Pagination `json:"pagination"`
}

type MyJobView struct {
	Applied []JobView `json:"applied"`
	Created []JobView `json:"created"`
}

func JobModelToView(job model.Job, lat float64, long float64) JobView {
	distance := helper.Distance(job.Latitude, job.Longitude, lat, long)

	return JobView{
		ID:                  job.ID,
		UserID:              job.UserID,
		Name:                job.Name,
		Description:         job.Description,
		Image:               job.Image,
		ValidUntil:          job.ValidUntil,
		IsPremium:           job.IsPremium,
		IsEquipmentRequired: job.IsEquipmentRequired,
		CategoryID:          job.CategoryID,
		TransactionID:       job.TransactionID,
		Latitude:            job.Latitude,
		Longitude:           job.Longitude,
		Distance:            distance,
		IsContractSigned:    job.IsContractSigned,
	}
}
