package bootstrap

import (
	config "hiringo/config"
	controller "hiringo/controller"
	category_controller "hiringo/controller/category"
	contract_controller "hiringo/controller/contract"
	job_controller "hiringo/controller/job"
	location_controller "hiringo/controller/location"
	rating_controller "hiringo/controller/rating"
	statistics_controller "hiringo/controller/statistics"
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
	app.GET("/locations/coordinates", location_controller.GetCoordinatesUser)
	app.POST("/locations/search", location_controller.GetLocation)

	// Category
	access_route.GET("/categories/:id", category_controller.GetCategoryDetail)
	access_route.GET("/categories/all", category_controller.GetAllCategories)
	access_route.POST("/categories/new", category_controller.CreateCategory)
	access_route.DELETE("/categories/:id", category_controller.DeleteCategory)

	// Transaction
	access_route.GET("/transactions/my", transaction_controller.GetMyTransactions)
	access_route.GET("/transactions/:id", transaction_controller.GetTransactionDetail)
	access_route.POST("/transactions/new", transaction_controller.CreateTransaction)

	// Rating
	access_route.GET("/ratings/my", rating_controller.GetMyRatings)
	access_route.GET("/ratings/:id", rating_controller.GetRatingDetail)
	access_route.POST("/ratings/new", rating_controller.CreateRating)

	// Job
	access_route.GET("/jobs/my", job_controller.GetMyJobs)
	access_route.GET("/jobs/:id", job_controller.GetJobDetail)
	access_route.GET("/jobs/:id/contracts", job_controller.GetJobContracts)
	access_route.POST("/jobs/search", job_controller.SearchJobs)
	access_route.POST("/jobs/new", job_controller.CreateJob)
	access_route.PUT("/jobs/:id", job_controller.UpdateJob)
	access_route.PUT("/jobs/:id/image", job_controller.AddOrUpdateJobImage)
	access_route.DELETE("/jobs/:id", job_controller.DeleteJob)

	// User Details
	access_route.GET("/user-details/my", user_detail_controller.MyUserDetail)
	access_route.GET("/user-details/:id/rating", user_detail_controller.GetUserRating)
	access_route.POST("/user-details/:id/reveal", user_detail_controller.RevealUserDetail)
	access_route.POST("/user-details/new", user_detail_controller.CreateUserDetail)
	access_route.PUT("/user-details/edit", user_detail_controller.UpdateUserDetail)

	// Contracts
	access_route.GET("/contracts/my", contract_controller.GetJobContracts)
	access_route.GET("/contracts/:id", contract_controller.GetContractDetail)
	access_route.GET("/contracts/:id/ratings", contract_controller.GetContractRatings)
	access_route.POST("/contracts/new", contract_controller.CreateContract)
	access_route.POST("/contracts/:id/sign", contract_controller.SignContract)
	access_route.PUT("/contracts/:id", contract_controller.UpdateContractDetail)
	access_route.DELETE("/contracts/:id", contract_controller.DeleteContract)

	// Statistics
	access_route.GET("/statistics/category", statistics_controller.CategoryStatistics)
	access_route.GET("/statistics/job", statistics_controller.JobStatistics)
	access_route.GET("/statistics/rating", statistics_controller.RatingStatistics)
	access_route.GET("/statistics/transaction", statistics_controller.TransactionStatistics)
}
