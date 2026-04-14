package baselib

import "math"

const (
	degreesPerTithy = 12.0
	maxTithy        = 30
)

// normalizeAngle returns an angle normalized into the range [0, 360).
// If the input is NaN or infinite, it returns NaN.
func normalizeAngle(angle float64) float64 {
	if math.IsNaN(angle) || math.IsInf(angle, 0) {
		return math.NaN()
	}

	a := math.Mod(angle, 360.0)
	if a < 0 {
		a += 360.0
	}

	// Mod can return 360.0 for exact multiples; force wrap to 0.0
	if a >= 360.0 {
		a -= 360.0
	}

	return a
}

// CalcTithy calculates the tithy (1..30) given the longitudes of the moon and the sun.
// - 1..15 : Sukla Paksha (waxing)
// - 16..30: Krishna Paksha (waning)
//
// The function normalizes input longitudes to [0,360) and computes the angular
// separation (moon - sun) also in [0,360). Each tithy is a 12 degree slice:
// tithy = floor(delta / 12) + 1
//
// Input guards:
// - If either input is NaN or infinite the function returns 0 to indicate an invalid result.
func CalcTithy(moon, sun float64) int {
	// Guard against invalid floating values early.
	if math.IsNaN(moon) || math.IsNaN(sun) || math.IsInf(moon, 0) || math.IsInf(sun, 0) {
		return 0
	}

	moonLon := normalizeAngle(moon)
	sunLon := normalizeAngle(sun)

	// If normalization produced NaN (shouldn't happen because we checked),
	// return 0 as an invalid signal.
	if math.IsNaN(moonLon) || math.IsNaN(sunLon) {
		return 0
	}

	// Angular separation from sun to moon in [0,360)
	delta := moonLon - sunLon
	if delta < 0 {
		delta += 360.0
	}

	// Numerical safety: clamp small negative zeros to 0
	if delta < 0 {
		delta = 0
	}

	tithy := int(math.Floor(delta/degreesPerTithy)) + 1

	// Ensure result is within the expected 1..30 range.
	if tithy < 1 {
		return 1
	}
	if tithy > maxTithy {
		return maxTithy
	}
	return tithy
}
