package utils

import (
	"errors"
	"sort"
	"time"

	db "github.com/danielronalds/messenger-server/db/dbtypes"
)

type MockedUserProvider struct {
	db        map[string]db.User
	passwords map[string]string
}

func NewMockedUserProvider(db map[string]db.User) MockedUserProvider {
	passwords := make(map[string]string, 0)

	for username := range db {
		passwords[username] = "password"
	}

	return MockedUserProvider{db, passwords}
}

func (p MockedUserProvider) GetUsers() ([]db.User, error) {
	users := make([]db.User, 0)

	for _, val := range p.db {
		users = append(users, val)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].UserName > users[j].UserName
	})

	return users, nil
}

func (p MockedUserProvider) GetUserWithPass(username string, password string) (db.User, error) {
	actualPassword := p.passwords[username]

	if actualPassword != password {
		return db.User{}, errors.New("incorrect password!")
	}

	return p.db[username], nil
}

func (p MockedUserProvider) CreateUser(username, displayName string, hashedPassword, salt []byte) (db.User, error) {
	return db.User{UserName: username, DisplayName: displayName}, nil
}

func (p MockedUserProvider) IsUsernameTaken(username string) bool {
	_, ok := p.db[username]

	return ok
}

type MockedMessageProvider struct {
	// Key is the receiver of the message
	db map[string][]db.Message
	users map[string]bool
}

func NewMockedMessageProvider(db map[string][]db.Message, users map[string]bool) MockedMessageProvider {
	return MockedMessageProvider{db,users}
}

func (p MockedMessageProvider) SendMessage(from string, to string, content string) (db.Message, error) {
	_, ok := p.users[to]
	if !ok {
		return db.Message{}, errors.New("Invalid user")
	}

	return db.Message{
		Id:        1,
		Sender:    from,
		Receiver:  to,
		Content:   content,
		Delivered: time.Time{},
		IsRead:    false,
	}, nil
}

func (p MockedMessageProvider) GetUnreadMessages(to string) ([]db.Message, error) {
	unreadMessages := make([]db.Message, 0)

	for receiver, messages := range p.db {
		if receiver != to {
			continue
		}

		for _, message := range messages {
			if !message.IsRead {
				unreadMessages = append(unreadMessages, message)
			}
		}
	}

	return unreadMessages, nil
}
