package ocpp16types

import (
	"testing"
)

func TestReservationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !ReservationStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ReservationStatusAccepted",
		)
	}
}

func TestReservationStatus_IsValid_Faulted(t *testing.T) {
	t.Parallel()

	if !ReservationStatusFaulted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ReservationStatusFaulted",
		)
	}
}

func TestReservationStatus_IsValid_Occupied(t *testing.T) {
	t.Parallel()

	if !ReservationStatusOccupied.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ReservationStatusOccupied",
		)
	}
}

func TestReservationStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	if !ReservationStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ReservationStatusRejected",
		)
	}
}

func TestReservationStatus_IsValid_Unavailable(t *testing.T) {
	t.Parallel()

	if !ReservationStatusUnavailable.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ReservationStatusUnavailable",
		)
	}
}

func TestReservationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := ReservationStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ReservationStatus(\"\")",
		)
	}
}

func TestReservationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := ReservationStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ReservationStatus(\"Unknown\")",
		)
	}
}

func TestReservationStatus_String(t *testing.T) {
	t.Parallel()

	got := ReservationStatusAccepted.String()
	want := string(ReservationStatusAccepted)

	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"ReservationStatus.String()",
			got,
			want,
		)
	}
}
