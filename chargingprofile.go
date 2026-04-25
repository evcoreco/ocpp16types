package ocpp16types

import (
	"errors"
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*ChargingProfile)(nil)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
	// stackLevelMin is the minimum valid stack level.
	stackLevelMin = 0
)

// ChargingProfileInput represents the raw input data for creating a
// ChargingProfile. The constructor NewChargingProfile validates all fields
// automatically.
type ChargingProfileInput struct {
	// Required: Unique identifier of this profile
	ChargingProfileID int
	// Optional: Only valid if ChargingProfilePurpose is TxProfile
	TransactionID *int
	// Required: Value determining level in hierarchy stack of profiles
	StackLevel int
	// Required: Defines the purpose of the schedule
	ChargingProfilePurpose string
	// Required: Indicates the kind of schedule
	ChargingProfileKind string
	// Optional: Only for Recurring profiles
	RecurrencyKind *string
	// Optional: Point in time at which the profile starts to be valid
	// (RFC3339 format)
	ValidFrom *string
	// Optional: Point in time at which the profile stops to be valid
	// (RFC3339 format)
	ValidTo *string
	// Required: Contains limits for the available power or current over time
	ChargingSchedule ChargingScheduleInput
}

// ChargingProfile represents a charging profile as defined in OCPP 1.6.
type ChargingProfile struct {
	chargingProfileID      Integer
	transactionID          *Integer
	stackLevel             Integer
	chargingProfilePurpose ChargingProfilePurposeType
	chargingProfileKind    ChargingProfileKindType
	recurrencyKind         *RecurrencyKindType
	validFrom              *DateTime
	validTo                *DateTime
	chargingSchedule       ChargingSchedule
}

// NewChargingProfile creates a new ChargingProfile from input.
// Returns an error if:
//   - ChargingProfileID is negative or exceeds uint16 max value (65535)
//   - TransactionID (if provided) is negative or exceeds uint16 max (65535)
//   - StackLevel is negative or exceeds uint16 max value (65535)
//   - ChargingProfilePurpose is not a valid value
//   - ChargingProfileKind is not a valid value
//   - RecurrencyKind (if provided) is not a valid value
//   - ValidFrom (if provided) is not a valid RFC3339 timestamp
//   - ValidTo (if provided) is not a valid RFC3339 timestamp
//   - ChargingSchedule is invalid
func NewChargingProfile(input ChargingProfileInput) (ChargingProfile, error) {
	var errs []error

	chargingProfileID, err := validateChargingProfileID(input.ChargingProfileID)
	if err != nil {
		errs = append(errs, err)
	}

	var transactionID *Integer

	if input.TransactionID != nil {
		txID, err := NewInteger(*input.TransactionID)
		if err != nil {
			errs = append(errs, fmt.Errorf("transactionID: %w", err))
		} else {
			transactionID = &txID
		}
	}

	stackLevel, err := validateStackLevel(input.StackLevel)
	if err != nil {
		errs = append(errs, err)
	}

	purpose, err := validateChargingProfilePurpose(input.ChargingProfilePurpose)
	if err != nil {
		errs = append(errs, err)
	}

	kind, err := validateChargingProfileKind(input.ChargingProfileKind)
	if err != nil {
		errs = append(errs, err)
	}

	var recurrencyKind *RecurrencyKindType

	if input.RecurrencyKind != nil {
		rk := RecurrencyKindType(*input.RecurrencyKind)
		if !rk.IsValid() {
			errs = append(errs, fmt.Errorf("recurrencyKind: %w", ErrInvalidValue))
		} else {
			recurrencyKind = &rk
		}
	}

	var validFrom *DateTime

	if input.ValidFrom != nil {
		vf, err := NewDateTime(*input.ValidFrom)
		if err != nil {
			errs = append(errs, fmt.Errorf("validFrom: %w", err))
		} else {
			validFrom = &vf
		}
	}

	var validTo *DateTime

	if input.ValidTo != nil {
		vt, err := NewDateTime(*input.ValidTo)
		if err != nil {
			errs = append(errs, fmt.Errorf("validTo: %w", err))
		} else {
			validTo = &vt
		}
	}

	chargingSchedule, err := NewChargingSchedule(input.ChargingSchedule)
	if err != nil {
		errs = append(errs, fmt.Errorf("chargingSchedule: %w", err))
	}

	if len(errs) > errCountZero {
		return ChargingProfile{}, errors.Join(errs...)
	}

	return ChargingProfile{
		chargingProfileID:      chargingProfileID,
		transactionID:          transactionID,
		stackLevel:             stackLevel,
		chargingProfilePurpose: purpose,
		chargingProfileKind:    kind,
		recurrencyKind:         recurrencyKind,
		validFrom:              validFrom,
		validTo:                validTo,
		chargingSchedule:       chargingSchedule,
	}, nil
}

