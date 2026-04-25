//go:build bench

package benchmark

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

const (
	worstCaseSampledCount = 50
)

func BenchmarkNewSampledValue_WorstCase_AllOptionals(
	b *testing.B,
) {
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

func BenchmarkNewChargingProfile_WorstCase_AllOptionals(
	b *testing.B,
) {
	b.ReportAllocs()

	transactionID := profileID
	recurrencyKind := st.RecurrencyKindDaily.String()
	validFrom := sampleTimestamp
	validTo := "2025-01-02T16:00:00Z"

	duration := scheduleDuration
	scheduleStart := sampleTimestamp
	minChargingRate := scheduleMinRate

	periods := []st.ChargingSchedulePeriodInput{
		{
			StartPeriod: periodStartZero,
			Limit:       periodLimitSixteen,
		},
	}

	scheduleInput := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := st.ChargingProfileInput{
		ChargingProfileID:      profileID,
		TransactionID:          &transactionID,
		StackLevel:             stackLevelZero,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind: st.ChargingProfileKindRecurring.String(),
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       scheduleInput,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewChargingProfile(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewMeterValue_WorstCase_50Sampled(
	b *testing.B,
) {
	b.ReportAllocs()

	context := st.ReadingContextSamplePeriodic.String()
	format := st.ValueFormatRaw.String()
	measurand := st.MeasurandEnergyActiveImportRegister.String()
	phase := st.PhaseL1.String()
	location := st.LocationOutlet.String()
	unit := st.UnitWh.String()

	var samples []st.SampledValueInput
	for i := 0; i < worstCaseSampledCount; i++ {
		samples = append(samples, st.SampledValueInput{
			Value:     sampleValue,
			Context:   &context,
			Format:    &format,
			Measurand: &measurand,
			Phase:     &phase,
			Location:  &location,
			Unit:      &unit,
		})
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
