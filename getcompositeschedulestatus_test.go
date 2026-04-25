package ocpp16types

import (
	"testing"
)

func TestGetCompSchedStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !GetCompositeScheduleStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"GetCompositeScheduleStatusAccepted",
		)
	}
}

func TestGetCompSchedStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !GetCompositeScheduleStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"GetCompositeScheduleStatusRejected",
		)
	}
}

func TestGetCompSchedStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := GetCompositeScheduleStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"GetCompositeScheduleStatus(\"\")",
		)
	}
}

func TestGetCompSchedStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := GetCompositeScheduleStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"GetCompositeScheduleStatus(\"Unknown\")",
		)
	}
}

func TestGetCompSchedStatus_String(t *testing.T) {
	t.Parallel()

	got := GetCompositeScheduleStatusAccepted.String()
	want := string(GetCompositeScheduleStatusAccepted)

	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"GetCompositeScheduleStatus.String()",
			got,
			want,
		)
	}
}
