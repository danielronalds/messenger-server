package stores

import (
	"time"

	"github.com/danielronalds/messenger-server/security"
)

// A struct to represents a user's logged in session
type UserSession struct {
	SessionKey string
	Username   string
	StartTime  time.Time
}

var userStore *UserStore

type UserStore struct {
	// Key is the user's session key, keeps UserSessions
	sessions map[string]UserSession
}

func GetUserStore() *UserStore {
	if userStore != nil {
		return userStore
	}

	userStore = &UserStore{
		sessions: make(map[string]UserSession),
	}

	return userStore
}

func (s UserStore) uniqueSessionKey(key string) bool {
	session := s.sessions[key]
	return session.SessionKey == ""
}

// Creates a session for the given user.
//
// NOTE: This does not communicate with the database, it is expected that the caller will do this
func (s UserStore) CreateSession(username string) (string, error) {
	sessionKey, err := security.GenerateSessionKey()

	if err != nil {
		return "", err
	}

	for !s.uniqueSessionKey(sessionKey) {
		sessionKey, err = security.GenerateSessionKey()
		if err != nil {
			return "", err
		}
	}

	session := UserSession{
		SessionKey: sessionKey,
		Username:   username,
		StartTime:  time.Now(),
	}

	s.sessions[sessionKey] = session

	return sessionKey, nil
}

// Gets the user session associated with the given key, returning nil if the key is invalid
func (s UserStore) GetSession(sessionKey string) *UserSession {
	session := s.sessions[sessionKey]

	if session.SessionKey == "" {
		return nil
	}

	return &session
}

// Removes a sesion from the user store
func (s UserStore) DeleteSession(sessionKey string) {
	delete(s.sessions, sessionKey)
}
