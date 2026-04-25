package ocpp16types

import (
	"errors"
	"testing"
)

const (
	testIntValue        = 100
	testIntZero         = 0
	testIntMax          = 65535
	testIntOverflow     = 65536
	errValueMismatchInt = "value mismatch: want %d, got %d"
)

func TestNewInteger(t *testing.T) {
	t.Parallel()

	input := testIntValue

	intVal, err := NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if intVal.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, intVal.Value())
	}
}

func TestNewInteger_Zero(t *testing.T) {
	t.Parallel()

	input := testIntZero

	intVal, err := NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if intVal.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, intVal.Value())
	}
}

func TestNewInteger_Max(t *testing.T) {
	t.Parallel()

	input := testIntMax

	intVal, err := NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if intVal.Value() != uint16(input) {
		t.Errorf(errValueMismatchInt, input, intVal.Value())
	}
}

func TestNewInteger_Negative(t *testing.T) {
	t.Parallel()

	_, err := NewInteger(-1)
	if err == nil {
		t.Fatal("negative integer should return error, got nil")
	}

	if !errors.Is(err, ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestNewInteger_Overflow(t *testing.T) {
	t.Parallel()

	_, err := NewInteger(testIntOverflow)
	if err == nil {
		t.Fatal("integer overflow should return error, got nil")
	}

	if !errors.Is(err, ErrInvalidValue) {
		t.Errorf("expected ErrInvalidValue, got: %v", err)
	}
}

func TestInteger_String(t *testing.T) {
	t.Parallel()

	input := testIntValue

	intVal, err := NewInteger(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "100"
	output := intVal.String()

	if output != expected {
		t.Errorf("value mismatch: want %q, got %q", expected, output)
	}
}
