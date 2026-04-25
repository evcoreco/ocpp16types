package ocpp16types

import (
	"testing"
)

const (
	testKeyName  = "MyKey"
	testKeyValue = "someValue"
)

func TestNewKeyValue_ValidWithValue(t *testing.T) {
	t.Parallel()

	val := testKeyValue
	input := KeyValueInput{
		Key:      testKeyName,
		Readonly: true,
		Value:    &val,
	}

	keyVal, err := NewKeyValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if keyVal.Key().Value() != testKeyName {
		t.Errorf(
			ErrorMethodMismatch,
			"Key",
			keyVal.Key().Value(),
			testKeyName,
		)
	}

	if !keyVal.Readonly() {
		t.Errorf(
			ErrorMethodMismatch,
			"Readonly",
			keyVal.Readonly(),
			true,
		)
	}

	if keyVal.Value() == nil {
		t.Fatalf(ErrorWantNonNil, "Value")
	}

	if keyVal.Value().Value() != testKeyValue {
		t.Errorf(
			ErrorMethodMismatch,
			"Value",
			keyVal.Value().Value(),
			testKeyValue,
		)
	}
}

func TestNewKeyValue_ValidNilValue(t *testing.T) {
	t.Parallel()

	input := KeyValueInput{
		Key:      testKeyName,
		Readonly: false,
		Value:    nil,
	}

	keyVal, err := NewKeyValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	if keyVal.Value() != nil {
		t.Errorf(
			"Value() = %v, want nil",
			keyVal.Value(),
		)
	}
}

func TestNewKeyValue_EmptyKey(t *testing.T) {
	t.Parallel()

	input := KeyValueInput{
		Key:      "",
		Readonly: false,
		Value:    nil,
	}

	_, err := NewKeyValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "empty key")
	}
}

func TestNewKeyValue_InvalidValue(t *testing.T) {
	t.Parallel()

	val := "\x01bad"
	input := KeyValueInput{
		Key:      "GoodKey",
		Readonly: false,
		Value:    &val,
	}

	_, err := NewKeyValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "invalid value")
	}
}

func TestKeyValue_String_WithValue(t *testing.T) {
	t.Parallel()

	val := testKeyValue
	input := KeyValueInput{
		Key:      testKeyName,
		Readonly: true,
		Value:    &val,
	}

	keyVal, err := NewKeyValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "KeyValue{Key: MyKey, Readonly: true, Value: someValue}"
	if keyVal.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"KeyValue.String()",
			keyVal.String(),
			expected,
		)
	}
}

func TestKeyValue_String_NilValue(t *testing.T) {
	t.Parallel()

	input := KeyValueInput{
		Key:      testKeyName,
		Readonly: false,
		Value:    nil,
	}

	keyVal, err := NewKeyValue(input)
	if err != nil {
		t.Fatalf(errUnexpectedFmt, err)
	}

	expected := "KeyValue{Key: MyKey, Readonly: false}"
	if keyVal.String() != expected {
		t.Errorf(
			ErrorMethodMismatch,
			"KeyValue.String()",
			keyVal.String(),
			expected,
		)
	}
}

func TestNewKeyValue_BothErrors(t *testing.T) {
	t.Parallel()

	val := "\x01bad"
	input := KeyValueInput{
		Key:      "",
		Readonly: false,
		Value:    &val,
	}

	_, err := NewKeyValue(input)
	if err == nil {
		t.Fatalf(ErrorWantNil, "both key and value errors")
	}
}
