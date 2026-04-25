# ocpp16types

[![CI](https://github.com/aasanchez/ocpp16types/actions/workflows/ci.yml/badge.svg)](https://github.com/aasanchez/ocpp16types/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/aasanchez/ocpp16types/branch/master/graph/badge.svg)](https://codecov.io/gh/aasanchez/ocpp16types)
[![Go Reference](https://pkg.go.dev/badge/github.com/aasanchez/ocpp16types.svg)](https://pkg.go.dev/github.com/aasanchez/ocpp16types)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Strictly validated Go types for the **Open Charge Point Protocol (OCPP) 1.6** specification. This package is the single source of truth for shared ocpp types consumed by other modules in the OCPP ecosystem.

**Zero external dependencies** — uses only the Go standard library.

## Requirements

- Go **1.23** or later

## Installation

```bash
go get github.com/aasanchez/ocpp16types
```

## Design Principles

- **Validation at construction time.** Every type is validated through a `New*` constructor. If the constructor returns without error, the value is guaranteed to be spec-compliant.
- **Immutability.** All struct fields are unexported and accessed via value receivers. Constructors defensively copy pointers and slices.
- **Thread safety.** Value receivers and immutable fields make all types safe for concurrent use without synchronization.
- **Sentinel errors with context.** All validation failures wrap `ErrInvalidValue` or `ErrEmptyValue`, enabling `errors.Is()` checks while carrying rich diagnostic messages.
- **`fmt.Stringer` everywhere.** Every public type implements `fmt.Stringer` via compile-time interface assertions.

## Type Inventory

### Value Types

| Type | Description | Constraints |
|------|-------------|-------------|
| `CiString20Type` | Case-insensitive string | Printable ASCII, max 20 chars |
| `CiString25Type` | Case-insensitive string | Printable ASCII, max 25 chars |
| `CiString50Type` | Case-insensitive string | Printable ASCII, max 50 chars |
| `CiString255Type` | Case-insensitive string | Printable ASCII, max 255 chars |
| `CiString500Type` | Case-insensitive string | Printable ASCII, max 500 chars |
| `DateTime` | RFC 3339 timestamp | Must be UTC |
| `Integer` | Unsigned 16-bit integer | 0–65535 |

### Enumeration Types

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

All enumerations expose an `IsValid() bool` method and a `String() string` method.

### Composite Types

| Type | Description | Pattern |
|------|-------------|---------|
| `IdToken` | Identifier token wrapping `CiString20Type` | Constructor |
| `IdTagInfo` | Authorization info (status + optional expiry/parent) | Constructor + builder (`WithExpiryDate`, `WithParentIdTag`) |
| `SampledValue` | Single meter value sample with optional context metadata | Constructor |
| `MeterValue` | Timestamped meter reading containing sampled values | Constructor |
| `ChargingSchedulePeriod` | Single period within a charging schedule | Constructor |
| `ChargingSchedule` | Full schedule with rate unit and periods | Constructor |
| `ChargingProfile` | Complete charging profile with nested schedule | Input struct + constructor |
| `AuthorizationData` | Authorization cache/list entry | Constructor |
| `KeyValue` | Configuration key–value pair | Constructor |
| `ListVersionNumber` | Versioned list identifier | Constructor |

## Usage

This package is imported with the `types` alias across the EVCore ecosystem (e.g., in `ocpp16messages`):

```go
import types "github.com/aasanchez/ocpp16types"
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

### Building an IdTagInfo with Optional Fields

```go
dt, _ := types.NewDateTime("2026-12-31T23:59:59Z")
parent, _ := types.NewIdToken("ParentTag001")

info, err := types.NewIdTagInfo(types.AuthorizationStatusAccepted)
if err != nil {
    log.Fatal(err)
}

info = info.
    WithExpiryDate(dt).
    WithParentIdTag(parent)

fmt.Println(info)
// IdTagInfo{Status: Accepted, ExpiryDate: 2026-12-31T23:59:59Z, ParentIdTag: ParentTag001}
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

This is how `ocpp16messages/authorize` uses `ocpp16types` to build a validated OCPP message:

```go
import types "github.com/aasanchez/ocpp16types"

// ReqMessage represents an OCPP 1.6 Authorize.req message.
type ReqMessage struct {
    IdTag types.IdToken
}

func Req(input ReqInput) (ReqMessage, error) {
    str, err := types.NewCiString20Type(input.IdTag)
    if err != nil {
        return ReqMessage{}, fmt.Errorf("idTag: %w", err)
    }

    idToken := types.NewIdToken(str)

    return ReqMessage{IdTag: idToken}, nil
}
```

### Real-World: Authorize.conf (from ocpp16messages)

Building a response with validated status and optional fields:

```go
import types "github.com/aasanchez/ocpp16types"

// ConfMessage represents an OCPP 1.6 Authorize.conf message.
type ConfMessage struct {
    IdTagInfo types.IdTagInfo
}

func Conf(input ConfInput) (ConfMessage, error) {
    // Validate status (required)
    info, err := types.NewIdTagInfo(
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
        info = info.WithParentIdTag(types.NewIdToken(ciStr))
    }

    return ConfMessage{IdTagInfo: info}, nil
}
```

### Real-World: BootNotification.req struct (from ocpp16messages)

Shows how types compose into a full OCPP message with required and optional fields:

```go
import types "github.com/aasanchez/ocpp16types"

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

Fuzz tests live in `tests_fuzz/` and benchmark tests in `tests_benchmark/`. Both directories use build tags to keep them out of default test runs.

Coverage is uploaded to [Codecov](https://codecov.io/gh/aasanchez/ocpp16types) on every CI push.

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
├── MIGRATION.md          # Migration guide from embedded types
└── LICENSE               # MIT
```

## Documentation

Godoc is the primary documentation surface. Run a local pkgsite server:

```bash
make pkgsite
```

This starts a server at `http://localhost:8080/github.com/aasanchez/ocpp16types`.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on pull requests, OCPP compliance changes, and opt-in test suites.

## License

[MIT](LICENSE) — Copyright (c) 2026 aasanchez
