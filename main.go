package main

import (
	"fmt"
	"log"
	"os"

	"github.com/danielronalds/messenger-server/resources/auth"
	"github.com/danielronalds/messenger-server/resources/user"
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

	e.GET("/users", user.GetUsers)
	e.POST("/users", user.CreateUser)

	e.POST("/auth", auth.Login)

	port := fmt.Sprintf(":%v", os.Getenv("SERVER_PORT"))
	fmt.Println(port)
	e.Logger.Fatal(e.Start(port))
}
