package userDetail

import (
	"net/http"

	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get user details
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get user's details
// @Tags user-detail
// @Description Get user's details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.UserDetailView}
// @Failure 400,401,403,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /user-details/my [get]
// @Security JWT
func MyUserDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	userDetail := &model.UserDetail{
		UserID: claims.User.ID,
	}

	userDetail_result := db.Where("user_id = ?", claims.User.ID).First(&userDetail)

	if userDetail_result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: userDetail_result.Error.Error(),
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
