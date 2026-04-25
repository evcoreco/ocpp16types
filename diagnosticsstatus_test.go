package ocpp16types

import (
	"testing"
)

const (
	diagStatusIdleStr      = "Idle"
	diagStatusUploadedStr  = "Uploaded"
	diagStatusUpFailedStr  = "UploadFailed"
	diagStatusUploadingStr = "Uploading"
	diagStatusMethodString = "DiagnosticsStatus.String()"
)

func TestDiagnosticsStatus_IsValid_Idle(
	t *testing.T,
) {
	t.Parallel()

	if !DiagnosticsStatusIdle.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DiagnosticsStatusIdle",
		)
	}
}

func TestDiagnosticsStatus_IsValid_Uploaded(
	t *testing.T,
) {
	t.Parallel()

	if !DiagnosticsStatusUploaded.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DiagnosticsStatusUploaded",
		)
	}
}

func TestDiagnosticsStatus_IsValid_UploadFailed(
	t *testing.T,
) {
	t.Parallel()

	if !DiagnosticsStatusUploadFailed.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DiagnosticsStatusUploadFailed",
		)
	}
}

func TestDiagnosticsStatus_IsValid_Uploading(
	t *testing.T,
) {
	t.Parallel()

	if !DiagnosticsStatusUploading.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DiagnosticsStatusUploading",
		)
	}
}

func TestDiagnosticsStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := DiagnosticsStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DiagnosticsStatus(\"\")",
		)
	}
}

func TestDiagnosticsStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := DiagnosticsStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DiagnosticsStatus(\"Unknown\")",
		)
	}
}

func TestDiagnosticsStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := DiagnosticsStatus("idle")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DiagnosticsStatus(\"idle\")",
		)
	}
}

func TestDiagnosticsStatus_String_Idle(
	t *testing.T,
) {
	t.Parallel()

	got := DiagnosticsStatusIdle.String()
	if got != diagStatusIdleStr {
		t.Errorf(
			ErrorMethodMismatch,
			diagStatusMethodString,
			got,
			diagStatusIdleStr,
		)
	}
}

func TestDiagnosticsStatus_String_Uploaded(
	t *testing.T,
) {
	t.Parallel()

	got := DiagnosticsStatusUploaded.String()
	if got != diagStatusUploadedStr {
		t.Errorf(
			ErrorMethodMismatch,
			diagStatusMethodString,
			got,
			diagStatusUploadedStr,
		)
	}
}

func TestDiagnosticsStatus_String_UploadFailed(
	t *testing.T,
) {
	t.Parallel()

	got := DiagnosticsStatusUploadFailed.String()
	if got != diagStatusUpFailedStr {
		t.Errorf(
			ErrorMethodMismatch,
			diagStatusMethodString,
			got,
			diagStatusUpFailedStr,
		)
	}
}

func TestDiagnosticsStatus_String_Uploading(
	t *testing.T,
) {
	t.Parallel()

	got := DiagnosticsStatusUploading.String()
	if got != diagStatusUploadingStr {
		t.Errorf(
			ErrorMethodMismatch,
			diagStatusMethodString,
			got,
			diagStatusUploadingStr,
		)
	}
}
