// Package ocpp16types provides strictly validated Go types for the
// Open Charge Point Protocol (OCPP) 1.6 specification.
//
// # What this package means
//
// ocpp16types is the single source of truth for OCPP 1.6 data types in the
// EVCore ecosystem. Every type is validated at construction time through a New*
// constructor. A value that passes construction is guaranteed to be
// spec-compliant, immutable, and thread-safe. The package has zero external
// dependencies beyond the Go standard library.
//
// # When to use this package
//
// Import ocpp16types when building OCPP 1.6 message payloads, validating
// incoming field values against the specification, or sharing type-safe OCPP
// values across multiple modules. The standard import alias across the EVCore
// ecosystem is:
//
//	import types "github.com/evcoreco/ocpp16types"
//
// # What this package is not
//
// ocpp16types is not a message layer, a transport implementation, or a JSON
// codec. It does not send or receive OCPP messages, parse WebSocket frames,
// or route calls. It defines, validates, and exposes data types only.
//
// # Where to look for adjacent concepts
//
//   - OCPP 1.6 message structs and call constructors: github.com/evcoreco/ocpp16messages
//   - OCPP 1.6 specification: Open Charge Alliance (openchargealliance.org)
//   - Sentinel errors for validation failures: [ErrEmptyValue], [ErrInvalidValue]
//
// # Value types
//
// [CiString20Type], [CiString25Type], [CiString50Type], [CiString255Type],
// and [CiString500Type] are case-insensitive strings restricted to printable
// ASCII (32–126) with bounded lengths as defined in OCPP 1.6 appendix 4.
//
// What they mean: a string field that the protocol treats as case-insensitive
// and limits to a specific maximum character count.
//
// When to use them: wherever OCPP 1.6 specifies a CiString field — vendor IDs,
// firmware versions, identifier tokens, configuration keys, and similar
// bounded text values.
//
// What they are not: general-purpose strings. Empty values, Unicode characters,
// and values that exceed the maximum length are all rejected at construction.
// They do not perform case folding; the "case-insensitive" label is a
// specification convention, not a runtime behavior of these types.
//
// See also: [IDToken] wraps [CiString20Type]; [KeyValue] uses [CiString50Type]
// for its key and [CiString500Type] for its value; [SampledValue] uses
// [CiString500Type] for the raw measurement string.
//
// # DateTime
//
// [DateTime] is an RFC 3339 UTC timestamp.
//
// What it means: a moment in time expressed as a UTC RFC 3339 string, as
// required by the OCPP 1.6 specification for all date-time fields.
//
// When to use it: expiry dates, transaction start and stop times, meter
// reading timestamps, and any other OCPP field typed as dateTime.
//
// What it is not: a timezone-aware or local-time value. Non-UTC offsets are
// rejected at construction. It is also not a duration or an interval.
//
// See also: [MeterValue] embeds a [DateTime] timestamp; [IDTagInfo.WithExpiryDate]
// accepts a [DateTime]; [ChargingProfile] uses optional [DateTime] fields for
// ValidFrom and ValidTo.
//
// # Integer
//
// [Integer] is a validated uint16 in the range 0–65535.
//
// What it means: the OCPP 1.6 integer primitive (appendix 4), which maps to an
// unsigned 16-bit range.
//
// When to use it: stack levels, connector identifiers, and other OCPP fields
// typed as integer in the specification.
//
// What it is not: a general Go int. Negative values and values above 65535 are
// rejected. It is not a counter or a sequence number managed by this package.
//
// See also: [ChargingSchedulePeriod] uses [Integer] for StartPeriod;
// [ChargingProfile] uses [Integer] for StackLevel and TransactionID.
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
// What they mean: typed string constants that represent the closed set of
// allowed values for a specific OCPP 1.6 field. Each enumeration is the
// direct Go mapping of the corresponding OCPP 1.6 enumeration.
//
// When to use them: assign an enumeration constant (e.g.
// [AuthorizationStatusAccepted]) to any OCPP field that the specification
// describes as an enum. Call IsValid() before accepting an enum value received
// from an external source.
//
// What they are not: open-ended string fields. An unrecognized string value
// fails IsValid() even if it resembles a valid value. They are not extensible
// without a specification change.
//
// See also: composite types [IDTagInfo], [SampledValue], and [ChargingProfile]
// embed enumeration types as required or optional fields.
//
// # Composite types
//
// [IDToken] wraps a [CiString20Type] for use as an identifier token.
//
// What it means: the OCPP 1.6 idToken data type — an identifier that a charge
// point sends to the Central System for authorization.
//
// When to use it: Authorize.req, StartTransaction.req, StopTransaction.req,
// and any OCPP message that carries an identifier token.
//
// What it is not: an authorization result. The result of checking a token is
// [IDTagInfo], not [IDToken].
//
// See also: [IDTagInfo] carries the authorization response for a given token;
// [CiString20Type] is the underlying validated string primitive.
//
// ---
//
// [IDTagInfo] carries authorization status with optional expiry date and parent
// tag. Use [IDTagInfo.WithExpiryDate] and [IDTagInfo.WithParentIDTag] to set
// optional fields after construction.
//
// What it means: the OCPP 1.6 idTagInfo data type — the Central System's
// authorization response for a presented idToken.
//
// When to use it: Authorize.conf, StartTransaction.conf, and StopTransaction.req.
//
// What it is not: the token being authorized. The token itself is [IDToken].
//
// See also: [AuthorizationStatus] is the required field; [IDToken] is the token
// that triggers authorization; [DateTime] is used for the optional expiry date.
//
// ---
//
// [SampledValue] represents a single meter value sample; [MeterValue] groups
// sampled values under a single UTC timestamp.
//
// What they mean: the OCPP 1.6 sampledValue and meterValue data types, used to
// report energy and power measurements from a charge point.
//
// When to use them: MeterValues.req and StopTransaction.req transaction meter
// values.
//
// What they are not: real-time streaming values. These types represent a
// discrete snapshot, already parsed and validated.
//
// See also: [Measurand], [ReadingContext], [Phase], [Location], [UnitOfMeasure],
// and [ValueFormat] are the enumeration fields within [SampledValue].
//
// ---
//
// [AuthorizationData] is an entry in the local authorization list.
//
// What it means: a single record in the charge point's local authorization
// cache, pairing an [IDToken] with optional [IDTagInfo].
//
// When to use it: SendLocalList.req payload construction.
//
// What it is not: a live authorization check. It is a static cache entry
// stored on the charge point.
//
// See also: [UpdateType] and [ListVersionNumber] are used alongside
// [AuthorizationData] in SendLocalList.req.
//
// ---
//
// [KeyValue] is a configuration key-value pair.
//
// What it means: a single configuration entry returned in GetConfiguration.conf,
// pairing a bounded key name with a string value and a read-only flag.
//
// When to use it: GetConfiguration.conf response construction.
//
// What it is not: a generic map entry or environment variable. It is scoped
// exclusively to OCPP configuration management.
//
// See also: [CiString50Type] (key) and [CiString500Type] (value) are the
// underlying validated string types.
//
// ---
//
// [ChargingSchedulePeriod], [ChargingSchedule], and [ChargingProfile] model
// the OCPP 1.6 smart-charging hierarchy.
//
// What they mean: the three levels of the OCPP 1.6 smart-charging data model.
// A [ChargingSchedulePeriod] is one time window with a power or current limit;
// a [ChargingSchedule] assembles periods into a complete schedule;
// a [ChargingProfile] wraps a schedule with purpose, kind, recurrence, and
// validity metadata.
//
// When to use them: SetChargingProfile.req, GetCompositeSchedule.conf, and
// StartTransaction.conf when a TxProfile is attached.
//
// What they are not: real-time power commands. They define a plan that a charge
// point executes autonomously once set.
//
// See also: [ChargingProfilePurposeType], [ChargingProfileKindType],
// [ChargingRateUnit], and [RecurrencyKindType] are the enumeration fields
// within the hierarchy.
//
// ---
//
// [ListVersionNumber] tracks the version of the local authorization list.
//
// What it means: the integer version number used by the Central System and
// charge point to stay synchronized on the local list state. The constants
// [ListVersionUnsupported] (−1) and [ListVersionEmpty] (0) represent
// well-known sentinel states defined by the specification.
//
// When to use it: SendLocalList.req and GetLocalListVersion.conf.
//
// What it is not: a sequence number for OCPP call IDs or a transaction counter.
//
// See also: [AuthorizationData] and [UpdateType] are the other types used in
// the SendLocalList.req payload.
//
// # Errors
//
// [ErrEmptyValue] and [ErrInvalidValue] are sentinel errors returned by all
// constructors that fail validation.
//
// What they mean: [ErrEmptyValue] indicates the input was absent when a value
// was required; [ErrInvalidValue] indicates the input was present but did not
// satisfy the specification constraints (wrong format, out of range, or
// unrecognized enum value).
//
// When to use them: check errors.Is(err, types.ErrEmptyValue) or
// errors.Is(err, types.ErrInvalidValue) to branch on the class of validation
// failure without matching the full diagnostic message.
//
// What they are not: detailed parse errors. The wrapped error message carries
// the diagnostic context; the sentinel identifies the failure category.
//
// See also: [errors.Is] and [errors.As] in the Go standard library for
// error chain inspection.
package ocpp16types
