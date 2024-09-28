package message

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	db "github.com/danielronalds/messenger-server/db/dbtypes"
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
)

func TestSendMessagePassing(t *testing.T) {
	messageJson := `{"to":"johnsmith","from":"janesmith","content":"But I really don't like the curtains"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messageJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB)
	handler := NewMessageHandler(provider)

	if assert.NoError(t, handler.SendMessage(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		message := db.Message{}
		json.Unmarshal(rec.Body.Bytes(), &message)

		assert.Equal(t, message.Receiver, "johnsmith")
		assert.Equal(t, message.Sender, "janesmith")
		assert.Equal(t, message.Content, "But I really don't like the curtains")
		assert.LessOrEqual(t, message.Delivered, time.Now())
		assert.Equal(t, message.IsRead, false)
	}
}

func TestSendMessageMissingField(t *testing.T) {
	messageJson := `{"to":"johnsmith","from":"janesmith"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messageJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB)
	handler := NewMessageHandler(provider)

	if assert.NoError(t, handler.SendMessage(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestSendMessageBlankContent(t *testing.T) {
	messageJson := `{"to":"johnsmith","from":"janesmith","content":""}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messageJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB)
	handler := NewMessageHandler(provider)

	if assert.NoError(t, handler.SendMessage(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

// FIXME: This test currently fails
func TestSendMessageInvalidUser(t *testing.T) {
	messageJson := `{"to":"johnsmith","from":"jonsnow","content":"testing"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(messageJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	provider := utils.NewMockedMessageProvider(mockDB)
	handler := NewMessageHandler(provider)

	if assert.NoError(t, handler.SendMessage(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}
