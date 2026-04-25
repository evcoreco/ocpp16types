package ocpp16types

import (
	"testing"
)

const (
	unitOfMeasureWhStr      = "Wh"
	unitOfMeasureKWhStr     = "kWh"
	unitOfMeasureWStr       = "W"
	unitOfMeasureKWStr      = "kW"
	unitOfMeasureVAStr      = "VA"
	unitOfMeasureKVAStr     = "kVA"
	unitOfMeasureVARStr     = "var"
	unitOfMeasureKVARStr    = "kvar"
	unitOfMeasureAStr       = "A"
	unitOfMeasureVStr       = "V"
	unitOfMeasureKStr       = "K"
	unitOfMeasureCelsiusStr = "Celsius"
	unitOfMeasurePercentStr = "Percent"
	unitMethodString        = "UnitOfMeasure.String()"
)

func TestUnitOfMeasure_IsValid_Wh(t *testing.T) {
	t.Parallel()

	if !UnitWh.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureWh")
	}
}

func TestUnitOfMeasure_IsValid_KWh(t *testing.T) {
	t.Parallel()

	if !UnitKWh.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureKWh")
	}
}

func TestUnitOfMeasure_IsValid_W(t *testing.T) {
	t.Parallel()

	if !UnitW.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureW")
	}
}

func TestUnitOfMeasure_IsValid_KW(t *testing.T) {
	t.Parallel()

	if !UnitKW.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureKW")
	}
}

func TestUnitOfMeasure_IsValid_VA(t *testing.T) {
	t.Parallel()

	if !UnitVA.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureVA")
	}
}

func TestUnitOfMeasure_IsValid_KVA(t *testing.T) {
	t.Parallel()

	if !UnitKVA.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureKVA")
	}
}

func TestUnitOfMeasure_IsValid_VAR(t *testing.T) {
	t.Parallel()

	if !UnitVar.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureVAR")
	}
}

func TestUnitOfMeasure_IsValid_KVAR(t *testing.T) {
	t.Parallel()

	if !UnitKvar.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureKVAR")
	}
}

func TestUnitOfMeasure_IsValid_A(t *testing.T) {
	t.Parallel()

	if !UnitA.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureA")
	}
}

func TestUnitOfMeasure_IsValid_V(t *testing.T) {
	t.Parallel()

	if !UnitV.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureV")
	}
}

func TestUnitOfMeasure_IsValid_K(t *testing.T) {
	t.Parallel()

	if !UnitK.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureK")
	}
}

func TestUnitOfMeasure_IsValid_Celsius(t *testing.T) {
	t.Parallel()

	if !UnitCelsius.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitOfMeasureCelsius")
	}
}

func TestUnitOfMeasure_IsValid_Percent(t *testing.T) {
	t.Parallel()

	if !UnitPercent.IsValid() {
		t.Errorf(ErrorIsValidFalse, "UnitPercent")
	}
}

func TestUnitOfMeasure_IsValid_Empty(t *testing.T) {
	t.Parallel()

	measure := UnitOfMeasure("")
	if measure.IsValid() {
		t.Errorf(ErrorIsValidTrue, "UnitOfMeasure(\"\")")
	}
}

func TestUnitOfMeasure_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	measure := UnitOfMeasure("Unknown")
	if measure.IsValid() {
		t.Errorf(ErrorIsValidTrue, "UnitOfMeasure(\"Unknown\")")
	}
}

func TestUnitOfMeasure_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	measure := UnitOfMeasure("wh")
	if measure.IsValid() {
		t.Errorf(ErrorIsValidTrue, "UnitOfMeasure(\"wh\")")
	}
}

func TestUnitOfMeasure_String_Wh(t *testing.T) {
	t.Parallel()

	got := UnitWh.String()
	if got != unitOfMeasureWhStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureWhStr,
		)
	}
}

func TestUnitOfMeasure_String_KWh(t *testing.T) {
	t.Parallel()

	got := UnitKWh.String()
	if got != unitOfMeasureKWhStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKWhStr,
		)
	}
}

func TestUnitOfMeasure_String_W(t *testing.T) {
	t.Parallel()

	got := UnitW.String()
	if got != unitOfMeasureWStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureWStr,
		)
	}
}

func TestUnitOfMeasure_String_KW(t *testing.T) {
	t.Parallel()

	got := UnitKW.String()
	if got != unitOfMeasureKWStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKWStr,
		)
	}
}

func TestUnitOfMeasure_String_VA(t *testing.T) {
	t.Parallel()

	got := UnitVA.String()
	if got != unitOfMeasureVAStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVAStr,
		)
	}
}

func TestUnitOfMeasure_String_KVA(t *testing.T) {
	t.Parallel()

	got := UnitKVA.String()
	if got != unitOfMeasureKVAStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKVAStr,
		)
	}
}

func TestUnitOfMeasure_String_VAR(t *testing.T) {
	t.Parallel()

	got := UnitVar.String()
	if got != unitOfMeasureVARStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVARStr,
		)
	}
}

func TestUnitOfMeasure_String_KVAR(t *testing.T) {
	t.Parallel()

	got := UnitKvar.String()
	if got != unitOfMeasureKVARStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKVARStr,
		)
	}
}

func TestUnitOfMeasure_String_A(t *testing.T) {
	t.Parallel()

	got := UnitA.String()
	if got != unitOfMeasureAStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureAStr,
		)
	}
}

func TestUnitOfMeasure_String_V(t *testing.T) {
	t.Parallel()

	got := UnitV.String()
	if got != unitOfMeasureVStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureVStr,
		)
	}
}

func TestUnitOfMeasure_String_K(t *testing.T) {
	t.Parallel()

	got := UnitK.String()
	if got != unitOfMeasureKStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureKStr,
		)
	}
}

func TestUnitOfMeasure_String_Celsius(t *testing.T) {
	t.Parallel()

	got := UnitCelsius.String()
	if got != unitOfMeasureCelsiusStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasureCelsiusStr,
		)
	}
}

func TestUnitOfMeasure_String_Percent(t *testing.T) {
	t.Parallel()

	got := UnitPercent.String()
	if got != unitOfMeasurePercentStr {
		t.Errorf(
			ErrorMethodMismatch,
			unitMethodString,
			got,
			unitOfMeasurePercentStr,
		)
	}
}
