package user

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	db "github.com/danielronalds/messenger-server/db/dbtypes"
	"github.com/danielronalds/messenger-server/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]db.User{
		"johnsmith": {UserName: "johnsmith", DisplayName: "John Smith"},
		"janesmith": {UserName: "janesmith", DisplayName: "Jane Smith"},
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

func TestCreateUserPassing(t *testing.T) {
	newUserJson := `{"username":"jonsnow","displayname":"Jon Snow","password":"winterIsComing"}`
	expectedRecJson := `{"username":"jonsnow","displayname":"Jon Snow"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewUserHandler(provider)

	if assert.NoError(t, handler.CreateUser(c)) {
		trimmedRecBody := strings.TrimRight(rec.Body.String(), "\n")

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedRecJson, trimmedRecBody)
	}
}

func TestCreateUserMissingField(t *testing.T) {
	newUserJson := `{"username":"jonsnow","password":"winterIsComing"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewUserHandler(provider)

	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestCreateUserDuplicateUsername(t *testing.T) {
	newUserJson := `{"username":"johnsmith","displayname":"Jon Snow","password":"winterIsComing"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(newUserJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewUserHandler(provider)

	if assert.NoError(t, handler.CreateUser(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}
