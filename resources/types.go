package resources

import "strings"

// A struct to represent the JSON posted to the Login endpoint
type PostedUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// A method to check whether the user object recieved in the POST request is valid
func (u PostedUser) IsValid() bool {
	trimmedUserName := strings.TrimSpace(u.UserName)
	trimmedPassword := strings.TrimSpace(u.Password)

	return len(trimmedUserName) > 0 && len(trimmedPassword) > 0
}

// A struct to represent the JSON posted to the CreateUser endpoint
type PostedNewUser struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayname"`
	Password    string `json:"password"`
}

// A method to check whether the user object recieved in the POST request is valid
func (u PostedNewUser) IsValid() bool {
	trimmedUserName := strings.TrimSpace(u.UserName)
	trimmedDisplayName := strings.TrimSpace(u.DisplayName)
	trimmedPassword := strings.TrimSpace(u.Password)

	return len(trimmedUserName) > 0 && len(trimmedDisplayName) > 0 && len(trimmedPassword) > 0
}
