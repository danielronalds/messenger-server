package auth

import (
	"log"
	"net/http"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/resources"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	loginDetails := new(resources.LoginAttempt)

	if err := c.Bind(&loginDetails); err != nil {
		log.Printf("Failed to bind login details: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !loginDetails.IsValid() {
		return c.String(http.StatusBadRequest, "Login details were not valid!")
	}

	dbUser, err := db.GetDatabase().GetUserWithPass(loginDetails.Id, loginDetails.Password)

	if err != nil {
		log.Printf("Failed to fetch user with login: %v", err.Error())
		return c.String(http.StatusNotFound, "Could not find user")
	}

	// Create a session
	sessionKey, err := stores.GetUserStore().CreateSession(dbUser.Id, dbUser.DisplayName)

	if err != nil {
		log.Printf("Failed to create login session")
		return c.String(http.StatusInternalServerError, "Failed to create session")
	}

	// Return session key to user
	return c.String(http.StatusOK, sessionKey)
}
