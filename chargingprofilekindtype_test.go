package ocpp16types

import (
	"testing"
)

const (
	profKindAbsoluteStr  = "Absolute"
	profKindRecurringStr = "Recurring"
	profKindRelativeStr  = "Relative"
	profKindMethodString = "ChargingProfileKindType.String()"
)

func TestProfileKind_IsValid_Absolute(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileKindAbsolute.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileKindAbsolute",
		)
	}
}

func TestProfileKind_IsValid_Recurring(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileKindRecurring.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileKindRecurring",
		)
	}
}

func TestProfileKind_IsValid_Relative(
	t *testing.T,
) {
	t.Parallel()

	if !ChargingProfileKindRelative.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ChargingProfileKindRelative",
		)
	}
}

func TestProfileKind_IsValid_Empty(t *testing.T) {
	t.Parallel()

	k := ChargingProfileKindType("")
	if k.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileKindType(\"\")",
		)
	}
}

func TestProfileKind_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	k := ChargingProfileKindType("Unknown")
	if k.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileKindType(\"Unknown\")",
		)
	}
}

func TestProfileKind_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	k := ChargingProfileKindType("absolute")
	if k.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfileKindType(\"absolute\")",
		)
	}
}

func TestProfileKind_String_Absolute(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileKindAbsolute.String()
	if got != profKindAbsoluteStr {
		t.Errorf(
			ErrorMethodMismatch,
			profKindMethodString,
			got,
			profKindAbsoluteStr,
		)
	}
}

func TestProfileKind_String_Recurring(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileKindRecurring.String()
	if got != profKindRecurringStr {
		t.Errorf(
			ErrorMethodMismatch,
			profKindMethodString,
			got,
			profKindRecurringStr,
		)
	}
}

func TestProfileKind_String_Relative(
	t *testing.T,
) {
	t.Parallel()

	got := ChargingProfileKindRelative.String()
	if got != profKindRelativeStr {
		t.Errorf(
			ErrorMethodMismatch,
			profKindMethodString,
			got,
			profKindRelativeStr,
		)
	}
}
