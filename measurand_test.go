package ocpp16types

import (
	"testing"
)

const (
	measurandCurrentExportStr                = "Current.Export"
	measurandCurrentImportStr                = "Current.Import"
	measurandCurrentOfferedStr               = "Current.Offered"
	measurandEnergyActiveExportRegisterStr   = "Energy.Active.Export.Register"
	measurandEnergyActiveImportRegisterStr   = "Energy.Active.Import.Register"
	measurandEnergyReactiveExportRegisterStr = "Energy.Reactive.Export.Register"
	measurandEnergyReactiveImportRegisterStr = "Energy.Reactive.Import.Register"
	measurandEnergyActiveExportIntervalStr   = "Energy.Active.Export.Interval"
	measurandEnergyActiveImportIntervalStr   = "Energy.Active.Import.Interval"
	measurandEnergyReactiveExportIntervalStr = "Energy.Reactive.Export.Interval"
	measurandEnergyReactiveImportIntervalStr = "Energy.Reactive.Import.Interval"
	measurandFrequencyStr                    = "Frequency"
	measurandPowerActiveExportStr            = "Power.Active.Export"
	measurandPowerActiveImportStr            = "Power.Active.Import"
	measurandPowerFactorStr                  = "Power.Factor"
	measurandPowerOfferedStr                 = "Power.Offered"
	measurandPowerReactiveExportStr          = "Power.Reactive.Export"
	measurandPowerReactiveImportStr          = "Power.Reactive.Import"
	measurandRPMStr                          = "RPM"
	measurandSoCStr                          = "SoC"
	measurandTemperatureStr                  = "Temperature"
	measurandVoltageStr                      = "Voltage"
	measurandMethodString                    = "Measurand.String()"
)

func TestMeasurand_IsValid_CurrentExport(t *testing.T) {
	t.Parallel()

	if !MeasurandCurrentExport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandCurrentExport")
	}
}

func TestMeasurand_IsValid_CurrentImport(t *testing.T) {
	t.Parallel()

	if !MeasurandCurrentImport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandCurrentImport")
	}
}

func TestMeasurand_IsValid_CurrentOffered(t *testing.T) {
	t.Parallel()

	if !MeasurandCurrentOffered.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandCurrentOffered")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyActiveExportRegister.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyActiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyActiveImportRegister.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyActiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportRegister(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyReactiveExportRegister.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyReactiveExportRegister")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportRegister(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyReactiveImportRegister.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyReactiveImportRegister")
	}
}

func TestMeasurand_IsValid_EnergyActiveExportInterval(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyActiveExportInterval.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyActiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyActiveImportInterval(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyActiveImportInterval.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyActiveImportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveExportInterval(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyReactiveExportInterval.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyReactiveExportInterval")
	}
}

func TestMeasurand_IsValid_EnergyReactiveImportInterval(t *testing.T) {
	t.Parallel()

	if !MeasurandEnergyReactiveImportInterval.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandEnergyReactiveImportInterval")
	}
}

func TestMeasurand_IsValid_Frequency(t *testing.T) {
	t.Parallel()

	if !MeasurandFrequency.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandFrequency")
	}
}

func TestMeasurand_IsValid_PowerActiveExport(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerActiveExport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerActiveExport")
	}
}

func TestMeasurand_IsValid_PowerActiveImport(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerActiveImport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerActiveImport")
	}
}

func TestMeasurand_IsValid_PowerFactor(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerFactor.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerFactor")
	}
}

func TestMeasurand_IsValid_PowerOffered(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerOffered.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerOffered")
	}
}

func TestMeasurand_IsValid_PowerReactiveExport(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerReactiveExport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerReactiveExport")
	}
}

func TestMeasurand_IsValid_PowerReactiveImport(t *testing.T) {
	t.Parallel()

	if !MeasurandPowerReactiveImport.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandPowerReactiveImport")
	}
}

func TestMeasurand_IsValid_RPM(t *testing.T) {
	t.Parallel()

	if !MeasurandRPM.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandRPM")
	}
}

func TestMeasurand_IsValid_SoC(t *testing.T) {
	t.Parallel()

	if !MeasurandSoC.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandSoC")
	}
}

func TestMeasurand_IsValid_Temperature(t *testing.T) {
	t.Parallel()

	if !MeasurandTemperature.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandTemperature")
	}
}

func TestMeasurand_IsValid_Voltage(t *testing.T) {
	t.Parallel()

	if !MeasurandVoltage.IsValid() {
		t.Errorf(ErrorIsValidFalse, "MeasurandVoltage")
	}
}

