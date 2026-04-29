# ocpp16types

[![CI](https://github.com/evcoreco/ocpp16types/actions/workflows/ci.yml/badge.svg)](https://github.com/evcoreco/ocpp16types/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16types/branch/master/graph/badge.svg)](https://codecov.io/gh/aasanchez/ocpp16types)
[![Go Reference](https://pkg.go.dev/badge/github.com/evcoreco/ocpp16types.svg)](https://pkg.go.dev/github.com/evcoreco/ocpp16types)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Strictly validated Go types for the **Open Charge Point Protocol (OCPP) 1.6** specification.

**What it means:** ocpp16types is the single source of truth for OCPP 1.6 data
types in the EVCore ecosystem. Every type is validated at construction time. A
value that passes construction is guaranteed to be spec-compliant, immutable,
and thread-safe.

**When to use it:** import this package when building OCPP 1.6 message payloads,
validating incoming field values against the specification, or sharing validated
OCPP types across multiple modules.

**What it is not:** ocpp16types is not a message layer, a transport, or a JSON
codec. It does not send or receive OCPP messages, parse WebSocket frames, or
manage connections. It defines and validates data types only.

**Where to look for adjacent concepts:**

- OCPP 1.6 message structs and call constructors: `github.com/evcoreco/ocpp16messages`
- OCPP 1.6 specification: Open Charge Alliance

**Zero external dependencies** — uses only the Go standard library.

## Requirements

- Go **1.25** or later

## Installation

```bash
go get github.com/evcoreco/ocpp16types@v1.0.1
```

## Design Principles

| Principle | What it means |
|-----------|--------------|
| **Validation at construction time** | Every type is validated through a `New*` constructor. If the constructor returns without error, the value is guaranteed to be spec-compliant. |
| **Immutability** | All struct fields are unexported and accessed via value receivers. Constructors defensively copy pointers and slices. |
| **Thread safety** | Value receivers and immutable fields make all types safe for concurrent use without synchronization. |
| **Sentinel errors with context** | All validation failures wrap `ErrInvalidValue` or `ErrEmptyValue`, enabling `errors.Is()` checks while carrying rich diagnostic messages. |
| **`fmt.Stringer` everywhere** | Every public type implements `fmt.Stringer` via compile-time interface assertions. |

## Type Inventory

### Value Types

Value types are the primitive building blocks of OCPP 1.6 messages. They
validate their input at construction and cannot hold an invalid value after
creation.

#### CiString Types

**What they mean:** case-insensitive strings bounded to a specific maximum
length, as defined in OCPP 1.6 appendix 4. The "case-insensitive" label is a
specification convention; these types do not perform case folding at runtime.

**When to use them:** wherever OCPP 1.6 specifies a `CiString` field — vendor
IDs, firmware versions, identifier tokens, configuration keys, and similar
bounded text values.

**What they are not:** general-purpose strings. Empty values, Unicode characters,
and values exceeding the maximum length are all rejected at construction. Do not
use them as free-form text fields.

**Where to look for adjacent concepts:** `IDToken` wraps `CiString20Type`;
`KeyValue` uses `CiString50Type` (key) and `CiString500Type` (value);
`SampledValue` uses `CiString500Type` for the raw measurement string.

| Type | Constraints |
|------|-------------|
| `CiString20Type` | Printable ASCII, max 20 chars |
| `CiString25Type` | Printable ASCII, max 25 chars |
| `CiString50Type` | Printable ASCII, max 50 chars |
| `CiString255Type` | Printable ASCII, max 255 chars |
| `CiString500Type` | Printable ASCII, max 500 chars |

#### DateTime

**What it means:** a moment in time expressed as an RFC 3339 UTC string, as
required by the OCPP 1.6 specification throughout all date-time fields.

**When to use it:** expiry dates, transaction start and stop times, meter
reading timestamps, and any other OCPP field typed as `dateTime`.

**What it is not:** a timezone-aware or local-time value. Non-UTC offsets are
rejected at construction. It is not a duration or an interval.

**Where to look for adjacent concepts:** `MeterValue` embeds a `DateTime`
timestamp; `IDTagInfo.WithExpiryDate` accepts a `DateTime`; `ChargingProfile`
uses optional `DateTime` fields for `ValidFrom` and `ValidTo`.

| Type | Constraints |
|------|-------------|
| `DateTime` | RFC 3339, must be UTC |

#### Integer

**What it means:** the OCPP 1.6 integer primitive (appendix 4), which maps to
an unsigned 16-bit range (0–65535).

**When to use it:** stack levels, connector identifiers, and other OCPP fields
typed as `integer` in the specification.

**What it is not:** a general Go `int`. Negative values and values above 65535
are rejected. It is not a counter managed by this package.

**Where to look for adjacent concepts:** `ChargingSchedulePeriod` uses `Integer`
for `StartPeriod`; `ChargingProfile` uses `Integer` for `StackLevel` and
`TransactionID`.

| Type | Constraints |
|------|-------------|
| `Integer` | 0–65535 (uint16) |

### Enumeration Types

**What they mean:** typed string constants that represent the closed set of
allowed values for a specific OCPP 1.6 field. Each enumeration is the direct
Go mapping of the corresponding OCPP 1.6 enumeration.

**When to use them:** assign an enumeration constant (e.g.
`AuthorizationStatusAccepted`) to any OCPP field that the specification
describes as an enum. Call `IsValid()` before accepting an enum value received
from an external source.

**What they are not:** open-ended string fields. An unrecognized string value
fails `IsValid()` even if it resembles a valid value. They are not extensible
without a specification change.

**Where to look for adjacent concepts:** composite types `IDTagInfo`,
`SampledValue`, and `ChargingProfile` embed enumeration types as required or
optional fields.

All enumerations expose an `IsValid() bool` method and a `String() string`
method.

| Type | Values |
|------|--------|
| `AuthorizationStatus` | Accepted, Blocked, Expired, Invalid, ConcurrentTx |
| `AvailabilityStatus` | Accepted, Rejected, Scheduled |
| `AvailabilityType` | Inoperative, Operative |
| `CancelReservationStatus` | Accepted, Rejected |
| `ChargePointErrorCode` | ConnectorLockFailure, EVCommunicationError, GroundFailure, … (16 values) |
| `ChargePointStatus` | Available, Preparing, Charging, SuspendedEVSE, SuspendedEV, Finishing, Reserved, Unavailable, Faulted |
| `ChargingProfileKindType` | Absolute, Recurring, Relative |
| `ChargingProfilePurposeType` | ChargePointMaxProfile, TxDefaultProfile, TxProfile |
| `ChargingProfileStatus` | Accepted, Rejected, NotSupported |
| `ChargingRateUnit` | W, A |
| `ClearCacheStatus` | Accepted, Rejected |
| `ClearChargingProfileStatus` | Accepted, Unknown |
| `ConfigurationStatus` | Accepted, Rejected, RebootRequired, NotSupported |
| `DataTransferStatus` | Accepted, Rejected, UnknownMessageId, UnknownVendor |
| `DiagnosticsStatus` | Idle, Uploaded, UploadFailed, Uploading |
| `FirmwareStatus` | Downloaded, DownloadFailed, Downloading, Idle, InstallationFailed, Installing, Installed |
| `GetCompositeScheduleStatus` | Accepted, Rejected |
| `Location` | Body, Cable, EV, Inlet, Outlet |
| `Measurand` | Energy.Active.Export.Register, Power.Active.Import, … (22 values) |
| `MessageTrigger` | BootNotification, DiagnosticsStatusNotification, FirmwareStatusNotification, Heartbeat, MeterValues, StatusNotification |
| `Phase` | L1, L2, L3, N, L1-N, L2-N, L3-N, L1-L2, L2-L3, L3-L1 |
| `ReadingContext` | Interruption.Begin, Interruption.End, Other, Sample.Clock, Sample.Periodic, Transaction.Begin, Transaction.End |
| `RecurrencyKindType` | Daily, Weekly |
| `RegistrationStatus` | Accepted, Pending, Rejected |
| `RemoteStartTransactionStatus` | Accepted, Rejected |
| `RemoteStopTransactionStatus` | Accepted, Rejected |
| `ReservationStatus` | Accepted, Faulted, Occupied, Rejected, Unavailable |
| `ResetStatus` | Accepted, Rejected |
| `ResetType` | Hard, Soft |
| `StopReason` | EmergencyStop, EVDisconnected, HardReset, Local, Other, PowerLoss, Reboot, Remote, SoftReset, UnlockCommand, DeAuthorized |
| `TriggerMessageStatus` | Accepted, Rejected, NotImplemented |
| `UnlockStatus` | Unlocked, UnlockFailed, NotSupported |
| `UpdateStatus` | Accepted, Failed, NotSupported, VersionMismatch |
| `UpdateType` | Differential, Full |
| `UnitOfMeasure` | Wh, kWh, varh, kvarh, W, kW, VA, kVA, var, kvar, A, V, Celsius, Fahrenheit, K, Percent |
| `ValueFormat` | Raw, SignedData |

### Composite Types

Composite types combine validated primitives and enumerations into the
structured OCPP 1.6 data objects used in message payloads.

#### IDToken

**What it means:** the OCPP 1.6 `idToken` data type — an identifier that a
charge point presents to the Central System for authorization.

**When to use it:** `Authorize.req`, `StartTransaction.req`,
`StopTransaction.req`, and any OCPP message that carries an identifier token.

**What it is not:** an authorization result. The result of checking a token is
`IDTagInfo`, not `IDToken`.

**Where to look for adjacent concepts:** `IDTagInfo` carries the authorization
response for a given token; `CiString20Type` is the underlying validated string
primitive.

#### IDTagInfo

**What it means:** the OCPP 1.6 `idTagInfo` data type — the Central System's
authorization response for a presented `idToken`.

**When to use it:** `Authorize.conf`, `StartTransaction.conf`, and
`StopTransaction.req`.

**What it is not:** the token being authorized. The token itself is `IDToken`.

**Where to look for adjacent concepts:** `AuthorizationStatus` is the required
field; `IDToken` is the token that triggers the check; `DateTime` is used for
the optional expiry date.

#### SampledValue and MeterValue

**What they mean:** the OCPP 1.6 `sampledValue` and `meterValue` data types,
used to report energy and power measurements from a charge point. A
`SampledValue` is one measurement reading; a `MeterValue` groups one or more
readings under a single UTC timestamp.

**When to use them:** `MeterValues.req` and `StopTransaction.req` transaction
meter values.

**What they are not:** real-time streaming values. These types represent a
discrete snapshot, already parsed and validated.

**Where to look for adjacent concepts:** `Measurand`, `ReadingContext`, `Phase`,
`Location`, `UnitOfMeasure`, and `ValueFormat` are the enumeration fields
within `SampledValue`.

#### AuthorizationData

**What it means:** a single record in the charge point's local authorization
cache, pairing an `IDToken` with optional `IDTagInfo`.

**When to use it:** `SendLocalList.req` payload construction.

**What it is not:** a live authorization check. It is a static cache entry
stored on the charge point.

**Where to look for adjacent concepts:** `UpdateType` and `ListVersionNumber`
are used alongside `AuthorizationData` in `SendLocalList.req`.

#### KeyValue

**What it means:** a single configuration entry returned in
`GetConfiguration.conf`, pairing a bounded key name with a string value and a
read-only flag.

**When to use it:** `GetConfiguration.conf` response construction.

**What it is not:** a generic map entry or environment variable. It is scoped
exclusively to OCPP configuration management.

**Where to look for adjacent concepts:** `CiString50Type` (key) and
`CiString500Type` (value) are the underlying validated string types.

#### ChargingSchedulePeriod, ChargingSchedule, ChargingProfile

**What they mean:** the three levels of the OCPP 1.6 smart-charging data model.
A `ChargingSchedulePeriod` is one time window with a power or current limit; a
`ChargingSchedule` assembles periods into a complete schedule; a
`ChargingProfile` wraps a schedule with purpose, kind, recurrence, and validity
metadata.

**When to use them:** `SetChargingProfile.req`, `GetCompositeSchedule.conf`,
and `StartTransaction.conf` when a `TxProfile` is attached.

**What they are not:** real-time power commands. They define a plan that a
charge point executes autonomously once set.

**Where to look for adjacent concepts:** `ChargingProfilePurposeType`,
`ChargingProfileKindType`, `ChargingRateUnit`, and `RecurrencyKindType` are the
enumeration fields within the hierarchy.

#### ListVersionNumber

**What it means:** the integer version number used by the Central System and
charge point to stay synchronized on the local list state. The constants
`ListVersionUnsupported` (−1) and `ListVersionEmpty` (0) represent well-known
sentinel states defined by the specification.

**When to use it:** `SendLocalList.req` and `GetLocalListVersion.conf`.

**What it is not:** a sequence number for OCPP call IDs or a transaction counter.

**Where to look for adjacent concepts:** `AuthorizationData` and `UpdateType`
are the other types used in the `SendLocalList.req` payload.

| Type | Description | Pattern |
|------|-------------|---------|
| `IDToken` | Identifier token wrapping `CiString20Type` | Constructor |
| `IDTagInfo` | Authorization info (status + optional expiry/parent) | Constructor + builder (`WithExpiryDate`, `WithParentIDTag`) |
| `SampledValue` | Single meter value sample with optional context metadata | Constructor |
| `MeterValue` | Timestamped meter reading containing sampled values | Constructor |
| `ChargingSchedulePeriod` | Single period within a charging schedule | Constructor |
| `ChargingSchedule` | Full schedule with rate unit and periods | Constructor |
| `ChargingProfile` | Complete charging profile with nested schedule | Input struct + constructor |
| `AuthorizationData` | Authorization cache/list entry | Constructor |
| `KeyValue` | Configuration key–value pair | Constructor |
| `ListVersionNumber` | Versioned list identifier | Constructor |

## Errors

**What they mean:** `ErrEmptyValue` indicates the input was absent when a value
was required; `ErrInvalidValue` indicates the input was present but did not
satisfy the specification constraints (wrong format, out of range, or
unrecognized enum value).

**When to use them:** check `errors.Is(err, types.ErrEmptyValue)` or
`errors.Is(err, types.ErrInvalidValue)` to branch on the class of validation
failure without matching the full diagnostic message.

**What they are not:** detailed parse errors. The wrapped error message carries
diagnostic context; the sentinel identifies the failure category.

**Where to look for adjacent concepts:** `errors.Is` and `errors.As` in the Go
standard library for error chain inspection.

## Usage

This package is imported with the `types` alias across the EVCore ecosystem
(e.g., in `ocpp16messages`):

```go
import types "github.com/evcoreco/ocpp16types"
```

### Creating a CiString

```go
cistr, err := types.NewCiString20Type("ABC123")
if err != nil {
    log.Fatal(err)
}
fmt.Println(cistr.Value()) // "ABC123"
```

### Working with Enumerations

```go
status := types.AuthorizationStatusAccepted
fmt.Println(status.IsValid()) // true
fmt.Println(status.String())  // "Accepted"
```

### Building an IDTagInfo with Optional Fields

```go
dt, _ := types.NewDateTime("2026-12-31T23:59:59Z")
parentStr, _ := types.NewCiString20Type("ParentTag001")
parent := types.NewIDToken(parentStr)

info, err := types.NewIDTagInfo(types.AuthorizationStatusAccepted)
if err != nil {
    log.Fatal(err)
}

info = info.
    WithExpiryDate(dt).
    WithParentIDTag(parent)

fmt.Println(info)
// IDTagInfo{Status: Accepted, ExpiryDate: 2026-12-31T23:59:59Z, ParentIDTag: ParentTag001}
```

### Error Handling with Sentinel Errors

```go
_, err := types.NewCiString20Type("")
if errors.Is(err, types.ErrEmptyValue) {
    // handle empty input
}

_, err = types.NewInteger(-1)
if errors.Is(err, types.ErrInvalidValue) {
    // handle out-of-range value
}
```

### Real-World: Authorize.req (from ocpp16messages)

This is how `ocpp16messages/authorize` uses `ocpp16types` to build a validated
OCPP message:

```go
import types "github.com/evcoreco/ocpp16types"

// ReqMessage represents an OCPP 1.6 Authorize.req message.
type ReqMessage struct {
    IdTag types.IdToken
}

func Req(input ReqInput) (ReqMessage, error) {
    str, err := types.NewCiString20Type(input.IdTag)
    if err != nil {
        return ReqMessage{}, fmt.Errorf("idTag: %w", err)
    }

    idToken := types.NewIDToken(str)

    return ReqMessage{IdTag: idToken}, nil
}
```

### Real-World: Authorize.conf (from ocpp16messages)

Building a response with validated status and optional fields:

```go
import types "github.com/evcoreco/ocpp16types"

// ConfMessage represents an OCPP 1.6 Authorize.conf message.
type ConfMessage struct {
    IdTagInfo types.IdTagInfo
}

func Conf(input ConfInput) (ConfMessage, error) {
    // Validate status (required)
    info, err := types.NewIDTagInfo(
        types.AuthorizationStatus(input.Status),
    )
    if err != nil {
        return ConfMessage{}, fmt.Errorf("status: %w", err)
    }

    // Apply optional fields
    if input.ExpiryDate != nil {
        expiryDate, err := types.NewDateTime(*input.ExpiryDate)
        if err != nil {
            return ConfMessage{}, fmt.Errorf("expiryDate: %w", err)
        }
        info = info.WithExpiryDate(expiryDate)
    }

    if input.ParentIdTag != nil {
        ciStr, err := types.NewCiString20Type(*input.ParentIdTag)
        if err != nil {
            return ConfMessage{}, fmt.Errorf("parentIdTag: %w", err)
        }
        info = info.WithParentIDTag(types.NewIDToken(ciStr))
    }

    return ConfMessage{IdTagInfo: info}, nil
}
```

### Real-World: BootNotification.req struct (from ocpp16messages)

Shows how types compose into a full OCPP message with required and optional
fields:

```go
import types "github.com/evcoreco/ocpp16types"

// ReqMessage represents an OCPP 1.6 BootNotification.req message.
type ReqMessage struct {
    ChargePointVendor       types.CiString20Type
    ChargePointModel        types.CiString20Type
    ChargePointSerialNumber *types.CiString25Type  // optional
    ChargeBoxSerialNumber   *types.CiString25Type  // optional
    FirmwareVersion         *types.CiString50Type  // optional
    Iccid                   *types.CiString20Type  // optional
    Imsi                    *types.CiString20Type  // optional
    MeterType               *types.CiString25Type  // optional
    MeterSerialNumber       *types.CiString25Type  // optional
}
```

## Testing

The project includes four test suites orchestrated via `make`:

```bash
make test           # Unit tests with coverage report
make test-example   # Testable examples (godoc)
make test-fuzz      # Fuzz testing (Go 1.20+, ./tests_fuzz)
make test-race      # Race detector
make test-all       # All of the above + lint
```

Fuzz tests live in `tests_fuzz/` and benchmark tests in `tests_benchmark/`.
Both directories use build tags to keep them out of default test runs.

Coverage is uploaded to [Codecov](https://codecov.io/gh/aasanchez/ocpp16types)
on every CI push.

## CI Pipeline

GitHub Actions runs on every push and pull request to `master`:

1. `go vet`
2. `staticcheck`
3. `golangci-lint` (revive, staticcheck, gofmt, goimports)
4. Unit tests with coverage
5. Example tests
6. Race detector
7. Codecov upload

## Project Structure

```
ocpp16types/
├── *.go                  # One file per type (source + _test.go)
├── doc.go                # Package-level godoc
├── errors.go             # Sentinel errors and format strings
├── tests_fuzz/           # Fuzz tests (build tag: fuzz)
├── tests_benchmark/      # Benchmark tests
├── .github/workflows/    # CI configuration
├── Makefile              # Build, test, lint, format targets
├── golangci.yml          # Linter configuration
├── codecov.yml           # Coverage thresholds
├── CONTRIBUTING.md       # Contribution guidelines
├── CLA.md                # Contributor License Agreement
└── LICENSE               # MIT
```

## Documentation

Godoc is the primary documentation surface. Run a local pkgsite server:

```bash
make pkgsite
```

This starts a server at `http://localhost:8080/github.com/evcoreco/ocpp16types`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on pull requests, OCPP
compliance changes, and opt-in test suites.

## License

[MIT](LICENSE) — Copyright (c) 2026 aasanchez
