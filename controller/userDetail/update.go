package userDetail

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UpdateUserDetailRequest struct {
	Email     string  `json:"email" validate:"required,email"`
	Telephone string  `json:"telephone" validate:"required"`
	Bio       string  `json:"bio" validate:"required"`
	Latitude  float64 `json:"latitude" validate:"required,numeric"`
	Longitude float64 `json:"longitude" validate:"required,numeric"`
}

/*
   |--------------------------------------------------------------------------
   | Update user details
   | @JWT via Acess Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Update user details
// @Tags user-detail
// @Description Update user details
// @Accept  json
// @Produce  json
// @Param id path string true "User Detail id"
// @Param user body UpdateUserDetailRequest true "User details"
// @Success 200 {object} view.Response{payload=view.UserDetailEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /user-details/{id} [put]
// @Security JWT
func UpdateUserDetail(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(UpdateUserDetailRequest)

	/*
	   |--------------------------------------------------------------------------
	   | Bind request
	   |--------------------------------------------------------------------------
	*/
	if err := config.BindAndValidate(ctx, req); err != nil {
		config.CloseDB(db).Close()

		return ctx.JSON(http.StatusBadRequest, &view.Response{
			Success: false,
			Message: config.GetMessageFromError(err.Error()),
			Payload: nil,
		})
	}

	userDetail := &model.UserDetail{
		UserID: ctx.Param("id"),
	}

	db.First(&userDetail, "user_id = ?", claims.User.ID)

	/*
	   |--------------------------------------------------------------------------
	   | Check if user's id the same as the logged in user
	   |--------------------------------------------------------------------------
	*/
	if userDetail.UserID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	/*
	   |--------------------------------------------------------------------------
	   | Check rquired fields
	   |--------------------------------------------------------------------------
	*/
	result := db.Model(&userDetail).Updates(model.UserDetail{
		Email:     req.Email,
		Telephone: req.Telephone,
		Bio:       req.Bio,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	})

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.UserDetailEmptyView{
			ID: userDetail.ID,
		},
	}

	/*
	   |--------------------------------------------------------------------------
	   | Main Error
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

		return view.ApiView(http.StatusInternalServerError, ctx, resp)
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
