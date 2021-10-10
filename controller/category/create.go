package category

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

/*
   |--------------------------------------------------------------------------
   | Create Category
   | @JWT via Access Token
   |--------------------------------------------------------------------------
*/
// Create Category
// @Tags category
// @Description Create Category
// @Accept  json
// @Produce  json
// @Param user body CreateCategoryRequest true "Category for Job"
// @Success 200 {object} view.Response{payload=view.CategoryView}
// @Failure 400,401,403,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /categories [post]
// @Security JWT
func CreateCategory(ctx echo.Context) error {
	claims := ctx.Get("user").(*jwt.Token).Claims.(*view.JwtCustomClaims)

	db := config.GetDB()
	req := new(CreateCategoryRequest)

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

	category := &model.Category{
		CreatedBy:   claims.User.ID,
		Name:        req.Name,
		Description: req.Description,
	}

	result := db.Create(&category)
	/*
	   |--------------------------------------------------------------------------
	   | DB relation error
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

		return view.ApiView(http.StatusCreated, ctx, resp)
	}

	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryView{
			ID:          category.ID,
			CreatedBy:   category.CreatedBy,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusInternalServerError, ctx, resp)
}
