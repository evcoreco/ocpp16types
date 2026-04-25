package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleAuthorizationStatus_IsValid() {
	status := st.AuthorizationStatusAccepted
	fmt.Println(status.IsValid())
	// Output:
	// true
}

func ExampleAuthorizationStatus_String() {
	status := st.AuthorizationStatusAccepted
	fmt.Println(status.String())
	// Output:
	// Accepted
}
