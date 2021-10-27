package contract

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateContractRequest struct {
	StartTime string `json:"start_time" validate:"required"`
	EndTime   string `json:"end_time" validate:"required"`
	JobID     string `json:"job_id" validate:"required"`
}

/*
   |--------------------------------------------------------------------------
   | Create Contract
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Contract
// @Tags contract
// @Description Create Contract
// @Accept  json
// @Produce  json
// @Param user body CreateContractRequest true "Contract for Job"
// @Success 200 {object} view.Response{payload=view.ContractView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/new [post]
// @Security JWT
func CreateContract(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateContractRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: "Bad request",
			Payload: nil,
		})
	}

	// Time Validation Start Time
	startTime, err := time.Parse("2006-01-02 15:04", req.StartTime)
	if err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "StartTime validation failed, example 2006-01-02 15:04",
			Payload: nil,
		})
	}

	endTime, err := time.Parse("2006-01-02 15:04", req.EndTime)
	if err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "EndTime validation failed, example 2006-01-02 15:04",
			Payload: nil,
		})
	}

	// Find a job
	job := model.Job{
		ID: req.JobID,
	}
	resultJob := db.First(&job)

	if resultJob.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Job not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	contract := &model.Contract{
		StartTime:                startTime,
		EndTime:                  endTime,
		JobID:                    req.JobID,
		SignedByProfessionalTime: time.Now(),
		ProfessionalID:           claims.User.ID,
		RecruiterID:              job.UserID,
	}

	result := db.Create(&contract)
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
		Payload: &view.ContractView{
			ID:                       contract.ID,
			StartTime:                contract.StartTime,
			EndTime:                  contract.EndTime,
			JobID:                    contract.JobID,
			SignedByProfessionalTime: contract.SignedByProfessionalTime,
			SignedByRecruiterTime:    contract.SignedByRecruiterTime,
			ProfessionalID:           contract.ProfessionalID,
			RecruiterID:              contract.RecruiterID,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
