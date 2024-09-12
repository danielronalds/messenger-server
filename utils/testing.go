package utils

import (
	"encoding/json"
	"testing"
)

// This function handles an error during testing
func HandleTestingError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("An error occured: %v", err.Error())
	}
}

// This function takes a struct and marshalls it into a "pretty" string
//
// NOTE: if this function fails, you will get an empty string to avoid returning an err
func PrettyString(object any) string {
	json, _ := json.MarshalIndent(object, "", "  ")

	return string(json)
}
