package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleNewSampledValue() {
	context := "Sample.Periodic"
	measurand := "Energy.Active.Import.Register"
	location := "Outlet"
	unit := "Wh"

	sv, err := st.NewSampledValue(st.SampledValueInput{
		Value:     "1234.5",
		Context:   &context,
		Measurand: &measurand,
		Location:  &location,
		Unit:      &unit,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(sv.Value().Value())
	fmt.Println(sv.Measurand().String())
	fmt.Println(sv.Unit().String())
	// Output:
	// 1234.5
	// Energy.Active.Import.Register
	// Wh
}

func ExampleNewSampledValue_minimal() {
	sv, err := st.NewSampledValue(st.SampledValueInput{
		Value: "42",
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(sv.Value().Value())
	fmt.Println(sv.Context() == nil)
	// Output:
	// 42
	// true
}

func ExampleNewSampledValue_invalid() {
	_, err := st.NewSampledValue(st.SampledValueInput{
		Value: "",
	})
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// NewSampledValue: Value: value cannot be empty
}
