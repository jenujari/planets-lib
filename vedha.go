package baselib

const (
	LEFT_VEDHA  = "left"
	RIGHT_VEDHA = "right"
	FRONT_VEDHA = "front"
	NO_VEDHA    = "no"
)

var (
	LeftVedhaMap = map[string]string{
		NAKSHATRA_ASHWINI:           NAKSHATRA_ROHINI,
		NAKSHATRA_BHARANI:           NAKSHATRA_KRITTIKA,
		NAKSHATRA_KRITTIKA:          NAKSHATRA_VISHAKHA,
		NAKSHATRA_ROHINI:            NAKSHATRA_SWATI,
		NAKSHATRA_MRIGASHIRSHA:      NAKSHATRA_CHITRA,
		NAKSHATRA_ARDRA:             NAKSHATRA_HASTA,
		NAKSHATRA_PUNARVASU:         NAKSHATRA_UTTARA_PHALGUNI,
		NAKSHATRA_PUSHYA:            NAKSHATRA_PURVA_PHALGUNI,
		NAKSHATRA_ASHLESHA:          NAKSHATRA_MAGHA,
		NAKSHATRA_MAGHA:             NAKSHATRA_SHRAVANA,
		NAKSHATRA_PURVA_PHALGUNI:    NAKSHATRA_ABHIJIT,
		NAKSHATRA_UTTARA_PHALGUNI:   NAKSHATRA_UTTARA_ASHADHA,
		NAKSHATRA_HASTA:             NAKSHATRA_PURVA_ASHADHA,
		NAKSHATRA_CHITRA:            NAKSHATRA_MOOLA,
		NAKSHATRA_SWATI:             NAKSHATRA_JYESTHA,
		NAKSHATRA_VISHAKHA:          NAKSHATRA_ANURADHA,
		NAKSHATRA_ANURADHA:          NAKSHATRA_BHARANI,
		NAKSHATRA_JYESTHA:           NAKSHATRA_ASHWINI,
		NAKSHATRA_MOOLA:             NAKSHATRA_REVATI,
		NAKSHATRA_PURVA_ASHADHA:     NAKSHATRA_UTTARA_BHADRAPADA,
		NAKSHATRA_UTTARA_ASHADHA:    NAKSHATRA_PURVA_BHADRAPADA,
		NAKSHATRA_ABHIJIT:           NAKSHATRA_SATABHISHA,
		NAKSHATRA_SHRAVANA:          NAKSHATRA_DHANISHTHA,
		NAKSHATRA_DHANISHTHA:        NAKSHATRA_ASHLESHA,
		NAKSHATRA_SATABHISHA:        NAKSHATRA_PUSHYA,
		NAKSHATRA_PURVA_BHADRAPADA:  NAKSHATRA_PUNARVASU,
		NAKSHATRA_UTTARA_BHADRAPADA: NAKSHATRA_ARDRA,
		NAKSHATRA_REVATI:            NAKSHATRA_MRIGASHIRSHA,
	}
	RightVedhaMap = map[string]string{
		NAKSHATRA_ASHWINI:           NAKSHATRA_JYESTHA,
		NAKSHATRA_BHARANI:           NAKSHATRA_ANURADHA,
		NAKSHATRA_KRITTIKA:          NAKSHATRA_BHARANI,
		NAKSHATRA_ROHINI:            NAKSHATRA_ASHWINI,
		NAKSHATRA_MRIGASHIRSHA:      NAKSHATRA_REVATI,
		NAKSHATRA_ARDRA:             NAKSHATRA_UTTARA_BHADRAPADA,
		NAKSHATRA_PUNARVASU:         NAKSHATRA_PURVA_BHADRAPADA,
		NAKSHATRA_PUSHYA:            NAKSHATRA_SATABHISHA,
		NAKSHATRA_ASHLESHA:          NAKSHATRA_DHANISHTHA,
		NAKSHATRA_MAGHA:             NAKSHATRA_ASHLESHA,
		NAKSHATRA_PURVA_PHALGUNI:    NAKSHATRA_PUSHYA,
		NAKSHATRA_UTTARA_PHALGUNI:   NAKSHATRA_PUNARVASU,
		NAKSHATRA_HASTA:             NAKSHATRA_ARDRA,
		NAKSHATRA_CHITRA:            NAKSHATRA_MRIGASHIRSHA,
		NAKSHATRA_SWATI:             NAKSHATRA_ROHINI,
		NAKSHATRA_VISHAKHA:          NAKSHATRA_KRITTIKA,
		NAKSHATRA_ANURADHA:          NAKSHATRA_VISHAKHA,
		NAKSHATRA_JYESTHA:           NAKSHATRA_SWATI,
		NAKSHATRA_MOOLA:             NAKSHATRA_CHITRA,
		NAKSHATRA_PURVA_ASHADHA:     NAKSHATRA_HASTA,
		NAKSHATRA_UTTARA_ASHADHA:    NAKSHATRA_UTTARA_PHALGUNI,
		NAKSHATRA_ABHIJIT:           NAKSHATRA_PURVA_PHALGUNI,
		NAKSHATRA_SHRAVANA:          NAKSHATRA_MAGHA,
		NAKSHATRA_DHANISHTHA:        NAKSHATRA_SHRAVANA,
		NAKSHATRA_SATABHISHA:        NAKSHATRA_ABHIJIT,
		NAKSHATRA_PURVA_BHADRAPADA:  NAKSHATRA_UTTARA_ASHADHA,
		NAKSHATRA_UTTARA_BHADRAPADA: NAKSHATRA_PURVA_ASHADHA,
		NAKSHATRA_REVATI:            NAKSHATRA_MOOLA,
	}
	FrontVedhaMap = map[string]string{
		NAKSHATRA_ASHWINI:           NAKSHATRA_PURVA_PHALGUNI,
		NAKSHATRA_BHARANI:           NAKSHATRA_MAGHA,
		NAKSHATRA_KRITTIKA:          NAKSHATRA_SHRAVANA,
		NAKSHATRA_ROHINI:            NAKSHATRA_ABHIJIT,
		NAKSHATRA_MRIGASHIRSHA:      NAKSHATRA_UTTARA_ASHADHA,
		NAKSHATRA_ARDRA:             NAKSHATRA_PURVA_ASHADHA,
		NAKSHATRA_PUNARVASU:         NAKSHATRA_MOOLA,
		NAKSHATRA_PUSHYA:            NAKSHATRA_JYESTHA,
		NAKSHATRA_ASHLESHA:          NAKSHATRA_ANURADHA,
		NAKSHATRA_MAGHA:             NAKSHATRA_BHARANI,
		NAKSHATRA_PURVA_PHALGUNI:    NAKSHATRA_ASHWINI,
		NAKSHATRA_UTTARA_PHALGUNI:   NAKSHATRA_REVATI,
		NAKSHATRA_HASTA:             NAKSHATRA_UTTARA_BHADRAPADA,
		NAKSHATRA_CHITRA:            NAKSHATRA_PURVA_BHADRAPADA,
		NAKSHATRA_SWATI:             NAKSHATRA_SATABHISHA,
		NAKSHATRA_VISHAKHA:          NAKSHATRA_DHANISHTHA,
		NAKSHATRA_ANURADHA:          NAKSHATRA_ASHLESHA,
		NAKSHATRA_JYESTHA:           NAKSHATRA_PUSHYA,
		NAKSHATRA_MOOLA:             NAKSHATRA_PUNARVASU,
		NAKSHATRA_PURVA_ASHADHA:     NAKSHATRA_ARDRA,
		NAKSHATRA_UTTARA_ASHADHA:    NAKSHATRA_MRIGASHIRSHA,
		NAKSHATRA_ABHIJIT:           NAKSHATRA_ROHINI,
		NAKSHATRA_SHRAVANA:          NAKSHATRA_KRITTIKA,
		NAKSHATRA_DHANISHTHA:        NAKSHATRA_VISHAKHA,
		NAKSHATRA_SATABHISHA:        NAKSHATRA_SWATI,
		NAKSHATRA_PURVA_BHADRAPADA:  NAKSHATRA_CHITRA,
		NAKSHATRA_UTTARA_BHADRAPADA: NAKSHATRA_HASTA,
		NAKSHATRA_REVATI:            NAKSHATRA_UTTARA_PHALGUNI,
	}
)

