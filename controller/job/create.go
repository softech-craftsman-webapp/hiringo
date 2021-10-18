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
	Name                string    `json:"name" validate:"required"`
	Description         string    `json:"descriptions" validate:"required"`
	Image               string    `json:"image"`
	IsEquipmentRequired bool      `json:"is_equipment_required"`
	ValidUntil          time.Time `json:"valid_until"`
	CategoryID          string    `json:"category_id"`
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
// @Success 200 {object} view.Response{payload=view.RatingView}
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

	job := &model.Job{
		UserID:              claims.User.ID,
		Name:                req.Name,
		Image:               req.Image,
		Description:         req.Description,
		IsEquipmentRequired: req.IsEquipmentRequired,
		ValidUntil:          req.ValidUntil,
		CategoryID:          req.CategoryID,
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
			Message: result.Error.Error(),
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusCreated, ctx, resp)
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
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
