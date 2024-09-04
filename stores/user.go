package stores

import (
	"time"

	"github.com/danielronalds/messenger-server/utils/security"
)

// A struct to represents a user's logged in session
type UserSession struct {
	sessionKey string
	userId     int
	userName   string
	startTime  time.Time
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
	return session.sessionKey == ""
}

// Creates a session for the given user.
//
// NOTE: This does not communicate with the database, it is expected that the endpoint will do this
func (s UserStore) CreateSession(userId int, username string) (string, error) {
	sessionKey, err := security.GenerateSessionKey()

	if err != nil {
		return "", err;
	}

	for !s.uniqueSessionKey(sessionKey) {
		sessionKey, err = security.GenerateSessionKey()
		if err != nil {
			return "", err;
		}
	}

	session := UserSession{
		sessionKey: sessionKey,
		userId:     userId,
		userName:   username,
		startTime:  time.Now(),
	}

	s.sessions[sessionKey] = session

	return sessionKey, nil
}
