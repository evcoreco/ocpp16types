package ocpp16types

import (
	"testing"
)

func TestRecurrencyKind_IsValid_Daily(t *testing.T) {
	t.Parallel()

	if !RecurrencyKindDaily.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RecurrencyKindDaily",
		)
	}
}

func TestRecurrencyKind_IsValid_Weekly(t *testing.T) {
	t.Parallel()

	if !RecurrencyKindWeekly.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RecurrencyKindWeekly",
		)
	}
}

func TestRecurrencyKind_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := RecurrencyKindType("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RecurrencyKindType(\"\")",
		)
	}
}

func TestRecurrencyKind_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := RecurrencyKindType("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RecurrencyKindType(\"Unknown\")",
		)
	}
}

func TestRecurrencyKind_String(t *testing.T) {
	t.Parallel()

	got := RecurrencyKindDaily.String()
	want := "Daily"

	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"RecurrencyKindType.String()",
			got,
			want,
		)
	}
}
