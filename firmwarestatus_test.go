package ocpp16types

import (
	"testing"
)

func TestFirmwareStatus_IsValid_Downloaded(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusDownloaded.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusDownloaded",
		)
	}
}

func TestFirmwareStatus_IsValid_DlFailed(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusDownloadFailed.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusDownloadFailed",
		)
	}
}

func TestFirmwareStatus_IsValid_Downloading(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusDownloading.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusDownloading",
		)
	}
}

func TestFirmwareStatus_IsValid_Idle(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusIdle.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusIdle",
		)
	}
}

func TestFirmwareStatus_IsValid_InstFailed(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusInstallationFailed.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusInstallationFailed",
		)
	}
}

func TestFirmwareStatus_IsValid_Installing(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusInstalling.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusInstalling",
		)
	}
}

func TestFirmwareStatus_IsValid_Installed(t *testing.T) {
	t.Parallel()

	if !FirmwareStatusInstalled.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"FirmwareStatusInstalled",
		)
	}
}

func TestFirmwareStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := FirmwareStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"FirmwareStatus(\"\")",
		)
	}
}

func TestFirmwareStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := FirmwareStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"FirmwareStatus(\"Unknown\")",
		)
	}
}

func TestFirmwareStatus_String(t *testing.T) {
	t.Parallel()

	got := FirmwareStatusDownloaded.String()
	want := "Downloaded"

	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"FirmwareStatus.String()",
			got,
			want,
		)
	}
}
