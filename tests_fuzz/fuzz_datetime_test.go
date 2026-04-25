//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzDateTime(f *testing.F) {
	f.Add("")
	f.Add("2024-01-15T10:30:00Z")
	f.Add("2024-01-15T10:30:00+05:00")
	f.Add("2024-01-15T10:30:00.123456789Z")
	f.Add("not-a-date")
	f.Add("2024-13-01T00:00:00Z")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		dt, err := st.NewDateTime(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.DateTime
			if dt != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		str := dt.String()
		if len(str) == 0 {
			t.Fatal(
				"String() returned empty",
			)
		}

		dt2, err := st.NewDateTime(str)
		if err != nil {
			t.Fatalf(
				"round-trip failed: %v",
				err,
			)
		}

		if dt != dt2 {
			t.Fatalf(
				"round-trip mismatch: %v != %v",
				dt, dt2,
			)
		}
	})
}
