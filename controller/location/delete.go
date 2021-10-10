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
   | Delete location
   | @JWT via Access Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Delete Location
// @Tags location
// @Description Delete Location
// @Accept  json
// @Produce  json
// @Param id path string true "Location id"
// @Success 200 {object} view.Response{payload=view.LocationEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /locations{id} [delete]
// @Security JWT
func DeleteLocation(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	location := &model.Location{
		ID: ctx.Param("id"),
	}

	db.First(&location, "id = ? AND user_id = ?", location.ID, claims.User.ID)

	if location.UserID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	result := db.Delete(&location)

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.LocationEmptyView{
			ID: location.ID,
		},
	}

	/*
	   |--------------------------------------------------------------------------
	   | Main Error
	   |--------------------------------------------------------------------------
	*/
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

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
