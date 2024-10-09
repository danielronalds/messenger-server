package inbox

import (
	"log"
	"net/http"
	"strings"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/db/dbtypes"
	"github.com/danielronalds/messenger-server/resources"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/danielronalds/messenger-server/utils"
	"github.com/danielronalds/slicetools"
	"github.com/labstack/echo/v4"
)

type (
	// The endpoint handler for messaging
	InboxHandler struct {
		db db.MessageProvider
	}

	GetMessagesBody struct {
		Key string `json:"key"`
		// Referred to as contact, as messages can be to or from
		Contact string `json:"contact"`
	}
)

func NewInboxHandler(db db.MessageProvider) InboxHandler {
	return InboxHandler{db}
}

func (h InboxHandler) GetMessages(c echo.Context) error {
	resBody := GetMessagesBody{}

	if err := c.Bind(&resBody); err != nil {
		log.Printf("Failed to bind GetMessageBody: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	session := stores.GetUserStore().GetSession(resBody.Key)
	if session == nil {
		return c.String(http.StatusUnauthorized, "Invalid key")
	}

	if len(strings.TrimSpace(resBody.Contact)) == 0 {
		return c.String(http.StatusBadRequest, "Invalid contact")
	}

	messages, err := h.db.GetMessages(session.Username, resBody.Contact)

	if err != nil {
		log.Printf("Error when fetching messages: %v", err.Error())
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, messages)
}

func (h InboxHandler) GetUnreadMessages(c echo.Context) error {
	resBody := resources.KeyStruct{}

	if err := c.Bind(&resBody); err != nil {
		log.Printf("Failed to retrieve key: %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	session := stores.GetUserStore().GetSession(resBody.Key)
	if session == nil {
		return c.String(http.StatusUnauthorized, "Invalid key")
	}

	unreadMessages, err := h.db.GetUnreadMessages(session.Username)

	if err != nil {
		log.Printf("Error when fetching unread messages: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, "An error occured reading messages")
	}

	if len(unreadMessages) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	messageIds := slicetools.Map(unreadMessages, func(m dbtypes.Message) int {
		return m.Id
	})

	messages, err := h.db.ReadMessages(messageIds)
	log.Printf("Read messages: %v", utils.PrettyString(messages))

	if err != nil {
		log.Printf("An error occured attempting to read messages: %v", err.Error())
		return c.JSON(http.StatusInternalServerError, "An error occured reading messages")
	}

	return c.JSON(http.StatusOK, unreadMessages)
}
