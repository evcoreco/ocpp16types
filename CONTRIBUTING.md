# Contributing

Thanks for your interest in improving `ocpp16types`.

## What contributing means

A contribution to this repository means adding or correcting a Go type, its
validation logic, or its documentation in a way that keeps the package
strictly aligned with the OCPP 1.6 specification. Every change must leave
the package at least as correct, immutable, and thread-safe as it was before.

## When to contribute

Contribute when:

- An OCPP 1.6 type, constraint, or enumeration value is missing or incorrect.
- A validated constructor, getter, or builder method does not behave as the
  specification requires.
- A godoc comment or example is missing, misleading, or does not follow the
  four-part documentation pattern (what it means, when it is used, what it is
  not, where to look for adjacent concepts).
- A test is absent for a specified behavior.

## What contributions are not

A contribution to this package is not:

- A general-purpose utility or helper that is not defined by OCPP 1.6.
- A refactor made solely for style without a spec or correctness reason.
- A new external dependency. This package must remain dependency-free.
- A transport, message codec, or protocol implementation — those belong in
  `ocpp16messages` or a dedicated module.

## Where to look for adjacent concepts

- OCPP 1.6 message constructors: `github.com/evcoreco/ocpp16messages`
- OCPP 1.6 specification: Open Charge Alliance (openchargealliance.org)
- Sentinel error definitions and format strings: `errors.go`
- Package-level documentation pattern: `doc.go`

---

## Quick start

- Install prerequisites: `golangci-lint`, `staticcheck`, `gci`, `gofumpt`,
  `golines`, `pkgsite`.
- Run formatting and quality checks before opening a PR:
  - `make format`
  - `make lint`
  - `make test`

## What to include in PRs

### OCPP compliance changes

**What they mean:** a PR that adds, corrects, or removes a type, constant,
validation rule, or enumeration value to match the OCPP 1.6 specification.

**When to open one:** when the package diverges from the specification, when a
new field or message type is missing, or when an existing constraint is wrong.

**What they are not:** a place to add convenience wrappers, aliases, or
types that the specification does not define.

**Where to look:** OCPP 1.6 specification appendix 4 (data types) and the
relevant message section. Include in the PR:

- The OCPP 1.6 type name and section reference.
- The spec reference (section/page) or a short excerpt.
- A minimal failing test demonstrating the previous behavior.
- Updated tests proving the new behavior and confirming no regressions.

### Adding new public API

**What it means:** a PR that exposes a new exported type, constructor,
method, or constant.

**When to open one:** when an OCPP 1.6 data type or field is not yet
represented in the package.

**What it is not:** an opportunity to add helper types or convenience
methods beyond the specification's data model.

**Where to look:** follow the constructor + validation pattern (`New*`).
Add public API tests in `package ocpp16types_test`. Add examples only when
they improve discoverability.

### Documentation changes

**What they mean:** a PR that updates godoc comments, the README, or
CONTRIBUTING to be more accurate or complete.

**When to open one:** when a comment is missing, incorrect, or does not follow
the four-part pattern (what it means, when it is used, what it is not, where to
look for adjacent concepts).

**What they are not:** cosmetic rewrites without substantive improvement.

**Where to look:** `doc.go` for the package-level pattern; individual source
files for per-type godoc. Run `make pkgsite` to preview rendered documentation.

### Concurrency and immutability changes

**What they mean:** a PR that touches pointer or slice fields in a type.

**When to open one:** when a new optional field is added or an existing field
changes from value to pointer semantics.

**What they are not:** an exception to the immutability requirement. Every
exported type must remain safe for concurrent use without synchronization.

**Where to look:** ensure constructors do not store caller-owned references.
Verify defensive copies exist in all getters that return pointers or slices.

---

## Opt-in test suites

This repo has opt-in suites that run weekly in CI and on demand locally:

| Suite | Command | Build tag |
|-------|---------|-----------|
| Fuzz | `make test-fuzz` | `fuzz` |
| Race | `make test-race` | `race` |

**What they mean:** extended validation runs that probe the package beyond
unit tests — fuzz tests exercise random inputs against constructors; the race
detector verifies concurrent access is safe.

**When to run them:** before any PR that changes validation logic, adds a new
type, or touches pointer or slice fields.

**What they are not:** a substitute for unit tests. Unit tests must still cover
all specified behaviors explicitly.

**Where to look:** `tests_fuzz/` for fuzz corpus and harnesses;
`tests_benchmark/` for performance baselines.

---

## Pull requests

- Keep PRs focused and small where possible.
- Ensure Markdown changes pass `markdownlint`.
- Prefer atomic tests (one behavior per test function).

## Code of Conduct

Be professional and constructive in all interactions. Contributions are
reviewed collaboratively — keep feedback focused on the code, not the
contributor.
