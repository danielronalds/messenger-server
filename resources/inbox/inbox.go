package inbox

import (
	"log"
	"net/http"

	"github.com/danielronalds/messenger-server/db"
	"github.com/danielronalds/messenger-server/db/dbtypes"
	"github.com/danielronalds/messenger-server/resources"
	"github.com/danielronalds/messenger-server/stores"
	"github.com/danielronalds/slicetools"
	"github.com/labstack/echo/v4"
)

type (
	// The endpoint handler for messaging
	InboxHandler struct {
		db db.MessageProvider
	}
)

func NewInboxHandler(db db.MessageProvider) InboxHandler {
	return InboxHandler{db}
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
		return c.JSON(http.StatusNoContent, make([]dbtypes.Message, 0))
	}

	messageIds := slicetools.Map(unreadMessages, func(m dbtypes.Message) int {
		return m.Id
	})

	h.db.ReadMessages(messageIds)

	return c.JSON(http.StatusOK, unreadMessages)
}
