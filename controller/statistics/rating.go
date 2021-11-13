package statistics

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
   | Rating statistics
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Rating statistics
// @Tags statistics
// @Description Rating statistics
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.RatingStatistics}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /statistics/rating [get]
// @Security JWT
func RatingStatistics(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	// get user ratings
	userRatings := []model.Rating{}
	db.Where("user_id = ?", claims.User.ID).Find(&userRatings)

	// ratings
	ratings := []model.Rating{}
	db.Find(&ratings)

	// get latest rating
	latestRating := model.Rating{
		UserID: claims.User.ID,
	}

	db.Where("user_id = ?", claims.User.ID).Order("created_at desc").First(&latestRating)

	// result
	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.RatingStatistics{
			LatestRating: &view.RatingView{
				ID:            latestRating.ID,
				SubmittedByID: latestRating.SubmittedByID,
				ContractID:    latestRating.ContractID,
				UserID:        latestRating.UserID,
				Points:        latestRating.Points,
				Comment:       latestRating.Comment,
			},
			UserRatingCount: len(userRatings),
			Total:           len(ratings),
			Time:            latestRating.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
