package user

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/resources"
	"github.com/danielronalds/messenger-server/security"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db db.UserProvider
}

func NewUserHandler(db db.UserProvider) UserHandler {
	return UserHandler { db }
}

func (h UserHandler) GetUsers(c echo.Context) error {
	users, err := h.db.GetUsers()

	if err != nil {
		log.Printf("Failed to get users: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to fetch Users")
	}

	return c.JSON(http.StatusOK, users)
}

func (h UserHandler) CreateUser(c echo.Context) error {
	user := new(resources.PostedNewUser)

	if err := c.Bind(&user); err != nil {
		log.Printf("Failed to bind posted user: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !user.IsValid() {
		return c.String(http.StatusBadRequest, "A field was either missing or blank!")
	}

	hasher := security.DefaultHash()

	hashedPassword, err := hasher.GenerateNewHash([]byte(strings.TrimSpace(user.Password)))

	if err != nil {
		log.Printf("Failed to generate a hash: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to generate a hash")
	}

	newUser, err := h.db.CreateUser(user.UserName, user.DisplayName, hashedPassword.Hash(), hashedPassword.Salt())
	if err != nil {
		log.Printf("Failed to create user: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusCreated, newUser)
}
