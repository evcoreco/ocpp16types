package ocpp16types

import (
	"testing"
)

const (
	readingContextInterruptionBeginStr = "Interruption.Begin"
	readingContextInterruptionEndStr   = "Interruption.End"
	readingContextOtherStr             = "Other"
	readingContextSampleClockStr       = "Sample.Clock"
	readingContextSamplePeriodicStr    = "Sample.Periodic"
	readingContextTransactionBeginStr  = "Transaction.Begin"
	readingContextTransactionEndStr    = "Transaction.End"
	readingContextTriggerStr           = "Trigger"
	readingCtxMethodString             = "ReadingContext.String()"
)

func TestReadingContext_IsValid_InterruptionBegin(t *testing.T) {
	t.Parallel()

	if !ReadingContextInterruptionBegin.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextInterruptionBegin")
	}
}

func TestReadingContext_IsValid_InterruptionEnd(t *testing.T) {
	t.Parallel()

	if !ReadingContextInterruptionEnd.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextInterruptionEnd")
	}
}

func TestReadingContext_IsValid_Other(t *testing.T) {
	t.Parallel()

	if !ReadingContextOther.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextOther")
	}
}

func TestReadingContext_IsValid_SampleClock(t *testing.T) {
	t.Parallel()

	if !ReadingContextSampleClock.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextSampleClock")
	}
}

func TestReadingContext_IsValid_SamplePeriodic(t *testing.T) {
	t.Parallel()

	if !ReadingContextSamplePeriodic.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextSamplePeriodic")
	}
}

func TestReadingContext_IsValid_TransactionBegin(t *testing.T) {
	t.Parallel()

	if !ReadingContextTransactionBegin.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextTransactionBegin")
	}
}

func TestReadingContext_IsValid_TransactionEnd(t *testing.T) {
	t.Parallel()

	if !ReadingContextTransactionEnd.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextTransactionEnd")
	}
}

func TestReadingContext_IsValid_Trigger(t *testing.T) {
	t.Parallel()

	if !ReadingContextTrigger.IsValid() {
		t.Errorf(ErrorIsValidFalse, "ReadingContextTrigger")
	}
}

func TestReadingContext_IsValid_Empty(t *testing.T) {
	t.Parallel()

	context := ReadingContext("")
	if context.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ReadingContext(\"\")")
	}
}

func TestReadingContext_IsValid_Unknown(t *testing.T) {
	t.Parallel()

	context := ReadingContext("Unknown")
	if context.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ReadingContext(\"Unknown\")")
	}
}

func TestReadingContext_IsValid_Lowercase(t *testing.T) {
	t.Parallel()

	context := ReadingContext("interruption.begin")
	if context.IsValid() {
		t.Errorf(ErrorIsValidTrue, "ReadingContext(\"interruption.begin\")")
	}
}

func TestReadingContext_String_InterruptionBegin(t *testing.T) {
	t.Parallel()

	got := ReadingContextInterruptionBegin.String()
	if got != readingContextInterruptionBeginStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextInterruptionBeginStr,
		)
	}
}

func TestReadingContext_String_InterruptionEnd(t *testing.T) {
	t.Parallel()

	got := ReadingContextInterruptionEnd.String()
	if got != readingContextInterruptionEndStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextInterruptionEndStr,
		)
	}
}

func TestReadingContext_String_Other(t *testing.T) {
	t.Parallel()

	got := ReadingContextOther.String()
	if got != readingContextOtherStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextOtherStr,
		)
	}
}

func TestReadingContext_String_SampleClock(t *testing.T) {
	t.Parallel()

	got := ReadingContextSampleClock.String()
	if got != readingContextSampleClockStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextSampleClockStr,
		)
	}
}

func TestReadingContext_String_SamplePeriodic(t *testing.T) {
	t.Parallel()

	got := ReadingContextSamplePeriodic.String()
	if got != readingContextSamplePeriodicStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextSamplePeriodicStr,
		)
	}
}

func TestReadingContext_String_TransactionBegin(t *testing.T) {
	t.Parallel()

	got := ReadingContextTransactionBegin.String()
	if got != readingContextTransactionBeginStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTransactionBeginStr,
		)
	}
}

func TestReadingContext_String_TransactionEnd(t *testing.T) {
	t.Parallel()

	got := ReadingContextTransactionEnd.String()
	if got != readingContextTransactionEndStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTransactionEndStr,
		)
	}
}

func TestReadingContext_String_Trigger(t *testing.T) {
	t.Parallel()

	got := ReadingContextTrigger.String()
	if got != readingContextTriggerStr {
		t.Errorf(
			ErrorMethodMismatch,
			readingCtxMethodString,
			got,
			readingContextTriggerStr,
		)
	}
}
