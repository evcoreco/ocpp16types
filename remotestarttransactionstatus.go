package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = RemoteStartTransactionStatus("")

// RemoteStartTransactionStatus represents the result of a
// RemoteStartTransaction request.
type RemoteStartTransactionStatus string

// Alias for shorter constant declarations.
type rstt = RemoteStartTransactionStatus

// RemoteStartTransactionStatus enumeration values as defined in OCPP 1.6.
const (
	RemoteStartTransactionStatusAccepted rstt = "Accepted"
	RemoteStartTransactionStatusRejected rstt = "Rejected"
)

// IsValid checks if the RemoteStartTransactionStatus value is valid
// per OCPP 1.6.
func (r RemoteStartTransactionStatus) IsValid() bool {
	switch r {
	case RemoteStartTransactionStatusAccepted,
		RemoteStartTransactionStatusRejected:
		return true
	default:
		return false
	}
}

// String returns the string representation of RemoteStartTransactionStatus.
func (r RemoteStartTransactionStatus) String() string {
	return string(r)
}
