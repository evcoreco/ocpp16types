package ocpp16types

import (
	"testing"
)

func TestUnlockStatus_IsValid_Unlocked(t *testing.T) {
	t.Parallel()

	if !UnlockStatusUnlocked.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UnlockStatusUnlocked",
		)
	}
}

func TestUnlockStatus_IsValid_UnlockFailed(t *testing.T) {
	t.Parallel()

	if !UnlockStatusUnlockFailed.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UnlockStatusUnlockFailed",
		)
	}
}

func TestUnlockStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	if !UnlockStatusNotSupported.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UnlockStatusNotSupported",
		)
	}
}

func TestUnlockStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := UnlockStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UnlockStatus(\"\")",
		)
	}
}

func TestUnlockStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := UnlockStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UnlockStatus(\"Unknown\")",
		)
	}
}

func TestUnlockStatus_String(t *testing.T) {
	t.Parallel()

	got := UnlockStatusUnlocked.String()

	want := string(UnlockStatusUnlocked)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"UnlockStatus.String()",
			got,
			want,
		)
	}
}
