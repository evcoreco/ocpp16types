// Package ocpp16types provides strictly validated Go types for the
// Open Charge Point Protocol (OCPP) 1.6 specification.
//
// Every type is validated at construction time: if a constructor returns
// without error, the value is guaranteed to be spec-compliant.
// All types are immutable, thread-safe, and depend only on the Go
// standard library.
//
// # Value types
//
// [CiString20Type], [CiString25Type], [CiString50Type], [CiString255Type],
// and [CiString500Type] are case-insensitive strings restricted to
// printable ASCII (32–126) with their respective maximum lengths.
//
// [DateTime] is an RFC 3339 timestamp that must already be in UTC.
//
// [Integer] is a validated uint16 value in the range 0–65535.
//
// # Enumeration types
//
// [AuthorizationStatus], [AvailabilityStatus], [AvailabilityType],
// [CancelReservationStatus], [ChargePointErrorCode], [ChargePointStatus],
// [ChargingProfileKindType], [ChargingProfilePurposeType],
// [ChargingProfileStatus], [ChargingRateUnit], [ClearCacheStatus],
// [ClearChargingProfileStatus], [ConfigurationStatus],
// [DataTransferStatus], [DiagnosticsStatus], [FirmwareStatus],
// [GetCompositeScheduleStatus], [Location], [Measurand],
// [MessageTrigger], [Phase], [ReadingContext], [RecurrencyKindType],
// [RegistrationStatus], [RemoteStartTransactionStatus],
// [RemoteStopTransactionStatus], [ReservationStatus], [ResetStatus],
// [ResetType], [StopReason], [TriggerMessageStatus], [UnlockStatus],
// [UpdateStatus], [UpdateType], [UnitOfMeasure], and [ValueFormat].
//
// Every enumeration exposes IsValid() bool and String() string methods.
//
// # Composite types
//
// [IDToken] wraps a [CiString20Type] for use as an identifier token.
//
// [IDTagInfo] carries authorization status with optional expiry date
// and parent tag. Use [IDTagInfo.WithExpiryDate] and
// [IDTagInfo.WithParentIDTag] to set optional fields after construction.
//
// [SampledValue] represents a single meter value sample; [MeterValue]
// groups sampled values under a single UTC timestamp.
//
// [AuthorizationData] is an entry in the local authorization list.
//
// [KeyValue] is a configuration key-value pair used in
// GetConfiguration.conf responses.
//
// [ChargingSchedulePeriod], [ChargingSchedule], and [ChargingProfile]
// model the OCPP smart-charging hierarchy.
//
// [ListVersionNumber] tracks the version of the local authorization
// list. The constants [ListVersionUnsupported] (-1) and
// [ListVersionEmpty] (0) have dedicated predicates on the type.
//
// # Errors
//
// All validation failures wrap [ErrInvalidValue] or [ErrEmptyValue],
// enabling errors.Is checks while carrying diagnostic context.
package ocpp16types
