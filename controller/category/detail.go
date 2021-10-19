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
   | Get Category Detail
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Get category Details
// @Tags category
// @Description Get category Details
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.CategoryView}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /categories/{id} [get]
// @Security JWT
func GetCategoryDetail(ctx echo.Context) error {
	db := config.GetDB()

	category := model.Category{
		ID: ctx.Param("id"),
	}
	result := db.First(&category)

	if result.Error != nil {
		resp := &view.Response{
			Success: true,
			Message: "Category not found",
			Payload: nil,
		}
		// close db
		config.CloseDB(db).Close()

		return view.ApiView(http.StatusNotFound, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryView{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			CreatedByID: category.CreatedByID,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
