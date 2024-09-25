package auth

import (
	"log"
	"net/http"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/resources"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/labstack/echo/v4"
)

type LoginReturn struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayname"`
}

func Login(c echo.Context) error {
	postedUser := new(resources.PostedUser)

	if err := c.Bind(&postedUser); err != nil {
		log.Printf("Failed to bind login details: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !postedUser.IsValid() {
		return c.String(http.StatusBadRequest, "Login details were not valid!")
	}

	dbUser, err := db.GetDatabase().GetUserWithPass(postedUser.UserName, postedUser.Password)

	if err != nil {
		log.Printf("Failed to fetch user with login: %v", err.Error())
		return c.String(http.StatusNotFound, "Could not find user")
	}

	// Create a session
	sessionKey, err := stores.GetUserStore().CreateSession(dbUser.UserName)

	if err != nil {
		log.Printf("Failed to create login session")
		return c.String(http.StatusInternalServerError, "Failed to create session")
	}

	// Return session key to user
	return c.JSON(http.StatusOK, LoginReturn{
		Key:         sessionKey,
		DisplayName: dbUser.DisplayName,
	})
}

type LogoutStruct struct {
	Key string `json:"key"`
}

func Logout(c echo.Context) error {
	// Key is in the body of the request instead of path as it is still a secret, despite it about
	// to be deleted
	logoutDetails := new(LogoutStruct)

	if err := c.Bind(&logoutDetails); err != nil {
		log.Printf("Failed to bind logout details: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	userStore := stores.GetUserStore()
	session := userStore.GetSession(logoutDetails.Key)

	if session == nil {
		log.Printf("Users session not found")
		return c.String(http.StatusNotFound, "Could not find user session")
	}

	userStore.DeleteSession(logoutDetails.Key)

	// Return session key to user
	return c.String(http.StatusOK, "Session removed")
}
