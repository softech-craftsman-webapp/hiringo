package location

import (
	config "hiringo/config"
	view "hiringo/view"

	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateLocationRequest struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required"`
}

func reverseGeoCode(lat float64, lng float64) (view.LocationView, error) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	address := view.LocationView{}
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f&zoom=18&addressdetails=1", lat, lng)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return address, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return address, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return address, err
	}

	json.Unmarshal(bodyBytes, &address)
	return address, nil
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
			Message: err.Error(),
			Payload: nil,
		})
	}

	address, error := reverseGeoCode(req.Latitude, req.Longitude)

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
			PlaceID:     address.PlaceID,
			OsmType:     address.OsmType,
			OsmID:       address.OsmID,
			Lat:         address.Lat,
			Lon:         address.Lon,
			DisplayName: address.DisplayName,
			Address:     address.Address,
			Boundingbox: address.Boundingbox,
		},
	}

	return view.ApiView(http.StatusOK, ctx, resp)
}
