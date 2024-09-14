package stores

import (
	"testing"

	"github.com/danielronalds/messenger-server/utils"
)

func TestUserStoreCreateAndReadSession(t *testing.T) {
	userName := "TestUser"

	store := GetUserStore()

	sessionKey, err := store.CreateSession(userName)
	utils.HandleTestingError(t, err)

	// Fetching a new user store to test singleton implementation
	newStore := GetUserStore()

	session := newStore.GetSession(sessionKey)

	if session == nil {
		t.Fatal("Failed to retrieve created session")
	}

	if session.Username != userName {
		t.Fatalf("Session details did not match\n%v", utils.PrettyString(session))
	}
}
