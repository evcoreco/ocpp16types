package ocpp16types

import "fmt"

// Compile-time interface verification.
var _ fmt.Stringer = (*ChargingSchedulePeriod)(nil)

const (
	// minNumberPhases is the minimum valid number of phases for charging.
	minNumberPhases = 1
	// maxNumberPhases is the maximum valid number of phases for charging.
	maxNumberPhases = 3
	// limitZero is the minimum valid limit value.
	limitZero = 0
)

// ChargingSchedulePeriodInput represents the raw input data for creating a
// ChargingSchedulePeriod. The constructor NewChargingSchedulePeriod validates
// all fields automatically.
type ChargingSchedulePeriodInput struct {
	// Required: Start of the period relative to schedule start (seconds)
	StartPeriod int
	// Required: Power limit during the schedule period (W or A depending on
	// chargingRateUnit)
	Limit float64
	// Optional: Number of phases (1-3) to use for charging
	NumberPhases *int
}

// ChargingSchedulePeriod represents a period within a charging schedule
// as defined in OCPP 1.6.
type ChargingSchedulePeriod struct {
	startPeriod  Integer
	limit        float64
	numberPhases *Integer
}

// NewChargingSchedulePeriod creates a new ChargingSchedulePeriod from input.
// Returns an error if:
//   - StartPeriod is negative or exceeds uint16 max value (65535)
//   - Limit is negative
//   - NumberPhases (if provided) is not between 1 and 3
func NewChargingSchedulePeriod(
	input ChargingSchedulePeriodInput,
) (ChargingSchedulePeriod, error) {
	startPeriod, err := NewInteger(input.StartPeriod)
	if err != nil {
		return ChargingSchedulePeriod{}, fmt.Errorf("startPeriod: %w", err)
	}

	if input.Limit < limitZero {
		return ChargingSchedulePeriod{}, fmt.Errorf(
			"limit: %w",
			ErrInvalidValue,
		)
	}

	var numberPhases *Integer

	if input.NumberPhases != nil {
		if *input.NumberPhases < minNumberPhases || *input.NumberPhases > maxNumberPhases {
			return ChargingSchedulePeriod{}, fmt.Errorf("numberPhases: %w", ErrInvalidValue)
		}

		// NewInteger cannot fail here: values 1-3 are always valid for uint16
		np, _ := NewInteger(*input.NumberPhases)
		numberPhases = &np
	}

	return ChargingSchedulePeriod{
		startPeriod:  startPeriod,
		limit:        input.Limit,
		numberPhases: numberPhases,
	}, nil
}

// StartPeriod returns the start period in seconds relative to schedule start.
func (c ChargingSchedulePeriod) StartPeriod() Integer {
	return c.startPeriod
}

// Limit returns the power limit during this period.
func (c ChargingSchedulePeriod) Limit() float64 {
	return c.limit
}

// NumberPhases returns the number of phases to use for charging, or nil if not
// specified.
func (c ChargingSchedulePeriod) NumberPhases() *Integer {
	if c.numberPhases == nil {
		return nil
	}

	copiedNumberPhases := *c.numberPhases

	return &copiedNumberPhases
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the ChargingSchedulePeriod for debugging purposes.
func (c ChargingSchedulePeriod) String() string {
	result := "ChargingSchedulePeriod{StartPeriod: " + c.startPeriod.String()

	result += fmt.Sprintf(", Limit: %g", c.limit)

	if c.numberPhases != nil {
		result += ", NumberPhases: " + c.numberPhases.String()
	}

	result += "}"

	return result
}