// validateChargingProfileID validates the charging profile ID.
func validateChargingProfileID(id int) (Integer, error) {
	profileID, err := NewInteger(id)
	if err != nil {
		return Integer{}, fmt.Errorf("chargingProfileID: %w", err)
	}

	return profileID, nil
}

// validateStackLevel validates the stack level.
func validateStackLevel(level int) (Integer, error) {
	if level < stackLevelMin {
		return Integer{}, fmt.Errorf("stackLevel: %w", ErrInvalidValue)
	}

	stackLevel, err := NewInteger(level)
	if err != nil {
		return Integer{}, fmt.Errorf("stackLevel: %w", err)
	}

	return stackLevel, nil
}

// validateChargingProfilePurpose validates the charging profile purpose.
func validateChargingProfilePurpose(
	purpose string,
) (ChargingProfilePurposeType, error) {
	purposeType := ChargingProfilePurposeType(purpose)
	if !purposeType.IsValid() {
		return "", fmt.Errorf("chargingProfilePurpose: %w", ErrInvalidValue)
	}

	return purposeType, nil
}

// validateChargingProfileKind validates the charging profile kind.
func validateChargingProfileKind(kind string) (ChargingProfileKindType, error) {
	kindType := ChargingProfileKindType(kind)
	if !kindType.IsValid() {
		return "", fmt.Errorf("chargingProfileKind: %w", ErrInvalidValue)
	}

	return kindType, nil
}

// ChargingProfileID returns the unique identifier of this profile.
func (c ChargingProfile) ChargingProfileID() Integer {
	return c.chargingProfileID
}

// TransactionID returns the transaction ID, or nil if not specified.
func (c ChargingProfile) TransactionID() *Integer {
	if c.transactionID == nil {
		return nil
	}

	copiedTransactionID := *c.transactionID

	return &copiedTransactionID
}

// StackLevel returns the stack level of this profile.
func (c ChargingProfile) StackLevel() Integer {
	return c.stackLevel
}

// ChargingProfilePurpose returns the purpose of this profile.
func (c ChargingProfile) ChargingProfilePurpose() ChargingProfilePurposeType {
	return c.chargingProfilePurpose
}

// ChargingProfileKind returns the kind of this profile.
func (c ChargingProfile) ChargingProfileKind() ChargingProfileKindType {
	return c.chargingProfileKind
}

// RecurrencyKind returns the recurrency kind, or nil if not specified.
func (c ChargingProfile) RecurrencyKind() *RecurrencyKindType {
	if c.recurrencyKind == nil {
		return nil
	}

	copiedRecurrencyKind := *c.recurrencyKind

	return &copiedRecurrencyKind
}

// ValidFrom returns the valid from timestamp, or nil if not specified.
func (c ChargingProfile) ValidFrom() *DateTime {
	if c.validFrom == nil {
		return nil
	}

	copiedValidFrom := *c.validFrom

	return &copiedValidFrom
}

// ValidTo returns the valid to timestamp, or nil if not specified.
func (c ChargingProfile) ValidTo() *DateTime {
	if c.validTo == nil {
		return nil
	}

	copiedValidTo := *c.validTo

	return &copiedValidTo
}

// ChargingSchedule returns the charging schedule for this profile.
func (c ChargingProfile) ChargingSchedule() ChargingSchedule {
	return c.chargingSchedule
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the ChargingProfile for debugging purposes.
func (c ChargingProfile) String() string {
	result := "ChargingProfile{Id: " + c.chargingProfileID.String()
	result += ", Purpose: " + c.chargingProfilePurpose.String()
	result += ", Kind: " + c.chargingProfileKind.String()
	result += ", StackLevel: " + c.stackLevel.String()

	if c.transactionID != nil {
		result += ", TransactionId: " + c.transactionID.String()
	}

	if c.recurrencyKind != nil {
		result += ", RecurrencyKind: " + c.recurrencyKind.String()
	}

	if c.validFrom != nil {
		result += ", ValidFrom: " + c.validFrom.String()
	}

	if c.validTo != nil {
		result += ", ValidTo: " + c.validTo.String()
	}

	result += "}"

	return result
}
