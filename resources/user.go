package resources

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/utils/security"
	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := db.GetDatabase().GetUsers()

	if err != nil {
		log.Printf("Failed to get users: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to fetch Users")
	}

	return c.JSON(http.StatusOK, users)
}

// A struct to represent the JSON posted to the CreateUser endpoint
type postedUser struct {
	UserName    string
	DisplayName string
	Password    string
}

// A method to check whether the user object recieved in the POST request is valid
func (u postedUser) isValid() bool {
	trimmedUserName := strings.TrimSpace(u.UserName)
	trimmedDisplayName := strings.TrimSpace(u.DisplayName)
	trimmedPassword := strings.TrimSpace(u.Password)

	return len(trimmedUserName) > 0 && len(trimmedDisplayName) > 0 && len(trimmedPassword) > 0
}

func CreateUser(c echo.Context) error {
	user := new(postedUser)

	if err := c.Bind(&user); err != nil {
		log.Printf("Failed to bind posted user: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !user.isValid() {
		return c.String(http.StatusBadRequest, "A field was either missing or blank!")
	}

	hasher := security.DefaultHash()

	hashedPassword, err := hasher.GenerateNewHash([]byte(user.Password))

	if err != nil {
		log.Printf("Failed to generate a hash: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to generate a hash")
	}

	pg := db.GetDatabase()

	newUser, err := pg.CreateUser(user.UserName, user.DisplayName, hashedPassword.Hash(), hashedPassword.Salt())
	if err != nil {
		log.Printf("Failed to create user: %v", err.Error())
		return c.String(http.StatusInternalServerError, "Failed to create user")
	}

	return c.JSON(http.StatusOK, newUser)
}
