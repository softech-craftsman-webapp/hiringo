package userDetail

import (
	"net/http"

	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type RevealUserDetailRequest struct {
	ContractID string `json:"contract_id"`
}

/*
   |--------------------------------------------------------------------------
   | Get user details by contract id
   | @JWT via Acess Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Get user details by contract id
// @Tags user-detail
// @Description Get user details by contract id
// @Accept  json
// @Produce  json
// @Param id path string true "User Detail id"
// @Param user body RevealUserDetailRequest true "Contract details"
// @Success 200 {object} view.Response{payload=view.UserDetailView}
// @Failure 400,401,403,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /user-details/{id}/reveal [post]
// @Security JWT
func RevealUserDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(RevealUserDetailRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: err.Error(),
			Payload: nil,
		})
	}

	userDetail := &model.UserDetail{
		ID: ctx.Param("id"),
	}

	userDetail_result := db.First(&userDetail)
	if userDetail_result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "User detail not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	contract := &model.Contract{
		ID: req.ContractID,
	}

	contract_result := db.First(&contract)
	if contract_result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Contract not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	if userDetail.UserID != claims.User.ID {
		if contract.RecruiterID != claims.User.ID {
			if contract.ProfessionalID != claims.User.ID {
				resp := &view.Response{
					Success: true,
					Message: "You are not allowed to access this contract",
					Payload: nil,
				}
				// close db
				config.CloseDB(db).Close()

				return view.ApiView(http.StatusForbidden, ctx, resp)
			}
		}
	}

	if contract.SignedByRecruiterTime.IsZero() || contract.SignedByProfessionalTime.IsZero() {
		resp := &view.Response{
			Success: true,
			Message: "Contract not yet signed",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.UserDetailView{
			ID:        userDetail.ID,
			UserID:    userDetail.UserID,
			Email:     userDetail.Email,
			Telephone: userDetail.Telephone,
			Bio:       userDetail.Bio,
			Latitude:  userDetail.Latitude,
			Longitude: userDetail.Longitude,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
