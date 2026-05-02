package bal

import (
	baselib "github.com/jenujari/planets-lib"

	"math"
)

// exaltationMapping stores the maximum exaltation degrees for each planet.
var exaltationMapping = map[string]float64{
	baselib.SUN:     10,
	baselib.MOON:    33,
	baselib.MARS:    298,
	baselib.MERCURY: 165,
	baselib.JUPITER: 95,
	baselib.VENUS:   357,
	baselib.SATURN:  20,
	baselib.RAHU:    80,
	baselib.KETU:    260,
}

// UchhBal calculates the Uchch Bala (exaltation strength) of a planet.
// pl_long: longitude of the planet in degrees.
// pl_name: name of the planet (use constants like baselib.SUN, baselib.MOON, etc.).
func UchhBal(pl_long float64, pl_name string) float64 {
	exalt_deg_max, ok := exaltationMapping[pl_name]
	if !ok {
		return 0
	}

	planetDeg := math.Abs(exalt_deg_max - pl_long)
	planetDeg = baselib.NormalizeAngle(planetDeg)

	// Convert degrees to radians for math.Cos
	rad := planetDeg * math.Pi / 180
	power := (1 + math.Cos(rad)) / 2
	uchhBal := ((1 + power) / 2) * 100

	return uchhBal
}
