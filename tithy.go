package baselib

import "math"

const (
	degreesPerTithy = 12.0
	maxTithy        = 30
)

const (
	SUNDAY    = "Sun"
	MONDAY    = "Mon"
	TUESDAY   = "Tue"
	WEDNESDAY = "Wed"
	THURSDAY  = "Thu"
	FRIDAY    = "Fri"
	SATURDAY  = "Sat"
)

const (
	NANDA  = "Nanda"
	BHADRA = "Bhadra"
	JAYA   = "Jaya"
	RIKTA  = "Rikta"
	POORNA = "Poorna"
)

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
	moonLon, okMoon := ValidAngle(moon)
	sunLon, okSun := ValidAngle(sun)
	if !okMoon || !okSun {
		return 0
	}

	// Angular separation from sun to moon in [0,360)
	delta := moonLon - sunLon
	if delta < 0 {
		delta += 360.0
	}

	// Numerical safety: clamp small negative zeros to 0
	delta = max(0, delta)

	tithy := int(math.Floor(delta/degreesPerTithy)) + 1

	// Ensure result is within the expected 1..30 range.
	return max(1, min(tithy, maxTithy))
}
