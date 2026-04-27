//go:build bench

package benchmark

import (
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

const (
	sampleValue     = "12345"
	measurandEnergy = "Energy.Active.Import.Register"
	unitWh          = "Wh"
	sampledCount10  = 10
)

func BenchmarkNewSampledValue_Minimal(b *testing.B) {
	b.ReportAllocs()

	input := st.SampledValueInput{
		Value: sampleValue,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewSampledValue_AllOptionals(b *testing.B) {
	b.ReportAllocs()

	context := st.ReadingContextSamplePeriodic.String()
	format := st.ValueFormatRaw.String()
	measurand := st.MeasurandEnergyActiveImportRegister.String()
	phase := st.PhaseL1.String()
	location := st.LocationOutlet.String()
	unit := st.UnitWh.String()

	input := st.SampledValueInput{
		Value:     sampleValue,
		Context:   &context,
		Format:    &format,
		Measurand: &measurand,
		Phase:     &phase,
		Location:  &location,
		Unit:      &unit,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewSampledValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_Single(b *testing.B) {
	b.ReportAllocs()

	input := st.MeterValueInput{
		Timestamp: sampleTimestamp,
		SampledValue: []st.SampledValueInput{
			{Value: sampleValue},
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_ManySampled(b *testing.B) {
	b.ReportAllocs()

	var samples []st.SampledValueInput
	for i := 0; i < sampledCount10; i++ {
		samples = append(
			samples,
			st.SampledValueInput{Value: sampleValue},
		)
	}

	input := st.MeterValueInput{
		Timestamp:    sampleTimestamp,
		SampledValue: samples,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewMeterValue(input); err != nil {
			b.Fatal(err)
		}
	}
}
