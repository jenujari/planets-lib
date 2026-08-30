package baselib

import (
	"sort"
	"strings"
)

const (
	NAKSHATRA_ASHWINI           = "Ashwini"
	NAKSHATRA_BHARANI           = "Bharani"
	NAKSHATRA_KRITTIKA          = "Krittika"
	NAKSHATRA_ROHINI            = "Rohini"
	NAKSHATRA_MRIGASHIRSHA      = "Mrigashirsha"
	NAKSHATRA_ARDRA             = "Ardra"
	NAKSHATRA_PUNARVASU         = "Punarvasu"
	NAKSHATRA_PUSHYA            = "Pushya"
	NAKSHATRA_ASHLESHA          = "Ashlesha"
	NAKSHATRA_MAGHA             = "Magha"
	NAKSHATRA_PURVA_PHALGUNI    = "Purva Phalguni"
	NAKSHATRA_UTTARA_PHALGUNI   = "Uttara Phalguni"
	NAKSHATRA_HASTA             = "Hasta"
	NAKSHATRA_CHITRA            = "Chitra"
	NAKSHATRA_SWATI             = "Swati"
	NAKSHATRA_VISHAKHA          = "Vishakha"
	NAKSHATRA_ANURADHA          = "Anuradha"
	NAKSHATRA_JYESTHA           = "Jyestha"
	NAKSHATRA_MOOLA             = "Moola"
	NAKSHATRA_PURVA_ASHADHA     = "Purva Ashadha"
	NAKSHATRA_UTTARA_ASHADHA    = "Uttara Ashadha"
	NAKSHATRA_ABHIJIT           = "Abhijit"
	NAKSHATRA_SHRAVANA          = "Shravana"
	NAKSHATRA_DHANISHTHA        = "Dhanishtha"
	NAKSHATRA_SATABHISHA        = "Shatabhisha"
	NAKSHATRA_PURVA_BHADRAPADA  = "Purva Bhadrapada"
	NAKSHATRA_UTTARA_BHADRAPADA = "Uttara Bhadrapada"
	NAKSHATRA_REVATI            = "Revati"
)

var NAKSHATRA_NAMES = []string{
	NAKSHATRA_ASHWINI,
	NAKSHATRA_BHARANI,
	NAKSHATRA_KRITTIKA,
	NAKSHATRA_ROHINI,
	NAKSHATRA_MRIGASHIRSHA,
	NAKSHATRA_ARDRA,
	NAKSHATRA_PUNARVASU,
	NAKSHATRA_PUSHYA,
	NAKSHATRA_ASHLESHA,
	NAKSHATRA_MAGHA,
	NAKSHATRA_PURVA_PHALGUNI,
	NAKSHATRA_UTTARA_PHALGUNI,
	NAKSHATRA_HASTA,
	NAKSHATRA_CHITRA,
	NAKSHATRA_SWATI,
	NAKSHATRA_VISHAKHA,
	NAKSHATRA_ANURADHA,
	NAKSHATRA_JYESTHA,
	NAKSHATRA_MOOLA,
	NAKSHATRA_PURVA_ASHADHA,
	NAKSHATRA_UTTARA_ASHADHA,
	NAKSHATRA_ABHIJIT,
	NAKSHATRA_SHRAVANA,
	NAKSHATRA_DHANISHTHA,
	NAKSHATRA_SATABHISHA,
	NAKSHATRA_PURVA_BHADRAPADA,
	NAKSHATRA_UTTARA_BHADRAPADA,
	NAKSHATRA_REVATI,
}

