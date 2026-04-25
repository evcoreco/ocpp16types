package ocpp16types

import (
	"testing"
)

func TestResetType_IsValid_Hard(t *testing.T) {
	t.Parallel()

	if !ResetTypeHard.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ResetTypeHard",
		)
	}
}

func TestResetType_IsValid_Soft(t *testing.T) {
	t.Parallel()

	if !ResetTypeSoft.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ResetTypeSoft",
		)
	}
}

func TestResetType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := ResetType("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ResetType(\"\")",
		)
	}
}

func TestResetType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := ResetType("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ResetType(\"Unknown\")",
		)
	}
}

func TestResetType_String(t *testing.T) {
	t.Parallel()

	got := ResetTypeHard.String()

	want := "Hard"
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"ResetType.String()",
			got,
			want,
		)
	}
}
