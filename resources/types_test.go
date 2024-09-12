package resources

import (
	"testing"

	"github.com/danielronalds/messenger-server/utils"
)

var defaultPostedNewUser PostedNewUser = PostedNewUser{
	UserName:    "unicornr14",
	DisplayName: "Unicorn",
	Password:    "p@$$w0rd!",
}

func TestValidNewUserTypeValid(t *testing.T) {
	user := defaultPostedNewUser

	if !user.isValid() {
		t.Fatalf("Struct failed validation when it should've passed\n %v", utils.PrettyString(user))
	}
}

func TestValidNewUserTypeMissingUserName(t *testing.T) {
	user := defaultPostedNewUser
	user.UserName = ""

	if user.isValid() {
		t.Fatalf("Struct passed validation when it should've failed\n %v", utils.PrettyString(user))
	}
}

func TestValidNewUserTypeMissingDisplayName(t *testing.T) {
	user := defaultPostedNewUser
	user.DisplayName = ""

	if user.isValid() {
		t.Fatalf("Struct passed validation when it should've failed\n %v", utils.PrettyString(user))
	}
}

func TestValidNewUserTypeMissingPassword(t *testing.T) {
	user := defaultPostedNewUser
	user.Password = ""

	if user.isValid() {
		t.Fatalf("Struct passed validation when it should've failed\n %v", utils.PrettyString(user))
	}
}

var defaultPostedUser PostedUser = PostedUser{
	UserName: "unicornr14",
	Password: "p@$$w0rd!",
}

func TestValidUserTypeValid(t *testing.T) {
	user := defaultPostedUser

	if !user.isValid() {
		t.Fatalf("Struct failed validation when it should've passed\n %v", utils.PrettyString(user))
	}
}

func TestValidUserTypeMissingUserName(t *testing.T) {
	user := defaultPostedUser
	user.UserName = ""

	if user.isValid() {
		t.Fatalf("Struct passed validation when it should've failed\n %v", utils.PrettyString(user))
	}
}

func TestValidUserTypeMissingPassword(t *testing.T) {
	user := defaultPostedUser
	user.Password = ""

	if user.isValid() {
		t.Fatalf("Struct passed validation when it should've failed\n %v", utils.PrettyString(user))
	}
}
