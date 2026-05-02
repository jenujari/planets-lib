# planets-lib

[![Go Test](https://github.com/jenujari/planets-lib/actions/workflows/go-test.yml/badge.svg?branch=main)](https://github.com/jenujari/planets-lib/actions/workflows/go-test.yml)

A small, focused Go library for basic astronomical/astrological utilities:
- Calculation of tithy (lunar day)
- Mapping longitudes to zodiac signs, navamsha signs, and lords
- Mapping longitudes to nakshatra (lunar mansion) and pada
- Planetary relationship calculations (Graha Maitri)
- Planetary speed categorization and Vedha calculation
- Comprehensive planetary strength calculations (Uchh Bal, Uday Bal, Vakra Bal, Kshetra Bal, Navansh Bal) through the `bal` package
- Lightweight planet coordinate container with derived computations
- DMS (degrees/minutes/seconds) formatting helpers

This repository is organized logically — core production code lives at the module root (e.g. `tithy.go`, `sign.go`, `nakshatra.go`, `planet.go`, `util.go`), while strength calculations reside in the `bal/` package. Tests are colocated with the code in `*_test.go` files.

## Highlights / Design decisions

- Angle normalization is centralized via `NormalizeAngle` to consistently map arbitrary angles into `[0, 360)`.
- Functions are defensive about invalid floats: NaN and ±Inf inputs are handled predictably instead of panicking.
  - `CalcTithy` returns `0` for invalid inputs.
  - `GetSignFrmDegree` returns an empty string for invalid inputs.
  - `GetNakshatraPadaFromDegree` returns a zero-value `NakshatraPada` for invalid inputs.
- `PlanetCord.CalculateDerivedValues()` computes derived fields (DMS, sign, nakshatra, retrograde) defensively and uses normalized longitude for lookups without mutating the stored longitude.
- Unit tests cover normalization, boundary values, and invalid inputs.

## Quick usage

- Compute tithy given moon & sun longitudes (degrees):
  - Use `CalcTithy(moonLong, sunLong)` — returns an `int` in `1..30` (or `0` when input is invalid).

- Get zodiac sign from a longitude:
  - Use `GetSignFrmDegree(longitude)` — returns sign name like `"Aries"` or `""` if input invalid.

- Get nakshatra and pada from a longitude:
  - Use `GetNakshatraPadaFromDegree(longitude)` — returns `NakshatraPada{Name string, Pada int}` (zero-value if invalid).

- Planet coordinates helper:
  - Use the `PlanetCord` struct; call `CalculateDerivedValues()` to populate derived fields like `Sign`, `Nakshatra`, DMS fields, and `IsRetro`.

- Calculate planetary strength:
  - Import `"github.com/jenujari/planets-lib/bal"` and call functions like `bal.KshetraBal(plLong, plName)` to evaluate positional strength, returning a normalized score from `0` to `100`.

Example (conceptual):
- To run existing tests locally: `go test ./...`
- To format code: `go fmt ./...`

## Running tests

- Run the full test suite:
  - `go test ./... -v`

- Run a focused test (example):
  - `go test -run TestCalcTithy ./...`

- Check coverage:
  - `go test ./... -cover`

The test suite includes cases for:
- Angle normalization (values outside 0..360, negatives, large positives)
- Edge boundaries near sign/nakshatra/pada/tithy transitions
- Invalid float inputs (NaN/Inf)
- PlanetCord derived value behavior

## Continuous Integration

A GitHub Actions workflow is included at `.github/workflows/go-test.yml`. It runs `go test ./...` for `push` and `pull_request` events.

To require CI to pass before merging to `main`:
1. On GitHub, go to Repository → Settings → Branches → Branch protection rules.
2. Add/Edit rule for `main`.
3. Enable "Require status checks to pass before merging".
4. Select the workflow job/check named `Go Test` (it will appear after a run).
5. Save the rule.

This ensures pull requests (and pushes) run the test suite; merges to `main` can be blocked until tests pass.

## Contributing

- Follow the project guidelines in `AGENTS.md`:
  - Use idiomatic Go formatting (`go fmt`).
  - Prefer table-driven tests for algorithmic logic.
  - Use short, imperative commit messages.
  - Run `go test ./...` before opening a PR.

- When updating calculations, add tests for numeric edge cases and normalization behavior.

## Files of interest

- `tithy.go` — tithy calculation and `NormalizeAngle`.
- `sign.go` — sign names, lords, and `GetSignFrmDegree`.
- `nakshatra.go` — nakshatra/pada mapping and vowel-based mapping helper.
- `planet.go` — speed classification, Vedha logic, planetary relationship calculation (Graha Maitri Chakra), `PlanetCord` struct and `CalculateDerivedValues`.
- `navanshRashi.go` — Navamsha block tracking.
- `util.go` — `DMS` utilities and formatting/parsing helpers.
- `bal/` — package encapsulating various planetary strength calculations (`UchhBal`, `UdayBal`, `VakraBal`, `KshetraBal`, `NavanshBal`).
- Tests: e.g., `tithy_test.go`, `bal/navanshbal_test.go`.

## Development environment

- The module declares `go 1.25.6` in `go.mod`; use a compatible Go toolchain (>= 1.18 works for modules, but module Go version is 1.25.6).
- Dependencies: `github.com/stretchr/testify` is used for assertions in tests.
