package message

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/labstack/echo/v4"
)

type (
	// The endpoint handler for messaging
	MessageHandler struct {
		db db.MessageProvider
	}

	PostedMessage struct {
		Key    string `json:"key"`
		To      string `json:"to"`
		Content string `json:"content"`
	}
)

func (m PostedMessage) IsValid() bool {
	trimmedTo := strings.TrimSpace(m.To)
	trimmedContent := strings.TrimSpace(m.Content)

	return len(trimmedTo) > 0 && len(trimmedContent) > 0
}

func NewMessageHandler(db db.MessageProvider) MessageHandler {
	return MessageHandler{db}
}

func (h MessageHandler) SendMessage(c echo.Context) error {
	postedMessage := new(PostedMessage)

	if err := c.Bind(&postedMessage); err != nil {
		log.Printf("Failed to bind message details: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	if !postedMessage.IsValid() {
		return c.String(http.StatusBadRequest, "Malformed request")
	}

	session := stores.GetUserStore().GetSession(postedMessage.Key)
	if session == nil {
		return c.String(http.StatusUnauthorized, "Invalid key");
	}

	if session.Username == postedMessage.To {
		return c.String(http.StatusBadRequest, "Cannot send message to self")
	}

	message, err := h.db.SendMessage(session.Username, postedMessage.To, postedMessage.Content)

	// Assuming that if an error happens from the DB call, then to username key was not acceptable
	if err != nil {
		log.Printf("Failed to send message: %v", err.Error())
		return c.String(http.StatusConflict, "Cannot find receiver")
	}

	return c.JSON(http.StatusCreated, message)
}
