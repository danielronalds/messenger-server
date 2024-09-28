package db

import (
	"time"

	. "github.com/danielronalds/messenger-server/db/dbtypes"
)

func (pg Postgres) SendMessage(from string, to string, content string) (Message, error) {
	newMessage := Message{}

	query := `INSERT INTO api.Messages (Sender, Receiver, Content, Delivered, IsRead) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	err := pg.connection.Get(&newMessage, query, from, to, content, time.Now(), false)

	return newMessage, err
}

func (pg Postgres) GetUnreadMessages(to string) ([]Message, error) {
	messages := []Message{}

	err := pg.connection.Select(&messages, "SELECT * FROM api.Messages WHERE Receiver = $1 AND IsRead = FALSE")

	return messages, err
}
