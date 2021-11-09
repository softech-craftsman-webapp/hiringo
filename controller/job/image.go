package job

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UpdateJobImageRequest struct {
	Image string `json:"image" validate:"required"`
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
// @Param user body UpdateJobImageRequest true "Job related informations"
// @Success 200 {object} view.Response{payload=view.JobView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/{id}/image [put]
// @Security JWT
func AddOrUpdateJobImage(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(UpdateJobImageRequest)

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

	job := &model.Job{
		ID: ctx.Param("id"),
	}

	db.First(&job, "id = ? AND user_id", job.ID, claims.User.ID)

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
		Image: req.Image,
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
			Message: result.Error.Error(),
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
