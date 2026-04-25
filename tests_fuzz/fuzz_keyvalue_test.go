//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzKeyValue(f *testing.F) {
	f.Add("MeterInterval", true, false, "")
	f.Add("MeterInterval", false, true, "300")
	f.Add("", false, false, "")
	f.Add("\x00", true, true, "\x00")
	f.Add("key", false, true, "")

	f.Fuzz(func(
		t *testing.T,
		key string,
		readonly bool,
		hasValue bool, value string,
	) {
		if len(key) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(value) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.KeyValueInput{
			Key:      key,
			Readonly: readonly,
		}

		if hasValue {
			input.Value = &value
		}

		kv, err := st.NewKeyValue(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.KeyValue
			if kv != zero {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		if kv.Key().Value() != key {
			t.Fatalf(
				"Key = %q, want %q",
				kv.Key().Value(), key,
			)
		}

		if kv.Readonly() != readonly {
			t.Fatalf(
				"Readonly = %v, want %v",
				kv.Readonly(), readonly,
			)
		}

		if hasValue {
			got := kv.Value()
			if got == nil {
				t.Fatal("Value nil after set")
			}

			if got.Value() != value {
				t.Fatalf(
					"Value = %q, want %q",
					got.Value(), value,
				)
			}
		} else {
			if kv.Value() != nil {
				t.Fatal("Value non-nil")
			}
		}
	})
}
