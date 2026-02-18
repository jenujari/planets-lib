package baseLib

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

func (p *PlanetCord) CalculateDerivedValues() {
	p.LongitudeDMS = NewDMS(p.Longitude)
	p.LatitudeDMS = NewDMS(p.Latitude)
	p.SpeedLongDMS = NewDMS(p.SpeedLong)
	p.Sign = GetSignFrmDegree(p.Longitude)
	p.Nakshatra = GetNakshatraPadaFromDegree(p.Longitude)
	p.IsRetro = p.SpeedLong < 0
}
