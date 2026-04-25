package ocpp16types_test

import (
	"fmt"

	st "github.com/aasanchez/ocpp16types"
)

func ExampleNewAuthorizationData() {
	expiryDate := "2027-12-31T23:59:59Z"

	authData, err := st.NewAuthorizationData(st.AuthorizationDataInput{
		IDTag: "TAG001",
		IDTagInfo: &st.IDTagInfoInput{
			Status:     "Accepted",
			ExpiryDate: &expiryDate,
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(authData.IDTag().Value())
	fmt.Println(authData.IDTagInfo().Status().String())
	// Output:
	// TAG001
	// Accepted
}

func ExampleNewAuthorizationData_deleteEntry() {
	authData, err := st.NewAuthorizationData(st.AuthorizationDataInput{
		IDTag:     "TAG002",
		IDTagInfo: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(authData.IDTag().Value())
	fmt.Println(authData.IDTagInfo() == nil)
	// Output:
	// TAG002
	// true
}