func VedhaTarget(f, d string) string {
	if f == "" || d == "" {
		return ""
	}

	switch d {
	case NO_VEDHA:
		return ""
	case LEFT_VEDHA:
		if v, ok := LeftVedhaMap[f]; ok {
			return v
		}
	case RIGHT_VEDHA:
		if v, ok := RightVedhaMap[f]; ok {
			return v
		}
	case FRONT_VEDHA:
		if v, ok := FrontVedhaMap[f]; ok {
			return v
		}
	}

	return ""
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

	return vedhaFromSpeedCategory(planet, speedCat), nil
}

// vedhaFromSpeedCategory applies the SBCLRF vedha rules to an already-classified
// speed category. Callers that have run PlanetSpeedCategory use this directly to
// avoid classifying the same speed twice.
func vedhaFromSpeedCategory(planet, speedCat string) string {
	if planet == RAHU || planet == KETU {
		return LEFT_VEDHA
	}

	if planet == SUN {
		switch speedCat {
		case SAMA, SHEEGHRA, ATI_SHEEGHRA:
			return LEFT_VEDHA
		case MAND, MADHYAM:
			return FRONT_VEDHA
		default:
			return NO_VEDHA
		}
	}

	if planet == MOON {
		if speedCat == SAMA || speedCat == SHEEGHRA || speedCat == ATI_SHEEGHRA {
			return LEFT_VEDHA
		}
		return FRONT_VEDHA
	}

	// When Vakri, Ati Vakri or Kutil Gati - Right Vedha. Left Vedha when "Ati Sheeghra" only. No Left Vedha in Sheeghra.
	if speedCat == VAKRA || speedCat == ATI_VAKRA || speedCat == KUTIL {
		return RIGHT_VEDHA
	}

	if speedCat == ATI_SHEEGHRA {
		return LEFT_VEDHA
	}

	return FRONT_VEDHA
}
