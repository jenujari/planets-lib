package baselib

import "errors"

const (
	KUTIL        = "kutil"
	ATI_VAKRA    = "ati-vakra"
	VAKRA        = "vakra"
	ATI_MAND     = "ati-mand"
	MAND         = "mand"
	MADHYAM      = "madhyam"
	SAMA         = "sama"
	SHEEGHRA     = "sheeghra"
	ATI_SHEEGHRA = "ati-sheeghra"
)

// PlanetSpeedCategory classifies a planet's longitudinal speed into the traditional
// categories using the thresholds provided in the table.
//   - For Rahu/Ketu a simple retro negative thresholding is used (kept from previous logic).
//   - For the planets in the table (Sun, Moon, Mars, Mercury, Jupiter, Venus, Saturn)
//     the negative (retro) and positive thresholds are applied exactly as provided.
//   - For planets not in the table (outer planets) we default to MADHYAM when unknown.
func PlanetSpeedCategory(planet string, speed float64) (string, error) {
	if isInvalidFloat(speed) {
		return "", errors.New("invalid speed")
	}

	// Special handling for Rahu/Ketu (kept as before)
	if planet == RAHU || planet == KETU {
		// For Rahu/Ketu use specific speed thresholds:
		// speed <= -0.2145833 => kutil
		// -0.2145833 < speed <= -0.1716667 => ati-vakra
		// speed > -0.1716667 => vakra
		if speed <= -0.2145833 {
			return KUTIL, nil
		}
		if speed <= -0.1716667 {
			return ATI_VAKRA, nil
		}
		return VAKRA, nil
	}

	switch planet {
	case SUN:
		// Sun: only positive thresholds (no retro in table). If negative appears, return MADHYAM.
		if speed < 0 {
			return ATI_MAND, nil
		}
		return classifyPos(speed,
			0.9639352, // ati-mand
			0.9750926, // mand
			0.98625,   // madhyam
			0.9974074, // sama
			1.0085648, // sheeghra
			1.0197222, // ati-sheeghra (max)
		)

	case MOON:
		// Moon: only positive thresholds
		if speed < 0 {
			return ATI_MAND, nil
		}
		return classifyPos(speed,
			12.3662037,
			12.9715741,
			13.5769444,
			14.1823148,
			14.7876852,
			15.3930556,
		)

	case MARS:
		// Mars: has retro thresholds and positive thresholds
		if speed < 0 {
			return classifyNeg(speed,
				-0.3638889, // kutil
				-0.2911111, // ati-vakra
				// -0.2183333, // vakra (upper bound for vakra is less negative)
			)
		}
		return classifyPos(speed,
			0.1318981,
			0.2637963,
			0.3956944,
			0.5275926,
			0.6594907,
			0.7913889,
		)

	case MERCURY:
		if speed < 0 {
			return classifyNeg(speed,
				-1.25,
				-1,
				// -0.75,
			)
		}
		return classifyPos(speed,
			0.3670833,
			0.7341667,
			1.10125,
			1.4683333,
			1.8354167,
			2.2025,
		)

	case JUPITER:
		if speed < 0 {
			return classifyNeg(speed,
				-0.1138889,
				-0.0911111,
				// -0.0683333,
			)
		}
		return classifyPos(speed,
			0.0404167,
			0.0808333,
			0.12125,
			0.1616667,
			0.2020833,
			0.2425,
		)

	case VENUS:
		if speed < 0 {
			return classifyNeg(speed,
				-0.5722222,
				-0.4577778,
				// -0.3433333,
			)
		}
		return classifyPos(speed,
			0.2098148,
			0.4196296,
			0.6294444,
			0.8392593,
			1.0490741,
			1.2588889,
		)

	case SATURN:
		if speed < 0 {
			return classifyNeg(speed,
				-0.0694444,
				-0.0555556,
				// -0.0416667,
			)
		}
		return classifyPos(speed,
			0.021713,
			0.0434259,
			0.0651389,
			0.0868519,
			0.1085648,
			0.1302778,
		)

	default:
		// Unknown/outer planets: we don't have table thresholds. Use a safe default.
		// If retro, classify as VAKRA. Otherwise return MADHYAM.
		if speed < 0 {
			return VAKRA, nil
		}
		return MADHYAM, nil
	}
}

// Helper for positive speed classification (increasing thresholds)
func classifyPos(s float64, atiMand, mand, madhyam, sama, sheeghra, atiSheeghra float64) (string, error) {
	switch {
	case s <= atiMand:
		return ATI_MAND, nil
	case s <= mand:
		return MAND, nil
	case s <= madhyam:
		return MADHYAM, nil
	case s <= sama:
		return SAMA, nil
	case s <= sheeghra:
		return SHEEGHRA, nil
	default:
		return ATI_SHEEGHRA, nil
	}
}

// Helper for negative (retrograde) classification (threshold values are negative,
// arranged from more-negative (kutil) to less-negative (vakra)).
func classifyNeg(s float64, kutilTh, atiVakraTh float64) (string, error) {
	// If thresholds are zeroed or not applicable (0), treat as unknown and return KUTIL for retro
	// if we get here; but table explicitly gives NA for some planets (we won't call classifyNeg
	// for those planets).
	if s <= kutilTh {
		return KUTIL, nil
	}
	if s <= atiVakraTh {
		return ATI_VAKRA, nil
	}
	// any remaining negative speeds (s < 0) map to vakra
	if s < 0 {
		return VAKRA, nil
	}
	// fallback
	return MADHYAM, nil
}