var NAKSHATRA_COUNT = map[string]int{
	NAKSHATRA_ASHWINI:           1,
	NAKSHATRA_BHARANI:           2,
	NAKSHATRA_KRITTIKA:          3,
	NAKSHATRA_ROHINI:            4,
	NAKSHATRA_MRIGASHIRSHA:      5,
	NAKSHATRA_ARDRA:             6,
	NAKSHATRA_PUNARVASU:         7,
	NAKSHATRA_PUSHYA:            8,
	NAKSHATRA_ASHLESHA:          9,
	NAKSHATRA_MAGHA:             10,
	NAKSHATRA_PURVA_PHALGUNI:    11,
	NAKSHATRA_UTTARA_PHALGUNI:   12,
	NAKSHATRA_HASTA:             13,
	NAKSHATRA_CHITRA:            14,
	NAKSHATRA_SWATI:             15,
	NAKSHATRA_VISHAKHA:          16,
	NAKSHATRA_ANURADHA:          17,
	NAKSHATRA_JYESTHA:           18,
	NAKSHATRA_MOOLA:             19,
	NAKSHATRA_PURVA_ASHADHA:     20,
	NAKSHATRA_UTTARA_ASHADHA:    21,
	NAKSHATRA_ABHIJIT:           22,
	NAKSHATRA_SHRAVANA:          23,
	NAKSHATRA_DHANISHTHA:        24,
	NAKSHATRA_SATABHISHA:        25,
	NAKSHATRA_PURVA_BHADRAPADA:  26,
	NAKSHATRA_UTTARA_BHADRAPADA: 27,
	NAKSHATRA_REVATI:            28,
}

type NakshatraPada struct {
	Name string `json:"name"`
	Pada int    `json:"pada"`
}

// nakshatraStarts holds the 113 pada boundaries covering the zodiac. Pada i spans
// the half-open interval [nakshatraStarts[i], nakshatraStarts[i+1]), so the table is
// contiguous and the final entry (360) is the exclusive end of the last pada.
//
// Name and pada are not stored: the ranges run in strict order, four padas per
// nakshatra, so entry i belongs to NAKSHATRA_NAMES[i/4] with pada i%4+1. Keeping only
// the boundaries lets the whole lookup table stay in L1 cache.
//
// Boundaries are the exact values from the original explicit ranges, including the
// irregular Uttara Ashadha, Abhijit and Shravana divisions.
var nakshatraStarts = []float64{
	// Ashwini
	0, 3.333333, 6.666666, 10,
	// Bharani
	13.333333, 16.666666, 20, 23.333333,
	// Krittika
	26.666666, 30, 33.333333, 36.666666,
	// Rohini
	40, 43.333333, 46.666666, 50,
	// Mrigashirsha
	53.333333, 56.666666, 60, 63.333333,
	// Ardra
	66.666666, 70, 73.333333, 76.666666,
	// Punarvasu
	80, 83.333333, 86.666666, 90,
	// Pushya
	93.333333, 96.666666, 100, 103.333333,
	// Ashlesha
	106.666666, 110, 113.333333, 116.666666,
	// Magha
	120, 123.333333, 126.666666, 130,
	// Purva Phalguni
	133.333333, 136.666666, 140, 143.333333,
	// Uttara Phalguni
	146.666666, 150, 153.333333, 156.666666,
	// Hasta
	160, 163.333333, 166.666666, 170,
	// Chitra
	173.333333, 176.666666, 180, 183.333333,
	// Swati
	186.666666, 190, 193.333333, 196.666666,
	// Vishakha
	200, 203.333333, 206.666666, 210,
	// Anuradha
	213.333333, 216.666666, 220, 223.333333,
	// Jyestha
	226.666666, 230, 233.333333, 236.666666,
	// Moola
	240, 243.333333, 246.666666, 250,
	// Purva Ashadha
	253.333333, 256.666666, 260, 263.333333,
	// Uttara Ashadha (2.5 degree padas)
	266.666666, 269.166666, 271.666666, 274.166666,
	// Abhijit (irregular small ranges)
	276.666666, 277.72222, 278.7777775, 279.83333325,
	// Shravana
	280.888889, 284, 287.111111, 290.222222,
	// Dhanishtha
	293.333333, 296.666666, 300, 303.333333,
	// Shatabhisha
	306.666666, 310, 313.333333, 316.666666,
	// Purva Bhadrapada
	320, 323.333333, 326.666666, 330,
	// Uttara Bhadrapada
	333.333333, 336.666666, 340, 343.333333,
	// Revati
	346.666666, 350, 353.333333, 356.666666,
	// exclusive end of the last pada
	360,
}

