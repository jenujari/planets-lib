package baselib

import (
	"math"
	"strconv"
)

// CalcUdayBal calculates the Udaybal (rising strength) of a planet based on its
// distance from the Sun, its motion (direct or retrograde), and its specific
// astance (combustion) values.
//
// Parameters:
//   - sun_long: Longitude of the Sun in degrees.
//   - pl_long: Longitude of the planet in degrees.
//   - pl_speed: Longitudinal speed of the planet (degrees per day).
//   - pl_name: Name of the planet (use constants like SUN, MOON, etc.).
//
// Returns:
//   - The calculated Udaybal value ranging from 0 to 100.
func CalcUdayBal(sun_long, pl_long, pl_speed float64, pl_name string) float64 {
	const MaxRetrunValue = 100.0
	const MinReturnValue = 0.0

	// If Planet is Sun then Udaybal is MaxRetrunValue
	if pl_name == SUN {
		return MaxRetrunValue
	}

	// If Planet is Rahu and Ketu then Udaybal is MinReturnValue
	if pl_name == RAHU || pl_name == KETU {
		return MinReturnValue
	}

	// Determine motion: direct or retro
	// speed >= 0 => Direct, speed < 0 => retro
	isRetro := pl_speed < 0

	// Decide planet astance (pl_ast)
	var pl_ast float64
	switch pl_name {
	case MOON:
		pl_ast = 12
	case MARS:
		pl_ast = 17
	case JUPITER:
		pl_ast = 11
	case SATURN, URANUS, NEPTUNE, PLUTO:
		pl_ast = 15
	case MERCURY:
		if isRetro {
			pl_ast = 12
		} else {
			pl_ast = 14
		}
	case VENUS:
		if isRetro {
			pl_ast = 8
		} else {
			pl_ast = 10
		}
	default:
		// Default fallback for any other planet
		pl_ast = 15
	}

	// Planet distance from Sun
	// pl_dist = planet_long - sun_long // take mod if negative or always use mod
	// We use the shortest angular distance in [0, 180].
	sLong := normalizeAngle(sun_long)
	pLong := normalizeAngle(pl_long)

	pl_dist := math.Abs(pLong - sLong)
	if pl_dist > 180 {
		pl_dist = 360 - pl_dist
	}

	// if pl_dist <= pl_ast then Udaybal is MinReturnValue
	if pl_dist <= pl_ast {
		return MinReturnValue
	}

	// Decide max_diff
	var max_diff float64
	switch pl_name {
	case MERCURY:
		max_diff = 27
	case VENUS:
		max_diff = 47
	default:
		max_diff = 180
	}

	// Calculate final Udaybal
	availabe_range := max_diff - pl_ast
	if availabe_range <= 0 {
		panic("check logic again this should not be coming for " + pl_name + " pl_ast=" + strconv.FormatFloat(pl_ast, 'f', -1, 64) + " pl_dist=" + strconv.FormatFloat(pl_dist, 'f', -1, 64))
	}

	per_deg := MaxRetrunValue / availabe_range
	out_of_ast := pl_dist - pl_ast

	pl_uday := per_deg * out_of_ast

	// Cap at MaxRetrunValue
	if pl_uday > MaxRetrunValue {
		pl_uday = MaxRetrunValue
	}

	return pl_uday
}
