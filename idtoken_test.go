package ocpp16types

import (
	"testing"
)

const (
	testRFIDTag123        = "RFID-TAG-123"
	testRFIDTag456        = "RFID-TAG-456"
	errUnexpectedCiString = "unexpected error creating CiString20Type: %v"
)

func TestNewIDToken(t *testing.T) {
	t.Parallel()

	token, err := NewCiString20Type(testRFIDTag123)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := NewIDToken(token)
	if idToken.String() != testRFIDTag123 {
		t.Errorf(
			ErrorMethodMismatch,
			"IDToken.String()",
			idToken.String(),
			testRFIDTag123,
		)
	}
}

func TestIDToken_Value(t *testing.T) {
	t.Parallel()

	token, err := NewCiString20Type(testRFIDTag456)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := NewIDToken(token)
	retrievedToken := idToken.Value()

	if retrievedToken.Value() != testRFIDTag456 {
		t.Errorf(
			ErrorMethodMismatch,
			"IDToken.Value().Value()",
			retrievedToken.Value(),
			testRFIDTag456,
		)
	}
}

func TestIDToken_String(t *testing.T) {
	t.Parallel()

	input := "RFID-TAG-789"

	token, err := NewCiString20Type(input)
	if err != nil {
		t.Fatalf(errUnexpectedCiString, err)
	}

	idToken := NewIDToken(token)
	if idToken.String() != input {
		t.Errorf(
			ErrorMethodMismatch,
			"IDToken.String()",
			idToken.String(),
			input,
		)
	}
}
