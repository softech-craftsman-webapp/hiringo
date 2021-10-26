package job

import (
	config "hiringo/config"
	helper "hiringo/helper"
	model "hiringo/model"
	view "hiringo/view"
	"sort"

	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// TODO: Somehow lat and long is required fields to calculate distance
// IDK, how to set default value for them
type SeachJobRequest struct {
	Name                string  `json:"name" validate:"required"`
	Description         string  `json:"description"`
	CategoryID          string  `json:"category_id"`
	IsEquipmentRequired bool    `json:"is_equipment_required"`
	Latitude            float64 `json:"latitude" validate:"required"`
	Longitude           float64 `json:"longitude" validate:"required"`
}

/*
   |--------------------------------------------------------------------------
   | Search jobs
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Search jobs
// @Tags job
// @Description Search jobs
// @Accept  json
// @Produce  json
// @Param user body SeachJobRequest true "Job related informations"
// @Success 200 {object} view.Response{payload=[]view.JobView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/search [post]
// @Security JWT
func SearchJobs(ctx echo.Context) error {
	db := config.GetDB()
	req := new(SeachJobRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "Bad request",
			Payload: nil,
		})
	}

	// Advanced search jobs
	// Get db result if request is not empty
	var jobs []model.Job
	var result *gorm.DB

	switch {
	case req.Name != "" && req.Description != "" && req.CategoryID != "" && req.IsEquipmentRequired:
		//
		// If all fields are not empty
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("description LIKE ?", helper.Format("%{{.}}%", req.Description)).
			Where("category_id = ?", req.CategoryID).
			Where("is_equipment_required = ?", req.IsEquipmentRequired).
			Find(&jobs)
	case req.Name != "" && req.CategoryID != "" && req.IsEquipmentRequired:
		//
		// If name, is_equipment_required and category_id are not
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("category_id = ?", req.CategoryID).
			Where("is_equipment_required = ?", req.IsEquipmentRequired).
			Find(&jobs)
	case req.Name != "" && req.Description != "" && req.CategoryID != "":
		//
		// If name, description and category_id are not empty
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("description LIKE ?", helper.Format("%{{.}}%", req.Description)).
			Where("category_id = ?", req.CategoryID).
			Find(&jobs)
	case req.Name != "" && req.CategoryID != "":
		//
		// If name and category_id are not empty
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("category_id = ?", req.CategoryID).
			Find(&jobs)
	case req.Name != "" && req.IsEquipmentRequired:
		//
		// If name, is_equipment_required and category_id are not
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("is_equipment_required = ?", req.IsEquipmentRequired).
			Find(&jobs)
	case req.Name != "" && req.Description != "":
		//
		// If name and description are not empty
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).
			Where("description LIKE ?", helper.Format("%{{.}}%", req.Description)).
			Find(&jobs)
	default:
		//
		// If name is not empty
		//
		result = db.Where("name LIKE ?", helper.Format("%{{.}}%", req.Name)).Find(&jobs)
	}

	// TODO: It can be optimized
	var formatted_jobs []view.JobView
	for _, job := range jobs {
		formatted_jobs = append(formatted_jobs, view.JobModelToView(job, req.Latitude, req.Longitude))
	}

	// sort by distance
	sort.Slice(formatted_jobs, func(i, j int) bool {
		return formatted_jobs[i].Distance < formatted_jobs[j].Distance
	})

	/*
	   |--------------------------------------------------------------------------
	   | Main Error
	   |--------------------------------------------------------------------------
	*/
	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Internal Server Error",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_jobs,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
