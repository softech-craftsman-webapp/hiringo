package job

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateJobRequest struct {
	Name                string  `json:"name" validate:"required"`
	Description         string  `json:"descriptions" validate:"required"`
	IsEquipmentRequired bool    `json:"is_equipment_required"`
	ValidUntil          string  `json:"valid_until" validate:"required"`
	CategoryID          string  `json:"category_id" validate:"required"`
	TransactionID       string  `json:"transaction_id" validate:"required"`
	Latitude            float64 `json:"latitude" validate:"required,numeric"`
	Longitude           float64 `json:"longitude" validate:"required,numeric"`
}

/*
   |--------------------------------------------------------------------------
   | Create Job
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Job
// @Tags job
// @Description Create Job
// @Accept  json
// @Produce  json
// @Param user body CreateJobRequest true "Job related informations"
// @Success 200 {object} view.Response{payload=view.JobView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/new [post]
// @Security JWT
func CreateJob(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateJobRequest)

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

	// Time Validation
	validUntil, err := time.Parse("2006-01-02 15:04", req.ValidUntil)
	if err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "Time validation failed, example 2006-01-02 15:04",
			Payload: nil,
		})
	}

	job := &model.Job{
		UserID:              claims.User.ID,
		Name:                req.Name,
		Description:         req.Description,
		IsEquipmentRequired: req.IsEquipmentRequired,
		ValidUntil:          validUntil,
		CategoryID:          req.CategoryID,
		TransactionID:       req.TransactionID,
		Latitude:            req.Latitude,
		Longitude:           req.Longitude,
	}

	result := db.Create(&job)

	/*
	   |--------------------------------------------------------------------------
	   | DB relation error
	   |--------------------------------------------------------------------------
	*/
	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: "Duplicate error",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.JobView{
			ID:                  job.ID,
			Name:                job.Name,
			Description:         job.Description,
			Image:               job.Image,
			IsEquipmentRequired: job.IsEquipmentRequired,
			ValidUntil:          job.ValidUntil,
			UserID:              job.UserID,
			CategoryID:          job.CategoryID,
			TransactionID:       job.TransactionID,
			Latitude:            job.Latitude,
			Longitude:           job.Longitude,
			IsContractSigned:    job.IsContractSigned,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
