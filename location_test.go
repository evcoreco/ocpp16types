package ocpp16types

import (
	"testing"
)

const (
	locationBodyStr      = "Body"
	locationCableStr     = "Cable"
	locationEVStr        = "EV"
	locationInletStr     = "Inlet"
	locationOutletStr    = "Outlet"
	locationMethodString = "Location.String()"
)

func TestLocation_IsValid_Body(t *testing.T) {
	t.Parallel()

	if !LocationBody.IsValid() {
		t.Errorf(ErrorIsValidFalse, "LocationBody")
	}
}

func TestLocation_IsValid_Cable(t *testing.T) {
	t.Parallel()

	if !LocationCable.IsValid() {
		t.Errorf(ErrorIsValidFalse, "LocationCable")
	}
}

func TestLocation_IsValid_EV(t *testing.T) {
	t.Parallel()

	if !LocationEV.IsValid() {
		t.Errorf(ErrorIsValidFalse, "LocationEV")
	}
}

func TestLocation_IsValid_Inlet(t *testing.T) {
	t.Parallel()

	if !LocationInlet.IsValid() {
		t.Errorf(ErrorIsValidFalse, "LocationInlet")
	}
}

func TestLocation_IsValid_Outlet(t *testing.T) {
	t.Parallel()

	if !LocationOutlet.IsValid() {
		t.Errorf(ErrorIsValidFalse, "LocationOutlet")
	}
}

func TestLocation_IsValid_Empty(t *testing.T) {
	t.Parallel()

	location := Location("")
	if location.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Location(\"\")")
	}
}

func TestLocation_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	location := Location("Unknown")
	if location.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Location(\"Unknown\")")
	}
}

func TestLocation_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	location := Location("body")
	if location.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Location(\"body\")")
	}
}

func TestLocation_String_Body(t *testing.T) {
	t.Parallel()

	got := LocationBody.String()
	if got != locationBodyStr {
		t.Errorf(
			ErrorMethodMismatch,
			locationMethodString,
			got,
			locationBodyStr,
		)
	}
}

func TestLocation_String_Cable(t *testing.T) {
	t.Parallel()

	got := LocationCable.String()
	if got != locationCableStr {
		t.Errorf(
			ErrorMethodMismatch,
			locationMethodString,
			got,
			locationCableStr,
		)
	}
}

func TestLocation_String_EV(t *testing.T) {
	t.Parallel()

	got := LocationEV.String()
	if got != locationEVStr {
		t.Errorf(
			ErrorMethodMismatch,
			locationMethodString,
			got,
			locationEVStr,
		)
	}
}

func TestLocation_String_Inlet(t *testing.T) {
	t.Parallel()

	got := LocationInlet.String()
	if got != locationInletStr {
		t.Errorf(
			ErrorMethodMismatch,
			locationMethodString,
			got,
			locationInletStr,
		)
	}
}

func TestLocation_String_Outlet(t *testing.T) {
	t.Parallel()

	got := LocationOutlet.String()
	if got != locationOutletStr {
		t.Errorf(
			ErrorMethodMismatch,
			locationMethodString,
			got,
			locationOutletStr,
		)
	}
}
