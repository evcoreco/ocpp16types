package ocpp16types

import (
	"testing"
)

func TestMsgTrigger_IsValid_BootNotif(t *testing.T) {
	t.Parallel()

	if !MessageTriggerBootNotification.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerBootNotification",
		)
	}
}

func TestMsgTrigger_IsValid_DiagStatus(t *testing.T) {
	t.Parallel()

	v := MessageTriggerDiagnosticsStatusNotification
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerDiagnosticsStatusNotification",
		)
	}
}

func TestMsgTrigger_IsValid_FwStatus(t *testing.T) {
	t.Parallel()

	v := MessageTriggerFirmwareStatusNotification
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerFirmwareStatusNotification",
		)
	}
}

func TestMsgTrigger_IsValid_Heartbeat(t *testing.T) {
	t.Parallel()

	if !MessageTriggerHeartbeat.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerHeartbeat",
		)
	}
}

func TestMsgTrigger_IsValid_MeterValues(t *testing.T) {
	t.Parallel()

	if !MessageTriggerMeterValues.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerMeterValues",
		)
	}
}

func TestMsgTrigger_IsValid_StatusNotif(t *testing.T) {
	t.Parallel()

	if !MessageTriggerStatusNotification.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"MessageTriggerStatusNotification",
		)
	}
}

func TestMsgTrigger_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := MessageTrigger("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"MessageTrigger(\"\")",
		)
	}
}

func TestMsgTrigger_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := MessageTrigger("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"MessageTrigger(\"Unknown\")",
		)
	}
}

func TestMsgTrigger_String(t *testing.T) {
	t.Parallel()

	got := MessageTriggerBootNotification.String()
	want := "BootNotification"

	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"MessageTrigger.String()",
			got,
			want,
		)
	}
}
