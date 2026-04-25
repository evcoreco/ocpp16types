package ocpp16types

import (
	"testing"
)

const (
	errCodeNoErrorStr              = "NoError"
	errCodeConnectorLockFailureStr = "ConnectorLockFailure"
	errCodeEVCommunicationErrorStr = "EVCommunicationError"
	errCodeGroundFailureStr        = "GroundFailure"
	errCodeHighTemperatureStr      = "HighTemperature"
	errCodeInternalErrorStr        = "InternalError"
	errCodeLocalListConflictStr    = "LocalListConflict"
	errCodeOtherErrorStr           = "OtherError"
	errCodeOverCurrentFailureStr   = "OverCurrentFailure"
	errCodeOverVoltageStr          = "OverVoltage"
	errCodePowerMeterFailureStr    = "PowerMeterFailure"
	errCodePowerSwitchFailureStr   = "PowerSwitchFailure"
	errCodeReaderFailureStr        = "ReaderFailure"
	errCodeResetFailureStr         = "ResetFailure"
	errCodeUnderVoltageStr         = "UnderVoltage"
	errCodeWeakSignalStr           = "WeakSignal"
	errCodeMethodString            = "ChargePointErrorCode.String()"
)

func TestChargePointErrorCode_IsValid_NoError(t *testing.T) {
	t.Parallel()

	if !ErrCodeNoError.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeNoError")
	}
}

func TestChargePointErrorCode_IsValid_ConnectorLockFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodeConnectorLockFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeConnectorLockFailure")
	}
}

func TestChargePointErrorCode_IsValid_EVCommunicationError(t *testing.T) {
	t.Parallel()

	if !ErrCodeEVCommunicationError.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeEVCommunicationError")
	}
}

func TestChargePointErrorCode_IsValid_GroundFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodeGroundFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeGroundFailure")
	}
}

func TestChargePointErrorCode_IsValid_HighTemperature(t *testing.T) {
	t.Parallel()

	if !ErrCodeHighTemperature.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeHighTemperature")
	}
}

func TestChargePointErrorCode_IsValid_InternalError(t *testing.T) {
	t.Parallel()

	if !ErrCodeInternalError.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeInternalError")
	}
}

func TestChargePointErrorCode_IsValid_LocalListConflict(t *testing.T) {
	t.Parallel()

	if !ErrCodeLocalListConflict.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeLocalListConflict")
	}
}

func TestChargePointErrorCode_IsValid_OtherError(t *testing.T) {
	t.Parallel()

	if !ErrCodeOtherError.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeOtherError")
	}
}

func TestChargePointErrorCode_IsValid_OverCurrentFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodeOverCurrentFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeOverCurrentFailure")
	}
}

func TestChargePointErrorCode_IsValid_OverVoltage(t *testing.T) {
	t.Parallel()

	if !ErrCodeOverVoltage.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeOverVoltage")
	}
}

func TestChargePointErrorCode_IsValid_PowerMeterFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodePowerMeterFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodePowerMeterFailure")
	}
}

func TestChargePointErrorCode_IsValid_PowerSwitchFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodePowerSwitchFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodePowerSwitchFailure")
	}
}

func TestChargePointErrorCode_IsValid_ReaderFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodeReaderFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeReaderFailure")
	}
}

func TestChargePointErrorCode_IsValid_ResetFailure(t *testing.T) {
	t.Parallel()

	if !ErrCodeResetFailure.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeResetFailure")
	}
}

func TestChargePointErrorCode_IsValid_UnderVoltage(t *testing.T) {
	t.Parallel()

	if !ErrCodeUnderVoltage.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeUnderVoltage")
	}
}

func TestChargePointErrorCode_IsValid_WeakSignal(t *testing.T) {
	t.Parallel()

	if !ErrCodeWeakSignal.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ErrCodeWeakSignal")
	}
}

func TestChargePointErrorCode_IsValid_Empty(t *testing.T) {
	t.Parallel()

	errCode := ChargePointErrorCode("")
	if errCode.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointErrorCode(\"\")")
	}
}

func TestChargePointErrorCode_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	errCode := ChargePointErrorCode("Unknown")
	if errCode.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointErrorCode(\"Unknown\")")
	}
}

func TestChargePointErrorCode_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	errCode := ChargePointErrorCode("noerror")
	if errCode.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ChargePointErrorCode(\"noerror\")")
	}
}

func TestChargePointErrorCode_String_NoError(t *testing.T) {
	t.Parallel()

	got := ErrCodeNoError.String()
	if got != errCodeNoErrorStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeNoErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_ConnectorLockFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodeConnectorLockFailure.String()
	if got != errCodeConnectorLockFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeConnectorLockFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_EVCommunicationError(t *testing.T) {
	t.Parallel()

	got := ErrCodeEVCommunicationError.String()
	if got != errCodeEVCommunicationErrorStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeEVCommunicationErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_GroundFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodeGroundFailure.String()
	if got != errCodeGroundFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeGroundFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_HighTemperature(t *testing.T) {
	t.Parallel()

	got := ErrCodeHighTemperature.String()
	if got != errCodeHighTemperatureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeHighTemperatureStr,
		)
	}
}

func TestChargePointErrorCode_String_InternalError(t *testing.T) {
	t.Parallel()

	got := ErrCodeInternalError.String()
	if got != errCodeInternalErrorStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeInternalErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_LocalListConflict(t *testing.T) {
	t.Parallel()

	got := ErrCodeLocalListConflict.String()
	if got != errCodeLocalListConflictStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeLocalListConflictStr,
		)
	}
}

func TestChargePointErrorCode_String_OtherError(t *testing.T) {
	t.Parallel()

	got := ErrCodeOtherError.String()
	if got != errCodeOtherErrorStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOtherErrorStr,
		)
	}
}

func TestChargePointErrorCode_String_OverCurrentFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodeOverCurrentFailure.String()
	if got != errCodeOverCurrentFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOverCurrentFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_OverVoltage(t *testing.T) {
	t.Parallel()

	got := ErrCodeOverVoltage.String()
	if got != errCodeOverVoltageStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeOverVoltageStr,
		)
	}
}

func TestChargePointErrorCode_String_PowerMeterFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodePowerMeterFailure.String()
	if got != errCodePowerMeterFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodePowerMeterFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_PowerSwitchFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodePowerSwitchFailure.String()
	if got != errCodePowerSwitchFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodePowerSwitchFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_ReaderFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodeReaderFailure.String()
	if got != errCodeReaderFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeReaderFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_ResetFailure(t *testing.T) {
	t.Parallel()

	got := ErrCodeResetFailure.String()
	if got != errCodeResetFailureStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeResetFailureStr,
		)
	}
}

func TestChargePointErrorCode_String_UnderVoltage(t *testing.T) {
	t.Parallel()

	got := ErrCodeUnderVoltage.String()
	if got != errCodeUnderVoltageStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeUnderVoltageStr,
		)
	}
}

func TestChargePointErrorCode_String_WeakSignal(t *testing.T) {
	t.Parallel()

	got := ErrCodeWeakSignal.String()
	if got != errCodeWeakSignalStr {
		t.Errorf(
			ErrorMethodMismatch,
			errCodeMethodString,
			got,
			errCodeWeakSignalStr,
		)
	}
}
