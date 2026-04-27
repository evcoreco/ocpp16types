package ocpp16types_test

import (
	"errors"
	"fmt"

	st "github.com/evcoreco/ocpp16types"
)

func ExampleErrEmptyValue() {
	_, err := st.NewCiString20Type("")
	fmt.Println(errors.Is(err, st.ErrEmptyValue))
	// Output:
	// true
}

func ExampleErrInvalidValue() {
	_, err := st.NewInteger(-1)
	fmt.Println(errors.Is(err, st.ErrInvalidValue))
	// Output:
	// true
}
