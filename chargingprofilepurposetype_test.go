package ocpp16types

import (
	"testing"
)

const (
	purposeChargePointMaxProfileStr = "ChargePointMaxProfile"
	purposeTxDefaultProfileStr      = "TxDefaultProfile"
	purposeTxProfileStr             = "TxProfile"
	purposeTypeMethodString         = "ChargingProfilePurposeType.String()"
)

func TestChargingProfilePurposeType_IsValid_CPMaxProfile(t *testing.T) {
	t.Parallel()

	if !ChargePointMaxProfile.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ChargePointMaxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	if !TxDefaultProfile.IsValid() {
		t.Errorf(ErrorIsValidFalse, "TxDefaultProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_TxProfile(t *testing.T) {
	t.Parallel()

	if !TxProfile.IsValid() {
		t.Errorf(ErrorIsValidFalse, "TxProfile")
	}
}

func TestChargingProfilePurposeType_IsValid_Empty(t *testing.T) {
	t.Parallel()

	purpose := ChargingProfilePurposeType("")
	if purpose.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargingProfilePurposeType(\"\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	purpose := ChargingProfilePurposeType("Unknown")
	if purpose.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargingProfilePurposeType(\"Unknown\")")
	}
}

func TestChargingProfilePurposeType_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	purpose := ChargingProfilePurposeType("chargepointa")
	if purpose.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ChargingProfilePurposeType(\"chargepointa\")",
		)
	}
}

func TestChargingProfilePurposeType_String_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	got := ChargePointMaxProfile.String()
	if got != purposeChargePointMaxProfileStr {
		t.Errorf(
			ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeChargePointMaxProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	got := TxDefaultProfile.String()
	if got != purposeTxDefaultProfileStr {
		t.Errorf(
			ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeTxDefaultProfileStr,
		)
	}
}

func TestChargingProfilePurposeType_String_TxProfile(t *testing.T) {
	t.Parallel()

	got := TxProfile.String()
	if got != purposeTxProfileStr {
		t.Errorf(
			ErrorMethodMismatch,
			purposeTypeMethodString,
			got,
			purposeTxProfileStr,
		)
	}
}
