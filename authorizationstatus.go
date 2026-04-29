package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = AuthorizationStatus("")

// AuthorizationStatus is the OCPP 1.6 authorization status enumeration,
// carried in the idTagInfo response to an authorization request.
//
// What it means: the Central System's decision about a presented idToken —
// whether the charge point should allow or deny the charging session.
//
// When to use it: populate the required status field when constructing
// [IDTagInfo] via [NewIDTagInfo]. Call [AuthorizationStatus.IsValid] before
// accepting a status string received from an external source.
//
// What it is not: a transport-level result code or an HTTP status. It is
// scoped exclusively to OCPP authorization responses.
//
// See also: [IDTagInfo] embeds AuthorizationStatus as its required field;
// [IDToken] is the token being evaluated; [AuthorizationData] pairs a token
// with an AuthorizationStatus in the local authorization cache.
type AuthorizationStatus string

// AuthorizationStatus enumeration values as defined in OCPP 1.6.
const (
	AuthorizationStatusAccepted     AuthorizationStatus = "Accepted"
	AuthorizationStatusBlocked      AuthorizationStatus = "Blocked"
	AuthorizationStatusExpired      AuthorizationStatus = "Expired"
	AuthorizationStatusInvalid      AuthorizationStatus = "Invalid"
	AuthorizationStatusConcurrentTx AuthorizationStatus = "ConcurrentTx"
)

// IsValid checks if the AuthorizationStatus value is valid per OCPP 1.6.
func (a AuthorizationStatus) IsValid() bool {
	switch a {
	case AuthorizationStatusAccepted,
		AuthorizationStatusBlocked,
		AuthorizationStatusExpired,
		AuthorizationStatusInvalid,
		AuthorizationStatusConcurrentTx:
		return true
	default:
		return false
	}
}

// String returns the string representation of AuthorizationStatus.
func (a AuthorizationStatus) String() string {
	return string(a)
}
