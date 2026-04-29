package ocpp16types

import (
	"fmt"
	"math"
	"strconv"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*Integer)(nil)

const (
	// decimalBase is the base-10 radix for string formatting.
	decimalBase = 10
	// integerMin is the minimum allowed value for Integer.
	integerMin = 0
)

// Integer is an OCPP 1.6 integer primitive — a validated uint16 in the range
// 0–65535, as defined in OCPP 1.6 appendix 4.
//
// What it means: the unsigned 16-bit integer type used throughout OCPP 1.6
// for fields that are constrained to non-negative whole numbers within the
// uint16 range.
//
// When to use it: stack levels, connector identifiers, and any other OCPP
// field the specification types as integer.
//
// What it is not: a general Go int. Negative values and values above 65535
// are rejected at construction. It is not a counter or sequence number
// managed by this package.
//
// See also: [ChargingSchedulePeriod] uses Integer for StartPeriod;
// [ChargingProfile] uses Integer for StackLevel and TransactionID.
type Integer struct {
	value uint16
}

// NewInteger creates a new Integer from an int, validating range.
func NewInteger(value int) (Integer, error) {
	if value < integerMin || value > math.MaxUint16 {
		return Integer{}, fmt.Errorf(
			"NewInteger: "+ErrorFieldFormat,
			"value",
			fmt.Errorf("%w: %d out of range (0-65535)", ErrInvalidValue, value),
		)
	}

	return Integer{value: uint16(value)}, nil
}

// Value returns the underlying uint16 value of the Integer.
func (integer Integer) Value() uint16 {
	return integer.value
}

// String implements the fmt.Stringer interface, returning the decimal string
// representation of the integer value.
func (integer Integer) String() string {
	return strconv.FormatUint(uint64(integer.value), decimalBase)
}
