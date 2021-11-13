package job

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get Job Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Job Details
// @Tags job
// @Description Get Job Details
// @Accept  json
// @Produce  json
// @Param id path string true "Job id"
// @Success 200 {object} view.Response{payload=view.JobView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/{id} [get]
// @Security JWT
func GetJobDetail(ctx echo.Context) error {
	db := config.GetDB()

	job := model.Job{
		ID: ctx.Param("id"),
	}
	result := db.Where("id ? = ", job.ID).First(&job)

	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: result.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.JobView{
			ID:                  job.ID,
			Name:                job.Name,
			Description:         job.Description,
			UserID:              job.UserID,
			Image:               job.Image,
			ValidUntil:          job.ValidUntil,
			IsEquipmentRequired: job.IsEquipmentRequired,
			CategoryID:          job.CategoryID,
			TransactionID:       job.TransactionID,
			Longitude:           job.Longitude,
			Latitude:            job.Latitude,
			IsContractSigned:    job.IsContractSigned,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
