//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

func FuzzChargingSchedulePeriod(f *testing.F) {
	f.Add(0, 32.0, false, 0)
	f.Add(0, 32.0, true, 1)
	f.Add(0, 32.0, true, 3)
	f.Add(0, 32.0, true, 0)
	f.Add(0, 32.0, true, 4)
	f.Add(-1, 0.0, false, 0)
	f.Add(65536, 0.0, false, 0)
	f.Add(0, -1.0, false, 0)

	f.Fuzz(func(
		t *testing.T,
		startPeriod int,
		limit float64,
		hasPhases bool,
		phases int,
	) {
		input := st.ChargingSchedulePeriodInput{
			StartPeriod: startPeriod,
			Limit:       limit,
		}

		if hasPhases {
			input.NumberPhases = &phases
		}

		period, err := st.NewChargingSchedulePeriod(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.ChargingSchedulePeriod
			if period != zero {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		sp := int(period.StartPeriod().Value())
		if sp != startPeriod {
			t.Fatalf(
				"StartPeriod() = %d, want %d",
				sp, startPeriod,
			)
		}

		if period.Limit() != limit {
			t.Fatalf(
				"Limit() = %f, want %f",
				period.Limit(), limit,
			)
		}

		np := period.NumberPhases()
		if hasPhases && np != nil {
			got := int(np.Value())
			if got != phases {
				t.Fatalf(
					"NumberPhases() = %d, want %d",
					got, phases,
				)
			}
		}
	})
}
