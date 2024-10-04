package db

import (
	"errors"
	"fmt"
	"time"

	. "github.com/danielronalds/messenger-server/db/dbtypes"
	"github.com/danielronalds/slicetools"
)

func (pg Postgres) SendMessage(from string, to string, content string) (Message, error) {
	newMessage := Message{}

	query := `INSERT INTO api.Messages (Sender, Receiver, Content, Delivered, IsRead) VALUES ($1, $2, $3, $4, $5) RETURNING *;`

	err := pg.connection.Get(&newMessage, query, from, to, content, time.Now(), false)

	return newMessage, err
}

func (pg Postgres) GetUnreadMessages(to string) ([]Message, error) {
	messages := []Message{}

	err := pg.connection.Select(&messages, "SELECT * FROM api.Messages WHERE Receiver = $1 AND IsRead = FALSE;")

	return messages, err
}

func (pg Postgres) ReadMessages(ids []int) ([]Message, error) {
	messages := []Message{}

	if len(ids) == 0 {
		return messages, errors.New("No messages to read!")
	}

	query := constructReadMessagesQuery(ids)


	err := pg.connection.Select(&messages, query)

	return messages, err
}

// Handles the logic for constructing the read SQL statement
//
// Split for the purpose of testing
func constructReadMessagesQuery(ids []int) string {
	idSelects := slicetools.MapWithIndex(ids, func (i int, id int) string {
		if i == 0 {
			return fmt.Sprintf("Id = %v", id)
		}
		return fmt.Sprintf(" AND Id = %v", id)
	})

	baseQuery := `UPDATE api.Messages SET IsRead = TRUE WHERE `
	query := slicetools.Foldl(idSelects, baseQuery, func(prev string, next string) string {
		return prev + next
	})

	query = query + " RETURNING *;"

	return query
}
