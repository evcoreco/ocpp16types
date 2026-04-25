package ocpp16types

import (
	"testing"
)

const (
	errUnexpectedIDTagInfo = "unexpected error creating IDTagInfo: %v"
	zeroLen                = 0
)

func TestNewIDTagInfo_Accepted(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.Status() != AuthorizationStatusAccepted {
		t.Errorf(
			ErrorMethodMismatch,
			"IDTagInfo.Status()",
			info.Status(),
			AuthorizationStatusAccepted,
		)
	}
}

func TestNewIDTagInfo_AllStatuses(t *testing.T) {
	t.Parallel()

	statuses := []AuthorizationStatus{
		AuthorizationStatusAccepted,
		AuthorizationStatusBlocked,
		AuthorizationStatusExpired,
		AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx,
	}

	for _, status := range statuses {
		info, err := NewIDTagInfo(status)
		if err != nil {
			t.Errorf("unexpected error for status %s: %v", status, err)

			continue
		}

		if info.Status() != status {
			t.Errorf(
				ErrorMethodMismatch,
				"IDTagInfo.Status()",
				info.Status(),
				status,
			)
		}
	}
}

func TestNewIDTagInfo_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := NewIDTagInfo(AuthorizationStatus("Bogus"))
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid authorization status")
	}
}

func TestIDTagInfo_WithExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	dt, err := NewDateTime("2024-12-31T23:59:59Z")
	if err != nil {
		t.Fatalf("unexpected error creating DateTime: %v", err)
	}

	info = info.WithExpiryDate(dt)
	expiryDate := info.ExpiryDate()

	if expiryDate == nil {
		t.Errorf(ErrorWantNonNil, "IDTagInfo.ExpiryDate()")
	}
}

func TestIDTagInfo_WithParentIDTag(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	token, err := NewCiString20Type("PARENT-TOKEN")
	if err != nil {
		t.Fatalf("unexpected error creating CiString20Type: %v", err)
	}

	parentToken := NewIDToken(token)
	info = info.WithParentIDTag(parentToken)
	retrievedParent := info.ParentIDTag()

	if retrievedParent == nil {
		t.Errorf(ErrorWantNonNil, "IDTagInfo.ParentIDTag()")
	}
}

func TestIDTagInfo_ExpiryDate_Nil(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	expiryDate := info.ExpiryDate()
	if expiryDate != nil {
		t.Errorf("IDTagInfo.ExpiryDate() = %v, want nil", expiryDate)
	}
}

func TestIDTagInfo_ParentIDTag_Nil(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	parentTag := info.ParentIDTag()
	if parentTag != nil {
		t.Errorf("IDTagInfo.ParentIDTag() = %v, want nil", parentTag)
	}
}

func TestIDTagInfo_String(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	strRepr := info.String()
	if !containsSubstring(strRepr, "IDTagInfo") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"IDTagInfo",
		)
	}
}

func TestIDTagInfo_String_WithExpiryDate(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	dt, err := NewDateTime("2024-12-31T23:59:59Z")
	if err != nil {
		t.Fatalf("unexpected error creating DateTime: %v", err)
	}

	info = info.WithExpiryDate(dt)
	strRepr := info.String()

	if !containsSubstring(strRepr, "ExpiryDate") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"ExpiryDate",
		)
	}
}

func TestIDTagInfo_String_WithParentIDTag(t *testing.T) {
	t.Parallel()

	info, err := NewIDTagInfo(AuthorizationStatusAccepted)
	if err != nil {
		t.Fatalf(errUnexpectedIDTagInfo, err)
	}

	token, err := NewCiString20Type("PARENT-TOKEN")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	parentToken := NewIDToken(token)
	info = info.WithParentIDTag(parentToken)
	strRepr := info.String()

	if !containsSubstring(strRepr, "ParentIDTag") {
		t.Errorf(
			ErrorWantContains,
			strRepr,
			"ParentIDTag",
		)
	}
}

func containsSubstring(str, substring string) bool {
	return len(str) > zeroLen && len(substring) > zeroLen &&
		findSubstring(str, substring)
}

func findSubstring(str, substring string) bool {
	for i := zeroLen; i <= len(str)-len(substring); i++ {
		if str[i:i+len(substring)] == substring {
			return true
		}
	}

	return false
}
