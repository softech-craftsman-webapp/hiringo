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
   | Get Contract ratings
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get Contract ratings
// @Tags contract
// @Description Get Contract ratings
// @Accept  json
// @Produce  json
// @Param id path string true "Contract id"
// @Success 200 {object} view.Response{payload=[]view.ContractRatingsView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /contracts/{id}/ratings [get]
// @Security JWT
func GetContractRatings(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	ratings := []model.Rating{}
	contract := &model.Contract{
		ID: ctx.Param("id"),
	}

	// get contract
	result := db.Where("id = ?", contract.ID).Find(&contract)

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

	// check if user is permitted for contract
	if contract.RecruiterID != claims.User.ID || contract.ProfessionalID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	// get ratings
	resultRatings := db.Where("contract_id = ?", contract.ID).Find(&ratings)

	if resultRatings.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: resultRatings.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	// TODO: It can be optimized
	var formatted_ratings []view.RatingView
	for _, rating := range ratings {
		formatted_ratings = append(formatted_ratings, view.RatingModelToView(rating))
	}

	// check is rating finished
	IsRatingFinished := false
	if len(formatted_ratings) >= 2 {
		IsRatingFinished = true
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.ContractRatingsView{
			ContractID:       contract.ID,
			JobID:            contract.JobID,
			RecruiterID:      contract.RecruiterID,
			ProfessionalID:   contract.ProfessionalID,
			IsRatingFinished: IsRatingFinished,
			Ratings:          formatted_ratings,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
