package ocpp16types

import (
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*IDTagInfo)(nil)

// IDTagInfo is the OCPP 1.6 idTagInfo data type — the Central System's
// authorization response for a presented idToken.
//
// What it means: a validated authorization record carrying a required
// [AuthorizationStatus] and optional expiry date and parent tag. It
// represents what the Central System says about a specific [IDToken].
//
// When to use it: Authorize.conf, StartTransaction.conf, and
// StopTransaction.req. Construct with [NewIDTagInfo], then attach optional
// fields using [IDTagInfo.WithExpiryDate] and [IDTagInfo.WithParentIDTag].
//
// What it is not: the token being authorized. The token itself is [IDToken].
// IDTagInfo is the response; IDToken is the input.
//
// See also: [AuthorizationStatus] is the required field; [IDToken] is the
// token that triggers the authorization check; [DateTime] is used for the
// optional expiry date; [AuthorizationData] stores IDTagInfo in the local
// authorization cache.
type IDTagInfo struct {
	status      AuthorizationStatus
	expiryDate  *DateTime // Optional
	parentIDTag *IDToken  // Optional
}

// NewIDTagInfo creates a new IDTagInfo with the given status.
// ExpiryDate and ParentIDTag are optional and can be set separately.
func NewIDTagInfo(status AuthorizationStatus) (IDTagInfo, error) {
	if !status.IsValid() {
		return IDTagInfo{}, fmt.Errorf(
			"NewIDTagInfo: "+ErrorFieldFormat,
			"AuthorizationStatus",
			ErrInvalidValue,
		)
	}

	return IDTagInfo{
		status:      status,
		expiryDate:  nil,
		parentIDTag: nil,
	}, nil
}

// Status returns the authorization status.
func (i IDTagInfo) Status() AuthorizationStatus {
	return i.status
}

// ExpiryDate returns a defensive copy of the expiry date, or nil if not set.
func (i IDTagInfo) ExpiryDate() *DateTime {
	if i.expiryDate == nil {
		return nil
	}

	copiedExpiryDate := *i.expiryDate

	return &copiedExpiryDate
}

// ParentIDTag returns a defensive copy of the parent ID tag, or nil if not set.
func (i IDTagInfo) ParentIDTag() *IDToken {
	if i.parentIDTag == nil {
		return nil
	}

	copiedParentIDTag := *i.parentIDTag

	return &copiedParentIDTag
}

// WithExpiryDate sets the expiry date and returns the IDTagInfo.
func (i IDTagInfo) WithExpiryDate(expiryDate DateTime) IDTagInfo {
	i.expiryDate = &expiryDate

	return i
}

// WithParentIDTag sets the parent ID tag and returns the IDTagInfo.
func (i IDTagInfo) WithParentIDTag(parentIDTag IDToken) IDTagInfo {
	i.parentIDTag = &parentIDTag

	return i
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the IDTagInfo for debugging purposes.
func (i IDTagInfo) String() string {
	result := "IDTagInfo{Status: " + i.status.String()

	if i.expiryDate != nil {
		result += ", ExpiryDate: " + i.expiryDate.String()
	}

	if i.parentIDTag != nil {
		result += ", ParentIDTag: " + i.parentIDTag.String()
	}

	result += "}"

	return result
}
