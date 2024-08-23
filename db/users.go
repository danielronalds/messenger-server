package db

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
