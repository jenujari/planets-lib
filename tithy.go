package baselib

import (
	"math"
	"time"
)

const (
	degreesPerTithy = 12.0
	maxTithy        = 30
)

// ist is Indian Standard Time, UTC+5:30, with no daylight saving.
var ist = time.FixedZone("IST", 5*60*60+30*60)

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

// TithyInstant is 06:30 IST on the calendar date of date, returned in UTC.
// The clock on date is ignored. 06:30 IST is 01:00 UTC on that same civil day.
func TithyInstant(date time.Time) time.Time {
	y, m, d := date.Date()
	return time.Date(y, m, d, 6, 30, 0, 0, ist).UTC()
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
