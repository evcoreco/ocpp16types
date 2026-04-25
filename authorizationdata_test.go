package ocpp16types

import (
	"testing"
)

const (
	testToken123  = "TOKEN123"
	testStatusAcc = "Accepted"
	testNotADate  = "not-a-date"
)

func TestNewAuthorizationData_ValidNoInfo(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IDTag:     testToken123,
		IDTagInfo: nil,
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IDTag().String() != testToken123 {
		t.Errorf(
			ErrorMethodMismatch,
			"AuthorizationData.IDTag()",
			authData.IDTag().String(),
			testToken123,
		)
	}

	if authData.IDTagInfo() != nil {
		t.Errorf(
			"AuthorizationData.IDTagInfo() = %v, want nil",
			authData.IDTagInfo(),
		)
	}
}

func TestNewAuthorizationData_ValidWithInfo(t *testing.T) {
	t.Parallel()

	expiry := "2024-12-31T23:59:59Z"
	parent := "PARENT-TAG"

	input := AuthorizationDataInput{
		IDTag: "TOKEN456",
		IDTagInfo: &IDTagInfoInput{
			Status:      testStatusAcc,
			ExpiryDate:  &expiry,
			ParentIDTag: &parent,
		},
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IDTagInfo() == nil {
		t.Fatalf(ErrorWantNonNil, "IDTagInfo")
	}

	if authData.IDTagInfo().Status() != AuthorizationStatusAccepted {
		t.Errorf(
			ErrorMethodMismatch,
			"IDTagInfo().Status()",
			authData.IDTagInfo().Status(),
			AuthorizationStatusAccepted,
		)
	}

	if authData.IDTagInfo().ExpiryDate() == nil {
		t.Errorf(ErrorWantNonNil, "ExpiryDate")
	}

	if authData.IDTagInfo().ParentIDTag() == nil {
		t.Errorf(ErrorWantNonNil, "ParentIDTag")
	}
}

func TestNewAuthorizationData_EmptyIDTag(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IDTag:     "",
		IDTagInfo: nil,
	}

	_, err := NewAuthorizationData(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty IDTag")
	}
}

func TestNewAuthorizationData_InvalidStatus(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IDTag: "TOKEN789",
		IDTagInfo: &IDTagInfoInput{
			Status:      "Bogus",
			ExpiryDate:  nil,
			ParentIDTag: nil,
		},
	}

	_, err := NewAuthorizationData(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid IDTagInfo status")
	}
}

func TestNewAuthorizationData_WithExpiryAndParent(
	t *testing.T,
) {
	t.Parallel()

	expiry := "2025-06-15T12:00:00Z"
	parent := "PARENT01"

	input := AuthorizationDataInput{
		IDTag: "CHILD01",
		IDTagInfo: &IDTagInfoInput{
			Status:      testStatusAcc,
			ExpiryDate:  &expiry,
			ParentIDTag: &parent,
		},
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if authData.IDTagInfo().ExpiryDate() == nil {
		t.Errorf(ErrorWantNonNil, "ExpiryDate")
	}

	if authData.IDTagInfo().ParentIDTag() == nil {
		t.Errorf(ErrorWantNonNil, "ParentIDTag")
	}
}

func TestBuildIDTagInfo_InvalidExpiryDate(t *testing.T) {
	t.Parallel()

	badDate := testNotADate
	input := IDTagInfoInput{
		Status:      testStatusAcc,
		ExpiryDate:  &badDate,
		ParentIDTag: nil,
	}

	_, err := buildIDTagInfo(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ExpiryDate")
	}
}

func TestBuildIDTagInfo_InvalidParentIDTag(t *testing.T) {
	t.Parallel()

	badParent := ""
	input := IDTagInfoInput{
		Status:      testStatusAcc,
		ExpiryDate:  nil,
		ParentIDTag: &badParent,
	}

	_, err := buildIDTagInfo(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ParentIDTag")
	}
}

func TestAuthorizationData_String_NoInfo(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IDTag:     testToken123,
		IDTagInfo: nil,
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "AuthorizationData{IDTag: TOKEN123}"
	if authData.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"AuthorizationData.String()",
			authData.String(),
			expected,
		)
	}
}

func TestAuthorizationData_String_WithInfo(t *testing.T) {
	t.Parallel()

	input := AuthorizationDataInput{
		IDTag: testToken123,
		IDTagInfo: &IDTagInfoInput{
			Status:      testStatusAcc,
			ExpiryDate:  nil,
			ParentIDTag: nil,
		},
	}

	authData, err := NewAuthorizationData(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	got := authData.String()
	if got != "AuthorizationData{IDTag: TOKEN123, IDTagInfo: IDTagInfo{Status: Accepted}}" {
		t.Errorf(
			ErrorMethodMismatch,
			"AuthorizationData.String()",
			got,
			"AuthorizationData{IDTag: TOKEN123, IDTagInfo: IDTagInfo{Status: Accepted}}",
		)
	}
}
