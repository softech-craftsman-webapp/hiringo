package statistics

import (
	config "hiringo/config"
	model "hiringo/model"
	view "hiringo/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
   |--------------------------------------------------------------------------
   | Category statistics
   | @JWT via Acess Token
   |--------------------------------------------------------------------------
*/
// Category statistics
// @Tags statistics
// @Description Category statistics
// @Accept  json
// @Produce  json
// @Success 200 {object} view.Response{payload=view.CategoryStatistics}
// @Failure 400,401,404,500 {object} view.Response
// @Failure default {object} view.Response
// @Router /statistics/category [get]
// @Security JWT
func CategoryStatistics(ctx echo.Context) error {
	db := config.GetDB()

	job := model.Job{}

	// get the most popular categories from the jobs table
	db.Table("jobs").Select("category_id, count(*) as total").Group("category_id").Order("total desc").Limit(1).Scan(&job)

	// most popular category
	category := model.Category{
		ID: job.CategoryID,
	}

	db.Where("id = ?", category.ID).First(&category)

	// get total list number of the categories
	categories := []model.Category{}
	db.Find(&categories)

	// result
	resp := &view.Response{
		Success: true,
		Message: "Success",
		Payload: &view.CategoryStatistics{
			PopularItem: &view.CategoryView{
				ID:          category.ID,
				CreatedByID: category.CreatedByID,
				Name:        category.Name,
				Description: category.Description,
			},
			Total: len(categories),
			Time:  category.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// close db
	config.CloseDB(db).Close()

	return view.ApiView(http.StatusOK, ctx, resp)
}
