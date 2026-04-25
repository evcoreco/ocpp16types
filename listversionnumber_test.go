package ocpp16types

import (
	"math"
	"testing"
)

const (
	testListVersion      = 5
	testListVersionZero  = 0
	testListVersionForty = 42
)

func TestNewListVersionNumber_ValidPositive(t *testing.T) {
	t.Parallel()

	lvn, err := NewListVersionNumber(testListVersion)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if lvn.Value() != testListVersion {
		t.Errorf(
			ErrorMethodMismatch,
			"Value",
			lvn.Value(),
			testListVersion,
		)
	}
}

func TestNewListVersionNumber_Unsupported(t *testing.T) {
	t.Parallel()

	lvn, err := NewListVersionNumber(-1)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if !lvn.IsUnsupported() {
		t.Errorf(
			ErrorMethodMismatch,
			"IsUnsupported",
			lvn.IsUnsupported(),
			true,
		)
	}

	if lvn.IsEmpty() {
		t.Errorf(
			ErrorMethodMismatch,
			"IsEmpty",
			lvn.IsEmpty(),
			false,
		)
	}
}

func TestNewListVersionNumber_Empty(t *testing.T) {
	t.Parallel()

	lvn, err := NewListVersionNumber(testListVersionZero)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if !lvn.IsEmpty() {
		t.Errorf(
			ErrorMethodMismatch,
			"IsEmpty",
			lvn.IsEmpty(),
			true,
		)
	}

	if lvn.IsUnsupported() {
		t.Errorf(
			ErrorMethodMismatch,
			"IsUnsupported",
			lvn.IsUnsupported(),
			false,
		)
	}
}

func TestNewListVersionNumber_OutOfRange(t *testing.T) {
	t.Parallel()

	_, err := NewListVersionNumber(math.MaxInt64)
	if err == nil {
		t.Fatalf(ErrorWantNil, "out of range value")
	}
}

func TestListVersionNumber_String(t *testing.T) {
	t.Parallel()

	lvn, err := NewListVersionNumber(testListVersionForty)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if lvn.String() != "42" {
		t.Errorf(
			ErrorMethodMismatch,
			"String",
			lvn.String(),
			"42",
		)
	}
}

func TestListVersionNumber_StringNegative(t *testing.T) {
	t.Parallel()

	lvn, err := NewListVersionNumber(-1)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if lvn.String() != "-1" {
		t.Errorf(
			ErrorMethodMismatch,
			"String",
			lvn.String(),
			"-1",
		)
	}
}
