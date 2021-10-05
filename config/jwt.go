package config

import (
	crypto "hiringo/crypto"
	view "hiringo/view"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

/*
   |--------------------------------------------------------------------------
   | JWT Middleware
   |--------------------------------------------------------------------------
*/
func Guard(app *echo.Echo) *echo.Group {
	// Keys
	// @Access
	access_key, error := crypto.AccessPublicKey()
	if error != nil {
		panic(error)
	}

	// Routes
	access_route := app.Group("")

	// Jwt middleware @Access
	access_route.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:        &view.JwtCustomClaims{},
		SigningMethod: crypto.SigningMethodName,
		SigningKey:    access_key,
	}))

	return access_route
}
