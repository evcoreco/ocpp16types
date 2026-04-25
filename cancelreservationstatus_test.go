package ocpp16types

import (
	"testing"
)

const (
	cancelResAcceptedStr  = "Accepted"
	cancelResRejectedStr  = "Rejected"
	cancelResMethodString = "CancelReservationStatus.String()"
)

func TestCancelResStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !CancelReservationStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"CancelReservationStatusAccepted",
		)
	}
}

func TestCancelResStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !CancelReservationStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"CancelReservationStatusRejected",
		)
	}
}

func TestCancelResStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := CancelReservationStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"CancelReservationStatus(\"\")",
		)
	}
}

func TestCancelResStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := CancelReservationStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"CancelReservationStatus(\"Unknown\")",
		)
	}
}

func TestCancelResStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := CancelReservationStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"CancelReservationStatus(\"accepted\")",
		)
	}
}

func TestCancelResStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := CancelReservationStatusAccepted.String()
	if got != cancelResAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			cancelResMethodString,
			got,
			cancelResAcceptedStr,
		)
	}
}

func TestCancelResStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := CancelReservationStatusRejected.String()
	if got != cancelResRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			cancelResMethodString,
			got,
			cancelResRejectedStr,
		)
	}
}
