package ocpp16types

import (
	"testing"
)

const (
	phaseL1Str        = "L1"
	phaseL2Str        = "L2"
	phaseL3Str        = "L3"
	phaseNStr         = "N"
	phaseL1L2Str      = "L1-L2"
	phaseL2L3Str      = "L2-L3"
	phaseL3L1Str      = "L3-L1"
	phaseMethodString = "Phase.String()"
)

func TestPhase_IsValid_L1(t *testing.T) {
	t.Parallel()

	if !PhaseL1.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL1")
	}
}

func TestPhase_IsValid_L2(t *testing.T) {
	t.Parallel()

	if !PhaseL2.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL2")
	}
}

func TestPhase_IsValid_L3(t *testing.T) {
	t.Parallel()

	if !PhaseL3.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL3")
	}
}

func TestPhase_IsValid_N(t *testing.T) {
	t.Parallel()

	if !PhaseN.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseN")
	}
}

func TestPhase_IsValid_L1L2(t *testing.T) {
	t.Parallel()

	if !PhaseL1L2.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL1L2")
	}
}

func TestPhase_IsValid_L2L3(t *testing.T) {
	t.Parallel()

	if !PhaseL2L3.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL2L3")
	}
}

func TestPhase_IsValid_L3L1(t *testing.T) {
	t.Parallel()

	if !PhaseL3L1.IsValid() {
		t.Errorf(ErrorIsValidFalse, "PhaseL3L1")
	}
}

func TestPhase_IsValid_Empty(t *testing.T) {
	t.Parallel()

	phase := Phase("")
	if phase.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Phase(\"\")")
	}
}

func TestPhase_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	phase := Phase("Unknown")
	if phase.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Phase(\"Unknown\")")
	}
}

func TestPhase_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	phase := Phase("single")
	if phase.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Phase(\"single\")")
	}
}

func TestPhase_String_L1(t *testing.T) {
	t.Parallel()

	got := PhaseL1.String()
	if got != phaseL1Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL1Str,
		)
	}
}

func TestPhase_String_L2(t *testing.T) {
	t.Parallel()

	got := PhaseL2.String()
	if got != phaseL2Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL2Str,
		)
	}
}

func TestPhase_String_L3(t *testing.T) {
	t.Parallel()

	got := PhaseL3.String()
	if got != phaseL3Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL3Str,
		)
	}
}

func TestPhase_String_N(t *testing.T) {
	t.Parallel()

	got := PhaseN.String()
	if got != phaseNStr {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseNStr,
		)
	}
}

func TestPhase_String_L1L2(t *testing.T) {
	t.Parallel()

	got := PhaseL1L2.String()
	if got != phaseL1L2Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL1L2Str,
		)
	}
}

func TestPhase_String_L2L3(t *testing.T) {
	t.Parallel()

	got := PhaseL2L3.String()
	if got != phaseL2L3Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL2L3Str,
		)
	}
}

func TestPhase_String_L3L1(t *testing.T) {
	t.Parallel()

	got := PhaseL3L1.String()
	if got != phaseL3L1Str {
		t.Errorf(
			ErrorMethodMismatch,
			phaseMethodString,
			got,
			phaseL3L1Str,
		)
	}
}
