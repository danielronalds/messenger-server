package resources

import (
	"log"
	"net/http"

	"github.com/danielronalds/messenger-server/db"
	"github.com/labstack/echo/v4"
)


func GetUsers(c echo.Context) error {
	pg := db.GetDatabase()

	users, err := pg.GetUsers()

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch Users")
		log.Fatalf("Failed to get users: %v", err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
