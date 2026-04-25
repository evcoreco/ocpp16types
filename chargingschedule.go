package ocpp16types

import (
	"errors"
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*ChargingSchedule)(nil)

const (
	// minChargingRateZero is the minimum valid charging rate value.
	minChargingRateZero = 0
	// periodsLenZero is the empty periods length.
	periodsLenZero = 0
)

// ChargingScheduleInput represents the raw input data for creating a
// ChargingSchedule. The constructor NewChargingSchedule validates all fields
// automatically.
type ChargingScheduleInput struct {
	// Optional: Duration of the charging schedule in seconds
	Duration *int
	// Required: Unit of measure for charging rate
	ChargingRateUnit string
	// Required: List of charging schedule periods (at least one)
	ChargingSchedulePeriod []ChargingSchedulePeriodInput
	// Optional: Minimum charging rate supported by the EV
	MinChargingRate *float64
	// Optional: Starting point of an absolute schedule (RFC3339 format)
	StartSchedule *string
}

// ChargingSchedule represents a composite charging schedule as defined in
// OCPP 1.6.
type ChargingSchedule struct {
	duration               *Integer
	startSchedule          *DateTime
	chargingRateUnit       ChargingRateUnit
	chargingSchedulePeriod []ChargingSchedulePeriod
	minChargingRate        *float64
}

// NewChargingSchedule creates a new ChargingSchedule from input.
// Returns an error if:
//   - Duration (if provided) is negative or exceeds uint16 max value (65535)
//   - ChargingRateUnit is not a valid value ("W" or "A")
//   - ChargingSchedulePeriod is empty or contains invalid periods
//   - MinChargingRate (if provided) is negative
//   - StartSchedule (if provided) is not a valid RFC3339 timestamp
func NewChargingSchedule(
	input ChargingScheduleInput,
) (ChargingSchedule, error) {
	var errs []error

	var duration *Integer

	if input.Duration != nil {
		d, err := NewInteger(*input.Duration)
		if err != nil {
			errs = append(errs, fmt.Errorf("duration: %w", err))
		} else {
			duration = &d
		}
	}

	var startSchedule *DateTime

	if input.StartSchedule != nil {
		ss, err := NewDateTime(*input.StartSchedule)
		if err != nil {
			errs = append(errs, fmt.Errorf("startSchedule: %w", err))
		} else {
			startSchedule = &ss
		}
	}

	chargingRateUnit, err := validateChargingRateUnit(input.ChargingRateUnit)
	if err != nil {
		errs = append(errs, err)
	}

	periods, periodErrs := validateSchedulePeriods(input.ChargingSchedulePeriod)
	errs = append(errs, periodErrs...)

	var minChargingRate *float64

	if input.MinChargingRate != nil {
		if *input.MinChargingRate < minChargingRateZero {
			errs = append(errs, fmt.Errorf("minChargingRate: %w", ErrInvalidValue))
		} else {
			copiedRate := *input.MinChargingRate
			minChargingRate = &copiedRate
		}
	}

	if len(errs) > errCountZero {
		return ChargingSchedule{}, errors.Join(errs...)
	}

	return ChargingSchedule{
		duration:               duration,
		startSchedule:          startSchedule,
		chargingRateUnit:       chargingRateUnit,
		chargingSchedulePeriod: periods,
		minChargingRate:        minChargingRate,
	}, nil
}

// validateChargingRateUnit validates the charging rate unit field.
func validateChargingRateUnit(unit string) (ChargingRateUnit, error) {
	chargingRateUnit := ChargingRateUnit(unit)
	if !chargingRateUnit.IsValid() {
		return "", fmt.Errorf("chargingRateUnit: %w", ErrInvalidValue)
	}

	return chargingRateUnit, nil
}

// validateSchedulePeriods validates the charging schedule periods.
func validateSchedulePeriods(
	periods []ChargingSchedulePeriodInput,
) ([]ChargingSchedulePeriod, []error) {
	if len(periods) == periodsLenZero {
		return nil, []error{
			fmt.Errorf("chargingSchedulePeriod: %w", ErrEmptyValue),
		}
	}

	var errs []error

	var validPeriods []ChargingSchedulePeriod

	for i, periodInput := range periods {
		period, err := NewChargingSchedulePeriod(periodInput)
		if err != nil {
			errs = append(
				errs,
				fmt.Errorf("chargingSchedulePeriod[%d]: %w", i, err),
			)
		} else {
			validPeriods = append(validPeriods, period)
		}
	}

	return validPeriods, errs
}

// Duration returns the duration in seconds, or nil if not specified.
func (c ChargingSchedule) Duration() *Integer {
	if c.duration == nil {
		return nil
	}

	copiedDuration := *c.duration

	return &copiedDuration
}

// StartSchedule returns the start schedule, or nil if not specified.
func (c ChargingSchedule) StartSchedule() *DateTime {
	if c.startSchedule == nil {
		return nil
	}

	copiedStartSchedule := *c.startSchedule

	return &copiedStartSchedule
}

// ChargingRateUnit returns the unit of measure for the charging rate.
func (c ChargingSchedule) ChargingRateUnit() ChargingRateUnit {
	return c.chargingRateUnit
}

// ChargingSchedulePeriod returns the list of charging schedule periods.
func (c ChargingSchedule) ChargingSchedulePeriod() []ChargingSchedulePeriod {
	if c.chargingSchedulePeriod == nil {
		return nil
	}

	copiedPeriods := make(
		[]ChargingSchedulePeriod,
		len(c.chargingSchedulePeriod),
	)
	copy(copiedPeriods, c.chargingSchedulePeriod)

	return copiedPeriods
}

// MinChargingRate returns the minimum charging rate, or nil if not specified.
func (c ChargingSchedule) MinChargingRate() *float64 {
	if c.minChargingRate == nil {
		return nil
	}

	copiedRate := *c.minChargingRate

	return &copiedRate
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the ChargingSchedule for debugging purposes.
func (c ChargingSchedule) String() string {
	result := "ChargingSchedule{RateUnit: " + c.chargingRateUnit.String()

	result += fmt.Sprintf(
		", Periods: [%d items]",
		len(c.chargingSchedulePeriod),
	)

	if c.duration != nil {
		result += ", Duration: " + c.duration.String()
	}

	if c.startSchedule != nil {
		result += ", StartSchedule: " + c.startSchedule.String()
	}

	if c.minChargingRate != nil {
		result += fmt.Sprintf(", MinChargingRate: %g", *c.minChargingRate)
	}

	result += "}"

	return result
}
