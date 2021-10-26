package userDetail

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get user's rating
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get average rating
// @Tags user-detail
// @Description Get user average rating
// @Accept  json
// @Produce  json
// @Param id path string true "Rating id"
// @Success 200 {object} view.Response{payload=view.UserRatingDetailView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /user-details/{id}/rating [get]
// @Security JWT
func GetUserRating(ctx echo.Context) error {
	db := config.GetDB()
	userID := ctx.Param("id")

	if userID == "" {
		resp := &view.Response{
			Success: true,
			Message: "User id is required",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	ratings := []model.Rating{}
	result := db.Where("user_id = ?", userID).Find(&ratings)

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
	total_rating := 0
	n_rating := 0
	average_rating := 0.00

	for _, rating := range ratings {
		total_rating += rating.Points
		n_rating++
	}

	// calculate average rating
	if n_rating > 0 {
		average_rating = float64(total_rating) / float64(n_rating)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: view.UserRatingDetailView{
			ID:         userID,
			Rating:     average_rating,
			TotalRates: n_rating,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
