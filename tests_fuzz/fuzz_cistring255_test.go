//go:build fuzz

package testsfuzz

import (
	"strings"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzCiString255(f *testing.F) {
	f.Add("")
	f.Add("a")
	f.Add("validstring")
	f.Add(strings.Repeat("x", 255))
	f.Add(strings.Repeat("x", 256))
	f.Add("\x00hidden")
	f.Add("\xff\xfe")

	f.Fuzz(func(t *testing.T, s string) {
		if len(s) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		cs, err := st.NewCiString255Type(s)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.CiString255Type
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
