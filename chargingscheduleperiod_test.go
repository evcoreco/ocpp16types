package ocpp16types

import (
	"testing"
)

const (
	testPhases             = 3
	testInvalidStartPeriod = -1
	testInvalidPhases      = 4
	testStartPeriod        = 60
	testStartPeriodZero    = 0
	testNegativeLimit      = -1.0
)

func intPtr(i int) *int {
	return &i
}

func TestNewChargingSchedulePeriod_Valid(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	_, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewChargingSchedulePeriod_WithPhases(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: intPtr(testPhases),
	}

	csp, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.NumberPhases() == nil {
		t.Errorf(ErrorWantNonNil, "ChargingSchedulePeriod.NumberPhases()")
	}
}

func TestNewChargingSchedulePeriod_InvalidStartPeriod(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testInvalidStartPeriod,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	_, err := NewChargingSchedulePeriod(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid StartPeriod")
	}
}

func TestNewChargingSchedulePeriod_InvalidPhases(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: intPtr(testInvalidPhases),
	}

	_, err := NewChargingSchedulePeriod(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid NumberPhases")
	}
}

func TestNewChargingSchedulePeriod_NilPhases(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	csp, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.NumberPhases() != nil {
		t.Errorf(
			"ChargingSchedulePeriod.NumberPhases() = %v, want nil",
			csp.NumberPhases(),
		)
	}
}

func TestNewChargingSchedulePeriod_NegativeLimit(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testNegativeLimit,
		NumberPhases: nil,
	}

	_, err := NewChargingSchedulePeriod(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "negative Limit")
	}
}

func TestChargingSchedulePeriod_String(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: nil,
	}

	csp, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingSchedulePeriod{StartPeriod: 0, Limit: 32}"
	if csp.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedulePeriod.String()",
			csp.String(),
			expected,
		)
	}
}

func TestChargingSchedulePeriod_String_WithPhases(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriodZero,
		Limit:        testLimitDefault,
		NumberPhases: intPtr(testPhases),
	}

	csp, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingSchedulePeriod{StartPeriod: 0, Limit: 32, NumberPhases: 3}"
	if csp.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedulePeriod.String()",
			csp.String(),
			expected,
		)
	}
}

func TestChargingSchedulePeriod_Getters(t *testing.T) {
	t.Parallel()

	input := ChargingSchedulePeriodInput{
		StartPeriod:  testStartPeriod,
		Limit:        testLimitSecondary,
		NumberPhases: nil,
	}

	csp, err := NewChargingSchedulePeriod(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if csp.StartPeriod().Value() != testStartPeriod {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedulePeriod.StartPeriod().Value()",
			csp.StartPeriod().Value(),
			testStartPeriod,
		)
	}

	if csp.Limit() != testLimitSecondary {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedulePeriod.Limit()",
			csp.Limit(),
			testLimitSecondary,
		)
	}
}
