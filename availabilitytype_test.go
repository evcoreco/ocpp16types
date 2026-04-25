package ocpp16types

import (
	"testing"
)

const (
	availTypeInoperativeStr = "Inoperative"
	availTypeOperativeStr   = "Operative"
	availTypeMethodString   = "AvailabilityType.String()"
)

func TestAvailabilityType_IsValid_Inoperative(
	t *testing.T,
) {
	t.Parallel()

	if !AvailabilityTypeInoperative.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"AvailabilityTypeInoperative",
		)
	}
}

func TestAvailabilityType_IsValid_Operative(
	t *testing.T,
) {
	t.Parallel()

	if !AvailabilityTypeOperative.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"AvailabilityTypeOperative",
		)
	}
}

func TestAvailabilityType_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	at := AvailabilityType("")
	if at.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityType(\"\")",
		)
	}
}

func TestAvailabilityType_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	at := AvailabilityType("Unknown")
	if at.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityType(\"Unknown\")",
		)
	}
}

func TestAvailabilityType_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	at := AvailabilityType("inoperative")
	if at.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"AvailabilityType(\"inoperative\")",
		)
	}
}

func TestAvailabilityType_String_Inoperative(
	t *testing.T,
) {
	t.Parallel()

	got := AvailabilityTypeInoperative.String()
	if got != availTypeInoperativeStr {
		t.Errorf(
			ErrorMethodMismatch,
			availTypeMethodString,
			got,
			availTypeInoperativeStr,
		)
	}
}

func TestAvailabilityType_String_Operative(
	t *testing.T,
) {
	t.Parallel()

	got := AvailabilityTypeOperative.String()
	if got != availTypeOperativeStr {
		t.Errorf(
			ErrorMethodMismatch,
			availTypeMethodString,
			got,
			availTypeOperativeStr,
		)
	}
}
