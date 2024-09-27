package dbtypes

// This struct represents a User object from the DB.
//
// NOTE: The password field is not included as this struct should not be used in password queries
type User struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayname"`
}
