package message

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielronalds/messenger-server/db"
	"github.com/labstack/echo/v4"
)

type (
	// The endpoint handler for messaging
	MessageHandler struct {
		db db.MessageProvider
	}

	PostedMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Content string `json:"content"`
	}
)

func (m PostedMessage) IsValid() bool {
	trimmedFrom := strings.TrimSpace(m.From)
	trimmedTo := strings.TrimSpace(m.To)
	trimmedContent := strings.TrimSpace(m.Content)

	return len(trimmedFrom) > 0 && len(trimmedTo) > 0 && len(trimmedContent) > 0
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

	message, err := h.db.SendMessage(postedMessage.From, postedMessage.To, postedMessage.Content)

	if err != nil {
		log.Printf("Failed to send message: %v", err.Error())
		// FIX: This might need to be 409?
		return c.String(http.StatusInternalServerError, "Unable to send message")
	}

	return c.JSON(http.StatusCreated, message)
}
