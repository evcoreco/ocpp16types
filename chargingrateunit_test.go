package ocpp16types

import (
	"testing"
)

const (
	unitWattsStr         = "W"
	unitAmperesStr       = "A"
	rateUnitMethodString = "ChargingRateUnit.String()"
)

func TestChargingRateUnit_IsValid_Watts(t *testing.T) {
	t.Parallel()

	if !ChargingRateUnitWatts.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargingRateUnitWatts")
	}
}

func TestChargingRateUnit_IsValid_Amperes(t *testing.T) {
	t.Parallel()

	if !ChargingRateUnitAmperes.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargingRateUnitAmperes")
	}
}

func TestChargingRateUnit_IsValid_Empty(t *testing.T) {
	t.Parallel()

	unit := ChargingRateUnit("")
	if unit.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargingRateUnit(\"\")")
	}
}

func TestChargingRateUnit_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	unit := ChargingRateUnit("Unknown")
	if unit.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargingRateUnit(\"Unknown\")")
	}
}

func TestChargingRateUnit_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	unit := ChargingRateUnit("w")
	if unit.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargingRateUnit(\"w\")")
	}
}

func TestChargingRateUnit_String_Watts(t *testing.T) {
	t.Parallel()

	got := ChargingRateUnitWatts.String()
	if got != unitWattsStr {
		t.Errorf(
			ErrorMethodMismatch,
			rateUnitMethodString,
			got,
			unitWattsStr,
		)
	}
}

func TestChargingRateUnit_String_Amperes(t *testing.T) {
	t.Parallel()

	got := ChargingRateUnitAmperes.String()
	if got != unitAmperesStr {
		t.Errorf(
			ErrorMethodMismatch,
			rateUnitMethodString,
			got,
			unitAmperesStr,
		)
	}
}
