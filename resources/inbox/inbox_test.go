package inbox

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"testing"
	"time"

	db "github.com/danielronalds/messenger-server/db/dbtypes"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/danielronalds/messenger-server/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string][]db.Message{
		"johnsmith": {{
			Id:        0,
			Sender:    "janesmith",
			Receiver:  "johnsmith",
			Content:   "Do you like the curtains?",
			Delivered: time.Time{},
			IsRead:    true,
		}, {
			Id:        2,
			Sender:    "janesmith",
			Receiver:  "johnsmith",
			Content:   "Too bad, we're keeping them",
			Delivered: time.Time{},
			IsRead:    false,
		}},
		"janesmith": {{
			Id:        1,
			Sender:    "johnsmith",
			Receiver:  "janesmith",
			Content:   "No, they're ugly",
			Delivered: time.Time{},
			IsRead:    true,
		}},
	}
	mockUsers = map[string]bool{
		"johnsmith": true,
		"janesmith": true,
	}
)

func TestGetMessagesPassing(t *testing.T) {
	key, err := stores.GetUserStore().CreateSession("johnsmith")
	utils.HandleTestingError(t, err)

	endpointJson := `{"key":"` + key + `","contact":"janesmith"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(endpointJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB, mockUsers)
	handler := NewInboxHandler(provider)

	if assert.NoError(t, handler.GetMessages(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		messages := make([]db.Message, 0)
		json.Unmarshal(rec.Body.Bytes(), &messages)


		expected := []db.Message {{
			Id:        0,
			Sender:    "janesmith",
			Receiver:  "johnsmith",
			Content:   "Do you like the curtains?",
			Delivered: time.Time{},
			IsRead:    true,
		}, {
			Id:        1,
			Sender:    "johnsmith",
			Receiver:  "janesmith",
			Content:   "No, they're ugly",
			Delivered: time.Time{},
			IsRead:    true,
		}}

		if !slices.Equal(messages, expected) {
			t.Fatalf("Slices weren't a match, actual: %v", utils.PrettyString(messages))
		}
	}
}

func TestGetMessagesMalformedRequest(t *testing.T) {
	key, err := stores.GetUserStore().CreateSession("johnsmith")
	utils.HandleTestingError(t, err)

	// Missing contact
	endpointJson := `{"key":"` + key + `"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(endpointJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB, mockUsers)
	handler := NewInboxHandler(provider)

	if assert.NoError(t, handler.GetMessages(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetMessagesNoMesssages(t *testing.T) {
	key, err := stores.GetUserStore().CreateSession("johnsmith")
	utils.HandleTestingError(t, err)

	endpointJson := `{"key":"` + key + `","contact":"janesmith"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(endpointJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(make(map[string][]db.Message), mockUsers)
	handler := NewInboxHandler(provider)

	if assert.NoError(t, handler.GetMessages(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

func TestGetUnreadMessagesPassing(t *testing.T) {
	key, err := stores.GetUserStore().CreateSession("johnsmith")
	utils.HandleTestingError(t, err)

	endpointJson := `{"key":"` + key + `"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(endpointJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB, mockUsers)
	handler := NewInboxHandler(provider)

	if assert.NoError(t, handler.GetUnreadMessages(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		messages := make([]db.Message, 0)
		json.Unmarshal(rec.Body.Bytes(), &messages)

		assert.Equal(t, messages[0].Receiver, "johnsmith")
		assert.Equal(t, messages[0].Sender, "janesmith")
		assert.Equal(t, messages[0].Content, "Too bad, we're keeping them")
		assert.Less(t, messages[0].Delivered, time.Now())
		assert.True(t, !messages[0].IsRead)
	}
}

func TestGetUnreadMessagesNoMessages(t *testing.T) {
	key, err := stores.GetUserStore().CreateSession("johnsmith")
	utils.HandleTestingError(t, err)

	endpointJson := `{"key":"` + key + `","contact":"janesmith"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(endpointJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(make(map[string][]db.Message), mockUsers)
	handler := NewInboxHandler(provider)

	if assert.NoError(t, handler.GetMessages(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
