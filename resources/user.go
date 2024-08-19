package resources

import (
	"log"
	"net/http"

	"github.com/danielronalds/messenger-server/db"
	"github.com/labstack/echo/v4"
)

type user struct {
	Id          int
	UserName    string
	DisplayName string
}

func GetUsers(c echo.Context) error {
	pg := db.GetDatabase()

	users := []user{}

	err := pg.Select(&users, "SELECT Id, UserName, DisplayName FROM api.Users")

	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to fetch Users")
		log.Fatalf("Failed to get users: %v", err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
