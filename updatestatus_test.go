package ocpp16types

import (
	"testing"
)

func TestUpdateStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !UpdateStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateStatusAccepted",
		)
	}
}

func TestUpdateStatus_IsValid_Failed(t *testing.T) {
	t.Parallel()

	if !UpdateStatusFailed.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateStatusFailed",
		)
	}
}

func TestUpdateStatus_IsValid_NotSupported(t *testing.T) {
	t.Parallel()

	if !UpdateStatusNotSupported.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateStatusNotSupported",
		)
	}
}

func TestUpdateStatus_IsValid_VersionMismatch(t *testing.T) {
	t.Parallel()

	if !UpdateStatusVersionMismatch.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateStatusVersionMismatch",
		)
	}
}

func TestUpdateStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := UpdateStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UpdateStatus(\"\")",
		)
	}
}

func TestUpdateStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := UpdateStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UpdateStatus(\"Unknown\")",
		)
	}
}

func TestUpdateStatus_String(t *testing.T) {
	t.Parallel()

	got := UpdateStatusAccepted.String()

	want := string(UpdateStatusAccepted)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"UpdateStatus.String()",
			got,
			want,
		)
	}
}
