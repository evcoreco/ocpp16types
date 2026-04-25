package ocpp16types

import (
	"testing"
)

func TestErrEmptyValue_NotNil(t *testing.T) {
	t.Parallel()

	if ErrEmptyValue == nil {
		t.Fatal("ErrEmptyValue should not be nil")
	}
}

func TestErrInvalidValue_NotNil(t *testing.T) {
	t.Parallel()

	if ErrInvalidValue == nil {
		t.Fatal("ErrInvalidValue should not be nil")
	}
}

func TestErrEmptyValue_Message(t *testing.T) {
	t.Parallel()

	expected := "value cannot be empty"
	if ErrEmptyValue.Error() != expected {
		t.Errorf("error message mismatch: want %q, got %q",
			expected, ErrEmptyValue.Error())
	}
}

func TestErrInvalidValue_Message(t *testing.T) {
	t.Parallel()

	expected := "invalid value"
	if ErrInvalidValue.Error() != expected {
		t.Errorf("error message mismatch: want %q, got %q",
			expected, ErrInvalidValue.Error())
	}
}
