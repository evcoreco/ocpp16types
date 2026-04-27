package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleNewIDToken() {
	cistr, err := st.NewCiString20Type("RFID-TAG-123")
	if err != nil {
		fmt.Println(err)

		return
	}

	token := st.NewIDToken(cistr)
	fmt.Println(token.String())
	// Output:
	// RFID-TAG-123
}
