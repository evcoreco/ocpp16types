//go:build fuzz

package testsfuzz

import (
	"math"
	"strconv"
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

func FuzzListVersionNumber(f *testing.F) {
	f.Add(-1)
	f.Add(0)
	f.Add(1)
	f.Add(math.MinInt32)
	f.Add(math.MaxInt32)
	f.Add(math.MaxInt32 + 1)
	f.Add(math.MinInt32 - 1)

	f.Fuzz(func(t *testing.T, value int) {
		lvn, err := st.NewListVersionNumber(value)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.ListVersionNumber
			if lvn != zero {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		got := int(lvn.Value())
		if got != value {
			t.Fatalf(
				"Value() = %d, want %d",
				got, value,
			)
		}

		if value == -1 && !lvn.IsUnsupported() {
			t.Fatal("IsUnsupported() false")
		}

		if value != -1 && lvn.IsUnsupported() {
			t.Fatal("IsUnsupported() true")
		}

		if value == 0 && !lvn.IsEmpty() {
			t.Fatal("IsEmpty() false")
		}

		if value != 0 && lvn.IsEmpty() {
			t.Fatal("IsEmpty() true")
		}

		want := strconv.FormatInt(
			int64(value), 10,
		)
		if lvn.String() != want {
			t.Fatalf(
				"String() = %q, want %q",
				lvn.String(), want,
			)
		}
	})
}
