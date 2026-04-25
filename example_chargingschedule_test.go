package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewChargingSchedule() {
	duration := 3600
	startSchedule := "2026-06-01T08:00:00Z"
	minRate := 6.0

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:         &duration,
		ChargingRateUnit: "A",
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{StartPeriod: 0, Limit: 32.0},
			{StartPeriod: 1800, Limit: 16.0},
		},
		MinChargingRate: &minRate,
		StartSchedule:   &startSchedule,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(schedule.ChargingRateUnit().String())
	fmt.Println(len(schedule.ChargingSchedulePeriod()))
	fmt.Println(schedule.Duration().String())
	// Output:
	// A
	// 2
	// 3600
}

func ExampleNewChargingSchedule_minimal() {
	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		ChargingRateUnit: "W",
		ChargingSchedulePeriod: []st.ChargingSchedulePeriodInput{
			{StartPeriod: 0, Limit: 7400.0},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(schedule.ChargingRateUnit().String())
	fmt.Println(schedule.Duration() == nil)
	// Output:
	// W
	// true
}
