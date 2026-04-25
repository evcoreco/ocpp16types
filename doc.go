// Package ocpp16types provides Open Charge Point Protocol (OCPP) 1.6
// core data types with strict validation for EV charging systems.
//
// This package implements validated types for the OCPP 1.6 protocol including:
//   - CiString types: case-insensitive strings with maximum length constraints
//     (CiString20Type, CiString25Type, CiString50Type,
//     CiString255Type, CiString500Type)
//   - DateTime: RFC3339-compliant timestamps that must already be UTC
//   - Integer: validated uint16 values (0–65535)
//
// Enumeration types:
//   - AuthorizationStatus, ChargePointErrorCode, ChargePointStatus
//   - ChargingProfilePurposeType, ChargingRateUnit
//   - Location, Measurand, Phase, ReadingContext
//   - RegistrationStatus, StopReason, UnitOfMeasure, ValueFormat
//
// Composite types:
//   - IDToken: wrapper around CiString20Type for identifier tokens
//   - IDTagInfo: authorization information with builder
//     pattern for optional fields
//   - SampledValue: single meter value sample with optional context metadata
//   - MeterValue: timestamped meter reading with sampled values
//   - AuthorizationData: entry in the local authorization list
//   - KeyValue: configuration key-value pair
//   - ChargingSchedulePeriod: single period within a charging schedule
//   - ChargingSchedule: schedule for charging with rate unit and periods
//   - ChargingProfile: complete charging profile with schedule and metadata
//   - ListVersionNumber: version number for the local authorization list
//
// All types enforce validation at construction time and are designed for
// thread-safe concurrent use with immutable fields and value receivers.
// Zero external dependencies — uses only Go standard library.
package ocpp16types
