package location

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get Location Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get location Details
// @Tags location
// @Description Get location Details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.LocationView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /locations/{id} [get]
// @Security JWT
func GetLocationDetail(ctx echo.Context) error {
	db := config.GetDB()

	location := model.Location{
		ID: ctx.Param("id"),
	}
	result := db.First(&location)

	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Location not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.LocationView{
			ID:            location.ID,
			UserID:        location.UserID,
			Longitude:     location.Longitude,
			Latitude:      location.Latitude,
			Address:       location.Address,
			Street:        location.Street,
			HouseNumber:   location.HouseNumber,
			Suburb:        location.Suburb,
			Postcode:      location.Postcode,
			State:         location.State,
			StateCode:     location.StateCode,
			StateDistrict: location.StateDistrict,
			County:        location.County,
			Country:       location.Country,
			CountryCode:   location.CountryCode,
			City:          location.City,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
