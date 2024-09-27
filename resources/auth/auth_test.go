package auth

import (
	"encoding/json"
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
)

func TestLoginPassing(t *testing.T) {
	loginJson := `{"username":"johnsmith","password":"password"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		session := LoginReturn{}
		json.Unmarshal(rec.Body.Bytes(), &session)

		assert.Equal(t, session.DisplayName, "John Smith")
	}
}

func TestLoginMissingField(t *testing.T) {
	loginJson := `{"password":"wrongpassword"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestLoginIncorrectPassword(t *testing.T) {
	loginJson := `{"username":"johnsmith","password":"wrongpassword"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestLoginInvalidUsername(t *testing.T) {
	loginJson := `{"username":"jonsnow","password":"password"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	if assert.NoError(t, handler.Login(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}

func TestLogoutPassing(t *testing.T) {
	// Creating a session by logging in
	loginJson := `{"username":"johnsmith","password":"password"}`

	e := echo.New()
	loginReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(loginJson))
	loginReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	loginRec := httptest.NewRecorder()
	c := e.NewContext(loginReq, loginRec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	utils.HandleTestingError(t, handler.Login(c))

	session := LoginReturn{}
	json.Unmarshal(loginRec.Body.Bytes(), &session)

	logoutJson := `{"key":"` + session.Key + `"}`

	logoutReq := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(logoutJson))
	logoutReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	logoutRec := httptest.NewRecorder()
	c = e.NewContext(logoutReq, logoutRec)

	if assert.NoError(t, handler.Logout(c)) {
		assert.Equal(t, http.StatusOK, logoutRec.Code)
	}
}

func TestLogoutInvalidKey(t *testing.T) {
	// Creating a session by logging in
	logoutJson := `{"key":"OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(logoutJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedUserProvider(mockDB)
	handler := NewAuthHandler(provider)

	if assert.NoError(t, handler.Logout(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
