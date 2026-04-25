package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewDateTime() {
	dateTime, err := st.NewDateTime("2024-01-15T10:30:00Z")
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(dateTime.String())
	// Output:
	// 2024-01-15T10:30:00Z
}

func ExampleNewDateTime_invalid() {
	_, err := st.NewDateTime("not-a-date")
	if err != nil {
		fmt.Println("Error: cannot parse date")

		return
	}
	// Output:
	// Error: cannot parse date
}
