package category

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Get categories
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get user details
// @Tags category
// @Description Get user details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=[]view.PublicCategoryView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /categories/all [get]
// @Security JWT
func GetAllCategories(ctx echo.Context) error {
	db := config.GetDB()

	categories := []model.Category{}
	result := db.Find(&categories)

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
	var formatted_categories []view.PublicCategoryView
	for _, category := range categories {
		formatted_categories = append(formatted_categories, view.CategoryModelToView(category))
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: formatted_categories,
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
