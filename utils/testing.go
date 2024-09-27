package utils

import (
	"encoding/json"
	"errors"
	"sort"
	"testing"

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

// This function handles an error during testing
func HandleTestingError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("An error occured: %v", err.Error())
	}
}

// This function takes a struct and marshalls it into a "pretty" string
//
// NOTE: if this function fails, you will get an empty string to avoid returning an err
func PrettyString(object any) string {
	json, _ := json.MarshalIndent(object, "", "  ")

	return string(json)
}
