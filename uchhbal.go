package baselib

import (
	"math"
)

// exaltationMapping stores the maximum exaltation degrees for each planet.
var exaltationMapping = map[string]float64{
	SUN:     10,
	MOON:    33,
	MARS:    298,
	MERCURY: 165,
	JUPITER: 95,
	VENUS:   357,
	SATURN:  20,
	RAHU:    80,
	KETU:    260,
}

// CalcUchhBal calculates the Uchch Bala (exaltation strength) of a planet.
// pl_long: longitude of the planet in degrees.
// pl_name: name of the planet (use constants like SUN, MOON, etc.).
func CalcUchhBal(pl_long float64, pl_name string) float64 {
	exalt_deg_max, ok := exaltationMapping[pl_name]
	if !ok {
		return 0
	}

	planetDeg := math.Abs(exalt_deg_max - pl_long)
	planetDeg = normalizeAngle(planetDeg)

	// Convert degrees to radians for math.Cos
	rad := planetDeg * math.Pi / 180
	power := (1 + math.Cos(rad)) / 2
	uchhBal := ((1 + power) / 2) * 100

	return uchhBal
}