// GetNakshatraPadaFromDegree returns the nakshatra and pada for a given longitude.
// It binary-searches the contiguous boundary table, so the ranges are half-open
// [start, end) exactly as before.
func GetNakshatraPadaFromDegree(d float64) NakshatraPada {
	var nakshatra NakshatraPada

	nd, ok := ValidAngle(d)
	if !ok {
		return nakshatra
	}

	// SearchFloat64s returns the first index whose boundary is >= nd. Step back one
	// unless nd landed exactly on a boundary, which is the start of its own pada.
	i := sort.SearchFloat64s(nakshatraStarts, nd)
	if i == len(nakshatraStarts) || nakshatraStarts[i] > nd {
		i--
	}

	// Fallback: out of table (shouldn't happen after normalization), return zero-value
	if i < 0 || i >= len(nakshatraStarts)-1 {
		return nakshatra
	}

	nakshatra.Name = NAKSHATRA_NAMES[i/4]
	nakshatra.Pada = i%4 + 1
	return nakshatra
}

// GetNakshatraFromVowel maps a phonetic vowel code to a NakshatraPada.
//
// Refactored to use a data-driven map for easier maintenance and clearer lookup.
// Keys are normalized to uppercase and trimmed.
// The original behavior is preserved: unrecognized keys return a zero-value NakshatraPada.
var nakshatraFromVowelMap = map[string]NakshatraPada{
	"CHU": {Name: NAKSHATRA_ASHWINI, Pada: 1},
	"CHE": {Name: NAKSHATRA_ASHWINI, Pada: 2},
	"CHO": {Name: NAKSHATRA_ASHWINI, Pada: 3},
	"LA":  {Name: NAKSHATRA_ASHWINI, Pada: 4},
	"LI":  {Name: NAKSHATRA_BHARANI, Pada: 1},
	"LU":  {Name: NAKSHATRA_BHARANI, Pada: 2},
	"LE":  {Name: NAKSHATRA_BHARANI, Pada: 3},
	"LO":  {Name: NAKSHATRA_BHARANI, Pada: 4},
	"AA":  {Name: NAKSHATRA_KRITTIKA, Pada: 1},
	"EE":  {Name: NAKSHATRA_KRITTIKA, Pada: 2},
	"OO":  {Name: NAKSHATRA_KRITTIKA, Pada: 3},
	"AE":  {Name: NAKSHATRA_KRITTIKA, Pada: 4},
	"O":   {Name: NAKSHATRA_ROHINI, Pada: 1},
	"V":   {Name: NAKSHATRA_ROHINI, Pada: 2},
	"B":   {Name: NAKSHATRA_ROHINI, Pada: 2},
	"VI":  {Name: NAKSHATRA_ROHINI, Pada: 3},
	"BI":  {Name: NAKSHATRA_ROHINI, Pada: 3},
	"VOO": {Name: NAKSHATRA_ROHINI, Pada: 4},
	"BOO": {Name: NAKSHATRA_ROHINI, Pada: 4},
	"VE":  {Name: NAKSHATRA_MRIGASHIRSHA, Pada: 1},
	"BE":  {Name: NAKSHATRA_MRIGASHIRSHA, Pada: 1},
	"VO":  {Name: NAKSHATRA_MRIGASHIRSHA, Pada: 2},
	"K":   {Name: NAKSHATRA_MRIGASHIRSHA, Pada: 3},
	"KI":  {Name: NAKSHATRA_MRIGASHIRSHA, Pada: 4},
	"KU":  {Name: NAKSHATRA_ARDRA, Pada: 1},
	"G":   {Name: NAKSHATRA_ARDRA, Pada: 2},
	"GHI": {Name: NAKSHATRA_ARDRA, Pada: 2},
	"GHO": {Name: NAKSHATRA_ARDRA, Pada: 3},
	"NG":  {Name: NAKSHATRA_ARDRA, Pada: 4},
	"CA":  {Name: NAKSHATRA_PUNARVASU, Pada: 1},
	"CHA": {Name: NAKSHATRA_PUNARVASU, Pada: 2},
	"JE":  {Name: NAKSHATRA_PUNARVASU, Pada: 3},
	"JO":  {Name: NAKSHATRA_PUNARVASU, Pada: 4},
	"TA":  {Name: NAKSHATRA_PUSHYA, Pada: 1},
	"TE":  {Name: NAKSHATRA_PUSHYA, Pada: 2},
	"TO":  {Name: NAKSHATRA_PUSHYA, Pada: 3},
	"NA":  {Name: NAKSHATRA_PUSHYA, Pada: 4},
	"NI":  {Name: NAKSHATRA_ASHLESHA, Pada: 1},
	"NU":  {Name: NAKSHATRA_ASHLESHA, Pada: 2},
	"NE":  {Name: NAKSHATRA_ASHLESHA, Pada: 3},
	"NO":  {Name: NAKSHATRA_ASHLESHA, Pada: 4},
	"BA":  {Name: NAKSHATRA_MAGHA, Pada: 1},
	"BHA": {Name: NAKSHATRA_MAGHA, Pada: 2},
	"MA":  {Name: NAKSHATRA_MAGHA, Pada: 3},
	"YA":  {Name: NAKSHATRA_MAGHA, Pada: 4},
	"RA":  {Name: NAKSHATRA_PURVA_PHALGUNI, Pada: 1},
	"RI":  {Name: NAKSHATRA_PURVA_PHALGUNI, Pada: 2},
	"RU":  {Name: NAKSHATRA_PURVA_PHALGUNI, Pada: 3},
	"RE":  {Name: NAKSHATRA_PURVA_PHALGUNI, Pada: 4},
	"RO":  {Name: NAKSHATRA_UTTARA_PHALGUNI, Pada: 1},
	"TAH": {Name: NAKSHATRA_UTTARA_PHALGUNI, Pada: 2},
	"TI":  {Name: NAKSHATRA_UTTARA_PHALGUNI, Pada: 3},
	"TU":  {Name: NAKSHATRA_UTTARA_PHALGUNI, Pada: 4},

	//nolint:misspell // 'TEH' is intentional (not a typo)
	"TEH": {Name: NAKSHATRA_HASTA, Pada: 1},

	"TOH":  {Name: NAKSHATRA_HASTA, Pada: 2},
	"NAH":  {Name: NAKSHATRA_HASTA, Pada: 3},
	"NEE":  {Name: NAKSHATRA_HASTA, Pada: 4},
	"PI":   {Name: NAKSHATRA_CHITRA, Pada: 1},
	"PU":   {Name: NAKSHATRA_CHITRA, Pada: 2},
	"PE":   {Name: NAKSHATRA_CHITRA, Pada: 3},
	"PO":   {Name: NAKSHATRA_CHITRA, Pada: 4},
	"RAA":  {Name: NAKSHATRA_SWATI, Pada: 1},
	"RIH":  {Name: NAKSHATRA_SWATI, Pada: 2},
	"RUH":  {Name: NAKSHATRA_SWATI, Pada: 3},
	"REH":  {Name: NAKSHATRA_SWATI, Pada: 4},
	"RII":  {Name: NAKSHATRA_VISHAKHA, Pada: 1},
	"RUU":  {Name: NAKSHATRA_VISHAKHA, Pada: 2},
	"REE":  {Name: NAKSHATRA_VISHAKHA, Pada: 3},
	"ROO":  {Name: NAKSHATRA_VISHAKHA, Pada: 4},
	"TAI":  {Name: NAKSHATRA_ANURADHA, Pada: 1},
	"TEI":  {Name: NAKSHATRA_ANURADHA, Pada: 2},
	"TOI":  {Name: NAKSHATRA_ANURADHA, Pada: 3},
	"NAI":  {Name: NAKSHATRA_ANURADHA, Pada: 4},
	"NAA":  {Name: NAKSHATRA_JYESTHA, Pada: 1},
	"NIH":  {Name: NAKSHATRA_JYESTHA, Pada: 2},
	"NUH":  {Name: NAKSHATRA_JYESTHA, Pada: 3},
	"NEH":  {Name: NAKSHATRA_JYESTHA, Pada: 4},
	"NAIY": {Name: NAKSHATRA_MOOLA, Pada: 1},
	"NII":  {Name: NAKSHATRA_MOOLA, Pada: 2},
	"NUU":  {Name: NAKSHATRA_MOOLA, Pada: 3},
	"NEE2": {Name: NAKSHATRA_MOOLA, Pada: 4},
	"BAH":  {Name: NAKSHATRA_PURVA_ASHADHA, Pada: 1},
	"BIH":  {Name: NAKSHATRA_PURVA_ASHADHA, Pada: 2},
	"BUH":  {Name: NAKSHATRA_PURVA_ASHADHA, Pada: 3},
	"BEH2": {Name: NAKSHATRA_PURVA_ASHADHA, Pada: 4},
	"BOH":  {Name: NAKSHATRA_UTTARA_ASHADHA, Pada: 1},
	"DA":   {Name: NAKSHATRA_UTTARA_ASHADHA, Pada: 2},
	"DEE":  {Name: NAKSHATRA_UTTARA_ASHADHA, Pada: 3},
	"DO":   {Name: NAKSHATRA_UTTARA_ASHADHA, Pada: 4},
	"DHI":  {Name: NAKSHATRA_ABHIJIT, Pada: 1},
	"DHE":  {Name: NAKSHATRA_ABHIJIT, Pada: 2},
	"DHO":  {Name: NAKSHATRA_ABHIJIT, Pada: 3},
	"NA2":  {Name: NAKSHATRA_ABHIJIT, Pada: 4},
	"NAH2": {Name: NAKSHATRA_SHRAVANA, Pada: 1},
	"NIH2": {Name: NAKSHATRA_SHRAVANA, Pada: 2},
	"NUH2": {Name: NAKSHATRA_SHRAVANA, Pada: 3},
	"NEH2": {Name: NAKSHATRA_SHRAVANA, Pada: 4},
	"PIH":  {Name: NAKSHATRA_DHANISHTHA, Pada: 1},
	"PUH":  {Name: NAKSHATRA_DHANISHTHA, Pada: 2},
	"PEH":  {Name: NAKSHATRA_DHANISHTHA, Pada: 3},
	"POH":  {Name: NAKSHATRA_DHANISHTHA, Pada: 4},
	"BA2":  {Name: NAKSHATRA_SATABHISHA, Pada: 1},
	"BHA2": {Name: NAKSHATRA_SATABHISHA, Pada: 2},
	"MA2":  {Name: NAKSHATRA_SATABHISHA, Pada: 3},
	"YA2":  {Name: NAKSHATRA_SATABHISHA, Pada: 4},
	"RA2":  {Name: NAKSHATRA_PURVA_BHADRAPADA, Pada: 1},
	"RI2":  {Name: NAKSHATRA_PURVA_BHADRAPADA, Pada: 2},
	"RU2":  {Name: NAKSHATRA_PURVA_BHADRAPADA, Pada: 3},
	"RE2":  {Name: NAKSHATRA_PURVA_BHADRAPADA, Pada: 4},
	"RO2":  {Name: NAKSHATRA_UTTARA_BHADRAPADA, Pada: 1},
	"TA2":  {Name: NAKSHATRA_UTTARA_BHADRAPADA, Pada: 2},
	"TE2":  {Name: NAKSHATRA_UTTARA_BHADRAPADA, Pada: 3},
	"TO2":  {Name: NAKSHATRA_UTTARA_BHADRAPADA, Pada: 4},
	"NA3":  {Name: NAKSHATRA_REVATI, Pada: 1},
	"NE3":  {Name: NAKSHATRA_REVATI, Pada: 2},
	"NO3":  {Name: NAKSHATRA_REVATI, Pada: 3},
	"NU3":  {Name: NAKSHATRA_REVATI, Pada: 4},
}

func GetNakshatraFromVowel(v string) NakshatraPada {
	key := strings.ToUpper(strings.TrimSpace(v))
	if val, ok := nakshatraFromVowelMap[key]; ok {
		return val
	}
	return NakshatraPada{}
}
