package ocpp16types

import (
	"testing"
)

const (
	reasonDeAuthorizedStr   = "DeAuthorized"
	reasonEmergencyStopStr  = "EmergencyStop"
	reasonEVDisconnectedStr = "EVDisconnected"
	reasonHardResetStr      = "HardReset"
	reasonLocalStr          = "Local"
	reasonOtherStr          = "Other"
	reasonPowerLossStr      = "PowerLoss"
	reasonRebootStr         = "Reboot"
	reasonRemoteStr         = "Remote"
	reasonSoftResetStr      = "SoftReset"
	reasonUnlockCommandStr  = "UnlockCommand"
	stopReasonMethodString  = "StopReason.String()"
)

func TestStopReason_IsValid_DeAuthorized(t *testing.T) {
	t.Parallel()

	if !ReasonDeAuthorized.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonDeAuthorized")
	}
}

func TestStopReason_IsValid_EmergencyStop(t *testing.T) {
	t.Parallel()

	if !ReasonEmergencyStop.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonEmergencyStop")
	}
}

func TestStopReason_IsValid_EVDisconnected(t *testing.T) {
	t.Parallel()

	if !ReasonEVDisconnected.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonEVDisconnected")
	}
}

func TestStopReason_IsValid_HardReset(t *testing.T) {
	t.Parallel()

	if !ReasonHardReset.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonHardReset")
	}
}

func TestStopReason_IsValid_Local(t *testing.T) {
	t.Parallel()

	if !ReasonLocal.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonLocal")
	}
}

func TestStopReason_IsValid_Other(t *testing.T) {
	t.Parallel()

	if !ReasonOther.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonOther")
	}
}

func TestStopReason_IsValid_PowerLoss(t *testing.T) {
	t.Parallel()

	if !ReasonPowerLoss.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonPowerLoss")
	}
}

func TestStopReason_IsValid_Reboot(t *testing.T) {
	t.Parallel()

	if !ReasonReboot.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonReboot")
	}
}

func TestStopReason_IsValid_Remote(t *testing.T) {
	t.Parallel()

	if !ReasonRemote.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonRemote")
	}
}

func TestStopReason_IsValid_SoftReset(t *testing.T) {
	t.Parallel()

	if !ReasonSoftReset.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonSoftReset")
	}
}

func TestStopReason_IsValid_UnlockCommand(t *testing.T) {
	t.Parallel()

	if !ReasonUnlockCommand.IsValid() {
		t.Errorf(ErrorIsValidFalse, "StopReasonUnlockCommand")
	}
}

func TestStopReason_IsValid_Empty(t *testing.T) {
	t.Parallel()

	reason := Reason("")
	if reason.IsValid() {
		t.Errorf(ErrorIsValidTrue, "StopReason(\"\")")
	}
}

func TestStopReason_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	reason := Reason("Unknown")
	if reason.IsValid() {
		t.Errorf(ErrorIsValidTrue, "StopReason(\"Unknown\")")
	}
}

func TestStopReason_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	reason := Reason("deauthorized")
	if reason.IsValid() {
		t.Errorf(ErrorIsValidTrue, "StopReason(\"deauthorized\")")
	}
}

func TestStopReason_String_DeAuthorized(t *testing.T) {
	t.Parallel()

	got := ReasonDeAuthorized.String()
	if got != reasonDeAuthorizedStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonDeAuthorizedStr,
		)
	}
}

func TestStopReason_String_EmergencyStop(t *testing.T) {
	t.Parallel()

	got := ReasonEmergencyStop.String()
	if got != reasonEmergencyStopStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonEmergencyStopStr,
		)
	}
}

func TestStopReason_String_EVDisconnected(t *testing.T) {
	t.Parallel()

	got := ReasonEVDisconnected.String()
	if got != reasonEVDisconnectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonEVDisconnectedStr,
		)
	}
}

func TestStopReason_String_HardReset(t *testing.T) {
	t.Parallel()

	got := ReasonHardReset.String()
	if got != reasonHardResetStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonHardResetStr,
		)
	}
}

func TestStopReason_String_Local(t *testing.T) {
	t.Parallel()

	got := ReasonLocal.String()
	if got != reasonLocalStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonLocalStr,
		)
	}
}

func TestStopReason_String_Other(t *testing.T) {
	t.Parallel()

	got := ReasonOther.String()
	if got != reasonOtherStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonOtherStr,
		)
	}
}

func TestStopReason_String_PowerLoss(t *testing.T) {
	t.Parallel()

	got := ReasonPowerLoss.String()
	if got != reasonPowerLossStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonPowerLossStr,
		)
	}
}

func TestStopReason_String_Reboot(t *testing.T) {
	t.Parallel()

	got := ReasonReboot.String()
	if got != reasonRebootStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonRebootStr,
		)
	}
}

func TestStopReason_String_Remote(t *testing.T) {
	t.Parallel()

	got := ReasonRemote.String()
	if got != reasonRemoteStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonRemoteStr,
		)
	}
}

func TestStopReason_String_SoftReset(t *testing.T) {
	t.Parallel()

	got := ReasonSoftReset.String()
	if got != reasonSoftResetStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonSoftResetStr,
		)
	}
}

func TestStopReason_String_UnlockCommand(t *testing.T) {
	t.Parallel()

	got := ReasonUnlockCommand.String()
	if got != reasonUnlockCommandStr {
		t.Errorf(
			ErrorMethodMismatch,
			stopReasonMethodString,
			got,
			reasonUnlockCommandStr,
		)
	}
}
