package ocpp16types

import (
	"testing"
)

func TestRemoteStartTxStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	v := RemoteStartTransactionStatusAccepted
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RemoteStartTransactionStatusAccepted",
		)
	}
}

func TestRemoteStartTxStatus_IsValid_Rejected(t *testing.T) {
	t.Parallel()

	v := RemoteStartTransactionStatusRejected
	if !v.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"RemoteStartTransactionStatusRejected",
		)
	}
}

func TestRemoteStartTxStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	s := RemoteStartTransactionStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"\")",
		)
	}
}

func TestRemoteStartTxStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	s := RemoteStartTransactionStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"RemoteStartTransactionStatus(\"Unknown\")",
		)
	}
}

func TestRemoteStartTxStatus_String(t *testing.T) {
	t.Parallel()

	v := RemoteStartTransactionStatusAccepted
	got := v.String()

	want := string(RemoteStartTransactionStatusAccepted)
	if got != want {
		t.Errorf(
			ErrorMethodMismatch,
			"RemoteStartTransactionStatus.String()",
			got,
			want,
		)
	}
}
