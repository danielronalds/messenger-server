package security

import "testing"

func handleError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("An error occured: %v", err.Error())
	}
}

func TestPasswordHashingRightPassword(t *testing.T) {
	password := "P@ssw0rd!"

	hasher := DefaultHash()

	hashedPass, err := hasher.GenerateNewHash([]byte(password))

	handleError(t, err)

	correctPass, err := hasher.Compare(hashedPass.hash, hashedPass.salt, []byte(password))

	handleError(t, err)

	if !correctPass {
		t.Fatalf("The password was marked not correct, when the two passwords were the same")
	}
}

func TestPasswordHashingWrongPassword(t *testing.T) {
	password := "P@ssw0rd!"

	hasher := DefaultHash()

	hashedPass, err := hasher.GenerateNewHash([]byte(password))

	handleError(t, err)

	correctPass, err := hasher.Compare(hashedPass.hash, hashedPass.salt, []byte("WrongPassword"))

	handleError(t, err)

	if correctPass {
		t.Fatalf("The password was marked as correct when the compared passwords were different")
	}
}
