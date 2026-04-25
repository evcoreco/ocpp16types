package ocpp16types

import (
	"testing"
)

const (
	clearCacheAcceptedStr  = "Accepted"
	clearCacheRejectedStr  = "Rejected"
	clearCacheMethodString = "ClearCacheStatus.String()"
)

func TestClearCacheStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !ClearCacheStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ClearCacheStatusAccepted",
		)
	}
}

func TestClearCacheStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !ClearCacheStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ClearCacheStatusRejected",
		)
	}
}

func TestClearCacheStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := ClearCacheStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearCacheStatus(\"\")",
		)
	}
}

func TestClearCacheStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := ClearCacheStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearCacheStatus(\"Unknown\")",
		)
	}
}

func TestClearCacheStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := ClearCacheStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ClearCacheStatus(\"accepted\")",
		)
	}
}

func TestClearCacheStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := ClearCacheStatusAccepted.String()
	if got != clearCacheAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			clearCacheMethodString,
			got,
			clearCacheAcceptedStr,
		)
	}
}

func TestClearCacheStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := ClearCacheStatusRejected.String()
	if got != clearCacheRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			clearCacheMethodString,
			got,
			clearCacheRejectedStr,
		)
	}
}
