package ocpp16types

import (
	"testing"
)

const (
	testProfileID       = 1
	testStackLevel      = 0
	testInvalidNegative = -1
	testPurpose         = "TxDefaultProfile"
	testKind            = "Absolute"
	testRecurrency      = "Daily"
	testValidFromStr    = "2024-01-01T00:00:00Z"
	testValidToStr      = "2024-12-31T23:59:59Z"
	testTxID            = 42
	testTxIDValid       = 100
	testStackOverflow   = 70000
	testBogus           = "Bogus"
	testTransactionID   = "TransactionID"
)

func validScheduleInput() ChargingScheduleInput {
	return ChargingScheduleInput{
		Duration:         nil,
		ChargingRateUnit: testRateUnitW,
		ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
			{
				StartPeriod:  testStartPeriodZero,
				Limit:        testLimitDefault,
				NumberPhases: nil,
			},
		},
		MinChargingRate: nil,
		StartSchedule:   nil,
	}
}

func TestNewChargingProfile_AllValid(t *testing.T) {
	t.Parallel()

	txID := testTxID
	recurrKind := testRecurrency
	validFrom := testValidFromStr
	validTo := testValidToStr

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          &txID,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.ChargingProfileID().Value() !=
		uint16(testProfileID) {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfileID",
			profile.ChargingProfileID().Value(),
			testProfileID,
		)
	}

	if profile.TransactionID() == nil {
		t.Errorf(
			ErrorWantNonNil,
			testTransactionID,
		)
	}

	if profile.RecurrencyKind() == nil {
		t.Errorf(ErrorWantNonNil, "RecurrencyKind")
	}

	if profile.ValidFrom() == nil {
		t.Errorf(ErrorWantNonNil, "ValidFrom")
	}

	if profile.ValidTo() == nil {
		t.Errorf(ErrorWantNonNil, "ValidTo")
	}
}

func TestNewChargingProfile_RequiredOnly(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.TransactionID() != nil {
		t.Errorf(
			"TransactionID() = %v, want nil",
			profile.TransactionID(),
		)
	}

	if profile.RecurrencyKind() != nil {
		t.Errorf(
			"RecurrencyKind() = %v, want nil",
			profile.RecurrencyKind(),
		)
	}

	if profile.ValidFrom() != nil {
		t.Errorf(
			"ValidFrom() = %v, want nil",
			profile.ValidFrom(),
		)
	}

	if profile.ValidTo() != nil {
		t.Errorf(
			"ValidTo() = %v, want nil",
			profile.ValidTo(),
		)
	}
}

func TestNewChargingProfile_InvalidProfileID(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testInvalidNegative,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfileID")
	}
}

func TestNewChargingProfile_InvalidStackLevel(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testInvalidNegative,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "negative StackLevel")
	}
}

func TestNewChargingProfile_InvalidPurpose(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testBogus,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfilePurpose")
	}
}

func TestNewChargingProfile_InvalidKind(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testBogus,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingProfileKind")
	}
}

func TestNewChargingProfile_InvalidRecurrencyKind(
	t *testing.T,
) {
	t.Parallel()

	recurrKind := testBogus
	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid RecurrencyKind")
	}
}

func TestNewChargingProfile_InvalidValidFrom(t *testing.T) {
	t.Parallel()

	validFrom := testNotADate
	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              &validFrom,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ValidFrom")
	}
}

func TestNewChargingProfile_InvalidValidTo(t *testing.T) {
	t.Parallel()

	validTo := testNotADate
	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ValidTo")
	}
}

func TestNewChargingProfile_InvalidSchedule(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
			StartSchedule:          nil,
		},
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid ChargingSchedule")
	}
}

func TestNewChargingProfile_ValidTransactionID(
	t *testing.T,
) {
	t.Parallel()

	txID := testTxIDValid
	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          &txID,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.TransactionID() == nil {
		t.Fatalf(
			ErrorWantNonNil,
			testTransactionID,
		)
	}

	if profile.TransactionID().Value() !=
		uint16(txID) {
		t.Errorf(
			ErrorMethodMismatch,
			testTransactionID,
			profile.TransactionID().Value(),
			txID,
		)
	}
}

func TestNewChargingProfile_InvalidTransactionID(
	t *testing.T,
) {
	t.Parallel()

	txID := testInvalidNegative
	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          &txID,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid TransactionID")
	}
}

func TestNewChargingProfile_StackLevelOverflow(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackOverflow,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "StackLevel overflow")
	}
}

func TestNewChargingProfile_InvalidSchedulePeriod(
	t *testing.T,
) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule: ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: testRateUnitW,
			ChargingSchedulePeriod: []ChargingSchedulePeriodInput{
				{
					StartPeriod:  testInvalidNegative,
					Limit:        testLimitDefault,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	}

	_, err := NewChargingProfile(input)
	if err == nil {
		t.Fatalf(
			ErrorWantNil,
			"invalid schedule period",
		)
	}
}

func TestChargingProfile_String_RequiredOnly(t *testing.T) {
	t.Parallel()

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          nil,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingProfile{Id: 1, Purpose: TxDefaultProfile, Kind: Absolute, StackLevel: 0}"
	if profile.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfile.String()",
			profile.String(),
			expected,
		)
	}
}

func TestChargingProfile_String_AllOptionals(t *testing.T) {
	t.Parallel()

	txID := testTxID
	recurrKind := testRecurrency
	validFrom := testValidFromStr
	validTo := testValidToStr

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          &txID,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "ChargingProfile{Id: 1, Purpose: TxDefaultProfile, Kind: Absolute, StackLevel: 0, TransactionId: 42, RecurrencyKind: Daily, ValidFrom: 2024-01-01T00:00:00Z, ValidTo: 2024-12-31T23:59:59Z}"
	if profile.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfile.String()",
			profile.String(),
			expected,
		)
	}
}

func TestChargingProfile_Getters(t *testing.T) {
	t.Parallel()

	txID := testTxID
	recurrKind := testRecurrency
	validFrom := testValidFromStr
	validTo := testValidToStr

	input := ChargingProfileInput{
		ChargingProfileID:      testProfileID,
		TransactionID:          &txID,
		StackLevel:             testStackLevel,
		ChargingProfilePurpose: testPurpose,
		ChargingProfileKind:    testKind,
		RecurrencyKind:         &recurrKind,
		ValidFrom:              &validFrom,
		ValidTo:                &validTo,
		ChargingSchedule:       validScheduleInput(),
	}

	profile, err := NewChargingProfile(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if profile.StackLevel().Value() !=
		uint16(testStackLevel) {
		t.Errorf(
			ErrorMethodMismatch,
			"StackLevel",
			profile.StackLevel().Value(),
			testStackLevel,
		)
	}

	purpose := ChargingProfilePurposeType(testPurpose)
	if profile.ChargingProfilePurpose() != purpose {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfilePurpose",
			profile.ChargingProfilePurpose(),
			purpose,
		)
	}

	kind := ChargingProfileKindType(testKind)
	if profile.ChargingProfileKind() != kind {
		t.Errorf(
			ErrorMethodMismatch,
			"ChargingProfileKind",
			profile.ChargingProfileKind(),
			kind,
		)
	}

	_ = profile.ChargingSchedule()
}
