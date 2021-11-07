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

type UpdateJobRequest struct {
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
// @Param id path string true "Job id"
// @Param user body UpdateJobRequest true "Job related informations"
// @Success 200 {object} view.Response{payload=view.JobView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/{id} [put]
// @Security JWT
func UpdateJob(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(UpdateJobRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: err.Error(),
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
		ID: ctx.Param("id"),
	}

	db.First(&job, "id = ? AND user_id = ?", job.ID, claims.User.ID)

	/*
	   |--------------------------------------------------------------------------
	   | Check if user's id the same as the logged in user
	   |--------------------------------------------------------------------------
	*/
	if job.UserID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check rquired fields
	   |--------------------------------------------------------------------------
	*/
	result := db.Model(&job).Updates(model.Job{
		Name:                req.Name,
		Description:         req.Description,
		IsEquipmentRequired: req.IsEquipmentRequired,
		ValidUntil:          validUntil,
		CategoryID:          req.CategoryID,
		TransactionID:       req.TransactionID,
		Latitude:            req.Latitude,
		Longitude:           req.Longitude,
	})

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.JobEmptyView{
			ID: job.ID,
		},
	}

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

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
