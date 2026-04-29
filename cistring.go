package ocpp16types

import (
	"errors"
	"fmt"
)

// Compile-time interface verification.
var (
	_ fmt.Stringer = (*CiString20Type)(nil)
	_ fmt.Stringer = (*CiString25Type)(nil)
	_ fmt.Stringer = (*CiString50Type)(nil)
	_ fmt.Stringer = (*CiString255Type)(nil)
	_ fmt.Stringer = (*CiString500Type)(nil)
)

// Maximum length constants for OCPP 1.6 CiString types.
const (
	CiString20Max  = 20
	CiString25Max  = 25
	CiString50Max  = 50
	CiString255Max = 255
	CiString500Max = 500
)

// ASCII printable character range boundaries.
const (
	asciiPrintableMin = 32  // Space character
	asciiPrintableMax = 126 // Tilde character
)

var (
	// errExceedsMaxLength is returned when a CiString value exceeds its
	// maximum length constraint.
	errExceedsMaxLength = errors.New("exceeds maximum length")

	// errInvalidASCII is returned when a CiString value contains characters
	// outside the printable ASCII range (32-126).
	errInvalidASCII = errors.New(
		"CiString: value contains non-printable ASCII characters",
	)
)

// ciString is the internal implementation of OCPP 1.6 case-insensitive string
// validation. It validates string length and ensures only printable ASCII
// characters (32-126) are present.
type ciString struct {
	value string
}

// newCiString creates a new ciString with the given maximum length constraint.
// Returns an error if the input is empty, exceeds maxLen, or contains
// non-printable ASCII characters.
func newCiString(input string, maxLen int) (ciString, error) {
	if input == "" {
		return ciString{}, ErrEmptyValue
	}

	cis := ciString{value: input}

	if len(cis.value) > maxLen {
		return ciString{}, fmt.Errorf(
			"%w: %w (len=%d, max=%d)",
			ErrInvalidValue,
			errExceedsMaxLength,
			len(cis.value),
			maxLen,
		)
	}

	for _, r := range input {
		if r < asciiPrintableMin || r > asciiPrintableMax {
			return ciString{}, fmt.Errorf(
				"%w: %w",
				ErrInvalidValue,
				errInvalidASCII,
			)
		}
	}

	return cis, nil
}

// val returns the underlying string value.
func (cis ciString) val() string {
	return cis.value
}

// CiString20Type is an OCPP 1.6 case-insensitive string with a maximum length
// of 20 characters.
//
// What it means: a bounded text field that OCPP 1.6 appendix 4 defines as
// CiString20Type — printable ASCII only, max 20 characters.
//
// When to use it: identifier tokens, ICCID, IMSI, and any other OCPP field
// the specification types as CiString20.
//
// What it is not: a general-purpose string. Empty values, Unicode characters,
// and strings longer than 20 characters are rejected at construction. The
// "case-insensitive" label is a specification convention; this type does not
// perform case folding.
//
// See also: [IDToken] wraps CiString20Type; [NewCiString25Type],
// [NewCiString50Type], [NewCiString255Type], [NewCiString500Type] for other
// length variants; [ErrEmptyValue] and [ErrInvalidValue] for error classes.
type CiString20Type struct {
	value ciString
}

// NewCiString20Type creates a new CiString20Type from a string value. Returns
// an error if the value exceeds 20 characters or contains non-printable ASCII.
func NewCiString20Type(value string) (CiString20Type, error) {
	cs, err := newCiString(value, CiString20Max)

	return CiString20Type{value: cs}, err
}

// Value returns the underlying string value of the CiString20Type.
func (c CiString20Type) Value() string {
	return c.value.val()
}

// String implements the fmt.Stringer interface, returning the string value.
func (c CiString20Type) String() string {
	return c.value.val()
}

// CiString25Type is an OCPP 1.6 case-insensitive string with a maximum length
// of 25 characters.
//
// What it means: a bounded text field that OCPP 1.6 appendix 4 defines as
// CiString25Type — printable ASCII only, max 25 characters.
//
// When to use it: charge point serial numbers, meter type, meter serial
// number, and any other OCPP field the specification types as CiString25.
//
// What it is not: a general-purpose string. Empty values, Unicode characters,
// and strings longer than 25 characters are rejected at construction.
//
// See also: [CiString20Type] for the 20-character variant; [CiString50Type],
// [CiString255Type], [CiString500Type] for longer variants.
type CiString25Type struct {
	value ciString
}

