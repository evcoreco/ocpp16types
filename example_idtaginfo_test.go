package ocpp16types_test

import (
	"fmt"

	st "github.com/evcoreco/ocpp16types"
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

func ExampleIDTagInfo_WithExpiryDate() {
	info, err := st.NewIDTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(err)

		return
	}

	dt, _ := st.NewDateTime("2027-12-31T23:59:59Z")
	info = info.WithExpiryDate(dt)

	fmt.Println(info.Status().String())
	fmt.Println(info.ExpiryDate().String())
	// Output:
	// Accepted
	// 2027-12-31T23:59:59Z
}

func ExampleIDTagInfo_WithParentIDTag() {
	info, err := st.NewIDTagInfo(st.AuthorizationStatusAccepted)
	if err != nil {
		fmt.Println(err)

		return
	}

	parentStr, _ := st.NewCiString20Type("PARENT-001")
	parent := st.NewIDToken(parentStr)
	info = info.WithParentIDTag(parent)

	fmt.Println(info.ParentIDTag().String())
	// Output:
	// PARENT-001
}
