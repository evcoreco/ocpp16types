//go:build fuzz

package testsfuzz

import (
	"reflect"
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

func FuzzChargingProfile(f *testing.F) {
	f.Add(
		1, false, 0, 0,
		"ChargePointMaxProfile", "Absolute",
		false, "",
		false, "", false, "",
		"W", 0, 32.0,
	)
	f.Add(
		100, true, 42, 1,
		"TxProfile", "Recurring",
		true, "Daily",
		true, "2024-01-15T10:30:00Z",
		true, "2024-12-31T23:59:59Z",
		"A", 0, 16.0,
	)
	f.Add(
		-1, false, 0, -1,
		"", "",
		false, "",
		false, "", false, "",
		"", 0, 0.0,
	)
	f.Add(
		65536, true, 65536, 65536,
		"bogus", "bogus",
		true, "bogus",
		true, "bad", true, "bad",
		"X", -1, -1.0,
	)

	f.Fuzz(func(
		t *testing.T,
		profileID int,
		hasTxID bool, txIDint,
		stackLevel int,
		purpose string, kind string,
		hasRecurr bool, recurr string,
		hasValidFrom bool, validFrom string,
		hasValidTo bool, validTo string,
		rateUnit string,
		startPeriod int, limit float64,
	) {
		if len(purpose) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(kind) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(recurr) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(validFrom) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(validTo) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		if len(rateUnit) > maxFuzzStringLen {
			t.Skip("input too large")
		}

		input := st.ChargingProfileInput{
			ChargingProfileID:      profileID,
			StackLevel:             stackLevel,
			ChargingProfilePurpose: purpose,
			ChargingProfileKind:    kind,
			ChargingSchedule: st.ChargingScheduleInput{
				ChargingRateUnit: rateUnit,
				ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
					{
						StartPeriod: startPeriod,
						Limit:       limit,
					},
				},
			},
		}

		if hasTxID {
			input.TransactionID = &txId
		}

		if hasRecurr {
			input.RecurrencyKind = &recurr
		}

		if hasValidFrom {
			input.ValidFrom = &validFrom
		}

		if hasValidTo {
			input.ValidTo = &validTo
		}

		cp, err := st.NewChargingProfile(input)
		if err != nil {
			if !isExpectedError(err) {
				t.Fatalf(
					"unexpected error: %v",
					err,
				)
			}

			var zero st.ChargingProfile
			if !reflect.DeepEqual(cp, zero) {
				t.Fatal(
					"expected zero on error",
				)
			}

			return
		}

		pid := int(cp.ChargingProfileID().Value())
		if pid != profileID {
			t.Fatalf(
				"ProfileId = %d, want %d",
				pid, profileID,
			)
		}

		sl := int(cp.StackLevel().Value())
		if sl != stackLevel {
			t.Fatalf(
				"StackLevel = %d, want %d",
				sl, stackLevel,
			)
		}

		if !cp.ChargingProfilePurpose().IsValid() {
			t.Fatal("Purpose invalid")
		}

		if !cp.ChargingProfileKind().IsValid() {
			t.Fatal("Kind invalid")
		}

		if hasTxID && cp.TransactionID() == nil {
			t.Fatal("TransactionID nil")
		}

		if hasRecurr && cp.RecurrencyKind() == nil {
			t.Fatal("RecurrencyKind nil")
		}

		if hasValidFrom && cp.ValidFrom() == nil {
			t.Fatal("ValidFrom nil")
		}

		if hasValidTo && cp.ValidTo() == nil {
			t.Fatal("ValidTo nil")
		}
	})
}
