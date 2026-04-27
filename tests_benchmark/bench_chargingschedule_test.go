//go:build bench

package benchmark

import (
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

const (
	scheduleDuration    = 60
	scheduleMinRate     = 0.0
	periodStartZero     = 0
	periodStartTen      = 10
	periodStartTwenty   = 20
	periodStartThirty   = 30
	periodLimitSixteen  = 16.0
	periodLimitEight    = 8.0
	periodLimitSix      = 6.0
	periodLimitFour     = 4.0
	numberPhasesThree   = 3
)

func BenchmarkNewChargingSchedule(b *testing.B) {
	b.ReportAllocs()

	duration := scheduleDuration
	startSchedule := sampleTimestamp
	minChargingRate := scheduleMinRate

	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: periodStartZero, Limit: periodLimitSixteen},
		{StartPeriod: periodStartTen, Limit: periodLimitEight},
	}

	input := st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &startSchedule,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewChargingSchedule(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewChargingSchedulePeriod_WithNumberPhases(
	b *testing.B,
) {
	b.ReportAllocs()

	phases := numberPhasesThree
	input := st.ChargingSchedulePeriodInput{
		StartPeriod:  periodStartZero,
		Limit:        periodLimitSixteen,
		NumberPhases: &phases,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewChargingSchedulePeriod(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChargingScheduleDurationGetter(b *testing.B) {
	b.ReportAllocs()

	duration := scheduleDuration
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: periodStartZero, Limit: periodLimitSixteen},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = schedule.Duration()
	}
}

func BenchmarkChargingScheduleStartScheduleGetter(b *testing.B) {
	b.ReportAllocs()

	duration := scheduleDuration
	startSchedule := sampleTimestamp
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: periodStartZero, Limit: periodLimitSixteen},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          &startSchedule,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = schedule.StartSchedule()
	}
}

func BenchmarkChargingScheduleMinChargingRateGetter(
	b *testing.B,
) {
	b.ReportAllocs()

	duration := scheduleDuration
	minChargingRate := 1.0
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: periodStartZero, Limit: periodLimitSixteen},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = schedule.MinChargingRate()
	}
}

func BenchmarkChargingSchedulePeriodGetter(b *testing.B) {
	b.ReportAllocs()

	duration := scheduleDuration
	periods := []st.ChargingSchedulePeriodInput{
		{StartPeriod: periodStartZero, Limit: periodLimitSixteen},
		{StartPeriod: periodStartTen, Limit: periodLimitEight},
		{StartPeriod: periodStartTwenty, Limit: periodLimitSix},
		{StartPeriod: periodStartThirty, Limit: periodLimitFour},
	}

	schedule, err := st.NewChargingSchedule(st.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       st.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        nil,
		StartSchedule:          nil,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		values := schedule.ChargingSchedulePeriod()
		_ = values[0].Limit()
	}
}
