package ocpp16types

import (
	"testing"
)

func TestTriggerMsgStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !TriggerMessageStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"TriggerMessageStatusAccepted",
		)
	}
}

func TestTriggerMsgStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !TriggerMessageStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"TriggerMessageStatusRejected",
		)
	}
}

func TestTriggerMsgStatus_IsValid_NotImpl(t *testing.T) {
	t.Parallel()

	v := TriggerMessageStatusNotImplemented
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"TriggerMessageStatusNotImplemented",
		)
	}
}

func TestTriggerMsgStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := TriggerMessageStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"TriggerMessageStatus(\"\")",
		)
	}
}

func TestTriggerMsgStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := TriggerMessageStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"TriggerMessageStatus(\"Unknown\")",
		)
	}
}

func TestTriggerMsgStatus_String(t *testing.T) {
	t.Parallel()

	got := TriggerMessageStatusAccepted.String()

	want := string(TriggerMessageStatusAccepted)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"TriggerMessageStatus.String()",
			got,
			want,
		)
	}
}
