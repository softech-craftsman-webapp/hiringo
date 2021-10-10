package bootstrap

import (
	config "hiringo/config"
	controller "hiringo/controller"
	category_controller "hiringo/controller/category"
	location_controller "hiringo/controller/location"
	transaction_controller "hiringo/controller/transaction"
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
	access_route.POST("/locations", location_controller.CreateLocation)
	access_route.DELETE("/locations/:id", location_controller.DeleteLocation)

	// Category
	access_route.POST("/categories", category_controller.CreateCategory)
	access_route.DELETE("/categories/:id", category_controller.DeleteCategory)

	// Transaction
	access_route.POST("/transactions", transaction_controller.CreateTransaction)
}
