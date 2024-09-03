package security

import (
	"testing"

	"github.com/danielronalds/messenger-server/utils"
)

// A simple unit test that generates 1000 keys to determine if they're all unique
func TestSessionKeyGenerationIsUnique(t *testing.T) {
	existingKeys := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		key, err := GenerateSessionKey()
		utils.HandleTestingError(t, err)

		if existingKeys[key] {
			t.Fatalf("Key already exists, it should be unique!")
		}

		existingKeys[key] = true
	}
}
