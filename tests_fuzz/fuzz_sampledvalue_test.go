//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzSampledValue(f *testing.F) {
	f.Add(
		"42.5",
		false, "", false, "",
		false, "", false, "",
		false, "", false, "",
	)
	f.Add(
		"100",
		true, "Sample.Periodic",
		true, "Raw",
		true, "Energy.Active.Import.Register",
		true, "L1",
		true, "Outlet",
		true, "Wh",
	)
	f.Add(
		"",
		false, "", false, "",
		false, "", false, "",
		false, "", false, "",
	)
	f.Add(
		"\x00",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
		true, "bogus",
	)

	f.Fuzz(func(
		t *testing.T,
		value string,
		hasCtx bool, ctx string,
		hasFmt bool, fmtVal string,
		hasMeas bool, meas string,
		hasPhase bool, phase string,
		hasLoc bool, loc string,
		hasUnit bool, unit string,
	) {
		if len(value) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.SampledValueInput{
			Value: value,
		}

		if hasCtx {
			input.Context = &ctx
		}

		if hasFmt {
			input.Format = &fmtVal
		}

		if hasMeas {
			input.Measurand = &meas
		}

		if hasPhase {
			input.Phase = &phase
		}

		if hasLoc {
			input.Location = &loc
		}

		if hasUnit {
			input.Unit = &unit
		}

		sv, err := st.NewSampledValue(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.SampledValue
			if sv != zero {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		if sv.Value().Value() != value {
			t.Fatalf(
				"Value = %q, want %q",
				sv.Value().Value(), value,
			)
		}

		if hasCtx && sv.Context() != nil {
			if !sv.Context().IsValid() {
				t.Fatal("Context invalid")
			}
		}

		if hasFmt && sv.Format() != nil {
			if !sv.Format().IsValid() {
				t.Fatal("Format invalid")
			}
		}

		if hasMeas && sv.Measurand() != nil {
			if !sv.Measurand().IsValid() {
				t.Fatal("Measurand invalid")
			}
		}

		if hasPhase && sv.Phase() != nil {
			if !sv.Phase().IsValid() {
				t.Fatal("Phase invalid")
			}
		}

		if hasLoc && sv.Location() != nil {
			if !sv.Location().IsValid() {
				t.Fatal("Location invalid")
			}
		}

		if hasUnit && sv.Unit() != nil {
			if !sv.Unit().IsValid() {
				t.Fatal("Unit invalid")
			}
		}
	})
}
