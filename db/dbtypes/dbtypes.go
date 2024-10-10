package dbtypes

import "time"

// This struct represents a User object from the DB.
//
// NOTE: The password field is not included as this struct should not be used in password queries
type User struct {
	UserName    string `json:"username"`
	DisplayName string `json:"displayname"`
}

// This struct represents a Message object from the DB
type Message struct {
	Id        int       `json:"id"`
	Sender    string    `json:"sender"`
	Receiver  string    `json:"receiver"`
	Content   string    `json:"content"`
	Delivered time.Time `json:"delivered"`
	IsRead    bool      `json:"isRead"`
}
