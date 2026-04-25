package ocpp16types

import (
	"testing"
)

const (
	statusAvailableStr     = "Available"
	statusPreparingStr     = "Preparing"
	statusChargingStr      = "Charging"
	statusSuspendedEVStr   = "SuspendedEV"
	statusSuspendedEVSEStr = "SuspendedEVSE"
	statusFinishingStr     = "Finishing"
	statusReservedStr      = "Reserved"
	statusUnavailableStr   = "Unavailable"
	statusFaultedStr       = "Faulted"
	cpStatusMethodString   = "ChargePointStatus.String()"
)

func TestChargePointStatus_IsValid_Available(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusAvailable.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusAvailable")
	}
}

func TestChargePointStatus_IsValid_Preparing(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusPreparing.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusPreparing")
	}
}

func TestChargePointStatus_IsValid_Charging(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusCharging.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusCharging")
	}
}

func TestChargePointStatus_IsValid_SuspendedEV(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusSuspendedEV.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusSuspendedEV")
	}
}

func TestChargePointStatus_IsValid_SuspendedEVSE(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusSuspendedEVSE.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusSuspendedEVSE")
	}
}

func TestChargePointStatus_IsValid_Finishing(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusFinishing.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusFinishing")
	}
}

func TestChargePointStatus_IsValid_Reserved(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusReserved.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusReserved")
	}
}

func TestChargePointStatus_IsValid_Unavailable(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusUnavailable.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusUnavailable")
	}
}

func TestChargePointStatus_IsValid_Faulted(t *testing.T) {
	t.Parallel()

	if !ChargePointStatusFaulted.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointStatusFaulted")
	}
}

func TestChargePointStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := ChargePointStatus("")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointStatus(\"\")")
	}
}

func TestChargePointStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := ChargePointStatus("Unknown")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointStatus(\"Unknown\")")
	}
}

func TestChargePointStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := ChargePointStatus("available")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointStatus(\"available\")")
	}
}

func TestChargePointStatus_String_Available(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusAvailable.String()
	if got != statusAvailableStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusAvailableStr,
		)
	}
}

func TestChargePointStatus_String_Preparing(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusPreparing.String()
	if got != statusPreparingStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusPreparingStr,
		)
	}
}

func TestChargePointStatus_String_Charging(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusCharging.String()
	if got != statusChargingStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusChargingStr,
		)
	}
}

func TestChargePointStatus_String_SuspendedEV(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusSuspendedEV.String()
	if got != statusSuspendedEVStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusSuspendedEVStr,
		)
	}
}

func TestChargePointStatus_String_SuspendedEVSE(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusSuspendedEVSE.String()
	if got != statusSuspendedEVSEStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusSuspendedEVSEStr,
		)
	}
}

func TestChargePointStatus_String_Finishing(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusFinishing.String()
	if got != statusFinishingStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusFinishingStr,
		)
	}
}

func TestChargePointStatus_String_Reserved(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusReserved.String()
	if got != statusReservedStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusReservedStr,
		)
	}
}

func TestChargePointStatus_String_Unavailable(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusUnavailable.String()
	if got != statusUnavailableStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusUnavailableStr,
		)
	}
}

func TestChargePointStatus_String_Faulted(t *testing.T) {
	t.Parallel()

	got := ChargePointStatusFaulted.String()
	if got != statusFaultedStr {
		t.Errorf(
			ErrorMethodMismatch,
			cpStatusMethodString,
			got,
			statusFaultedStr,
		)
	}
}
