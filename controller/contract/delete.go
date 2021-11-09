package contract

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
   | Delete contract
   | @JWT via Access Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Delete contract
// @Tags contract
// @Description Delete contract
// @Accept  json
// @Produce  json
// @Param id path string true "Contract id"
// @Success 200 {object} view.Response{payload=view.ContractEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/{id} [delete]
// @Security JWT
func DeleteContract(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	contract := &model.Contract{
		ID: ctx.Param("id"),
	}

	result := db.First(&contract, "id = ? AND  professional_id = ?", contract.ID, claims.User.ID)

	// db error
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

	// check permission
	if contract.ProfessionalID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	// delete
	job := model.Job{
		ID: contract.JobID,
	}
	resultJob := db.First(&job, "id = ?", job.ID)

	if resultJob.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: resultJob.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	if job.IsContractSigned {
		resp := &view.Response{
			Success: true,
			Message: "You are not permitted to delete contract. Contract is already signed.",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	// delete
	resultDelete := db.Delete(&contract)
	if resultDelete.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: resultDelete.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.ContractEmptyView{
			ID: contract.ID,
		},
	}

	/*
	   |--------------------------------------------------------------------------
	   | Main Error
	   |--------------------------------------------------------------------------
	*/

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
