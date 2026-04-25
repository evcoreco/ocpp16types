package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = (*IDToken)(nil)

// IDToken represents an OCPP 1.6 identifier token used for authorization.
// It wraps a CiString20Type value that must not be empty.
type IDToken struct {
	value CiString20Type
}

// NewIDToken creates a new IDToken from a CiString20Type value.
// CiString20Type already validates that the value is non-empty.
func NewIDToken(ciString20Type CiString20Type) IDToken {
	return IDToken{value: ciString20Type}
}

// Value returns the underlying CiString20Type value of the IDToken.
func (idtoken IDToken) Value() CiString20Type {
	return idtoken.value
}

// String returns the string representation of the IDToken.
func (idtoken IDToken) String() string {
	return idtoken.value.Value()
}
