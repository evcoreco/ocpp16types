package ocpp16types

import (
	"testing"
)

const (
	testSampledValueStr = "42.5"
	errUnexpectedFmt    = "unexpected error: %v"
)

func strPtr(s string) *string {
	return &s
}

func TestNewSampledValue_MinimalValid(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewSampledValue_AllFields(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   strPtr("Sample.Periodic"),
		Format:    strPtr("Raw"),
		Measurand: strPtr("Voltage"),
		Phase:     strPtr("L1"),
		Location:  strPtr("Outlet"),
		Unit:      strPtr("V"),
	}

	sampledVal, err := NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if sampledVal.Context() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Context()")
	}

	if sampledVal.Format() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Format()")
	}

	if sampledVal.Measurand() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Measurand()")
	}

	if sampledVal.Phase() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Phase()")
	}

	if sampledVal.Location() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Location()")
	}

	if sampledVal.Unit() == nil {
		t.Errorf(ErrorWantNonNil, "SampledValue.Unit()")
	}
}

func TestNewSampledValue_EmptyValue(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     "",
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty Value")
	}
}

func TestNewSampledValue_InvalidContext(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   strPtr(testBogus),
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Context")
	}
}

func TestNewSampledValue_InvalidFormat(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    strPtr(testBogus),
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Format")
	}
}

func TestNewSampledValue_InvalidMeasurand(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: strPtr(testBogus),
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Measurand")
	}
}

func TestNewSampledValue_InvalidPhase(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     strPtr(testBogus),
		Location:  nil,
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Phase")
	}
}

func TestNewSampledValue_InvalidLocation(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  strPtr(testBogus),
		Unit:      nil,
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Location")
	}
}

func TestNewSampledValue_InvalidUnit(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      strPtr(testBogus),
	}

	_, err := NewSampledValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Unit")
	}
}

func TestSampledValue_String_Minimal(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	sampledVal, err := NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "SampledValue{Value: 42.5}"
	if sampledVal.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"SampledValue.String()",
			sampledVal.String(),
			expected,
		)
	}
}

func TestSampledValue_String_AllFields(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   strPtr("Sample.Periodic"),
		Format:    strPtr("Raw"),
		Measurand: strPtr("Voltage"),
		Phase:     strPtr("L1"),
		Location:  strPtr("Outlet"),
		Unit:      strPtr("V"),
	}

	sampledVal, err := NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "SampledValue{Value: 42.5, Context: Sample.Periodic, Format: Raw, Measurand: Voltage, Phase: L1, Location: Outlet, Unit: V}"
	if sampledVal.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"SampledValue.String()",
			sampledVal.String(),
			expected,
		)
	}
}

func TestNewSampledValue_NilOptionals(t *testing.T) {
	t.Parallel()

	input := SampledValueInput{
		Value:     testSampledValueStr,
		Context:   nil,
		Format:    nil,
		Measurand: nil,
		Phase:     nil,
		Location:  nil,
		Unit:      nil,
	}

	sampledVal, err := NewSampledValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if sampledVal.Context() != nil {
		t.Errorf("SampledValue.Context() = %v, want nil", sampledVal.Context())
	}

	if sampledVal.Format() != nil {
		t.Errorf("SampledValue.Format() = %v, want nil", sampledVal.Format())
	}

	if sampledVal.Measurand() != nil {
		t.Errorf("SampledValue.Measurand() = %v, want nil",
			sampledVal.Measurand())
	}

	if sampledVal.Phase() != nil {
		t.Errorf("SampledValue.Phase() = %v, want nil", sampledVal.Phase())
	}

	if sampledVal.Location() != nil {
		t.Errorf("SampledValue.Location() = %v, want nil",
			sampledVal.Location())
	}

	if sampledVal.Unit() != nil {
		t.Errorf("SampledValue.Unit() = %v, want nil", sampledVal.Unit())
	}
}
