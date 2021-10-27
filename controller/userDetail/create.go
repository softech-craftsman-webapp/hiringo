package userDetail

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateUserDetailRequest struct {
	Email     string  `json:"email" validate:"required,email"`
	Telephone string  `json:"telephone" validate:"required,numeric"`
	Bio       string  `json:"bio" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required,numeric"`
	Longitude float64 `json:"longitude" validate:"required,numeric"`
}

/*
   |--------------------------------------------------------------------------
   | Create User Detail
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Category
// @Tags user-detail
// @Description Create Category
// @Accept  json
// @Produce  json
// @Param user body CreateUserDetailRequest true "User details for user"
// @Success 200 {object} view.Response{payload=view.UserDetailView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /user-details/new [post]
// @Security JWT
func CreateUserDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateUserDetailRequest)

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

	/*
	   |--------------------------------------------------------------------------
	   | Check user detail exist
	   |--------------------------------------------------------------------------
	*/
	if err := db.Where("user_id = ?", claims.User.ID).First(&model.UserDetail{}).Error; err == nil {
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusBadRequest, ctx, &view.Response{
			Success: false,
			Message: "User detail already exist",
			Payload: nil,
		})
	}

	userDetail := &model.UserDetail{
		Email:     req.Email,
		Telephone: req.Telephone,
		Bio:       req.Bio,
		UserID:    claims.User.ID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	result := db.Create(&userDetail)

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
		Payload: &view.UserDetailView{
			ID:        userDetail.ID,
			Email:     userDetail.Email,
			Telephone: userDetail.Telephone,
			Bio:       userDetail.Bio,
			UserID:    userDetail.UserID,
			Latitude:  userDetail.Latitude,
			Longitude: userDetail.Longitude,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
