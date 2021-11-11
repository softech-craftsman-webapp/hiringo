package job

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Delete job
   | @JWT via Access Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Delete Job
// @Tags job
// @Description Delete Job
// @Accept  json
// @Produce  json
// @Param id path string true "Job id"
// @Success 200 {object} view.Response{payload=view.JobEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/{id} [delete]
// @Security JWT
func DeleteJob(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	job := &model.Job{
		ID: ctx.Param("id"),
	}

	db.First(&job, "id = ? AND user_id = ?", job.ID, claims.User.ID)

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

	result := db.Delete(&job)

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
			Success: false,
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
