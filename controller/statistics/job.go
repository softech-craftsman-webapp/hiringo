package statistics

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
   | Job statistics
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Job statistics
// @Tags statistics
// @Description Job statistics
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.JobStatistics}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /statistics/job [get]
// @Security JWT
func JobStatistics(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	contract := model.Contract{}

	// get the most popular job which has the most contracts
	db.Table("contracts").Select("job_id, count(*) as total").Group("job_id").Order("total desc").Limit(1).Scan(&contract)

	// get job
	job := model.Job{
		ID: contract.JobID,
	}

	db.Where("id = ?", job.ID).First(&job)

	// get all jobs
	jobs := []model.Job{}
	db.Find(&jobs).Order("created_at")

	// user jobs
	userJobs := []model.Job{}

	db.Where("user_id = ?", claims.User.ID).Find(&userJobs).Order("created_at")

	// result
	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.JobStatistics{
			PopularItem: &view.JobView{
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
				Distance:            0,
				IsContractSigned:    job.IsContractSigned,
			},
			Total:        len(jobs),
			UserJobCount: len(userJobs),
			Time:         job.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
