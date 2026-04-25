package ocpp16types

import (
	"testing"
)

const (
	availStatusAcceptedStr  = "Accepted"
	availStatusRejectedStr  = "Rejected"
	availStatusScheduledStr = "Scheduled"
	availStatusMethodString = "AvailabilityStatus.String()"
)

func TestAvailabilityStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !AvailabilityStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"AvailabilityStatusAccepted",
		)
	}
}

func TestAvailabilityStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !AvailabilityStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"AvailabilityStatusRejected",
		)
	}
}

func TestAvailabilityStatus_IsValid_Scheduled(
	t *testing.T,
) {
	t.Parallel()

	if !AvailabilityStatusScheduled.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"AvailabilityStatusScheduled",
		)
	}
}

func TestAvailabilityStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	status := AvailabilityStatus("")
	if status.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityStatus(\"\")",
		)
	}
}

func TestAvailabilityStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	status := AvailabilityStatus("Unknown")
	if status.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityStatus(\"Unknown\")",
		)
	}
}

func TestAvailabilityStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	status := AvailabilityStatus("accepted")
	if status.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityStatus(\"accepted\")",
		)
	}
}

func TestAvailabilityStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := AvailabilityStatusAccepted.String()
	if got != availStatusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			availStatusMethodString,
			got,
			availStatusAcceptedStr,
		)
	}
}

func TestAvailabilityStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := AvailabilityStatusRejected.String()
	if got != availStatusRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			availStatusMethodString,
			got,
			availStatusRejectedStr,
		)
	}
}

func TestAvailabilityStatus_String_Scheduled(
	t *testing.T,
) {
	t.Parallel()

	got := AvailabilityStatusScheduled.String()
	if got != availStatusScheduledStr {
		t.Errorf(
			ErrorMethodMismatch,
			availStatusMethodString,
			got,
			availStatusScheduledStr,
		)
	}
}
