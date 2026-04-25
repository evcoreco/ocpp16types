package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewChargingProfile() {
	recurrency := "Daily"
	validFrom := "2026-06-01T00:00:00Z"

	profile, err := st.NewChargingProfile(st.ChargingProfileInput{
		ChargingProfileID:      1,
		StackLevel:             0,
		ChargingProfilePurpose: "TxDefaultProfile",
		ChargingProfileKind:    "Recurring",
		RecurrencyKind:         &recurrency,
		ValidFrom:              &validFrom,
		ChargingSchedule: st.ChargingScheduleInput{
			ChargingRateUnit: "A",
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{StartPeriod: 0, Limit: 32.0},
				{StartPeriod: 28800, Limit: 16.0},
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(profile.ChargingProfileID().String())
	fmt.Println(profile.ChargingProfilePurpose().String())
	fmt.Println(profile.ChargingProfileKind().String())
	fmt.Println(profile.RecurrencyKind().String())
	// Output:
	// 1
	// TxDefaultProfile
	// Recurring
	// Daily
}

func ExampleNewChargingProfile_txProfile() {
	txID := 42

	profile, err := st.NewChargingProfile(st.ChargingProfileInput{
		ChargingProfileID:      10,
		TransactionID:          &txID,
		StackLevel:             1,
		ChargingProfilePurpose: "TxProfile",
		ChargingProfileKind:    "Relative",
		ChargingSchedule: st.ChargingScheduleInput{
			ChargingRateUnit: "W",
			ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
				{StartPeriod: 0, Limit: 11000.0},
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(profile.TransactionID().String())
	fmt.Println(profile.ChargingProfilePurpose().String())
	// Output:
	// 42
	// TxProfile
}
