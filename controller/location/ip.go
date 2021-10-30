package location

import (
	view "hiringo/view"

	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func reverseIp(ip string) (view.CoordinatesView, error) {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	fmt.Println(ip)

	coordinates := view.CoordinatesView{}
	url := "http://ip-api.com/json/" + ip

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return coordinates, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return coordinates, err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return coordinates, err
	}

	json.Unmarshal(bodyBytes, &coordinates)
	return coordinates, nil
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
// @Success 200 {object} view.Response{payload=view.CoordinatesView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /locations/coordinates [get]
// @Security JWT
func GetCoordinatesUser(ctx echo.Context) error {
	ip := ctx.RealIP()
	coordinates, error := reverseIp(ip)

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
		Payload: &view.CoordinatesView{
			Ip:  ip,
			Lat: coordinates.Lat,
			Lon: coordinates.Lon,
		},
	}

	return view.ApiView(http.StatusOK, ctx, resp)
}
