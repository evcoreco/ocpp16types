package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = (*IDToken)(nil)

// IDToken is the OCPP 1.6 idToken data type — an identifier that a charge
// point presents to the Central System for authorization.
//
// What it means: a validated, bounded token value wrapping a [CiString20Type].
// It carries the identity that triggers an authorization check.
//
// When to use it: Authorize.req, StartTransaction.req, StopTransaction.req,
// and any other OCPP message that carries an identifier token.
//
// What it is not: an authorization result. The outcome of checking a token is
// [IDTagInfo], not IDToken. IDToken is the input; IDTagInfo is the response.
//
// See also: [IDTagInfo] carries the authorization result for a given token;
// [CiString20Type] is the underlying validated string primitive;
// [AuthorizationData] pairs an IDToken with optional IDTagInfo in the local
// authorization cache.
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
