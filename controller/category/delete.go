package category

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
   | Delete category
   | @JWT via Access Token
   | @Param id
   |--------------------------------------------------------------------------
*/
// Delete category
// @Tags category
// @Description Delete category
// @Accept  json
// @Produce  json
// @Param id path string true "Category id"
// @Success 200 {object} view.Response{payload=view.CategoryEmptyView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /categories/{id} [delete]
// @Security JWT
func DeleteCategory(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()

	category := &model.Category{
		ID: ctx.Param("id"),
	}

	db.First(&category, "id = ? AND user_id = ?", category.ID, claims.User.ID)

	if category.CreatedByID != claims.User.ID {
		resp := &view.Response{
			Success: true,
			Message: "Forbidden",
			Payload: nil,
		}

		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusForbidden, ctx, resp)
	}

	result := db.Delete(&category)

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryEmptyView{
			ID: category.ID,
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
