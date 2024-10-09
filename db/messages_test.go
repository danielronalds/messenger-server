package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructReadMessagesQuery(t *testing.T) {
	ids := []int{1, 2, 3, 4}

	query := constructReadMessagesQuery(ids)

	expected := "UPDATE api.Messages SET IsRead = TRUE WHERE Id = 1 OR Id = 2 OR Id = 3 OR Id = 4 RETURNING *;"

	assert.Equal(t, expected, query)
}

func TestConstructReadMessagesQueryOneId(t *testing.T) {
	ids := []int{1}

	query := constructReadMessagesQuery(ids)

	expected := "UPDATE api.Messages SET IsRead = TRUE WHERE Id = 1 RETURNING *;"

	assert.Equal(t, expected, query)
}
