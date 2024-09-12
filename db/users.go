package db

import (
	"github.com/danielronalds/messenger-server/utils/security"
)

// This file contains the db logics concerning the Users table

// This struct represents a User object from the DB.
//
// NOTE: The password field is not included as this struct should not be used in password queries
type User struct {
	Id          int
	UserName    string
	DisplayName string
}

func (pg Postgres) GetUsers() ([]User, error) {
	users := []User{}

	err := pg.connection.Select(&users, "SELECT Id, UserName, DisplayName FROM api.Users")

	return users, err
}

func (pg Postgres) GetUserWithPass(id int, password string) (User, error) {
	hasher := security.DefaultHash()

	salt, err := pg.getSalt(id)

	if err != nil {
		return User{}, err
	}

	hashedPass, err := hasher.GenerateHash([]byte(password), salt)

	if err != nil {
		return User{}, err
	}

	user := User{}

	query := `SELECT Id, UserName, DisplayName FROM api.Users WHERE Id = $1 AND Password = $2`

	err = pg.connection.Get(&user, query, id, hashedPass.Hash())

	return user, err
}

func (pg Postgres) CreateUser(username, displayName string, hashedPassword, salt []byte) (User, error) {
	newUser := User{}

	query := `INSERT INTO api.Users (UserName, DisplayName, Password, Salt) VALUES ($1, $2, $3, $4) RETURNING Id, UserName, DisplayName`

	err := pg.connection.Get(&newUser, query, username, displayName, hashedPassword, salt)

	return newUser, err
}

func (pg Postgres) getSalt(id int) ([]byte, error) {
	salt := []byte{}

	query := `SELECT Salt FROM api.Users WHERE Id = $1`

	err := pg.connection.Get(&salt, query, id)

	return salt, err
}

func (pg Postgres) DeleteUser(id int, password string) (int64, error) {
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
}
