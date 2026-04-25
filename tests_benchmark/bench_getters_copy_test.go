//go:build bench

package benchmark

import (
	"testing"

	st "github.com/aasanchez/ocpp16types"
)

var (
	sinkInteger      *st.Integer
	sinkDateTime     *st.DateTime
	sinkFloat64      *float64
	sinkCiString500  *st.CiString500Type
	sinkRecurrency   *st.RecurrencyKindType
	sinkScheduleList []st.ChargingSchedulePeriod
)

func buildChargingProfile(b *testing.B) st.ChargingProfile {
	b.Helper()

	transactionID := profileID
	recurrencyKind := st.RecurrencyKindDaily.String()
	validFrom := sampleTimestamp
	validTo := "2025-01-02T16:00:00Z"

	duration := scheduleDuration
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
		MinChargingRate:        nil,
		StartSchedule:          nil,
	}

	profile, err := st.NewChargingProfile(st.ChargingProfileInput{
		ChargingProfileID:      profileID,
		TransactionID:          &transactionID,
		StackLevel:             stackLevelZero,
		ChargingProfilePurpose: st.TxProfile.String(),
		ChargingProfileKind: st.ChargingProfileKindRecurring.String(),
		RecurrencyKind:         &recurrencyKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       scheduleInput,
	})
	if err != nil {
		b.Fatal(err)
	}

	return profile
}

func BenchmarkChargingProfileTransactionIDGetter_Copy(
	b *testing.B,
) {
	b.ReportAllocs()

	profile := buildChargingProfile(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkInteger = profile.TransactionID()
	}
}

func BenchmarkChargingProfileRecurrencyKindGetter_Copy(
	b *testing.B,
) {
	b.ReportAllocs()

	profile := buildChargingProfile(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkRecurrency = profile.RecurrencyKind()
	}
}

func BenchmarkChargingProfileValidFromGetter_Copy(
	b *testing.B,
) {
	b.ReportAllocs()

	profile := buildChargingProfile(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkDateTime = profile.ValidFrom()
	}
}

func BenchmarkChargingProfileValidToGetter_Copy(
	b *testing.B,
) {
	b.ReportAllocs()

	profile := buildChargingProfile(b)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkDateTime = profile.ValidTo()
	}
}

func BenchmarkKeyValueValueGetter_Copy(b *testing.B) {
	b.ReportAllocs()

	value := keyValueValue
	keyValue, err := st.NewKeyValue(st.KeyValueInput{
		Key:      keyValueKey,
		Readonly: false,
		Value:    &value,
	})
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkCiString500 = keyValue.Value()
	}
}
