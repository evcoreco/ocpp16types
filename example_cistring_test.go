package ocpp16types_test

import (
	"fmt"
	"strings"

	st "github.com/aasanchez/ocpp16types"
)

const (
	ciString20MaxLen  = 20
	ciString20OverLen = 21
)

func ExampleNewCiString20Type() {
	input := strings.Repeat("A", ciString20MaxLen)

	cistr, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(cistr.Value())
	// Output:
	// AAAAAAAAAAAAAAAAAAAA
}

func ExampleNewCiString20Type_invalid() {
	input := strings.Repeat("B", ciString20OverLen)

	_, err := st.NewCiString20Type(input)
	if err != nil {
		fmt.Println(err)

		return
	}
	// Output:
	// invalid value: exceeds maximum length (len=21, max=20)
}
