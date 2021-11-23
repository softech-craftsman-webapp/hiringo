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
   | Get Job Detail for authenticated user
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Job Detail for authenticated user
// @Tags job
// @Description Get Job Detail for authenticated user
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.MyJobView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/my [get]
// @Security JWT
func GetMyJobs(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	createdJob := []model.Job{}
	appliedJob := []model.Job{}
	jobIDs := []string{}

	// get job ids which user applied
	db.Model(&model.Contract{}).Where("professional_id = ?", claims.User.ID).Pluck("job_id", &jobIDs).Order("created_at DESC")
	db.Model(&model.Job{}).Where("id IN (?)", jobIDs).Find(&appliedJob).Order("created_at DESC")

	// get created jobs
	db.Model(&model.Job{}).Where("user_id = ?", claims.User.ID).Find(&createdJob).Order("created_at DESC")

	// TODO: It can be optimized
	var formatted_createdJob []view.JobView
	var formatted_appliedJob []view.JobView

	for _, job := range createdJob {
		formatted_createdJob = append(formatted_createdJob, view.JobModelToView(job, 0, 0))
	}

	for _, job := range appliedJob {
		formatted_appliedJob = append(formatted_appliedJob, view.JobModelToView(job, 0, 0))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.MyJobView{
			Created: formatted_createdJob,
			Applied: formatted_appliedJob,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
