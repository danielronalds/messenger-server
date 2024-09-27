package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]db.User{
		"johnsmith": {UserName: "johnsmith", DisplayName: "John Smith"},
		"janesmith": {UserName:"janesmith", DisplayName:"Jane Smith"},
	}
	dbJSON = `[{"username":"johnsmith","displayname":"John Smith"},{"username":"janesmith","displayname":"Jane Smith"}]`
)

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewUserHandler(provider)

	if assert.NoError(t, handler.GetUsers(c)) {
		trimmedRecBody := strings.TrimRight(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, dbJSON, trimmedRecBody)
	}
}
