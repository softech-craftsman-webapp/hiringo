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

/*
   |--------------------------------------------------------------------------
   | Sign Contract by Recruiter
   | @JWT via Acess Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Sign Contract by Recruiter
// @Tags contract
// @Description Sign Contract by Recruiter
// @Accept  json
// @Produce  json
// @Param id path string true "Contract id"
// @Success 200 {object} view.Response{payload=view.ContractEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/{id}/sign [post]
// @Security JWT
func SignContract(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	contract := &model.Contract{
		ID: ctx.Param("id"),
	}

	result := db.First(&contract, "id = ? AND recruiter_id = ?", contract.ID, claims.User.ID)

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

	job := &model.Job{
		ID: contract.JobID,
	}

	resultJob := db.First(&job, "id = ? AND user_id = ?", job.ID, contract.RecruiterID)

	if resultJob.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultJob.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	// is contract already signed
	if job.IsContractSigned {
		resp := &view.Response{
			Success: true,
			Message: "Contract already signed",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check required fields
	   |--------------------------------------------------------------------------
	*/
	resultJobUpdate := db.Model(&job).Updates(model.Job{
		IsContractSigned: true,
	})

	resultContractUpdate := db.Model(&contract).Updates(model.Contract{
		SignedByRecruiterTime: time.Now(),
	})

	if resultContractUpdate.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultContractUpdate.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	if resultJobUpdate.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultJobUpdate.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryEmptyView{
			ID: contract.ID,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
