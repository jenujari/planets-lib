# Repository Guidelines

## Project Structure & Module Organization
This repository is a small Go library module (`github.com/jenujari/planets-lib`) with a flat layout. Production code lives at the repository root in files such as `planet.go`, `nakshatra.go`, `sign.go`, `tithy.go`, and `util.go`. Tests sit beside the code they cover in `*_test.go` files such as `tithy_test.go` and `util_test.go`. There are no separate `cmd/`, `internal/`, or asset directories yet, so keep new packages intentional and only introduce subdirectories when the API surface grows.

## Build, Test, and Development Commands
Use the standard Go toolchain:

- `go test ./...` runs the full test suite for the module.
- `go test -run TestCalcTithy ./...` runs a focused test while iterating.
- `go fmt ./...` formats all Go files with canonical Go style.
- `go test -cover ./...` checks package coverage before opening a PR.

Run commands from the repository root.

## Coding Style & Naming Conventions
Follow idiomatic Go formatting and let `go fmt` manage indentation and spacing; do not align code manually. Keep package-level APIs clear and exported names in PascalCase (`PlanetCord`, `CalcTithy`). Use lowercase file names with underscores only when they improve readability; tests must end with `_test.go`. Prefer table-driven tests for calculation logic, matching the existing patterns in `tithy_test.go` and `util_test.go`.

## Testing Guidelines
Tests use Go’s `testing` package with `github.com/stretchr/testify/assert` for assertions. Name tests `TestXxx` and use descriptive subtest names with `t.Run(...)` for boundary cases and normalization behavior. Add coverage for numeric edge cases whenever changing planetary calculations, angle normalization, or formatting helpers.

## Commit & Pull Request Guidelines
Current history uses short, imperative commit subjects such as `lib setup`, `added is retro flag`, and `tity calculation`. Prefer concise, present-tense messages that describe one logical change; make them more precise than `changes`. Pull requests should include a brief summary, test evidence (`go test ./...` or coverage output), and any affected formulas or examples when behavior changes.
