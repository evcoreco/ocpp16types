# Contributing

Thanks for your interest in improving `ocpp16types`.

## Quick start

- Install prerequisites (see `AGENTS.md`).
- Run formatting and quality checks before opening a PR:
  - `make format`
  - `make lint`
  - `make test`

## What to include in PRs

### OCPP compliance changes

When changing validation or type definitions, include:

- The OCPP 1.6 type name and section reference.
- The spec reference (section/page) or a short excerpt.
- A minimal failing test demonstrating the previous behavior.
- Updated tests proving the new behavior (and no regressions).

### Adding new public API

- Add public API tests in `tests/` subdirectory using `package ocpp16types_test`.
- Add examples only when they improve discoverability.
- Follow the constructor + validation pattern (`New*`).

### Concurrency / immutability

If your change touches pointers or slices:

- Ensure constructors do not store caller-owned references.
- Verify defensive copies in getters for pointer and slice fields.

## Opt-in suites

This repo has opt-in suites that run weekly in CI and on demand locally:

- Fuzz: `make test-fuzz` (`./tests_fuzz`, build tag `fuzz`)
- Race: `make test-race` (build tag `race`)

## Pull requests

- Keep PRs focused and small where possible.
- Ensure Markdown changes pass `markdownlint`.
- Prefer atomic tests (one behavior per test function).

## Code of Conduct

This project follows `CODE_OF_CONDUCT.md`.
