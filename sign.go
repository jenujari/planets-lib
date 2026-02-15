package baseLib

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
	d = math.Abs(d)

	if d > 360 {
		d = math.Mod(d, 360)
	}

	signIndex := int(d / 30)

	return SIGNS[signIndex]
}

// RAHU and KETU are not considered in this function
// RAHU is lord of Aries and Virgo
// KETU is lord of Libra
func GetSignLord(sign string) string {
	switch sign {
	case SIGN_ARIES:
		return MARS
	case SIGN_TAURUS:
		return VENUS
	case SIGN_GEMINI:
		return MERCURY
	case SIGN_CANCER:
		return MOON
	case SIGN_LEO:
		return SUN
	case SIGN_VIRGO:
		return MERCURY
	case SIGN_LIBRA:
		return VENUS
	case SIGN_SCORPIO:
		return MARS
	case SIGN_SAGITTARIUS:
		return JUPITER
	case SIGN_CAPRICORN:
		return SATURN
	case SIGN_AQUARIUS:
		return SATURN
	case SIGN_PISCES:
		return JUPITER
	default:
		return ""
	}
}
