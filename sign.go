package baselib

import "math"

const (
	SIGN_ARIES       = "Aries"
	SIGN_TAURUS      = "Taurus"
	SIGN_GEMINI      = "Gemini"
	SIGN_CANCER      = "Cancer"
	SIGN_LEO         = "Leo"
	SIGN_VIRGO       = "Virgo"
	SIGN_LIBRA       = "Libra"
	SIGN_SCORPIO     = "Scorpio"
	SIGN_SAGITTARIUS = "Sagittarius"
	SIGN_CAPRICORN   = "Capricorn"
	SIGN_AQUARIUS    = "Aquarius"
	SIGN_PISCES      = "Pisces"
)

var SIGNS = []string{
	SIGN_ARIES,
	SIGN_TAURUS,
	SIGN_GEMINI,
	SIGN_CANCER,
	SIGN_LEO,
	SIGN_VIRGO,
	SIGN_LIBRA,
	SIGN_SCORPIO,
	SIGN_SAGITTARIUS,
	SIGN_CAPRICORN,
	SIGN_AQUARIUS,
	SIGN_PISCES,
}

var SIGN_COUNT = map[string]int{
	SIGN_ARIES:       1,
	SIGN_TAURUS:      2,
	SIGN_GEMINI:      3,
	SIGN_CANCER:      4,
	SIGN_LEO:         5,
	SIGN_VIRGO:       6,
	SIGN_LIBRA:       7,
	SIGN_SCORPIO:     8,
	SIGN_SAGITTARIUS: 9,
	SIGN_CAPRICORN:   10,
	SIGN_AQUARIUS:    11,
	SIGN_PISCES:      12,
}

func GetSignFrmDegree(d float64) string {
	// Guard against invalid floating-point inputs.
	if math.IsNaN(d) || math.IsInf(d, 0) {
		return ""
	}

	// Normalize to [0,360) using shared helper. This handles negative angles
	// and large values consistently with other functions in the package.
	nd := normalizeAngle(d)

	// Each zodiac sign spans 30 degrees.
	signIndex := int(nd / 30.0)

	// Defensive clamping to ensure index is always valid.
	signIndex = max(0, min(signIndex, len(SIGNS)-1))

	return SIGNS[signIndex]
}

// RAHU and KETU are not considered in this function
// RAHU is lord of Aries and Virgo
// KETU is lord of Libra
//
// Refactor: use a small map for lookups to make the mapping data-driven
// and easier to maintain.
var signLordMap = map[string]string{
	SIGN_ARIES:       MARS,
	SIGN_TAURUS:      VENUS,
	SIGN_GEMINI:      MERCURY,
	SIGN_CANCER:      MOON,
	SIGN_LEO:         SUN,
	SIGN_VIRGO:       MERCURY,
	SIGN_LIBRA:       VENUS,
	SIGN_SCORPIO:     MARS,
	SIGN_SAGITTARIUS: JUPITER,
	SIGN_CAPRICORN:   SATURN,
	SIGN_AQUARIUS:    SATURN,
	SIGN_PISCES:      JUPITER,
}

func GetSignLord(sign string) string {
	if sign == "" {
		return ""
	}
	if lord, ok := signLordMap[sign]; ok {
		return lord
	}
	return ""
}
