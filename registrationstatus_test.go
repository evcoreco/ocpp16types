package ocpp16types

import (
	"testing"
)

const (
	regStatusAcceptedStr  = "Accepted"
	regStatusPendingStr   = "Pending"
	regStatusRejectedStr  = "Rejected"
	regStatusMethodString = "RegistrationStatus.String()"
)

func TestRegistrationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !RegistrationStatusAccepted.IsValid() {
		t.Errorf(ErrorIsValidFalse, "RegistrationStatusAccepted")
	}
}

func TestRegistrationStatus_IsValid_Pending(t *testing.T) {
	t.Parallel()

	if !RegistrationStatusPending.IsValid() {
		t.Errorf(ErrorIsValidFalse, "RegistrationStatusPending")
	}
}

func TestRegistrationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !RegistrationStatusRejected.IsValid() {
		t.Errorf(ErrorIsValidFalse, "RegistrationStatusRejected")
	}
}

func TestRegistrationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := RegistrationStatus("")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "RegistrationStatus(\"\")")
	}
}

func TestRegistrationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := RegistrationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "RegistrationStatus(\"Unknown\")")
	}
}

func TestRegistrationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := RegistrationStatus("accepted")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "RegistrationStatus(\"accepted\")")
	}
}

func TestRegistrationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := RegistrationStatusAccepted.String()
	if got != regStatusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusAcceptedStr,
		)
	}
}

func TestRegistrationStatus_String_Pending(t *testing.T) {
	t.Parallel()

	got := RegistrationStatusPending.String()
	if got != regStatusPendingStr {
		t.Errorf(
			ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusPendingStr,
		)
	}
}

func TestRegistrationStatus_String_Rejected(t *testing.T) {
	t.Parallel()

	got := RegistrationStatusRejected.String()
	if got != regStatusRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			regStatusMethodString,
			got,
			regStatusRejectedStr,
		)
	}
}
