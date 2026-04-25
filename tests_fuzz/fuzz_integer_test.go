//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzInteger(f *testing.F) {
	f.Add(0)
	f.Add(-1)
	f.Add(1)
	f.Add(65535)
	f.Add(65536)
	f.Add(-100)

	f.Fuzz(func(t *testing.T, i int) {
		integer, err := st.NewInteger(i)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.Integer
			if integer != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		if int(integer.Value()) != i {
			t.Fatalf(
				"Value() = %d, want %d",
				integer.Value(),
				uint16(i),
			)
		}
	})
}
