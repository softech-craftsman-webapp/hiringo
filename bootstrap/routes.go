package bootstrap

import (
	config "hiringo/config"
	controller "hiringo/controller"
	category_controller "hiringo/controller/category"
	job_controller "hiringo/controller/job"
	location_controller "hiringo/controller/location"
	rating_controller "hiringo/controller/rating"
	transaction_controller "hiringo/controller/transaction"
	user_detail_controller "hiringo/controller/userDetail"
	_ "hiringo/docs"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

/*
	|--------------------------------------------------------------------------
	| Routes and its middleware
	|--------------------------------------------------------------------------
*/
func InitRoutes(app *echo.Echo) {
	// Access, Refresh Application Routes
	access_route := config.Guard(app)

	// enable validation
	app.Validator = &config.CustomValidator{Validator: validator.New()}

	// Swagger
	app.GET("/openapi/*", echoSwagger.WrapHandler)
	app.GET("/openapi", controller.SwaggerRedirect)

	// Location
	access_route.POST("/locations/search", location_controller.GetLocation)

	// Category
	access_route.GET("/categories/all", category_controller.GetAllCategories)
	access_route.POST("/categories/new", category_controller.CreateCategory)
	access_route.DELETE("/categories/:id", category_controller.DeleteCategory)
	access_route.GET("/categories/:id", category_controller.GetCategoryDetail)

	// Transaction
	access_route.POST("/transactions/my", transaction_controller.GetMyTransactions)
	access_route.POST("/transactions/new", transaction_controller.CreateTransaction)
	access_route.GET("/transactions/:id", transaction_controller.GetTransactionDetail)

	// Rating
	access_route.POST("/ratings/new", rating_controller.CreateRating)
	access_route.GET("/ratings/:id", rating_controller.GetRatingDetail)

	// Job
	access_route.POST("/jobs/new", job_controller.CreateJob)
	access_route.DELETE("/jobs/:id", job_controller.DeleteJob)
	access_route.POST("/jobs/search", job_controller.SearchJobs)
	access_route.GET("/jobs/:id", job_controller.GetJobDetail)
	access_route.PUT("/jobs/:id", job_controller.UpdateJob)
	access_route.PUT("/jobs/:id/image", job_controller.AddOrUpdateJobImage)

	// User Details
	access_route.POST("/user-detail/new", user_detail_controller.CreateUserDetail)
	access_route.GET("/user-detail/:id/rating", user_detail_controller.GetUserRating)
	access_route.PUT("/user-details/:id", user_detail_controller.UpdateUserDetail)
	access_route.POST("/user-details/:id/reveal", user_detail_controller.RevealUserDetail)
	access_route.GET("/user-details/my", user_detail_controller.MyUserDetail)
}
