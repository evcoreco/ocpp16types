package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewIDTagInfo() {
	info, err := st.NewIDTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(info.Status().String())
	// Output:
	// Accepted
}

func ExampleNewIDTagInfo_invalid() {
	_, err := st.NewIDTagInfo(st.AuthorizationStatus("Bogus"))
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// NewIDTagInfo: AuthorizationStatus: invalid value
}
