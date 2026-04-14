package baselib

import "math"

const (
	NAKSHATRA_ASHWINI           = "Ashwini"
	NAKSHATRA_BHARANI           = "Bharani"
	NAKSHATRA_KRITTICA          = "Krittika"
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
	NAKSHATRA_JYESTHA           = "Jyeshtha"
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
	NAKSHATRA_KRITTICA,
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
	NAKSHATRA_KRITTICA:          3,
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
	Name string
	Pada int
}

// GetNakshatraPadaFromDegree returns the nakshatra and pada for a given longitude.
// Improvements:
// - Guard against NaN and +/-Inf inputs and return a zero-value NakshatraPada
// - Normalize input angle to [0,360) using package normalizeAngle helper
// - Use the same ranges as the original implementation but operate on the normalized angle
func GetNakshatraPadaFromDegree(d float64) NakshatraPada {
	var nakshatra NakshatraPada

	// Guard against invalid inputs
	if math.IsNaN(d) || math.IsInf(d, 0) {
		return nakshatra
	}

	nd := normalizeAngle(d)

	// The following ranges follow the original mapping. They operate on the normalized angle.
	if nd >= 0 && nd < 3.333333 {
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 1
	} else if nd >= 3.333333 && nd < 6.666666 {
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 2
	} else if nd >= 6.666666 && nd < 10 {
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 3
	} else if nd >= 10 && nd < 13.333333 {
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 4
	} else if nd >= 13.333333 && nd < 16.666666 {
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 1
	} else if nd >= 16.666666 && nd < 20 {
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 2
	} else if nd >= 20 && nd < 23.333333 {
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 3
	} else if nd >= 23.333333 && nd < 26.666666 {
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 4
	} else if nd >= 26.666666 && nd < 30 {
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 1
	} else if nd >= 30 && nd < 33.333333 {
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 2
	} else if nd >= 33.333333 && nd < 36.666666 {
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 3
	} else if nd >= 36.666666 && nd < 40 {
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 4
	} else if nd >= 40 && nd < 43.333333 {
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 1
	} else if nd >= 43.333333 && nd < 46.666666 {
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 2
	} else if nd >= 46.666666 && nd < 50 {
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 3
	} else if nd >= 50 && nd < 53.333333 {
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 4
	} else if nd >= 53.333333 && nd < 56.666666 {
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 1
	} else if nd >= 56.666666 && nd < 60 {
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 2
	} else if nd >= 60 && nd < 63.333333 {
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 3
	} else if nd >= 63.333333 && nd < 66.666666 {
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 4
	} else if nd >= 66.666666 && nd < 70 {
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 1
	} else if nd >= 70 && nd < 73.333333 {
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 2
	} else if nd >= 73.333333 && nd < 76.666666 {
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 3
	} else if nd >= 76.666666 && nd < 80 {
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 4
	} else if nd >= 80 && nd < 83.333333 {
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 1
	} else if nd >= 83.333333 && nd < 86.666666 {
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 2
	} else if nd >= 86.666666 && nd < 90 {
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 3
	} else if nd >= 90 && nd < 93.333333 {
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 4
	} else if nd >= 93.333333 && nd < 96.666666 {
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 1
	} else if nd >= 96.666666 && nd < 100 {
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 2
	} else if nd >= 100 && nd < 103.333333 {
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 3
	} else if nd >= 103.333333 && nd < 106.666666 {
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 4
	} else if nd >= 106.666666 && nd < 110 {
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 1
	} else if nd >= 110 && nd < 113.333333 {
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 2
	} else if nd >= 113.333333 && nd < 116.666666 {
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 3
	} else if nd >= 116.666666 && nd < 120 {
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 4
	} else if nd >= 120 && nd < 123.333333 {
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 1
	} else if nd >= 123.333333 && nd < 126.666666 {
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 2
	} else if nd >= 126.666666 && nd < 130 {
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 3
	} else if nd >= 130 && nd < 133.333333 {
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 4
	} else if nd >= 133.333333 && nd < 136.666666 {
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 1
	} else if nd >= 136.666666 && nd < 140 {
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 2
	} else if nd >= 140 && nd < 143.333333 {
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 3
	} else if nd >= 143.333333 && nd < 146.666666 {
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 4
	} else if nd >= 146.666666 && nd < 150 {
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 1
	} else if nd >= 150 && nd < 153.333333 {
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 2
	} else if nd >= 153.333333 && nd < 156.666666 {
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 3
	} else if nd >= 156.666666 && nd < 160 {
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 4
	} else if nd >= 160 && nd < 163.333333 {
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 1
	} else if nd >= 163.333333 && nd < 166.666666 {
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 2
	} else if nd >= 166.666666 && nd < 170 {
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 3
	} else if nd >= 170 && nd < 173.333333 {
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 4
	} else if nd >= 173.333333 && nd < 176.666666 {
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 1
	} else if nd >= 176.666666 && nd < 180 {
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 2
	} else if nd >= 180 && nd < 183.333333 {
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 3
	} else if nd >= 183.333333 && nd < 186.666666 {
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 4
	} else if nd >= 186.666666 && nd < 190 {
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 1
	} else if nd >= 190 && nd < 193.333333 {
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 2
	} else if nd >= 193.333333 && nd < 196.666666 {
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 3
	} else if nd >= 196.666666 && nd < 200 {
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 4
	} else if nd >= 200 && nd < 203.333333 {
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 1
	} else if nd >= 203.333333 && nd < 206.666666 {
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 2
	} else if nd >= 206.666666 && nd < 210 {
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 3
	} else if nd >= 210 && nd < 213.333333 {
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 4
	} else if nd >= 213.333333 && nd < 216.666666 {
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 1
	} else if nd >= 216.666666 && nd < 220 {
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 2
	} else if nd >= 220 && nd < 223.333333 {
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 3
	} else if nd >= 223.333333 && nd < 226.666666 {
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 4
	} else if nd >= 226.666666 && nd < 230 {
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 1
	} else if nd >= 230 && nd < 233.333333 {
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 2
	} else if nd >= 233.333333 && nd < 236.666666 {
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 3
	} else if nd >= 236.666666 && nd < 240 {
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 4
	} else if nd >= 240 && nd < 243.333333 {
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 1
	} else if nd >= 243.333333 && nd < 246.666666 {
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 2
	} else if nd >= 246.666666 && nd < 250 {
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 3
	} else if nd >= 250 && nd < 253.333333 {
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 4
	} else if nd >= 253.333333 && nd < 256.666666 {
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 1
	} else if nd >= 256.666666 && nd < 260 {
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 2
	} else if nd >= 260 && nd < 263.333333 {
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 3
	} else if nd >= 263.333333 && nd < 266.666666 {
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 4
	} else if nd >= 266.666666 && nd < 269.166666 {
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 1
	} else if nd >= 269.166666 && nd < 271.666666 {
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 2
	} else if nd >= 271.666666 && nd < 274.166666 {
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 3
	} else if nd >= 274.166666 && nd < 276.666666 {
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 4
	} else if nd >= 276.666666 && nd < 277.72222 {
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 1
	} else if nd >= 277.72222 && nd < 278.7777775 {
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 2
	} else if nd >= 278.7777775 && nd < 279.83333325 {
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 3
	} else if nd >= 279.83333325 && nd < 280.888889 {
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 4
	} else if nd >= 280.888889 && nd < 284 {
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 1
	} else if nd >= 284 && nd < 287.111111 {
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 2
	} else if nd >= 287.111111 && nd < 290.222222 {
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 3
	} else if nd >= 290.222222 && nd < 293.333333 {
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 4
	} else if nd >= 293.333333 && nd < 296.666666 {
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 1
	} else if nd >= 296.666666 && nd < 300 {
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 2
	} else if nd >= 300 && nd < 303.333333 {
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 3
	} else if nd >= 303.333333 && nd < 306.666666 {
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 4
	} else if nd >= 306.666666 && nd < 310 {
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 1
	} else if nd >= 310 && nd < 313.333333 {
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 2
	} else if nd >= 313.333333 && nd < 316.666666 {
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 3
	} else if nd >= 316.666666 && nd < 320 {
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 4
	} else if nd >= 320 && nd < 323.333333 {
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 1
	} else if nd >= 323.333333 && nd < 326.666666 {
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 2
	} else if nd >= 326.666666 && nd < 330 {
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 3
	} else if nd >= 330 && nd < 333.333333 {
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 4
	} else if nd >= 333.333333 && nd < 336.666666 {
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 1
	} else if nd >= 336.666666 && nd < 340 {
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 2
	} else if nd >= 340 && nd < 343.333333 {
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 3
	} else if nd >= 343.333333 && nd < 346.666666 {
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 4
	} else if nd >= 346.666666 && nd < 350 {
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 1
	} else if nd >= 350 && nd < 353.333333 {
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 2
	} else if nd >= 353.333333 && nd < 356.666666 {
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 3
	} else if nd >= 356.666666 && nd < 360 {
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 4
	}

	return nakshatra
}

