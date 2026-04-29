//go:build bench

package benchmark

import (
	"testing"

	st "github.com/evcoreco/ocpp16types"
)

const (
	authDataTag       = "TAG-1"
	authDataParent    = "PARENT-1"
	profileID         = 1
	stackLevelZero    = 0
	listVersionSample = 42
	keyValueKey       = "HeartbeatInterval"
	keyValueValue     = "60"
)

func BenchmarkNewAuthorizationData_WithIDTagInfo(b *testing.B) {
	b.ReportAllocs()

	expiry := sampleTimestamp
	parentIDTag := authDataParent

	input := st.AuthorizationDataInput{
		IDTag: authDataTag,
		IDTagInfo: &st.IDTagInfoInput{
			Status:      st.AuthorizationStatusAccepted.String(),
			ExpiryDate:  &expiry,
			ParentIDTag: &parentIDTag,
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewAuthorizationData(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewAuthorizationData_WithoutIDTagInfo(
	b *testing.B,
) {
	b.ReportAllocs()

	input := st.AuthorizationDataInput{
		IDTag:     authDataTag,
		IDTagInfo: nil,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewAuthorizationData(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewChargingProfile_AllOptionals(b *testing.B) {
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

func BenchmarkNewKeyValue_WithValue(b *testing.B) {
	b.ReportAllocs()

	value := keyValueValue
	input := st.KeyValueInput{
		Key:      keyValueKey,
		Readonly: false,
		Value:    &value,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewKeyValue_WithoutValue(b *testing.B) {
	b.ReportAllocs()

	input := st.KeyValueInput{
		Key:      keyValueKey,
		Readonly: false,
		Value:    nil,
	}

	for i := 0; i < b.N; i++ {
		if _, err := st.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewListVersionNumber(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := st.NewListVersionNumber(
			listVersionSample,
		); err != nil {
			b.Fatal(err)
		}
	}
}
