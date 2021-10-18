package location

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
   | Get categories
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get user details
// @Tags users
// @Description Get user details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=[]view.LocationView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /locations/my [get]
// @Security JWT
func GetMyLocations(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	locations := []model.Location{}
	result := db.Where("user_id = ?", claims.User.ID).Find(&locations)

	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Internal Server Error",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	// TODO: It can be optimized
	var formatted_locations []view.LocationView
	for _, location := range locations {
		formatted_locations = append(formatted_locations, view.LocationModelToView(location))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_locations,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
