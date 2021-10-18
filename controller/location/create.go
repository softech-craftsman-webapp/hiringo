package location

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/codingsince1985/geo-golang/openstreetmap"
	"github.com/golang-jwt/jwt"
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
// @Router /locations/new [post]
// @Security JWT
func CreateLocation(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateLocationRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: "Bad request",
			Payload: nil,
		})
	}

	geocoder := openstreetmap.Geocoder()
	address, error := geocoder.ReverseGeocode(req.Latitude, req.Longitude)

	if error != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "Geocode could not find an Address",
			Payload: nil,
		})
	}

	location := &model.Location{
		UserID:        claims.User.ID,
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
	}

	result := db.Create(&location)
	/*
	   |--------------------------------------------------------------------------
	   | DB relation error
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

		return view.ApiView(http.StatusCreated, ctx, resp)
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
