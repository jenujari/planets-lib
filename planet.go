package baselib

import "math"

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
	Longitude    float64 `json:"longitude"`
	Latitude     float64 `json:"latitude"`
	Distance     float64 `json:"distance"`
	SpeedLong    float64 `json:"speedLong"`
	SpeedLat     float64 `json:"speedLat"`
	SpeedDist    float64 `json:"speedDist"`
	LongitudeDMS DMS
	LatitudeDMS  DMS
	SpeedLongDMS DMS
	Sign         string
	Nakshatra    NakshatraPada
	IsRetro      bool
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
	} else {
		normLon := normalizeAngle(p.Longitude)
		// The helper functions will also defensively handle edge cases if necessary.
		p.Sign = GetSignFrmDegree(normLon)
		p.Nakshatra = GetNakshatraPadaFromDegree(normLon)
	}

	// Determine retrograde flag only when speed is finite.
	if math.IsNaN(p.SpeedLong) || math.IsInf(p.SpeedLong, 0) {
		p.IsRetro = false
	} else {
		p.IsRetro = p.SpeedLong < 0
	}
}
