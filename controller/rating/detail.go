package rating

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get Rating Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get rating Details
// @Tags rating
// @Description Get rating Details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.RatingView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /ratings/{id} [get]
// @Security JWT
func GetRatingDetail(ctx echo.Context) error {
	db := config.GetDB()

	rating := model.Rating{
		ID: ctx.Param("id"),
	}
	result := db.First(&rating)

	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Rating not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.RatingView{
			ID:            rating.ID,
			UserID:        rating.UserID,
			ContractID:    rating.ContractID,
			Points:        rating.Points,
			SubmittedByID: rating.SubmittedByID,
			Comment:       rating.Comment,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
