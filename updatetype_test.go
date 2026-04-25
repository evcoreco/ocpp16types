package ocpp16types

import (
	"testing"
)

func TestUpdateType_IsValid_Full(t *testing.T) {
	t.Parallel()

	if !UpdateTypeFull.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateTypeFull",
		)
	}
}

func TestUpdateType_IsValid_Differential(t *testing.T) {
	t.Parallel()

	if !UpdateTypeDifferential.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"UpdateTypeDifferential",
		)
	}
}

func TestUpdateType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := UpdateType("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UpdateType(\"\")",
		)
	}
}

func TestUpdateType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := UpdateType("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"UpdateType(\"Unknown\")",
		)
	}
}

func TestUpdateType_String(t *testing.T) {
	t.Parallel()

	got := UpdateTypeFull.String()

	want := string(UpdateTypeFull)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"UpdateType.String()",
			got,
			want,
		)
	}
}
