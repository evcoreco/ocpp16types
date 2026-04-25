//go:build fuzz

package testsfuzz

import (
	"errors"
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const maxFuzzStringLen = 10000

func isExpectedError(err error) bool {
	return errors.Is(err, st.ErrEmptyValue) ||
		errors.Is(err, st.ErrInvalidValue)
}

func FuzzCiString20(f *testing.F) {
	f.Add("")
	f.Add("a")
	f.Add("validstring")
	f.Add(strings.Repeat("x", 20))
	f.Add(strings.Repeat("x", 21))
	f.Add("\x00hidden")
	f.Add("\xff\xfe")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString20Type(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.CiString20Type
			if cs != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		if cs.Value() != s {
			t.Fatalf(
				"Value() = %q, want %q",
				cs.Value(), s,
			)
		}

		if cs.String() != s {
			t.Fatalf(
				"String() = %q, want %q",
				cs.String(), s,
			)
		}
	})
}

func FuzzCiString25(f *testing.F) {
	f.Add("")
	f.Add("a")
	f.Add("validstring")
	f.Add(strings.Repeat("x", 25))
	f.Add(strings.Repeat("x", 26))
	f.Add("\x00hidden")
	f.Add("\xff\xfe")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString25Type(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.CiString25Type
			if cs != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		if cs.Value() != s {
			t.Fatalf(
				"Value() = %q, want %q",
				cs.Value(), s,
			)
		}

		if cs.String() != s {
			t.Fatalf(
				"String() = %q, want %q",
				cs.String(), s,
			)
		}
	})
}

func FuzzCiString50(f *testing.F) {
	f.Add("")
	f.Add("a")
	f.Add("validstring")
	f.Add(strings.Repeat("x", 50))
	f.Add(strings.Repeat("x", 51))
	f.Add("\x00hidden")
	f.Add("\xff\xfe")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString50Type(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.CiString50Type
			if cs != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		if cs.Value() != s {
			t.Fatalf(
				"Value() = %q, want %q",
				cs.Value(), s,
			)
		}

		if cs.String() != s {
			t.Fatalf(
				"String() = %q, want %q",
				cs.String(), s,
			)
		}
	})
}
