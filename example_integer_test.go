package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

const exampleIntegerValue = 42

func ExampleNewInteger() {
	num, err := st.NewInteger(exampleIntegerValue)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(num.String())
	// Output:
	// 42
}

func ExampleNewInteger_invalid() {
	_, err := st.NewInteger(-1)
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// NewInteger: value: invalid value: -1 out of range (0-65535)
}
