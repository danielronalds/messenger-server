package user

import "strings"

// A struct to represent the JSON posted to the Login endpoint
type postedUser struct {
	UserName    string
	Password    string
}

// A method to check whether the user object recieved in the POST request is valid
func (u postedUser) isValid() bool {
	trimmedUserName := strings.TrimSpace(u.UserName)
	trimmedPassword := strings.TrimSpace(u.Password)

	return len(trimmedUserName) > 0 && len(trimmedPassword) > 0
}

// A struct to represent the JSON posted to the CreateUser endpoint
type postedNewUser struct {
	UserName    string
	DisplayName string
	Password    string
}

// A method to check whether the user object recieved in the POST request is valid
func (u postedNewUser) isValid() bool {
	trimmedUserName := strings.TrimSpace(u.UserName)
	trimmedDisplayName := strings.TrimSpace(u.DisplayName)
	trimmedPassword := strings.TrimSpace(u.Password)

	return len(trimmedUserName) > 0 && len(trimmedDisplayName) > 0 && len(trimmedPassword) > 0
}
