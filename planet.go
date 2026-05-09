package baselib

import (
	"errors"
	"fmt"
	"math"
)

const (
	SUN     = "Sun"
	MOON    = "Moon"
	MERCURY = "Mercury"
	VENUS   = "Venus"
	MARS    = "Mars"
	JUPITER = "Jupiter"
	SATURN  = "Saturn"
	URANUS  = "Uranus"
	NEPTUNE = "Neptune"
	PLUTO   = "Pluto"
	RAHU    = "Rahu"
	KETU    = "Ketu"

	KUTIL        = "kutil"
	ATI_VAKRA    = "ati-vakra"
	VAKRA        = "vakra"
	ATI_MAND     = "ati-mand"
	MAND         = "mand"
	MADHYAM      = "madhyam"
	SAMA         = "sama"
	SHEEGHRA     = "sheeghra"
	ATI_SHEEGHRA = "ati-sheeghra"

	LEFT_VEDHA  = "left-vedha"
	RIGHT_VEDHA = "right-vedha"
	FRONT_VEDHA = "front-vedha"
	NO_VEDHA    = "no-vedha"

	FRIEND  = "Friend"
	NEUTRAL = "Neutral"
	ENEMY   = "Enemy"
	SELF    = "Self"
)

var PLANET_NAMES = []string{SUN, MOON, MERCURY, VENUS, MARS, JUPITER, SATURN, URANUS, NEPTUNE, PLUTO, RAHU, KETU}
var PLANET_LIB_MAP = map[string]int{
	SUN:     0,
	MOON:    1,
	MERCURY: 2,
	VENUS:   3,
	MARS:    4,
	JUPITER: 5,
	SATURN:  6,
	URANUS:  7,
	NEPTUNE: 8,
	PLUTO:   9,
	RAHU:    10,
	KETU:    10,
}

