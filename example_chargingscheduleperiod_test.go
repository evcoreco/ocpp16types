package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleNewChargingSchedulePeriod() {
	phases := 3

	period, err := st.NewChargingSchedulePeriod(st.ChargingSchedulePeriodInput{
		StartPeriod:  0,
		Limit:        32.0,
		NumberPhases: &phases,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(period.StartPeriod().String())
	fmt.Println(period.Limit())
	fmt.Println(period.NumberPhases().String())
	// Output:
	// 0
	// 32
	// 3
}

func ExampleNewChargingSchedulePeriod_minimal() {
	period, err := st.NewChargingSchedulePeriod(st.ChargingSchedulePeriodInput{
		StartPeriod: 900,
		Limit:       16.0,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(period.StartPeriod().String())
	fmt.Println(period.Limit())
	fmt.Println(period.NumberPhases() == nil)
	// Output:
	// 900
	// 16
	// true
}
