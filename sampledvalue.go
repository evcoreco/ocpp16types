package ocpp16types

import (
	"errors"
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*SampledValue)(nil)

// SampledValueInput represents the raw input data for creating a SampledValue.
type SampledValueInput struct {
	// Required: The measurement value as a string.
	Value string
	// Optional: The reading context.
	Context *string
	// Optional: The value format.
	Format *string
	// Optional: The type of measurement.
	Measurand *string
	// Optional: The phase of the measurement.
	Phase *string
	// Optional: The location of the measurement.
	Location *string
	// Optional: The unit of the measurement.
	Unit *string
}

// SampledValue represents a single sampled value as defined in OCPP 1.6.
type SampledValue struct {
	value     CiString500Type
	context   *ReadingContext
	format    *ValueFormat
	measurand *Measurand
	phase     *Phase
	location  *Location
	unit      *UnitOfMeasure
}

// Value returns the measurement value.
func (s SampledValue) Value() CiString500Type {
	return s.value
}

// Context returns a defensive copy of the reading context, or nil if not set.
func (s SampledValue) Context() *ReadingContext {
	if s.context == nil {
		return nil
	}

	copiedContext := *s.context

	return &copiedContext
}

// Format returns a defensive copy of the value format, or nil if not set.
func (s SampledValue) Format() *ValueFormat {
	if s.format == nil {
		return nil
	}

	copiedFormat := *s.format

	return &copiedFormat
}

// Measurand returns a defensive copy of the measurand, or nil if not set.
func (s SampledValue) Measurand() *Measurand {
	if s.measurand == nil {
		return nil
	}

	copiedMeasurand := *s.measurand

	return &copiedMeasurand
}

// Phase returns a defensive copy of the phase, or nil if not set.
func (s SampledValue) Phase() *Phase {
	if s.phase == nil {
		return nil
	}

	copiedPhase := *s.phase

	return &copiedPhase
}

// Location returns a defensive copy of the location, or nil if not set.
func (s SampledValue) Location() *Location {
	if s.location == nil {
		return nil
	}

	copiedLocation := *s.location

	return &copiedLocation
}

// Unit returns a defensive copy of the unit of measure, or nil if not set.
func (s SampledValue) Unit() *UnitOfMeasure {
	if s.unit == nil {
		return nil
	}

	copiedUnit := *s.unit

	return &copiedUnit
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the SampledValue for debugging purposes.
func (s SampledValue) String() string {
	result := "SampledValue{Value: " + s.value.String()

	if s.context != nil {
		result += ", Context: " + s.context.String()
	}

	if s.format != nil {
		result += ", Format: " + s.format.String()
	}

	if s.measurand != nil {
		result += ", Measurand: " + s.measurand.String()
	}

	if s.phase != nil {
		result += ", Phase: " + s.phase.String()
	}

	if s.location != nil {
		result += ", Location: " + s.location.String()
	}

	if s.unit != nil {
		result += ", Unit: " + s.unit.String()
	}

	result += "}"

	return result
}

// sampledValueValidation holds validated fields during construction.
type sampledValueValidation struct {
	value     CiString500Type
	context   *ReadingContext
	format    *ValueFormat
	measurand *Measurand
	phase     *Phase
	location  *Location
	unit      *UnitOfMeasure
}

// NewSampledValue creates a new SampledValue from the given input.
// It validates all fields and accumulates errors, returning them together.
func NewSampledValue(input SampledValueInput) (SampledValue, error) {
	validated, errs := validateSampledValueInput(input)
	if errs != nil {
		return SampledValue{}, fmt.Errorf(
			"NewSampledValue: %w",
			errors.Join(errs...),
		)
	}

	return SampledValue(validated), nil
}

func validateSampledValueInput(
	input SampledValueInput,
) (sampledValueValidation, []error) {
	var errs []error

	var validated sampledValueValidation

	validated.value, errs = validateSampledValueValue(input.Value, errs)
	validated.context, errs = validateSampledValueContext(input.Context, errs)
	validated.format, errs = validateSampledValueFormat(input.Format, errs)
	validated.measurand, errs = validateSampledValueMeasurand(
		input.Measurand,
		errs,
	)
	validated.phase, errs = validateSampledValuePhase(input.Phase, errs)
	validated.location, errs = validateSampledValueLocation(
		input.Location,
		errs,
	)
	validated.unit, errs = validateSampledValueUnit(input.Unit, errs)

	return validated, errs
}

func validateSampledValueValue(
	value string,
	errs []error,
) (CiString500Type, []error) {
	ciValue, err := NewCiString500Type(value)
	if err != nil {
		return CiString500Type{value: ciString{value: ""}}, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Value", err),
		)
	}

	return ciValue, errs
}

func validateSampledValueContext(
	context *string,
	errs []error,
) (*ReadingContext, []error) {
	if context == nil {
		return nil, errs
	}

	readingCtx := ReadingContext(*context)
	if !readingCtx.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Context", ErrInvalidValue),
		)
	}

	return &readingCtx, errs
}

func validateSampledValueFormat(
	format *string,
	errs []error,
) (*ValueFormat, []error) {
	if format == nil {
		return nil, errs
	}

	valueFormat := ValueFormat(*format)
	if !valueFormat.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Format", ErrInvalidValue),
		)
	}

	return &valueFormat, errs
}

func validateSampledValueMeasurand(
	measurand *string,
	errs []error,
) (*Measurand, []error) {
	if measurand == nil {
		return nil, errs
	}

	measurandVal := Measurand(*measurand)
	if !measurandVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Measurand", ErrInvalidValue),
		)
	}

	return &measurandVal, errs
}

func validateSampledValuePhase(
	phase *string,
	errs []error,
) (*Phase, []error) {
	if phase == nil {
		return nil, errs
	}

	phaseVal := Phase(*phase)
	if !phaseVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Phase", ErrInvalidValue),
		)
	}

	return &phaseVal, errs
}

func validateSampledValueLocation(
	location *string,
	errs []error,
) (*Location, []error) {
	if location == nil {
		return nil, errs
	}

	locationVal := Location(*location)
	if !locationVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Location", ErrInvalidValue),
		)
	}

	return &locationVal, errs
}

func validateSampledValueUnit(
	unit *string,
	errs []error,
) (*UnitOfMeasure, []error) {
	if unit == nil {
		return nil, errs
	}

	unitOfMeasure := UnitOfMeasure(*unit)
	if !unitOfMeasure.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf(ErrorFieldFormat, "Unit", ErrInvalidValue),
		)
	}

	return &unitOfMeasure, errs
}
