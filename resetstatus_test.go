package ocpp16types

import (
	"testing"
)

func TestResetStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !ResetStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ResetStatusAccepted",
		)
	}
}

func TestResetStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !ResetStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ResetStatusRejected",
		)
	}
}

func TestResetStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := ResetStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ResetStatus(\"\")",
		)
	}
}

func TestResetStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := ResetStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ResetStatus(\"Unknown\")",
		)
	}
}

func TestResetStatus_String(t *testing.T) {
	t.Parallel()

	got := ResetStatusAccepted.String()

	want := string(ResetStatusAccepted)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"ResetStatus.String()",
			got,
			want,
		)
	}
}
