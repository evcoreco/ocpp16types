package ocpp16types

import (
	"testing"
)

const (
	confStatusAcceptedStr  = "Accepted"
	confStatusRejectedStr  = "Rejected"
	confStatusRebootStr    = "RebootRequired"
	confStatusNotSuppStr   = "NotSupported"
	confStatusMethodString = "ConfigurationStatus.String()"
)

func TestConfigStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !ConfigurationStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ConfigurationStatusAccepted",
		)
	}
}

func TestConfigStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !ConfigurationStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ConfigurationStatusRejected",
		)
	}
}

func TestConfigStatus_IsValid_RebootRequired(
	t *testing.T,
) {
	t.Parallel()

	if !ConfigurationStatusRebootRequired.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ConfigurationStatusRebootRequired",
		)
	}
}

func TestConfigStatus_IsValid_NotSupported(
	t *testing.T,
) {
	t.Parallel()

	if !ConfigurationStatusNotSupported.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"ConfigurationStatusNotSupported",
		)
	}
}

func TestConfigStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := ConfigurationStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ConfigurationStatus(\"\")",
		)
	}
}

func TestConfigStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := ConfigurationStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ConfigurationStatus(\"Unknown\")",
		)
	}
}

func TestConfigStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := ConfigurationStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"ConfigurationStatus(\"accepted\")",
		)
	}
}

func TestConfigStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := ConfigurationStatusAccepted.String()
	if got != confStatusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			confStatusMethodString,
			got,
			confStatusAcceptedStr,
		)
	}
}

func TestConfigStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := ConfigurationStatusRejected.String()
	if got != confStatusRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			confStatusMethodString,
			got,
			confStatusRejectedStr,
		)
	}
}

func TestConfigStatus_String_RebootRequired(
	t *testing.T,
) {
	t.Parallel()

	got := ConfigurationStatusRebootRequired.String()
	if got != confStatusRebootStr {
		t.Errorf(
			ErrorMethodMismatch,
			confStatusMethodString,
			got,
			confStatusRebootStr,
		)
	}
}

func TestConfigStatus_String_NotSupported(
	t *testing.T,
) {
	t.Parallel()

	got := ConfigurationStatusNotSupported.String()
	if got != confStatusNotSuppStr {
		t.Errorf(
			ErrorMethodMismatch,
			confStatusMethodString,
			got,
			confStatusNotSuppStr,
		)
	}
}
