package main

import (
	"fmt"
	"os"

	bootstrap "hiringo/bootstrap"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Hiringo API
// @version 1.0
// @description Hiringo API Service.

// @host 127.0.0.1:8888
// @BasePath /

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(".env file is not imported, in production kindly ignore this message")
	}

	/*
	   |--------------------------------------------------------------------------
	   | Start Server
	   |--------------------------------------------------------------------------
	*/
	app := echo.New()
	port := os.Getenv("PORT")

	// Application
	bootstrap.Start(app, port)
}
