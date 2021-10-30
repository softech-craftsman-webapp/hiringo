package location

import (
	config "hiringo/config"
	view "hiringo/view"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/labstack/echo/v4"
)

type CreateLocationRequest struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
}

/*
   |--------------------------------------------------------------------------
   | Create Location
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Location
// @Tags location
// @Description Create Location
// @Accept  json
// @Produce  json
// @Param user body CreateLocationRequest true "Geolocation coordinates (longitude and latitude)"
// @Success 200 {object} view.Response{payload=view.LocationView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /locations/search [post]
// @Security JWT
func GetLocation(ctx echo.Context) error {
	req := new(CreateLocationRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: "Bad request",
			Payload: nil,
		})
	}

	geocoder := openstreetmap.Geocoder()
	address, error := geocoder.ReverseGeocode(req.Latitude, req.Longitude)

	if error != nil {
		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: error.Error(),
			Payload: nil,
		})
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.LocationView{
			Longitude:     req.Longitude,
			Latitude:      req.Latitude,
			Address:       address.FormattedAddress,
			Street:        address.Street,
			HouseNumber:   address.HouseNumber,
			Suburb:        address.Suburb,
			Postcode:      address.Postcode,
			State:         address.State,
			StateCode:     address.StateCode,
			StateDistrict: address.StateDistrict,
			County:        address.County,
			Country:       address.Country,
			CountryCode:   address.CountryCode,
			City:          address.City,
		},
	}

	return view.ApiView(http.StatusOK, ctx, resp)
}
