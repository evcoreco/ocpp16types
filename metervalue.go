package ocpp16types

import (
	"errors"
	"fmt"
	"time"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*MeterValue)(nil)

const (
	// meterValueErrCountZero is the empty error count.
	meterValueErrCountZero = 0
)

// MeterValueInput represents the raw input data for creating a MeterValue.
type MeterValueInput struct {
	// Required: The timestamp of the measurement in RFC3339 format.
	Timestamp string
	// Required: One or more sampled values.
	SampledValue []SampledValueInput
}

// MeterValue represents a collection of sampled values at a specific timestamp
// as defined in OCPP 1.6.
type MeterValue struct {
	timestamp    DateTime
	sampledValue []SampledValue
}

// Timestamp returns the measurement timestamp.
func (m MeterValue) Timestamp() DateTime {
	return m.timestamp
}

// SampledValue returns a defensive copy of the sampled values.
func (m MeterValue) SampledValue() []SampledValue {
	if m.sampledValue == nil {
		return nil
	}

	copiedValues := make([]SampledValue, len(m.sampledValue))
	copy(copiedValues, m.sampledValue)

	return copiedValues
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the MeterValue for debugging purposes.
func (m MeterValue) String() string {
	result := "MeterValue{Timestamp: " + m.timestamp.String()

	result += fmt.Sprintf(", SampledValue: [%d items]", len(m.sampledValue))

	result += "}"

	return result
}

// meterValueValidation holds validated fields during construction.
type meterValueValidation struct {
	timestamp    DateTime
	sampledValue []SampledValue
}

// NewMeterValue creates a new MeterValue from the given input.
// It validates all fields and accumulates errors, returning them together.
func NewMeterValue(input MeterValueInput) (MeterValue, error) {
	validated, errs := validateMeterValueInput(input)
	if errs != nil {
		return MeterValue{}, fmt.Errorf(
			"NewMeterValue: %w",
			errors.Join(errs...),
		)
	}

	return MeterValue(validated), nil
}

func validateMeterValueInput(
	input MeterValueInput,
) (meterValueValidation, []error) {
	var errs []error

	var validated meterValueValidation

	validated.timestamp, errs = validateMeterValueTimestamp(
		input.Timestamp,
		errs,
	)
	validated.sampledValue, errs = validateMeterValueSampledValues(
		input.SampledValue,
		errs,
	)

	return validated, errs
}

func validateMeterValueTimestamp(
	timestamp string,
	errs []error,
) (DateTime, []error) {
	dateTime, err := NewDateTime(timestamp)
	if err != nil {
		return DateTime{value: time.Time{}}, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Timestamp", err),
		)
	}

	return dateTime, errs
}

func validateMeterValueSampledValues(
	sampledValues []SampledValueInput,
	errs []error,
) ([]SampledValue, []error) {
	if len(sampledValues) == meterValueErrCountZero {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "SampledValue", ErrEmptyValue),
		)
	}

	var validValues []SampledValue

	for i, svInput := range sampledValues {
		sampledValue, err := NewSampledValue(svInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("sampledValue[%d]: %w", i, err))
		} else {
			validValues = append(validValues, sampledValue)
		}
	}

	return validValues, errs
}