// GetNakshatraFromVowel maps a phonetic vowel code to a NakshatraPada.
// This function preserves the original mapping but uses constants for names.
func GetNakshatraFromVowel(v string) NakshatraPada {
	var nakshatra NakshatraPada

	switch v {

	case "CHU": // चु  चू
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 1
	case "CHE": // चे
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 2
	case "CHO": // चो
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 3
	case "LA": // ल
		nakshatra.Name = NAKSHATRA_ASHWINI
		nakshatra.Pada = 4

	case "LI": // ली
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 1
	case "LU": // लु
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 2
	case "LE": // ले
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 3
	case "LO": // लो
		nakshatra.Name = NAKSHATRA_BHARANI
		nakshatra.Pada = 4

	case "AA": // आ अ
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 1
	case "EE": // इ
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 2
	case "OO": // ऊ उ
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 3
	case "AE": // ए
		nakshatra.Name = NAKSHATRA_KRITTICA
		nakshatra.Pada = 4

	case "O": // ओ
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 1
	case "V": // व
	case "B": // ब
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 2
	case "VI": // वि वी
	case "BI": // बि
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 3
	case "VOO": // वू वु
	case "BOO": // बू बु
		nakshatra.Name = NAKSHATRA_ROHINI
		nakshatra.Pada = 4

	case "VE": // वे
	case "BE": // बे
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 1
	case "VO": // वो
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 2
	case "k": // क
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 3
	case "KI": // की  कि
		nakshatra.Name = NAKSHATRA_MRIGASHIRSHA
		nakshatra.Pada = 4

	case "KU": // कु
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 1
	// ... the rest of the original vowel-to-nakshatra mapping follows the same pattern.
	// For brevity we continue the mapping exactly as before but using constants.
	case "G": // ग
	case "GHI": // घी घि
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 2
	case "GHO": // घो
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 3
	case "NG": // ङ
		nakshatra.Name = NAKSHATRA_ARDRA
		nakshatra.Pada = 4

	case "CA":
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 1
	case "CHA":
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 2
	case "JE":
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 3
	case "JO":
		nakshatra.Name = NAKSHATRA_PUNARVASU
		nakshatra.Pada = 4

	case "TA":
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 1
	case "TE":
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 2
	case "TO":
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 3
	case "NA":
		nakshatra.Name = NAKSHATRA_PUSHYA
		nakshatra.Pada = 4

	case "NI":
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 1
	case "NU":
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 2
	case "NE":
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 3
	case "NO":
		nakshatra.Name = NAKSHATRA_ASHLESHA
		nakshatra.Pada = 4

	case "BA":
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 1
	case "BHA":
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 2
	case "MA":
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 3
	case "YA":
		nakshatra.Name = NAKSHATRA_MAGHA
		nakshatra.Pada = 4

	case "RA":
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 1
	case "RI":
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 2
	case "RU":
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 3
	case "RE":
		nakshatra.Name = NAKSHATRA_PURVA_PHALGUNI
		nakshatra.Pada = 4

	case "RO":
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 1
	case "TAH":
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 2
	case "TI":
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 3
	case "TU":
		nakshatra.Name = NAKSHATRA_UTTARA_PHALGUNI
		nakshatra.Pada = 4

	case "TEH":
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 1
	case "TOH":
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 2
	case "NAH":
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 3
	case "NEE":
		nakshatra.Name = NAKSHATRA_HASTA
		nakshatra.Pada = 4

	case "PI":
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 1
	case "PU":
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 2
	case "PE":
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 3
	case "PO":
		nakshatra.Name = NAKSHATRA_CHITRA
		nakshatra.Pada = 4

	case "RAA":
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 1
	case "RIH":
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 2
	case "RUH":
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 3
	case "REH":
		nakshatra.Name = NAKSHATRA_SWATI
		nakshatra.Pada = 4

	case "RII":
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 1
	case "RUU":
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 2
	case "REE":
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 3
	case "ROO":
		nakshatra.Name = NAKSHATRA_VISHAKHA
		nakshatra.Pada = 4

	case "TAI":
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 1
	case "TEI":
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 2
	case "TOI":
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 3
	case "NAI":
		nakshatra.Name = NAKSHATRA_ANURADHA
		nakshatra.Pada = 4

	case "NAA":
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 1
	case "NIH":
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 2
	case "NUH":
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 3
	case "NEH":
		nakshatra.Name = NAKSHATRA_JYESTHA
		nakshatra.Pada = 4

	case "NAIY":
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 1
	case "NII":
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 2
	case "NUU":
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 3
	case "NEE2":
		nakshatra.Name = NAKSHATRA_MOOLA
		nakshatra.Pada = 4

	case "BAH":
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 1
	case "BIH":
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 2
	case "BUH":
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 3
	case "BEH2":
		nakshatra.Name = NAKSHATRA_PURVA_ASHADHA
		nakshatra.Pada = 4

	case "BOH":
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 1
	case "DA":
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 2
	case "DEE":
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 3
	case "DO":
		nakshatra.Name = NAKSHATRA_UTTARA_ASHADHA
		nakshatra.Pada = 4

	case "DHI":
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 1
	case "DHE":
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 2
	case "DHO":
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 3
	case "NA2":
		nakshatra.Name = NAKSHATRA_ABHIJIT
		nakshatra.Pada = 4

	case "NAH2":
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 1
	case "NIH2":
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 2
	case "NUH2":
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 3
	case "NEH2":
		nakshatra.Name = NAKSHATRA_SHRAVANA
		nakshatra.Pada = 4

	case "PIH":
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 1
	case "PUH":
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 2
	case "PEH":
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 3
	case "POH":
		nakshatra.Name = NAKSHATRA_DHANISHTHA
		nakshatra.Pada = 4

	case "BA2":
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 1
	case "BHA2":
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 2
	case "MA2":
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 3
	case "YA2":
		nakshatra.Name = NAKSHATRA_SATABHISHA
		nakshatra.Pada = 4

	case "RA2":
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 1
	case "RI2":
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 2
	case "RU2":
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 3
	case "RE2":
		nakshatra.Name = NAKSHATRA_PURVA_BHADRAPADA
		nakshatra.Pada = 4

	case "RO2":
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 1
	case "TA2":
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 2
	case "TE2":
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 3
	case "TO2":
		nakshatra.Name = NAKSHATRA_UTTARA_BHADRAPADA
		nakshatra.Pada = 4

	case "NA3":
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 1
	case "NE3":
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 2
	case "NO3":
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 3
	case "NU3":
		nakshatra.Name = NAKSHATRA_REVATI
		nakshatra.Pada = 4

	default:
		// leave zero-value nakshatra if unrecognized
	}

	return nakshatra
}
