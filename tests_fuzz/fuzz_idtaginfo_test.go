//go:build fuzz

package testsfuzz

import (
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

func FuzzIDTagInfo(f *testing.F) {
	f.Add("Accepted", false, "", false, "")
	f.Add("Blocked", false, "", false, "")
	f.Add("Expired", false, "", false, "")
	f.Add("Invalid", false, "", false, "")
	f.Add("ConcurrentTx", false, "", false, "")
	f.Add("", false, "", false, "")
	f.Add("bogus", false, "", false, "")
	f.Add(
		"Accepted",
		true, "2024-01-15T10:30:00Z",
		true, "PARENT01",
	)

	f.Fuzz(func(
		t *testing.T,
		statusStr string,
		hasExpiry bool, expiry string,
		hasParent bool, parent string,
	) {
		if len(statusStr) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(expiry) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(parent) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		status := st.AuthorizationStatus(statusStr)

		info, err := st.NewIDTagInfo(status)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.IDTagInfo
			if info != zero {
				t.Fatal(
					"expected zero value on error",
				)
			}

			return
		}

		if !info.Status().IsValid() {
			t.Fatalf(
				"Status().IsValid() = false for %q",
				status,
			)
		}

		if hasExpiry {
			dt, dtErr := st.NewDateTime(expiry)
			if dtErr == nil {
				info = info.WithExpiryDate(dt)
				if info.ExpiryDate() == nil {
					t.Fatal(
						"ExpiryDate nil after set",
					)
				}
			}
		}

		if hasParent {
			cs, csErr := st.NewCiString20Type(parent)
			if csErr == nil {
				tok := st.NewIDToken(cs)
				info = info.WithParentIDTag(tok)
				if info.ParentIDTag() == nil {
					t.Fatal(
						"ParentIDTag nil after set",
					)
				}
			}
		}
	})
}
