package baselib

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
)

var PLANET_NAMES = []string{SUN, MOON, MERCURY, VENUS, MARS, JUPITER, SATURN, URANUS, NEPTUNE, PLUTO, RAHU, KETU}

// PLANET_LIB_MAP maps planet names to Swiss Ephemeris body numbers (SE_SUN = 0
// through SE_PLUTO = 9, SE_MEAN_NODE = 10).
//
// RAHU and KETU deliberately share the value 10. Swiss Ephemeris defines no separate
// Ketu body: the lunar nodes are always opposite, so callers query the mean node (10)
// for Rahu and derive Ketu from the same result by adding 180 degrees. This is not a
// duplicate-key mistake — do not "fix" KETU to 11, which is SE_TRUE_NODE, a different
// node model (true rather than mean) and not Ketu.
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
	VedhaTarget   string        `json:"vedhaTarget"`
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
	if normLon, ok := ValidAngle(p.Longitude); !ok {
		// Invalid longitude -> clear sign and nakshatra to indicate unknown
		p.Sign = ""
		p.Nakshatra = NakshatraPada{}
		p.NavamsaSign = ""
		p.Vargottama = false
	} else {
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
	if isInvalidFloat(p.SpeedLong) {
		p.IsRetro = false
	} else {
		p.IsRetro = p.SpeedLong < 0
	}

	// Classify the speed once; the vedha rules reuse the same category.
	cat, err := PlanetSpeedCategory(p.Name, p.SpeedLong)
	if err == nil {
		p.SpeedCategory = cat
		p.Vedha = vedhaFromSpeedCategory(p.Name, cat)
	}

	p.VedhaTarget = VedhaTarget(p.Nakshatra.Name, p.Vedha)

	if p.Sign != "" {
		p.SignLord = GetSignLord(p.Sign)
		if p.SignLord != "" && p.Name != URANUS && p.Name != NEPTUNE && p.Name != PLUTO {
			p.SignLordship, _ = GetGrahaMaitri(p.Name, p.SignLord)
		}
	}
}
