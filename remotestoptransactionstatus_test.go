package ocpp16types

import (
	"testing"
)

func TestRemoteStopTxStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	v := RemoteStopTransactionStatusAccepted
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RemoteStopTransactionStatusAccepted",
		)
	}
}

func TestRemoteStopTxStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	v := RemoteStopTransactionStatusRejected
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RemoteStopTransactionStatusRejected",
		)
	}
}

func TestRemoteStopTxStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := RemoteStopTransactionStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"\")",
		)
	}
}

func TestRemoteStopTxStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := RemoteStopTransactionStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RemoteStopTransactionStatus(\"Unknown\")",
		)
	}
}

func TestRemoteStopTxStatus_String(t *testing.T) {
	t.Parallel()

	v := RemoteStopTransactionStatusAccepted
	got := v.String()

	want := string(RemoteStopTransactionStatusAccepted)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"RemoteStopTransactionStatus.String()",
			got,
			want,
		)
	}
}
