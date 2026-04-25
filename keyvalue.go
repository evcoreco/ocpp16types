package ocpp16types

import (
	"errors"
	"fmt"
)

// Compile-time interface verification.
var _ fmt.Stringer = (*KeyValue)(nil)

// KeyValueInput represents the raw input data for creating a KeyValue.
// The constructor NewKeyValue validates all fields automatically.
type KeyValueInput struct {
	// Required: Configuration key name (max 50 chars)
	Key string
	// Required: Whether the configuration value is read-only
	Readonly bool
	// Optional: Configuration value (max 500 chars), nil if key is known but
	// not set
	Value *string
}

// KeyValue represents a configuration key-value pair as defined in OCPP 1.6.
// It is used in GetConfiguration.conf to return configuration settings.
type KeyValue struct {
	key      CiString50Type
	readonly bool
	value    *CiString500Type
}

// NewKeyValue creates a new KeyValue from input.
// Returns an error if:
//   - Key is empty or exceeds 50 characters or contains invalid chars
//   - Value (if provided) exceeds 500 characters or contains invalid chars
func NewKeyValue(input KeyValueInput) (KeyValue, error) {
	var errs []error

	key, err := NewCiString50Type(input.Key)
	if err != nil {
		errs = append(errs, fmt.Errorf("key: %w", err))
	}

	var value *CiString500Type

	if input.Value != nil {
		val, err := NewCiString500Type(*input.Value)
		if err != nil {
			errs = append(errs, fmt.Errorf("value: %w", err))
		} else {
			value = &val
		}
	}

	if len(errs) > errCountZero {
		return KeyValue{}, errors.Join(errs...)
	}

	return KeyValue{
		key:      key,
		readonly: input.Readonly,
		value:    value,
	}, nil
}

// Key returns the configuration key name.
func (k KeyValue) Key() CiString50Type {
	return k.key
}

// Readonly returns whether the configuration value is read-only.
func (k KeyValue) Readonly() bool {
	return k.readonly
}

// Value returns the configuration value, or nil if not set.
func (k KeyValue) Value() *CiString500Type {
	if k.value == nil {
		return nil
	}

	copiedValue := *k.value

	return &copiedValue
}

// String implements the fmt.Stringer interface, returning a human-readable
// representation of the KeyValue for debugging purposes.
func (k KeyValue) String() string {
	result := "KeyValue{Key: " + k.key.String()

	result += fmt.Sprintf(", Readonly: %t", k.readonly)

	if k.value != nil {
		result += ", Value: " + k.value.String()
	}

	result += "}"

	return result
}