// NewCiString25Type creates a new CiString25Type from a string value. Returns
// an error if the value exceeds 25 characters or contains non-printable ASCII.
func NewCiString25Type(value string) (CiString25Type, error) {
	cs, err := newCiString(value, CiString25Max)

	return CiString25Type{value: cs}, err
}

// Value returns the underlying string value of the CiString25Type.
func (c CiString25Type) Value() string {
	return c.value.val()
}

// String implements the fmt.Stringer interface, returning the string value.
func (c CiString25Type) String() string {
	return c.value.val()
}

// CiString50Type is an OCPP 1.6 case-insensitive string with a maximum length
// of 50 characters.
//
// What it means: a bounded text field that OCPP 1.6 appendix 4 defines as
// CiString50Type — printable ASCII only, max 50 characters.
//
// When to use it: firmware version, configuration keys, and any other OCPP
// field the specification types as CiString50.
//
// What it is not: a general-purpose string. Empty values, Unicode characters,
// and strings longer than 50 characters are rejected at construction.
//
// See also: [KeyValue] uses CiString50Type for its key field; [CiString20Type],
// [CiString25Type] for shorter variants; [CiString255Type], [CiString500Type]
// for longer variants.
type CiString50Type struct {
	value ciString
}

// NewCiString50Type creates a new CiString50Type from a string value. Returns
// an error if the value exceeds 50 characters or contains non-printable ASCII.
func NewCiString50Type(value string) (CiString50Type, error) {
	cs, err := newCiString(value, CiString50Max)

	return CiString50Type{value: cs}, err
}

// Value returns the underlying string value of the CiString50Type.
func (c CiString50Type) Value() string {
	return c.value.val()
}

// String implements the fmt.Stringer interface, returning the string value.
func (c CiString50Type) String() string {
	return c.value.val()
}

// CiString255Type is an OCPP 1.6 case-insensitive string with a maximum length
// of 255 characters.
//
// What it means: a bounded text field that OCPP 1.6 appendix 4 defines as
// CiString255Type — printable ASCII only, max 255 characters.
//
// When to use it: any OCPP field the specification types as CiString255 —
// typically longer descriptive or diagnostic text fields.
//
// What it is not: a general-purpose string. Empty values, Unicode characters,
// and strings longer than 255 characters are rejected at construction.
//
// See also: [CiString500Type] for the 500-character variant; [CiString50Type]
// for shorter text fields.
type CiString255Type struct {
	value ciString
}

// NewCiString255Type creates a new CiString255Type from a string value. Returns
// an error if the value exceeds 255 characters or contains non-printable ASCII.
func NewCiString255Type(value string) (CiString255Type, error) {
	cs, err := newCiString(value, CiString255Max)

	return CiString255Type{value: cs}, err
}

// Value returns the underlying string value of the CiString255Type.
func (c CiString255Type) Value() string {
	return c.value.val()
}

// String implements the fmt.Stringer interface, returning the string value.
func (c CiString255Type) String() string {
	return c.value.val()
}

// CiString500Type is an OCPP 1.6 case-insensitive string with a maximum length
// of 500 characters.
//
// What it means: a bounded text field that OCPP 1.6 appendix 4 defines as
// CiString500Type — printable ASCII only, max 500 characters.
//
// When to use it: configuration values, raw meter reading strings, and any
// other OCPP field the specification types as CiString500.
//
// What it is not: a general-purpose string. Empty values, Unicode characters,
// and strings longer than 500 characters are rejected at construction.
//
// See also: [KeyValue] uses CiString500Type for its value field; [SampledValue]
// uses CiString500Type for the raw measurement string; [CiString255Type] for
// the 255-character variant.
type CiString500Type struct {
	value ciString
}

// NewCiString500Type creates a new CiString500Type from a string value. Returns
// an error if the value exceeds 500 characters or contains non-printable ASCII.
func NewCiString500Type(value string) (CiString500Type, error) {
	cis, err := newCiString(value, CiString500Max)

	return CiString500Type{value: cis}, err
}

// Value returns the underlying string value of the CiString500Type.
func (c CiString500Type) Value() string {
	return c.value.val()
}

// String implements the fmt.Stringer interface, returning the string value.
func (c CiString500Type) String() string {
	return c.value.val()
}
