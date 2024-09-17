package db

import (
	"github.com/danielronalds/messenger-server/utils/security"
)

// This file contains the db logics concerning the Users table

// This struct represents a User object from the DB.
//
// NOTE: The password field is not included as this struct should not be used in password queries
type User struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayname"`
}

func (pg Postgres) GetUsers() ([]User, error) {
	users := []User{}

	err := pg.connection.Select(&users, "SELECT UserName, DisplayName FROM api.Users")

	return users, err
}

func (pg Postgres) GetUserWithPass(username string, password string) (User, error) {
	hasher := security.DefaultHash()

	salt, err := pg.getSalt(username)

	if err != nil {
		return User{}, err
	}

	hashedPass, err := hasher.GenerateHash([]byte(password), salt)

	if err != nil {
		return User{}, err
	}

	user := User{}

	query := `SELECT UserName, DisplayName FROM api.Users WHERE Username = $1 AND Password = $2`

	err = pg.connection.Get(&user, query, username, hashedPass.Hash())

	return user, err
}

func (pg Postgres) CreateUser(username, displayName string, hashedPassword, salt []byte) (User, error) {
	newUser := User{}

	query := `INSERT INTO api.Users (UserName, DisplayName, Password, Salt) VALUES ($1, $2, $3, $4) RETURNING UserName, DisplayName`

	err := pg.connection.Get(&newUser, query, username, displayName, hashedPassword, salt)

	return newUser, err
}

func (pg Postgres) getSalt(username string) ([]byte, error) {
	salt := []byte{}

	query := `SELECT Salt FROM api.Users WHERE username = $1`

	err := pg.connection.Get(&salt, query, username)

	return salt, err
}

/* func (pg Postgres) DeleteUser(id int, password string) (int64, error) {
	hasher := security.DefaultHash()

	salt, err := pg.getSalt(id)

	if err != nil {
		return 0, err
	}

	hashedPass, err := hasher.GenerateHash([]byte(password), salt)

	if err != nil {
		return 0, err
	}

	query := `DELETE FROM api.Users WHERE Id = $1 AND Password = $2`

	result, err := pg.connection.Exec(query, id, hashedPass.Hash())

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
} */