// grahaMaitriChakra codifies the Permanent Planetary Relationship Table (Narpatijayacharya).
// The map is structured as [base_planet][target_planet] -> Relationship.
var grahaMaitriChakra = map[string]map[string]string{
	SUN: {
		MOON:    FRIEND,
		MARS:    FRIEND,
		JUPITER: FRIEND,
		MERCURY: NEUTRAL,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MOON: {
		SUN:     FRIEND,
		MERCURY: FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		VENUS:   NEUTRAL,
		SATURN:  NEUTRAL,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MARS: {
		SUN:     FRIEND,
		MOON:    FRIEND,
		JUPITER: FRIEND,
		VENUS:   NEUTRAL,
		SATURN:  NEUTRAL,
		MERCURY: ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	MERCURY: {
		SUN:     FRIEND,
		VENUS:   FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		SATURN:  NEUTRAL,
		MOON:    ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	JUPITER: {
		SUN:     FRIEND,
		MOON:    FRIEND,
		MARS:    FRIEND,
		SATURN:  NEUTRAL,
		MERCURY: ENEMY,
		VENUS:   ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	VENUS: {
		MERCURY: FRIEND,
		SATURN:  FRIEND,
		MARS:    NEUTRAL,
		JUPITER: NEUTRAL,
		SUN:     ENEMY,
		MOON:    ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	SATURN: {
		MERCURY: FRIEND,
		VENUS:   FRIEND,
		JUPITER: NEUTRAL,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		RAHU:    ENEMY,
		KETU:    ENEMY,
	},
	RAHU: {
		KETU:    FRIEND,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		MERCURY: ENEMY,
		JUPITER: ENEMY,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
	},
	KETU: {
		RAHU:    FRIEND,
		SUN:     ENEMY,
		MOON:    ENEMY,
		MARS:    ENEMY,
		MERCURY: ENEMY,
		JUPITER: ENEMY,
		VENUS:   ENEMY,
		SATURN:  ENEMY,
	},
}

type PlanetCord struct {
	Name          string        `json:"name"`
	Longitude     float64       `json:"longitude"`
	Latitude      float64       `json:"latitude"`
	Distance      float64       `json:"distance"`
	SpeedLong     float64       `json:"speedLong"`
	SpeedLat      float64       `json:"speedLat"`
	SpeedDist     float64       `json:"speedDist"`
	SpeedCategory string        `json:"speedCategory"`
	Vedha         string        `json:"vedha"`
	LongitudeDMS  DMS           `json:"longitudeDMS"`
	LatitudeDMS   DMS           `json:"latitudeDMS"`
	SpeedLongDMS  DMS           `json:"speedLongDMS"`
	Sign          string        `json:"sign"`
	NavamsaSign   string        `json:"navamsaSign"`
	Nakshatra     NakshatraPada `json:"nakshatra"`
	IsRetro       bool          `json:"isRetro"`
	SignLord      string        `json:"signLord"`
	SignLordship  string        `json:"signLordship"`
	Vargottama    bool          `json:"vargottama"`
}

// CalculateDerivedValues computes derived fields from raw numeric fields.
// Improvements:
// - Defensively handles NaN and +/-Inf inputs.
// - Uses a normalized longitude for sign and nakshatra mapping (without mutating the stored longitude).
// - Ensures IsRetro is only set when speed is a finite value.
func (p *PlanetCord) CalculateDerivedValues() {
	// Compute DMS representations. NewDMS / ParseFromDegree already handle NaN/Inf defensively,
	// but we still call them explicitly to populate the DMS fields consistently.
	p.LongitudeDMS = NewDMS(p.Longitude)
	p.LatitudeDMS = NewDMS(p.Latitude)
	p.SpeedLongDMS = NewDMS(p.SpeedLong)

	// Determine sign and nakshatra using a normalized longitude.
	// Do not mutate p.Longitude here so callers retain the original value.
	if math.IsNaN(p.Longitude) || math.IsInf(p.Longitude, 0) {
		// Invalid longitude -> clear sign and nakshatra to indicate unknown
		p.Sign = ""
		p.Nakshatra = NakshatraPada{}
		p.NavamsaSign = ""
		p.Vargottama = false
	} else {
		normLon := NormalizeAngle(p.Longitude)
		// The helper functions will also defensively handle edge cases if necessary.
		p.Sign = GetSignFrmDegree(normLon)
		p.Nakshatra = GetNakshatraPadaFromDegree(normLon)
		
		_, p.NavamsaSign = CalcNavanshRashi(normLon)
		if p.Sign != "" && p.Sign == p.NavamsaSign {
			p.Vargottama = true
		} else {
			p.Vargottama = false
		}
	}

	// Determine retrograde flag only when speed is finite.
	if math.IsNaN(p.SpeedLong) || math.IsInf(p.SpeedLong, 0) {
		p.IsRetro = false
	} else {
		p.IsRetro = p.SpeedLong < 0
	}

	cat, err := PlanetSpeedCategory(p.Name, p.SpeedLong)
	if err == nil {
		p.SpeedCategory = cat
	}

	vedha, err := PlanetSBCLRFVedha(p.Name, p.SpeedLong)
	if err == nil {
		p.Vedha = vedha
	}

	if p.Sign != "" {
		p.SignLord = GetSignLord(p.Sign)
		if p.SignLord != "" && p.Name != URANUS && p.Name != NEPTUNE && p.Name != PLUTO {
			p.SignLordship, _ = GetGrahaMaitri(p.Name, p.SignLord)
		}
	}
}

// GetGrahaMaitri returns the permanent relationship of the basePlanet towards the targetPlanet.
// It uses the Narpatijayacharya Planetary Friendship Table logic.
func GetGrahaMaitri(basePlanet, targetPlanet string) (string, error) {
	baseMap, ok := grahaMaitriChakra[basePlanet]
	if !ok {
		return "", fmt.Errorf("base planet %s not covered in Graha Maitri Chakra", basePlanet)
	}

	relationship, ok := baseMap[targetPlanet]
	if !ok {
		if basePlanet == targetPlanet {
			return SELF, nil
		}
		return "", fmt.Errorf("relationship from %s to %s not defined in the table", basePlanet, targetPlanet)
	}

	return relationship, nil
}

// PlanetSBCLRFVedha determines the Vedha (obstruction) type for a planet based on
// its longitudinal speed using traditional SBCLRF rules. It maps the speed to a
// speed category via PlanetSpeedCategory and then translates that category into
// one of the vedha constants (LEFT_VEDHA, RIGHT_VEDHA, FRONT_VEDHA, NO_VEDHA).
// Special rules:
//   - Rahu/Ketu always return LEFT_VEDHA.
//   - Sun and Moon have bespoke mappings for left/front/no vedha based on their
//     speed categories.
//   - For other planets Vakra/Ati-Vakra/Kutil map to RIGHT_VEDHA, and
//     Ati-Sheeghra maps to LEFT_VEDHA.
//
// Parameters:
//   - planet: Name of the planet (use provided constants like SUN, MOON, etc.).
//   - speed: Longitudinal speed (degrees per day).
//
// Returns:
//   - vedha string (one of the vedha constants or, in fallback, the speed category),
//     and an error if classification fails (e.g. invalid speed).
func PlanetSBCLRFVedha(planet string, speed float64) (string, error) {
	speedCat, err := PlanetSpeedCategory(planet, speed)
	if err != nil {
		return "", err
	}

	if planet == RAHU || planet == KETU {
		return LEFT_VEDHA, nil
	}

	if planet == SUN {
		switch speedCat {
		case SAMA:
			fallthrough
		case SHEEGHRA:
			fallthrough
		case ATI_SHEEGHRA:
			return LEFT_VEDHA, nil
		case MAND:
			fallthrough
		case MADHYAM:
			return FRONT_VEDHA, nil
		default:
			return NO_VEDHA, nil
		}
	}

	if planet == MOON {
		if speedCat == SAMA || speedCat == SHEEGHRA || speedCat == ATI_SHEEGHRA {
			return LEFT_VEDHA, nil
		}
		return FRONT_VEDHA, nil
	}

	// When Vakri, Ati Vakri or Kutil Gati - Right Vedha. Left Vedha when "Ati Sheeghra" only. No Left Vedha in Sheeghra.
	if speedCat == VAKRA || speedCat == ATI_VAKRA || speedCat == KUTIL {
		return RIGHT_VEDHA, nil
	}

	if speedCat == ATI_SHEEGHRA {
		return LEFT_VEDHA, nil
	}

	return FRONT_VEDHA, nil
}

// PlanetSpeedCategory classifies a planet's longitudinal speed into the traditional
// categories using the thresholds provided in the table.
//   - For Rahu/Ketu a simple retro negative thresholding is used (kept from previous logic).
//   - For the planets in the table (Sun, Moon, Mars, Mercury, Jupiter, Venus, Saturn)
//     the negative (retro) and positive thresholds are applied exactly as provided.
//   - For planets not in the table (outer planets) we default to MADHYAM when unknown.
func PlanetSpeedCategory(planet string, speed float64) (string, error) {
	if math.IsNaN(speed) || math.IsInf(speed, 0) {
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
