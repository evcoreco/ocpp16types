package ocpp16types

import (
	"testing"
)

const (
	valueFormatRawStr        = "Raw"
	valueFormatSignedDataStr = "SignedData"
	valueFormatMethodString  = "ValueFormat.String()"
)

func TestValueFormat_IsValid_Raw(t *testing.T) {
	t.Parallel()

	if !ValueFormatRaw.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ValueFormatRaw")
	}
}

func TestValueFormat_IsValid_SignedData(t *testing.T) {
	t.Parallel()

	if !ValueFormatSignedData.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ValueFormatSignedData")
	}
}

func TestValueFormat_IsValid_Empty(t *testing.T) {
	t.Parallel()

	format := ValueFormat("")
	if format.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ValueFormat(\"\")")
	}
}

func TestValueFormat_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	format := ValueFormat("Unknown")
	if format.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ValueFormat(\"Unknown\")")
	}
}

func TestValueFormat_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	format := ValueFormat("raw")
	if format.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ValueFormat(\"raw\")")
	}
}

func TestValueFormat_String_Raw(t *testing.T) {
	t.Parallel()

	got := ValueFormatRaw.String()
	if got != valueFormatRawStr {
		t.Errorf(
			ErrorMethodMismatch,
			valueFormatMethodString,
			got,
			valueFormatRawStr,
		)
	}
}

func TestValueFormat_String_SignedData(t *testing.T) {
	t.Parallel()

	got := ValueFormatSignedData.String()
	if got != valueFormatSignedDataStr {
		t.Errorf(
			ErrorMethodMismatch,
			valueFormatMethodString,
			got,
			valueFormatSignedDataStr,
		)
	}
}
