package ocpp16types

import (
	"testing"
)

const (
	statusAcceptedStr      = "Accepted"
	statusBlockedStr       = "Blocked"
	statusExpiredStr       = "Expired"
	statusInvalidStr       = "Invalid"
	statusConcurrentTxStr  = "ConcurrentTx"
	authStatusMethodString = "AuthorizationStatus.String()"
)

func TestAuthorizationStatus_IsValid_Accepted(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusAccepted.IsValid() {
		t.Errorf(ErrorIsValidFalse, "AuthorizationStatusAccepted")
	}
}

func TestAuthorizationStatus_IsValid_Blocked(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusBlocked.IsValid() {
		t.Errorf(ErrorIsValidFalse, "AuthorizationStatusBlocked")
	}
}

func TestAuthorizationStatus_IsValid_Expired(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusExpired.IsValid() {
		t.Errorf(ErrorIsValidFalse, "AuthorizationStatusExpired")
	}
}

func TestAuthorizationStatus_IsValid_Invalid(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusInvalid.IsValid() {
		t.Errorf(ErrorIsValidFalse, "AuthorizationStatusInvalid")
	}
}

func TestAuthorizationStatus_IsValid_ConcurrentTx(t *testing.T) {
	t.Parallel()

	if !AuthorizationStatusConcurrentTx.IsValid() {
		t.Errorf(ErrorIsValidFalse, "AuthorizationStatusConcurrentTx")
	}
}

func TestAuthorizationStatus_IsValid_Empty(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "AuthorizationStatus(\"\")")
	}
}

func TestAuthorizationStatus_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("Unknown")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "AuthorizationStatus(\"Unknown\")")
	}
}

func TestAuthorizationStatus_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	status := AuthorizationStatus("accepted")
	if status.IsValid() {
		t.Errorf(ErrorIsValidTrue, "AuthorizationStatus(\"accepted\")")
	}
}

func TestAuthorizationStatus_String_Accepted(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusAccepted.String()
	if got != statusAcceptedStr {
		t.Errorf(
			ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusAcceptedStr,
		)
	}
}

func TestAuthorizationStatus_String_Blocked(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusBlocked.String()
	if got != statusBlockedStr {
		t.Errorf(
			ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusBlockedStr,
		)
	}
}

func TestAuthorizationStatus_String_Expired(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusExpired.String()
	if got != statusExpiredStr {
		t.Errorf(
			ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusExpiredStr,
		)
	}
}

func TestAuthorizationStatus_String_Invalid(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusInvalid.String()
	if got != statusInvalidStr {
		t.Errorf(
			ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusInvalidStr,
		)
	}
}

func TestAuthorizationStatus_String_ConcurrentTx(t *testing.T) {
	t.Parallel()

	got := AuthorizationStatusConcurrentTx.String()
	if got != statusConcurrentTxStr {
		t.Errorf(
			ErrorMethodMismatch,
			authStatusMethodString,
			got,
			statusConcurrentTxStr,
		)
	}
}
