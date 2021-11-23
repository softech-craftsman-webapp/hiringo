package rating

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
   | Get ratings of user
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get user ratings
// @Tags rating
// @Description Get user ratings
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=[]view.RatingView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /ratings/my [get]
// @Security JWT
func GetMyRatings(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)
	db := config.GetDB()

	ratings := []model.Rating{}
	result := db.Where("user_id = ?", claims.User.ID).Find(&ratings).Order("created_at DESC")

	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: result.Error.Error(),
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	// TODO: It can be optimized
	var formatted_ratings []view.RatingView
	for _, rating := range ratings {
		formatted_ratings = append(formatted_ratings, view.RatingModelToView(rating))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_ratings,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
