//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzAuthorizationData(f *testing.F) {
	f.Add(
		"TAG001",
		false, "", false, "", false, "",
	)
	f.Add(
		"TAG001",
		true, "Accepted", false, "", false, "",
	)
	f.Add(
		"TAG001",
		true, "Accepted",
		true, "2024-01-15T10:30:00Z",
		true, "PARENT01",
	)
	f.Add(
		"",
		false, "", false, "", false, "",
	)
	f.Add(
		"TAG001",
		true, "bogus", false, "", false, "",
	)
	f.Add(
		"\x00",
		true, "Accepted", false, "", false, "",
	)

	f.Fuzz(func(
		t *testing.T,
		idTag string,
		hasInfo bool, status string,
		hasExpiry bool, expiry string,
		hasParent bool, parent string,
	) {
		if len(idTag) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(status) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(expiry) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(parent) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.AuthorizationDataInput{
			IDTag: idTag,
		}

		if hasInfo {
			info := &st.IDTagInfoInput{
				Status: status,
			}

			if hasExpiry {
				info.ExpiryDate = &expiry
			}

			if hasParent {
				info.ParentIDTag = &parent
			}

			input.IDTagInfo = info
		}

		ad, err := st.NewAuthorizationData(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.AuthorizationData
			if ad != zero {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		if ad.IDTag().String() != idTag {
			t.Fatalf(
				"IdTag = %q, want %q",
				ad.IDTag().String(), idTag,
			)
		}

		if hasInfo {
			if ad.IDTagInfo() == nil {
				t.Fatal("IDTagInfo nil")
			}

			if !ad.IDTagInfo().Status().IsValid() {
				t.Fatal("Status invalid")
			}

			if hasExpiry &&
				ad.IDTagInfo().ExpiryDate() == nil {
				t.Fatal("ExpiryDate nil")
			}

			if hasParent &&
				ad.IDTagInfo().ParentIDTag() == nil {
				t.Fatal("ParentIDTag nil")
			}
		}
	})
}
