package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleNewListVersionNumber() {
	version, err := st.NewListVersionNumber(5)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(version.String())
	fmt.Println(version.IsEmpty())
	// Output:
	// 5
	// false
}

func ExampleNewListVersionNumber_unsupported() {
	version, err := st.NewListVersionNumber(-1)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(version.IsUnsupported())
	// Output:
	// true
}

func ExampleNewListVersionNumber_empty() {
	version, err := st.NewListVersionNumber(0)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(version.IsEmpty())
	// Output:
	// true
}
