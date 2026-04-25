//go:build fuzz

package testsfuzz

import (
	"reflect"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzMeterValue(f *testing.F) {
	f.Add(
		"2024-01-15T10:30:00Z", "42.5",
		false, "", false, "",
		false, "", false, "",
		false, "", false, "",
	)
	f.Add(
		"2024-01-15T10:30:00Z", "100",
		true, "Sample.Periodic",
		true, "Raw",
		true, "Energy.Active.Import.Register",
		true, "L1",
		true, "Outlet",
		true, "Wh",
	)
	f.Add(
		"", "",
		false, "", false, "",
		false, "", false, "",
		false, "", false, "",
	)
	f.Add(
		"bad", "\x00",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
	)

	f.Fuzz(func(
		t *testing.T,
		timestamp string, value string,
		hasCtx bool, ctx string,
		hasFmt bool, fmt string,
		hasMeas bool, meas string,
		hasPhase bool, phase string,
		hasLoc bool, loc string,
		hasUnit bool, unit string,
	) {
		if len(timestamp) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(value) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		svInput := st.SampledValueInput{
			Value: value,
		}

		if hasCtx {
			svInput.Context = &ctx
		}

		if hasFmt {
			svInput.Format = &fmt
		}

		if hasMeas {
			svInput.Measurand = &meas
		}

		if hasPhase {
			svInput.Phase = &phase
		}

		if hasLoc {
			svInput.Location = &loc
		}

		if hasUnit {
			svInput.Unit = &unit
		}

		input := st.MeterValueInput{
			Timestamp: timestamp,
			SampledValue: []st.SampledValueInput{
				svInput,
			},
		}

		mv, err := st.NewMeterValue(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.MeterValue
			if !reflect.DeepEqual(mv, zero) {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		ts := mv.Timestamp().String()
		if len(ts) == 0 {
			t.Fatal(
				"Timestamp().String() empty",
			)
		}

		if len(mv.SampledValue()) == 0 {
			t.Fatal("SampledValue() empty")
		}
	})
}
