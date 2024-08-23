package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielronalds/messenger-server/resources"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load Env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/users", resources.GetUsers)
	e.POST("/users", resources.CreateUser)

	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println(port)
	e.Logger.Fatal(e.Start(port))
}
