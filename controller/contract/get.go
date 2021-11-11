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
   | Get Contract Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Contract Detail
// @Tags contract
// @Description Get Contract Detail
// @Accept  json
// @Produce  json
// @Param id path string true "Contract id"
// @Success 200 {object} view.Response{payload=view.ContractView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/{id} [get]
// @Security JWT
func GetContractDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	contract := model.Contract{
		ID: ctx.Param("id"),
	}
	result := db.First(&contract)

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

	// check if user has access to this contract
	if contract.RecruiterID != claims.User.ID && contract.ProfessionalID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "You are not allowed to access this contract",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.ContractView{
			StartTime:                contract.StartTime,
			EndTime:                  contract.EndTime,
			SignedByRecruiterTime:    contract.SignedByRecruiterTime,
			SignedByProfessionalTime: contract.SignedByProfessionalTime,
			RecruiterID:              contract.RecruiterID,
			ProfessionalID:           contract.ProfessionalID,
			JobID:                    contract.JobID,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
