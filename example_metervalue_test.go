package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleNewMeterValue() {
	mv, err := st.NewMeterValue(st.MeterValueInput{
		Timestamp: "2026-01-15T10:30:00Z",
		SampledValue: []st.SampledValueInput{
			{Value: "1500.0"},
			{Value: "230.1"},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(mv.Timestamp().String())
	fmt.Println(len(mv.SampledValue()))
	fmt.Println(mv.SampledValue()[0].Value().Value())
	// Output:
	// 2026-01-15T10:30:00Z
	// 2
	// 1500.0
}

func ExampleNewMeterValue_invalid() {
	_, err := st.NewMeterValue(st.MeterValueInput{
		Timestamp:    "2026-01-15T10:30:00Z",
		SampledValue: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// NewMeterValue: SampledValue: value cannot be empty
}