func TestMeasurand_IsValid_Empty(t *testing.T) {
	t.Parallel()

	measurand := Measurand("")
	if measurand.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Measurand(\"\")")
	}
}

func TestMeasurand_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	measurand := Measurand("Unknown")
	if measurand.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Measurand(\"Unknown\")")
	}
}

func TestMeasurand_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	measurand := Measurand("current.export")
	if measurand.IsValid() {
		t.Errorf(ErrorIsValidTrue, "Measurand(\"current.export\")")
	}
}

func TestMeasurand_String_CurrentExport(t *testing.T) {
	t.Parallel()

	got := MeasurandCurrentExport.String()
	if got != measurandCurrentExportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentExportStr,
		)
	}
}

func TestMeasurand_String_CurrentImport(t *testing.T) {
	t.Parallel()

	got := MeasurandCurrentImport.String()
	if got != measurandCurrentImportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentImportStr,
		)
	}
}

func TestMeasurand_String_CurrentOffered(t *testing.T) {
	t.Parallel()

	got := MeasurandCurrentOffered.String()
	if got != measurandCurrentOfferedStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandCurrentOfferedStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveExportRegister(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyActiveExportRegister.String()
	if got != measurandEnergyActiveExportRegisterStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveExportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveImportRegister(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyActiveImportRegister.String()
	if got != measurandEnergyActiveImportRegisterStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveImportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveExportRegister(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyReactiveExportRegister.String()
	if got != measurandEnergyReactiveExportRegisterStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveExportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveImportRegister(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyReactiveImportRegister.String()
	if got != measurandEnergyReactiveImportRegisterStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveImportRegisterStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveExportInterval(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyActiveExportInterval.String()
	if got != measurandEnergyActiveExportIntervalStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveExportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyActiveImportInterval(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyActiveImportInterval.String()
	if got != measurandEnergyActiveImportIntervalStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyActiveImportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveExportInterval(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyReactiveExportInterval.String()
	if got != measurandEnergyReactiveExportIntervalStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveExportIntervalStr,
		)
	}
}

func TestMeasurand_String_EnergyReactiveImportInterval(t *testing.T) {
	t.Parallel()

	got := MeasurandEnergyReactiveImportInterval.String()
	if got != measurandEnergyReactiveImportIntervalStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandEnergyReactiveImportIntervalStr,
		)
	}
}

func TestMeasurand_String_Frequency(t *testing.T) {
	t.Parallel()

	got := MeasurandFrequency.String()
	if got != measurandFrequencyStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandFrequencyStr,
		)
	}
}

func TestMeasurand_String_PowerActiveExport(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerActiveExport.String()
	if got != measurandPowerActiveExportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerActiveExportStr,
		)
	}
}

func TestMeasurand_String_PowerActiveImport(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerActiveImport.String()
	if got != measurandPowerActiveImportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerActiveImportStr,
		)
	}
}

func TestMeasurand_String_PowerFactor(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerFactor.String()
	if got != measurandPowerFactorStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerFactorStr,
		)
	}
}

func TestMeasurand_String_PowerOffered(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerOffered.String()
	if got != measurandPowerOfferedStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerOfferedStr,
		)
	}
}

func TestMeasurand_String_PowerReactiveExport(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerReactiveExport.String()
	if got != measurandPowerReactiveExportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerReactiveExportStr,
		)
	}
}

func TestMeasurand_String_PowerReactiveImport(t *testing.T) {
	t.Parallel()

	got := MeasurandPowerReactiveImport.String()
	if got != measurandPowerReactiveImportStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandPowerReactiveImportStr,
		)
	}
}

func TestMeasurand_String_RPM(t *testing.T) {
	t.Parallel()

	got := MeasurandRPM.String()
	if got != measurandRPMStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandRPMStr,
		)
	}
}

func TestMeasurand_String_SoC(t *testing.T) {
	t.Parallel()

	got := MeasurandSoC.String()
	if got != measurandSoCStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandSoCStr,
		)
	}
}

func TestMeasurand_String_Temperature(t *testing.T) {
	t.Parallel()

	got := MeasurandTemperature.String()
	if got != measurandTemperatureStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandTemperatureStr,
		)
	}
}

func TestMeasurand_String_Voltage(t *testing.T) {
	t.Parallel()

	got := MeasurandVoltage.String()
	if got != measurandVoltageStr {
		t.Errorf(
			ErrorMethodMismatch,
			measurandMethodString,
			got,
			measurandVoltageStr,
		)
	}
}
