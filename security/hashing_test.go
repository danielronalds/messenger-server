package security

import (
	"github.com/danielronalds/messenger-server/utils"
	"testing"
)

func TestPasswordHashingRightPassword(t *testing.T) {
	password := "P@ssw0rd!"

	hasher := DefaultHash()

	hashedPass, err := hasher.GenerateNewHash([]byte(password))

	utils.HandleTestingError(t, err)

	correctPass, err := hasher.Compare(hashedPass.hash, hashedPass.salt, []byte(password))

	utils.HandleTestingError(t, err)

	if !correctPass {
		t.Fatalf("The password was marked not correct, when the two passwords were the same")
	}
}

func TestPasswordHashingWrongPassword(t *testing.T) {
	password := "P@ssw0rd!"

	hasher := DefaultHash()

	hashedPass, err := hasher.GenerateNewHash([]byte(password))

	utils.HandleTestingError(t, err)

	correctPass, err := hasher.Compare(hashedPass.hash, hashedPass.salt, []byte("WrongPassword"))

	utils.HandleTestingError(t, err)

	if correctPass {
		t.Fatalf("The password was marked as correct when the compared passwords were different")
	}
}
