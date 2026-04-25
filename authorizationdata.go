package ocpp16types

import (
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*AuthorizationData)(nil)

// AuthorizationDataInput represents the raw input data for creating an
// AuthorizationData entry in the local authorization li
type AuthorizationDataInput struct {
	// Required: The identifier to be authorized.
	IDTag string
	// Optional: Authorization status information for this idTag.
	// When omitted in a Differential update, the entry is deleted.
	IDTagInfo *IDTagInfoInput
}

// IDTagInfoInput represents the raw input data for IDTagInfo.
type IDTagInfoInput struct {
	// Required: Authorization status.
	Status string
	// Optional: Expiry date in RFC3339 format.
	ExpiryDate *string
	// Optional: Parent identifier tag.
	ParentIDTag *string
}

// AuthorizationData represents an entry in the local authorization list.
type AuthorizationData struct {
	idTag     IDToken
	idTagInfo *IDTagInfo
}

// IDTag returns the identifier token.
func (a AuthorizationData) IDTag() IDToken {
	return a.idTag
}

// IDTagInfo returns a defensive copy of the authorization info, or nil if not set.
func (a AuthorizationData) IDTagInfo() *IDTagInfo {
	if a.idTagInfo == nil {
		return nil
	}

	copiedIDTagInfo := *a.idTagInfo

	return &copiedIDTagInfo
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the AuthorizationData for debugging purposes.
func (a AuthorizationData) String() string {
	result := "AuthorizationData{IDTag: " + a.idTag.String()

	if a.idTagInfo != nil {
		result += ", IDTagInfo: " + a.idTagInfo.String()
	}

	result += "}"

	return result
}

// NewAuthorizationData creates a new AuthorizationData from the given input.
// It validates all fields and returns an error if:
//   - IDTag is empty or invalid
//   - IDTagInfo.Status is invalid (when IDTagInfo is provided)
func NewAuthorizationData(
	input AuthorizationDataInput,
) (AuthorizationData, error) {
	ciString, err := NewCiString20Type(input.IDTag)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+ErrorFieldFormat,
			"IDTag",
			err,
		)
	}

	idToken := NewIDToken(ciString)

	if input.IDTagInfo == nil {
		return AuthorizationData{
			idTag:     idToken,
			idTagInfo: nil,
		}, nil
	}

	idTagInfo, err := buildIDTagInfo(*input.IDTagInfo)
	if err != nil {
		return AuthorizationData{}, fmt.Errorf(
			"NewAuthorizationData: "+ErrorFieldFormat,
			"IDTagInfo",
			err,
		)
	}

	return AuthorizationData{
		idTag:     idToken,
		idTagInfo: &idTagInfo,
	}, nil
}

func buildIDTagInfo(input IDTagInfoInput) (IDTagInfo, error) {
	status := AuthorizationStatus(input.Status)

	idTagInfo, err := NewIDTagInfo(status)
	if err != nil {
		return IDTagInfo{}, fmt.Errorf("buildIDTagInfo: %w", err)
	}

	if input.ExpiryDate != nil {
		expiryDate, err := NewDateTime(*input.ExpiryDate)
		if err != nil {
			return IDTagInfo{}, fmt.Errorf(
				ErrorFieldFormat,
				"ExpiryDate",
				err,
			)
		}

		idTagInfo = idTagInfo.WithExpiryDate(expiryDate)
	}

	if input.ParentIDTag != nil {
		ciString, err := NewCiString20Type(*input.ParentIDTag)
		if err != nil {
			return IDTagInfo{}, fmt.Errorf(
				ErrorFieldFormat,
				"ParentIDTag",
				err,
			)
		}

		parentIDToken := NewIDToken(ciString)
		idTagInfo = idTagInfo.WithParentIDTag(parentIDToken)
	}

	return idTagInfo, nil
}
