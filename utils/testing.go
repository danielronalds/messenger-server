package utils

import (
	"encoding/json"
	"testing"

	"github.com/danielronalds/messenger-server/db"
)

type MockedUserProvider struct {
	db map[string]db.User
}

func NewMockedUserProvider(db map[string]db.User) MockedUserProvider {
	return MockedUserProvider { db }
}

func (p MockedUserProvider) GetUsers() ([]db.User, error) {
	users := make([]db.User, 0);

	for _, val := range p.db {
		users = append(users, val)
	}

	return users, nil;
}

func (p MockedUserProvider) GetUserWithPass(username string, password string) (db.User, error) {
	return db.User{}, nil;
}

func (p MockedUserProvider) CreateUser(username, displayName string, hashedPassword, salt []byte) (db.User, error) {
	return db.User{}, nil;
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
