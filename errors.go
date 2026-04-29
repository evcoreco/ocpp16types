package ocpp16types

import (
	"errors"
)

// Error message format strings used throughout the types package.
const (
	// ErrorExpectedError is used in tests to format error assertion messages.
	// Parameter: error value to be displayed.
	ErrorExpectedError = "Expected Error: %v"

	// ErrorFieldFormat is the standard format string for wrapping errors
	// with field name context. The first parameter is the field name
	// (string), and the second parameter is the wrapped error. This format
	// enables error chain compatibility using fmt.Errorf with the %w verb.
	ErrorFieldFormat = "%s: %w"

	// ErrorMismatch is used in tests to format expected vs actual value
	// comparison failures. The first parameter is the expected value,
	// and the second parameter is the actual value.
	ErrorMismatch = "Expected %q, got %q"

	// ErrorMismatchValue is used in tests to format expected vs actual value
	// comparison failures using %v formatting. The first parameter is the
	// expected value, and the second parameter is the actual value.
	ErrorMismatchValue = "Expected %v, got %v"

	// ErrorUnexpectedError is used in tests to format unexpected error
	// messages. Parameter: error value that was not expected.
	ErrorUnexpectedError = "Unexpected Error: %v"

	// ErrorWantContains is used in tests to verify error messages contain
	// expected substrings. The first parameter is the actual error, and the
	// second parameter is the expected substring.
	ErrorWantContains = "error = %v, want error containing '%s'"

	// ErrorWantNil is used in tests when an error was expected but nil was
	// returned. Parameter: description of the condition that should have
	// caused an error.
	ErrorWantNil = "error = nil, want error for %s"

	// ErrorWantNonNil is used in tests when a value was expected but nil was
	// returned. Parameter: field name that should have been non-nil.
	ErrorWantNonNil = "%s = nil, want non-nil"

	// ErrorWrapping is used in tests to verify error wrapping behavior.
	// First parameter is the outer error, second parameter is the expected
	// wrapped error.
	ErrorWrapping = "%v, wants wrapping %v"

	// ErrorMethodMismatch is used in tests to format method return value
	// comparison failures. The first parameter is the type/method name,
	// the second parameter is the actual value, and the third parameter
	// is the expected value.
	ErrorMethodMismatch = "%s = %v, want %v"

	// ErrorIsValidTrue is used in tests when IsValid() returned true but
	// false was expected. Parameter: the type name and value tested.
	ErrorIsValidTrue = "%s.IsValid() = true, want false"

	// ErrorIsValidFalse is used in tests when IsValid() returned false but
	// true was expected. Parameter: the type name and value tested.
	ErrorIsValidFalse = "%s.IsValid() = false, want true"
)

var (
	// ErrEmptyValue is the sentinel error for empty-field validation failures.
	//
	// What it means: the input was absent (empty string or nil) when a
	// non-empty value was required by the OCPP 1.6 specification.
	//
	// When to use it: check errors.Is(err, ErrEmptyValue) to detect this
	// specific class of failure and surface a meaningful message to the caller.
	//
	// What it is not: a general "not found" or "nil pointer" error. It is
	// only returned by New* constructors in this package.
	//
	// See also: [ErrInvalidValue] for validation failures where the value was
	// present but did not satisfy specification constraints.
	ErrEmptyValue = errors.New("value cannot be empty")

	// ErrInvalidValue is the sentinel error for constraint-violation failures.
	//
	// What it means: the input was present but failed a specification
	// constraint — wrong format, out of range, or an unrecognized enumeration
	// value.
	//
	// When to use it: check errors.Is(err, ErrInvalidValue) to detect this
	// class of failure independently of the diagnostic message text.
	//
	// What it is not: a parse error or a transport error. It is only returned
	// by New* constructors and IsValid() callers in this package.
	//
	// See also: [ErrEmptyValue] for failures caused by a missing value.
	ErrInvalidValue = errors.New("invalid value")
)
