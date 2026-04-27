//go:build fuzz

package testsfuzz

import (
	"reflect"
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

func FuzzChargingSchedule(f *testing.F) {
	f.Add(
		"W", 0, 32.0,
		false, 0, false, 0.0,
		false, "",
	)
	f.Add(
		"A", 100, 16.0,
		true, 3600, true, 6.0,
		true, "2024-01-15T10:30:00Z",
	)
	f.Add(
		"", 0, -1.0,
		false, 0, false, 0.0,
		false, "",
	)
	f.Add(
		"X", -1, 0.0,
		true, -1, true, -1.0,
		true, "bad",
	)

	f.Fuzz(func(
		t *testing.T,
		rateUnit string,
		startPeriod int, limit float64,
		hasDuration bool, duration int,
		hasMinRate bool, minRate float64,
		hasStartSched bool, startSched string,
	) {
		if len(rateUnit) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(startSched) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.ChargingScheduleInput{
			ChargingRateUnit: rateUnit,
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{
					StartPeriod: startPeriod,
					Limit:       limit,
				},
			},
		}

		if hasDuration {
			input.Duration = &duration
		}

		if hasMinRate {
			input.MinChargingRate = &minRate
		}

		if hasStartSched {
			input.StartSchedule = &startSched
		}

		cs, err := st.NewChargingSchedule(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.ChargingSchedule
			if !reflect.DeepEqual(cs, zero) {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		if !cs.ChargingRateUnit().IsValid() {
			t.Fatal(
				"ChargingRateUnit invalid",
			)
		}

		periods := cs.ChargingSchedulePeriod()
		if len(periods) == 0 {
			t.Fatal(
				"no ChargingSchedulePeriod",
			)
		}

		if hasDuration && cs.Duration() == nil {
			t.Fatal("Duration nil after set")
		}

		if hasMinRate && cs.MinChargingRate() == nil {
			t.Fatal("MinChargingRate nil")
		}

		if hasStartSched && cs.StartSchedule() == nil {
			t.Fatal("StartSchedule nil")
		}
	})
}
