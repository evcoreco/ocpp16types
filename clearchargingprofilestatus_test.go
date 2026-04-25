package ocpp16types

import (
	"testing"
)

const (
	clearProfAcceptedStr  = "Accepted"
	clearProfUnknownStr   = "Unknown"
	clearProfMethodString = "ClearChargingProfileStatus.String()"
)

func TestClearProfStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !ClearChargingProfileStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ClearChargingProfileStatusAccepted",
		)
	}
}

func TestClearProfStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	if !ClearChargingProfileStatusUnknown.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ClearChargingProfileStatusUnknown",
		)
	}
}

func TestClearProfStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := ClearChargingProfileStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearChargingProfileStatus(\"\")",
		)
	}
}

func TestClearProfStatus_IsValid_Invalid(
	t *testing.T,
) {
	t.Parallel()

	s := ClearChargingProfileStatus("Invalid")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearChargingProfileStatus(\"Invalid\")",
		)
	}
}

func TestClearProfStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := ClearChargingProfileStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearChargingProfileStatus(\"accepted\")",
		)
	}
}

func TestClearProfStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := ClearChargingProfileStatusAccepted.String()
	if got != clearProfAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			clearProfMethodString,
			got,
			clearProfAcceptedStr,
		)
	}
}

func TestClearProfStatus_String_Unknown(
	t *testing.T,
) {
	t.Parallel()

	got := ClearChargingProfileStatusUnknown.String()
	if got != clearProfUnknownStr {
		t.Errorf(
			ErrorMethodMismatch,
			clearProfMethodString,
			got,
			clearProfUnknownStr,
		)
	}
}
