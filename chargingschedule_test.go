package ocpp16types

import (
	"testing"
)

const (
	testInvalidDuration = -1
	testRateUnitW       = "W"
	testLimitDefault    = 32.0
	testLimitSecondary  = 16.0
	testDurationSeconds = 3600
	testMinChargingRate = 6.0
	testNegativeRate    = -1.0
	testStartPeriodSec  = 900
	testExpectedPeriods = 2
)

func TestNewChargingSchedule_Valid(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	_, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}
}

func TestNewChargingSchedule_InvalidRateUnit(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: "X",
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingRateUnit")
	}
}

func TestNewChargingSchedule_EmptyPeriods(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:               nil,
		ChargingRateUnit:       testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{},
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty ChargingSchedulePeriod")
	}
}

func TestNewChargingSchedule_WithOptionals(t *testing.T) {
	t.Parallel()

	duration := testDurationSeconds
	minRate := testMinChargingRate
	startSchedule := testTimestamp

	input := ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: &minRate,
		StartSchedule:   &startSchedule,
	}

	schedule, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.Duration() == nil {
		t.Errorf(ErrorWantNonNil, "ChargingSchedule.Duration()")
	}

	if schedule.MinChargingRate() == nil {
		t.Errorf(ErrorWantNonNil, "ChargingSchedule.MinChargingRate()")
	}

	if schedule.StartSchedule() == nil {
		t.Errorf(ErrorWantNonNil, "ChargingSchedule.StartSchedule()")
	}
}

func TestNewChargingSchedule_NilOptionals(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.Duration() != nil {
		t.Errorf("ChargingSchedule.Duration() = %v, want nil",
			schedule.Duration())
	}

	if schedule.MinChargingRate() != nil {
		t.Errorf("ChargingSchedule.MinChargingRate() = %v, want nil",
			schedule.MinChargingRate())
	}

	if schedule.StartSchedule() != nil {
		t.Errorf("ChargingSchedule.StartSchedule() = %v, want nil",
			schedule.StartSchedule())
	}
}

func TestNewChargingSchedule_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:               nil,
		ChargingRateUnit:       "X",
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{},
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "multiple validation errors")
	}
}

func TestNewChargingSchedule_InvalidDuration(t *testing.T) {
	t.Parallel()

	duration := testInvalidDuration
	input := ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid Duration")
	}
}

func TestNewChargingSchedule_InvalidStartSchedule(
	t *testing.T,
) {
	t.Parallel()

	badSchedule := testNotADate
	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   &badSchedule,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid StartSchedule")
	}
}

func TestNewChargingSchedule_InvalidMinChargingRate(
	t *testing.T,
) {
	t.Parallel()

	negativeRate := testNegativeRate
	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: &negativeRate,
		StartSchedule:   nil,
	}

	_, err := NewChargingSchedule(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "negative MinChargingRate")
	}
}

func TestChargingSchedule_NilPeriods(t *testing.T) {
	t.Parallel()

	schedule := ChargingSchedule{
		duration:               nil,
		startSchedule:          nil,
		chargingRateUnit:       "",
		chargingSchedulePeriod: nil,
		minChargingRate:        nil,
	}
	if schedule.ChargingSchedulePeriod() != nil {
		t.Error(
			"ChargingSchedulePeriod() should be nil",
		)
	}
}

func TestChargingSchedule_String_Minimal(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingSchedule{RateUnit: W, Periods: [1 items]}"
	if schedule.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedule.String()",
			schedule.String(),
			expected,
		)
	}
}

func TestChargingSchedule_String_AllOptionals(t *testing.T) {
	t.Parallel()

	duration := testDurationSeconds
	minRate := testMinChargingRate
	startSchedule := testTimestamp

	input := ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: &minRate,
		StartSchedule:   &startSchedule,
	}

	schedule, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingSchedule{RateUnit: W, Periods: [1 items], Duration: 3600, StartSchedule: 2024-01-15T10:30:00Z, MinChargingRate: 6}"
	if schedule.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedule.String()",
			schedule.String(),
			expected,
		)
	}
}

func TestChargingSchedule_Getters(t *testing.T) {
	t.Parallel()

	input := ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
			{
				StartPeriod:  testStartPeriodSec,
				Limit:        testLimitSecondary,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}

	schedule, err := NewChargingSchedule(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if schedule.ChargingRateUnit() != ChargingRateUnit(testRateUnitW) {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingSchedule.ChargingRateUnit()",
			schedule.ChargingRateUnit(),
			ChargingRateUnit(testRateUnitW),
		)
	}

	periods := schedule.ChargingSchedulePeriod()
	if len(periods) != testExpectedPeriods {
		t.Errorf(
			ErrorMethodMismatch,
			"len(ChargingSchedule.ChargingSchedulePeriod())",
			len(periods),
			testExpectedPeriods,
		)
	}
}
