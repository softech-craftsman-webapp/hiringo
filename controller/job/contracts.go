package job

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

/*
   |--------------------------------------------------------------------------
   | Get Job Contracts
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Job Contracts
// @Tags contract
// @Description Get Job Contracts
// @Accept  json
// @Produce  json
// @Param id path string true "Job id"
// @Success 200 {object} view.Response{payload=[]view.ContractView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /jobs/{id}/contracts [get]
// @Security JWT
func GetJobContracts(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	job := model.Job{
		ID: ctx.Param("id"),
	}

	var contracts []model.Contract
	result := db.Where("id = ?", job.ID).First(&job)

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

	// check if user is the owner of the job contracts
	var resultContracts *gorm.DB
	if job.UserID == claims.User.ID {
		resultContracts = db.Where("recruiter_id = ? AND job_id = ?", claims.User.ID, job.ID).Find(&contracts)
	} else {
		resultContracts = db.Where("professional_id = ? AND job_id = ?", claims.User.ID, job.ID).Find(&contracts)
	}

	if resultContracts.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultContracts.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	// TODO: It can be optimized
	var formatted_contracts []view.ContractView
	for _, contract := range contracts {
		formatted_contracts = append(formatted_contracts, view.ContractModelToView(contract))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_contracts,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
