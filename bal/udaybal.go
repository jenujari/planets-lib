package bal

import (
	baselib "github.com/jenujari/planets-lib"

	"math"
	"strconv"
)

// UdayBal calculates the Udaybal (rising strength) of a planet based on its
// distance from the Sun, its motion (direct or retrograde), and its specific
// astance (combustion) values.
//
// Parameters:
//   - sun_long: Longitude of the Sun in degrees.
//   - pl_long: Longitude of the planet in degrees.
//   - isRetro: Whether the planet is in retrograde motion.
//   - pl_name: Name of the planet (use constants like baselib.SUN, baselib.MOON, etc.).
//
// Returns:
//   - The calculated Udaybal value ranging from 0 to 100.
func UdayBal(sun_long, pl_long float64, isRetro bool, pl_name string) float64 {
	const MaxRetrunValue = 100.0
	const MinReturnValue = 0.0

	// If Planet is Sun then Udaybal is MaxRetrunValue
	if pl_name == baselib.SUN {
		return MaxRetrunValue
	}

	// If Planet is Rahu and Ketu then Udaybal is MinReturnValue
	if pl_name == baselib.RAHU || pl_name == baselib.KETU {
		return MinReturnValue
	}



	// Decide planet astance (pl_ast)
	var pl_ast float64
	switch pl_name {
	case baselib.MOON:
		pl_ast = 12
	case baselib.MARS:
		pl_ast = 17
	case baselib.JUPITER:
		pl_ast = 11
	case baselib.SATURN, baselib.URANUS, baselib.NEPTUNE, baselib.PLUTO:
		pl_ast = 15
	case baselib.MERCURY:
		if isRetro {
			pl_ast = 12
		} else {
			pl_ast = 14
		}
	case baselib.VENUS:
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
	sLong := baselib.NormalizeAngle(sun_long)
	pLong := baselib.NormalizeAngle(pl_long)

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
	case baselib.MERCURY:
		max_diff = 27
	case baselib.VENUS:
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
	pl_uday = min(pl_uday, MaxRetrunValue)

	return pl_uday
}
