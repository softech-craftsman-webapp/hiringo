package main

import (
	"fmt"
	"log"
	"os"
	"time"

	bootstrap "hiringo/bootstrap"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

// @title Hiringo API
// @version 1.0
// @description Hiringo API Service.

// @host https://main-api.hiringo.tech
// @BasePath /

// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func main() {
	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env file is not imported, in production kindly ignore this message")
	}

	// Set timezone globally
	if tz := os.Getenv("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("error loading location '%s': %v\n", tz, err)
		}
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
