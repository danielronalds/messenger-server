package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructReadMessagesQuery(t *testing.T) {
	ids := []int{1, 2, 3, 4}

	query := constructReadMessagesQuery(ids)

	expected := "UPDATE api.Messages SET IsRead = TRUE WHERE Id = 1 AND Id = 2 AND Id = 3 AND Id = 4 RETURNING *;"

	assert.Equal(t, expected, query)
}

func TestConstructReadMessagesQueryOneId(t *testing.T) {
	ids := []int{1}

	query := constructReadMessagesQuery(ids)

	expected := "UPDATE api.Messages SET IsRead = TRUE WHERE Id = 1 RETURNING *;"

	assert.Equal(t, expected, query)
}
