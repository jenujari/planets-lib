package bal

import baselib "github.com/jenujari/planets-lib"

// VakraBal calculates the Vakra Bal (retrograde strength) of a planet.
// It is based on the planet's retrograde speed relative to its maximum possible retrograde speed.
//
// Parameters:
//   - pl_speed: The longitudinal speed of the planet (degrees per day).
//   - pl_name: The name of the planet (use constants like baselib.MARS, baselib.JUPITER, etc.).
//
// Returns:
//   - The calculated Vakra Bal value ranging from 0 to 100.
func VakraBal(pl_speed float64, pl_name string) float64 {
	const MaxRetrunValue = 100.0
	const MinReturnValue = 0.0

	// Sun and Moon are never retrograde.
	if pl_name == baselib.SUN || pl_name == baselib.MOON {
		return MinReturnValue
	}

	// Rahu and Ketu are always retrograde in traditional Vedic astrology.
	if pl_name == baselib.RAHU || pl_name == baselib.KETU {
		return MaxRetrunValue
	}

	// If speed is positive (direct motion), Vakra Bal is 0.
	if pl_speed >= 0 {
		return MinReturnValue
	}

	// Planet Max retro speed mapping
	var maxRetroSpeed float64
	switch pl_name {
	case baselib.MARS:
		maxRetroSpeed = 0.436666
	case baselib.MERCURY:
		maxRetroSpeed = 1.5
	case baselib.JUPITER:
		maxRetroSpeed = 0.136666
	case baselib.VENUS:
		maxRetroSpeed = 0.686666
	case baselib.SATURN:
		maxRetroSpeed = 0.0833333
	default:
		// Fallback for other planets or unknown names
		return MinReturnValue
	}

	vakraBal := (-pl_speed / maxRetroSpeed) * 100

	return min(vakraBal, MaxRetrunValue)
}
