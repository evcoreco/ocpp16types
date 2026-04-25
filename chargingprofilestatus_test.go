package ocpp16types

import (
	"testing"
)

const (
	profStatusAcceptedStr     = "Accepted"
	profStatusRejectedStr     = "Rejected"
	profStatusNotSupportedStr = "NotSupported"
	profStatusMethodString    = "ChargingProfileStatus.String()"
)

func TestProfileStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileStatusAccepted",
		)
	}
}

func TestProfileStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileStatusRejected",
		)
	}
}

func TestProfileStatus_IsValid_NotSupported(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileStatusNotSupported.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileStatusNotSupported",
		)
	}
}

func TestProfileStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := ChargingProfileStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileStatus(\"\")",
		)
	}
}

func TestProfileStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := ChargingProfileStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileStatus(\"Unknown\")",
		)
	}
}

func TestProfileStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := ChargingProfileStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileStatus(\"accepted\")",
		)
	}
}

func TestProfileStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileStatusAccepted.String()
	if got != profStatusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			profStatusMethodString,
			got,
			profStatusAcceptedStr,
		)
	}
}

func TestProfileStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileStatusRejected.String()
	if got != profStatusRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			profStatusMethodString,
			got,
			profStatusRejectedStr,
		)
	}
}

func TestProfileStatus_String_NotSupported(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileStatusNotSupported.String()
	if got != profStatusNotSupportedStr {
		t.Errorf(
			ErrorMethodMismatch,
			profStatusMethodString,
			got,
			profStatusNotSupportedStr,
		)
	}
}
