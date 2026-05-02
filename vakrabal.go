package baselib

// CalcVakraBal calculates the Vakra Bal (retrograde strength) of a planet.
// It is based on the planet's retrograde speed relative to its maximum possible retrograde speed.
//
// Parameters:
//   - pl_speed: The longitudinal speed of the planet (degrees per day).
//   - pl_name: The name of the planet (use constants like MARS, JUPITER, etc.).
//
// Returns:
//   - The calculated Vakra Bal value ranging from 0 to 100.
func CalcVakraBal(pl_speed float64, pl_name string) float64 {
	const MaxRetrunValue = 100.0
	const MinReturnValue = 0.0

	// Sun and Moon are never retrograde.
	if pl_name == SUN || pl_name == MOON {
		return MinReturnValue
	}

	// Rahu and Ketu are always retrograde in traditional Vedic astrology.
	if pl_name == RAHU || pl_name == KETU {
		return MaxRetrunValue
	}

	// If speed is positive (direct motion), Vakra Bal is 0.
	if pl_speed >= 0 {
		return MinReturnValue
	}

	// Planet Max retro speed mapping
	var maxRetroSpeed float64
	switch pl_name {
	case MARS:
		maxRetroSpeed = 0.8
	case MERCURY:
		maxRetroSpeed = 1.6
	case JUPITER:
		maxRetroSpeed = 0.35
	case VENUS:
		maxRetroSpeed = 0.8
	case SATURN:
		maxRetroSpeed = 0.25
	default:
		// Fallback for other planets or unknown names
		return MinReturnValue
	}

	vakraBal := (-pl_speed / maxRetroSpeed) * 100

	return min(vakraBal, MaxRetrunValue)
}
