package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewKeyValue() {
	value := "3600"

	kv, err := st.NewKeyValue(st.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(kv.Key().Value())
	fmt.Println(kv.Readonly())
	fmt.Println(kv.Value().Value())
	// Output:
	// HeartbeatInterval
	// false
	// 3600
}

func ExampleNewKeyValue_readonlyNoValue() {
	kv, err := st.NewKeyValue(st.KeyValueInput{
		Key:      "ChargePointModel",
		Readonly: true,
		Value:    nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(kv.Key().Value())
	fmt.Println(kv.Readonly())
	fmt.Println(kv.Value() == nil)
	// Output:
	// ChargePointModel
	// true
	// true
}
