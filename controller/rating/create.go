package rating

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateRatingRequest struct {
	UserID     string `json:"user_id" validate:"required"`
	ContractID string `json:"contract_id" validate:"required"`
	Points     int    `json:"points" validate:"required,gte=1,lte=5"`
	Comment    string `json:"comment" validate:"required"`
}

/*
   |--------------------------------------------------------------------------
   | Create Rating
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Rating
// @Tags rating
// @Description Create Rating
// @Accept  json
// @Produce  json
// @Param user body CreateRatingRequest true "User id (send points to user) and points"
// @Success 200 {object} view.Response{payload=view.RatingView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /ratings/new [post]
// @Security JWT
func CreateRating(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateRatingRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: err.Error(),
			Payload: nil,
		})
	}

	/*
	   |--------------------------------------------------------------------------
	   | User should not be able to rate himself
	   |--------------------------------------------------------------------------
	*/
	if claims.User.ID == req.UserID {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "You can't rate yourself",
			Payload: nil,
		})
	}

	/*
	   |--------------------------------------------------------------------------
	   | Find contract by id
	   |--------------------------------------------------------------------------
	*/
	contract := new(model.Contract)
	if err := db.Where("id = ?", req.ContractID).First(&contract).Error; err != nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: err.Error(),
			Payload: nil,
		})
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check if rating already exists
	   |--------------------------------------------------------------------------
	*/
	if err := db.Where("user_id = ? AND contract_id = ?", req.UserID, req.ContractID).First(&model.Rating{}).Error; err == nil {
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "Rating already exists",
			Payload: nil,
		})
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check contract
	   |--------------------------------------------------------------------------
	*/
	if contract.RecruiterID != claims.User.ID {
		if contract.ProfessionalID != claims.User.ID {
			config.CloseDB(db).Close()

			return view.ApiView(http.StatusForbidden, ctx, &view.Response{
				Success: false,
				Message: "You are not allowed to rate this contract",
				Payload: nil,
			})
		}
	}

	rating := &model.Rating{
		SubmittedByID: claims.User.ID,
		ContractID:    req.ContractID,
		UserID:        req.UserID,
		Points:        req.Points,
		Comment:       req.Comment,
	}

	result := db.Create(&rating)
	/*
	   |--------------------------------------------------------------------------
	   | DB relation error
	   |--------------------------------------------------------------------------
	*/
	if result.Error != nil {
		resp := &view.Response{
			Success: false,
			Message: "Duplicate error",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, resp)
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
