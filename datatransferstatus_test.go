package ocpp16types

import (
	"testing"
)

const (
	dtStatusAcceptedStr  = "Accepted"
	dtStatusRejectedStr  = "Rejected"
	dtStatusUnkMsgIDStr  = "UnknownMessageId"
	dtStatusUnkVendorStr = "UnknownVendor"
	dtStatusMethodString = "DataTransferStatus.String()"
)

func TestDataTransferStatus_IsValid_Accepted(
	t *testing.T,
) {
	t.Parallel()

	if !DataTransferStatusAccepted.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DataTransferStatusAccepted",
		)
	}
}

func TestDataTransferStatus_IsValid_Rejected(
	t *testing.T,
) {
	t.Parallel()

	if !DataTransferStatusRejected.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DataTransferStatusRejected",
		)
	}
}

func TestDataTransferStatus_IsValid_UnkMsgId(
	t *testing.T,
) {
	t.Parallel()

	if !DataTransferStatusUnknownMessageID.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DataTransferStatusUnknownMessageID",
		)
	}
}

func TestDataTransferStatus_IsValid_UnkVendor(
	t *testing.T,
) {
	t.Parallel()

	if !DataTransferStatusUnknownVendor.IsValid() {
		t.Errorf(
			ErrorIsValidFalse,
			"DataTransferStatusUnknownVendor",
		)
	}
}

func TestDataTransferStatus_IsValid_Empty(
	t *testing.T,
) {
	t.Parallel()

	s := DataTransferStatus("")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DataTransferStatus(\"\")",
		)
	}
}

func TestDataTransferStatus_IsValid_Unknown(
	t *testing.T,
) {
	t.Parallel()

	s := DataTransferStatus("Unknown")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DataTransferStatus(\"Unknown\")",
		)
	}
}

func TestDataTransferStatus_IsValid_Lowercase(
	t *testing.T,
) {
	t.Parallel()

	s := DataTransferStatus("accepted")
	if s.IsValid() {
		t.Errorf(
			ErrorIsValidTrue,
			"DataTransferStatus(\"accepted\")",
		)
	}
}

func TestDataTransferStatus_String_Accepted(
	t *testing.T,
) {
	t.Parallel()

	got := DataTransferStatusAccepted.String()
	if got != dtStatusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			dtStatusMethodString,
			got,
			dtStatusAcceptedStr,
		)
	}
}

func TestDataTransferStatus_String_Rejected(
	t *testing.T,
) {
	t.Parallel()

	got := DataTransferStatusRejected.String()
	if got != dtStatusRejectedStr {
		t.Errorf(
			ErrorMethodMismatch,
			dtStatusMethodString,
			got,
			dtStatusRejectedStr,
		)
	}
}

func TestDataTransferStatus_String_UnkMsgId(
	t *testing.T,
) {
	t.Parallel()

	got := DataTransferStatusUnknownMessageID.String()
	if got != dtStatusUnkMsgIDStr {
		t.Errorf(
			ErrorMethodMismatch,
			dtStatusMethodString,
			got,
			dtStatusUnkMsgIDStr,
		)
	}
}

func TestDataTransferStatus_String_UnkVendor(
	t *testing.T,
) {
	t.Parallel()

	got := DataTransferStatusUnknownVendor.String()
	if got != dtStatusUnkVendorStr {
		t.Errorf(
			ErrorMethodMismatch,
			dtStatusMethodString,
			got,
			dtStatusUnkVendorStr,
		)
	}
}
