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
   | Get User Contracts
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Job Contracts
// @Tags contract
// @Description Get Job Contracts
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=[]view.ContractView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/my [get]
// @Security JWT
func GetJobContracts(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	var contracts []model.Contract
	result := db.Where("professional_id = ?", claims.User.ID).Find(&contracts).Order("created_at DESC")

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
