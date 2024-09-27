package db

import . "github.com/danielronalds/messenger-server/db/dbtypes"

type UserProvider interface {
	GetUsers() ([]User, error)
	GetUserWithPass(username string, password string) (User, error)
	CreateUser(username, displayName string, hashedPassword, salt []byte) (User, error)
	IsUsernameTaken(username string) bool
}
