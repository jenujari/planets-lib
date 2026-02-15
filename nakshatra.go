package baseLib

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

func GetNakshatraPadaFromDegree(d float64) NakshatraPada {
	var nakshatra NakshatraPada

	if d >= 0 && d < 3.333333 {
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 1
	} else if d >= 3.333333 && d < 6.666666 {
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 2
	} else if d >= 6.666666 && d < 10 {
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 3
	} else if d >= 10 && d < 13.333333 {
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 4
	} else if d >= 13.333333 && d < 16.666666 {
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 1
	} else if d >= 16.666666 && d < 20 {
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 2
	} else if d >= 20 && d < 23.333333 {
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 3
	} else if d >= 23.333333 && d < 26.666666 {
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 4
	} else if d >= 26.666666 && d < 30 {
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 1
	} else if d >= 30 && d < 33.333333 {
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 2
	} else if d >= 33.333333 && d < 36.666666 {
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 3
	} else if d >= 36.666666 && d < 40 {
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 4
	} else if d >= 40 && d < 43.333333 {
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 1
	} else if d >= 43.333333 && d < 46.666666 {
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 2
	} else if d >= 46.666666 && d < 50 {
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 3
	} else if d >= 50 && d < 53.333333 {
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 4
	} else if d >= 53.333333 && d < 56.666666 {
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 1
	} else if d >= 56.666666 && d < 60 {
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 2
	} else if d >= 60 && d < 63.333333 {
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 3
	} else if d >= 63.333333 && d < 66.666666 {
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 4
	} else if d >= 66.666666 && d < 70 {
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 1
	} else if d >= 70 && d < 73.333333 {
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 2
	} else if d >= 73.333333 && d < 76.666666 {
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 3
	} else if d >= 76.666666 && d < 80 {
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 4
	} else if d >= 80 && d < 83.333333 {
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 1
	} else if d >= 83.333333 && d < 86.666666 {
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 2
	} else if d >= 86.666666 && d < 90 {
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 3
	} else if d >= 90 && d < 93.333333 {
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 4
	} else if d >= 93.333333 && d < 96.666666 {
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 1
	} else if d >= 96.666666 && d < 100 {
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 2
	} else if d >= 100 && d < 103.333333 {
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 3
	} else if d >= 103.333333 && d < 106.666666 {
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 4
	} else if d >= 106.666666 && d < 110 {
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 1
	} else if d >= 110 && d < 113.333333 {
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 2
	} else if d >= 113.333333 && d < 116.666666 {
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 3
	} else if d >= 116.666666 && d < 120 {
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 4
	} else if d >= 120 && d < 123.333333 {
		nakshatra.Name = "Magha"
		nakshatra.Pada = 1
	} else if d >= 123.333333 && d < 126.666666 {
		nakshatra.Name = "Magha"
		nakshatra.Pada = 2
	} else if d >= 126.666666 && d < 130 {
		nakshatra.Name = "Magha"
		nakshatra.Pada = 3
	} else if d >= 130 && d < 133.333333 {
		nakshatra.Name = "Magha"
		nakshatra.Pada = 4
	} else if d >= 133.333333 && d < 136.666666 {
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 1
	} else if d >= 136.666666 && d < 140 {
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 2
	} else if d >= 140 && d < 143.333333 {
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 3
	} else if d >= 143.333333 && d < 146.666666 {
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 4
	} else if d >= 146.666666 && d < 150 {
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 1
	} else if d >= 150 && d < 153.333333 {
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 2
	} else if d >= 153.333333 && d < 156.666666 {
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 3
	} else if d >= 156.666666 && d < 160 {
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 4
	} else if d >= 160 && d < 163.333333 {
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 1
	} else if d >= 163.333333 && d < 166.666666 {
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 2
	} else if d >= 166.666666 && d < 170 {
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 3
	} else if d >= 170 && d < 173.333333 {
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 4
	} else if d >= 173.333333 && d < 176.666666 {
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 1
	} else if d >= 176.666666 && d < 180 {
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 2
	} else if d >= 180 && d < 183.333333 {
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 3
	} else if d >= 183.333333 && d < 186.666666 {
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 4
	} else if d >= 186.666666 && d < 190 {
		nakshatra.Name = "Swati"
		nakshatra.Pada = 1
	} else if d >= 190 && d < 193.333333 {
		nakshatra.Name = "Swati"
		nakshatra.Pada = 2
	} else if d >= 193.333333 && d < 196.666666 {
		nakshatra.Name = "Swati"
		nakshatra.Pada = 3
	} else if d >= 196.666666 && d < 200 {
		nakshatra.Name = "Swati"
		nakshatra.Pada = 4
	} else if d >= 200 && d < 203.333333 {
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 1
	} else if d >= 203.333333 && d < 206.666666 {
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 2
	} else if d >= 206.666666 && d < 210 {
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 3
	} else if d >= 210 && d < 213.333333 {
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 4
	} else if d >= 213.333333 && d < 216.666666 {
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 1
	} else if d >= 216.666666 && d < 220 {
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 2
	} else if d >= 220 && d < 223.333333 {
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 3
	} else if d >= 223.333333 && d < 226.666666 {
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 4
	} else if d >= 226.666666 && d < 230 {
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 1
	} else if d >= 230 && d < 233.333333 {
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 2
	} else if d >= 233.333333 && d < 236.666666 {
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 3
	} else if d >= 236.666666 && d < 240 {
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 4
	} else if d >= 240 && d < 243.333333 {
		nakshatra.Name = "Moola"
		nakshatra.Pada = 1
	} else if d >= 243.333333 && d < 246.666666 {
		nakshatra.Name = "Moola"
		nakshatra.Pada = 2
	} else if d >= 246.666666 && d < 250 {
		nakshatra.Name = "Moola"
		nakshatra.Pada = 3
	} else if d >= 250 && d < 253.333333 {
		nakshatra.Name = "Moola"
		nakshatra.Pada = 4
	} else if d >= 253.333333 && d < 256.666666 {
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 1
	} else if d >= 256.666666 && d < 260 {
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 2
	} else if d >= 260 && d < 263.333333 {
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 3
	} else if d >= 263.333333 && d < 266.666666 {
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 4
	} else if d >= 266.666666 && d < 269.166666 {
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 1
	} else if d >= 269.166666 && d < 271.666666 {
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 2
	} else if d >= 271.666666 && d < 274.166666 {
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 3
	} else if d >= 274.166666 && d < 276.666666 {
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 4
	} else if d >= 276.666666 && d < 277.72222 {
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 1
	} else if d >= 277.72222 && d < 278.7777775 {
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 2
	} else if d >= 278.7777775 && d < 279.83333325 {
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 3
	} else if d >= 279.83333325 && d < 280.888889 {
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 4
	} else if d >= 280.888889 && d < 284 {
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 1
	} else if d >= 284 && d < 287.111111 {
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 2
	} else if d >= 287.111111 && d < 290.222222 {
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 3
	} else if d >= 290.222222 && d < 293.333333 {
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 4
	} else if d >= 293.333333 && d < 296.666666 {
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 1
	} else if d >= 296.666666 && d < 300 {
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 2
	} else if d >= 300 && d < 303.333333 {
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 3
	} else if d >= 303.333333 && d < 306.666666 {
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 4
	} else if d >= 306.666666 && d < 310 {
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 1
	} else if d >= 310 && d < 313.333333 {
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 2
	} else if d >= 313.333333 && d < 316.666666 {
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 3
	} else if d >= 316.666666 && d < 320 {
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 4
	} else if d >= 320 && d < 323.333333 {
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 1
	} else if d >= 323.333333 && d < 326.666666 {
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 2
	} else if d >= 326.666666 && d < 330 {
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 3
	} else if d >= 330 && d < 333.333333 {
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 4
	} else if d >= 333.333333 && d < 336.666666 {
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 1
	} else if d >= 336.666666 && d < 340 {
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 2
	} else if d >= 340 && d < 343.333333 {
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 3
	} else if d >= 343.333333 && d < 346.666666 {
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 4
	} else if d >= 346.666666 && d < 350 {
		nakshatra.Name = "Revati"
		nakshatra.Pada = 1
	} else if d >= 350 && d < 353.333333 {
		nakshatra.Name = "Revati"
		nakshatra.Pada = 2
	} else if d >= 353.333333 && d < 356.666666 {
		nakshatra.Name = "Revati"
		nakshatra.Pada = 3
	} else if d >= 356.666666 && d < 360 {
		nakshatra.Name = "Revati"
		nakshatra.Pada = 4
	}

	return nakshatra
}

func GetNakshatraFromVowel(v string) NakshatraPada {
	var nakshatra NakshatraPada

	switch v {

	case "CHU": // चु  चू
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 1
	case "CHE": // चे
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 2
	case "CHO": // चो
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 3
	case "LA": // ल
		nakshatra.Name = "Ashwini"
		nakshatra.Pada = 4

	case "LI": // ली
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 1
	case "LU": // लु
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 2
	case "LE": // ले
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 3
	case "LO": // लो
		nakshatra.Name = "Bharani"
		nakshatra.Pada = 4

	case "AA": // आ अ
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 1
	case "EE": // इ
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 2
	case "OO": // ऊ उ
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 3
	case "AE": // ए
		nakshatra.Name = "Krittika"
		nakshatra.Pada = 4

	case "O": // ओ
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 1
	case "V": // व
	case "B": // ब
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 2
	case "VI": // वि वी
	case "BI": // बि
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 3
	case "VOO": // वू वु
	case "BOO": // बू बु
		nakshatra.Name = "Rohini"
		nakshatra.Pada = 4

	case "VE": // वे
	case "BE": // बे
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 1
	case "VO": // वो
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 2
	case "k": // क
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 3
	case "KI": // की  कि
		nakshatra.Name = "Mrigashirsha"
		nakshatra.Pada = 4

	case "KU": // कु
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 1
	case "GHA": // घ
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 2
	case "ANG": //ग्ः
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 3
	case "CHA": // छ
		nakshatra.Name = "Ardra"
		nakshatra.Pada = 4

	case "KE": // के
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 1
	case "KO": // को
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 2
	case "HA": // ह
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 3
	case "HI": // हि
		nakshatra.Name = "Punarvasu"
		nakshatra.Pada = 4

	case "HU": // हु
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 1
	case "HE": // हे
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 2
	case "HO": // हो
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 3
	case "DD": // ड
		nakshatra.Name = "Pushya"
		nakshatra.Pada = 4

	case "DDI": // डी
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 1
	case "DDU": //डू
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 2
	case "DDE": // डे
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 3
	case "DDO": // डो
		nakshatra.Name = "Ashlesha"
		nakshatra.Pada = 4

	case "M": // म
		nakshatra.Name = "Magha"
		nakshatra.Pada = 1
	case "MI": // मि
		nakshatra.Name = "Magha"
		nakshatra.Pada = 2
	case "MU": // मु
		nakshatra.Name = "Magha"
		nakshatra.Pada = 3
	case "ME": // मे
		nakshatra.Name = "Magha"
		nakshatra.Pada = 4

	case "MO": // मो
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 1
	case "TT": // ट
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 2
	case "TTI": // टी
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 3
	case "TTU": //टू
		nakshatra.Name = "Purva Phalguni"
		nakshatra.Pada = 4

	case "TTE": // टे
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 1
	case "TTO": // टो
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 2
	case "PA": //प
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 3
	case "PI": // पी
		nakshatra.Name = "Uttara Phalguni"
		nakshatra.Pada = 4

	case "PU": // पू
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 1
	case "SHA": // ष
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 2
	case "ANA": // ण
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 3
	case "THA": // ठ
		nakshatra.Name = "Hasta"
		nakshatra.Pada = 4

	case "PE": // पे
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 1
	case "PO": // पो
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 2
	case "RA": // र
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 3
	case "RI": // री
		nakshatra.Name = "Chitra"
		nakshatra.Pada = 4

	case "RU": // रु
		nakshatra.Name = "Swati"
		nakshatra.Pada = 1
	case "RE": // रे
		nakshatra.Name = "Swati"
		nakshatra.Pada = 2
	case "RO": // रो
		nakshatra.Name = "Swati"
		nakshatra.Pada = 3
	case "TA": // ता त
		nakshatra.Name = "Swati"
		nakshatra.Pada = 4

	case "TI": // ति
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 1
	case "TU": // तु
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 2
	case "TE": // ते
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 3
	case "TO": // तो
		nakshatra.Name = "Vishakha"
		nakshatra.Pada = 4

	case "NA": // ना न
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 1
	case "NI": // नी
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 2
	case "NU": // नु
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 3
	case "NE": // ने
		nakshatra.Name = "Anuradha"
		nakshatra.Pada = 4

	case "NO": // नो
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 1
	case "YA": // य
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 2
	case "YI": // यि
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 3
	case "YU": // यु
		nakshatra.Name = "Jyeshtha"
		nakshatra.Pada = 4

	case "YE": // ये
		nakshatra.Name = "Moola"
		nakshatra.Pada = 1
	case "YO": // यो
		nakshatra.Name = "Moola"
		nakshatra.Pada = 2
	case "BH": // भ
		nakshatra.Name = "Moola"
		nakshatra.Pada = 3
	case "BHI": // भी
		nakshatra.Name = "Moola"
		nakshatra.Pada = 4

	case "BHU": // भू
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 1
	case "DHA": // ध
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 2
	case "PHA": // फ
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 3
	case "DDHA": // ढ़
		nakshatra.Name = "Purva Ashadha"
		nakshatra.Pada = 4

	case "BHE": // भे
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 1
	case "BHO": // भो
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 2
	case "JA": // जा ज
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 3
	case "JI": // जी
		nakshatra.Name = "Uttara Ashadha"
		nakshatra.Pada = 4

	case "JU": // जू
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 1
	case "JE": // जे
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 2
	case "JO": // जो
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 3
	case "KHA": // ख
		nakshatra.Name = "Abhijit"
		nakshatra.Pada = 4

	case "KHI": // खि
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 1
	case "KHU": // खु
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 2
	case "KHE": // खे
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 3
	case "KHO": // खो
		nakshatra.Name = "Shravana"
		nakshatra.Pada = 4

	case "GA": // ग
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 1
	case "GI": // गि
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 2
	case "GU": // गु
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 3
	case "GE": // गे
		nakshatra.Name = "Dhanishta"
		nakshatra.Pada = 4

	case "GO": // गो
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 1
	case "SA": // सा स श
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 2
	case "SI": // सी
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 3
	case "SU": // सु
		nakshatra.Name = "Shatabhisha"
		nakshatra.Pada = 4

	case "SE": // से
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 1
	case "SO": // सो
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 2
	case "DA": // द
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 3
	case "DI": // दी
		nakshatra.Name = "Purva Bhadrapada"
		nakshatra.Pada = 4

	case "DU": // दु
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 1
	case "THA_": // ठ थ
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 2
	case "JHA": // झ
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 3
	case "ANA_": // ण  (न)
		nakshatra.Name = "Uttara Bhadrapada"
		nakshatra.Pada = 4

	case "DE": // दे
		nakshatra.Name = "Revati"
		nakshatra.Pada = 1
	case "DO": // दो
		nakshatra.Name = "Revati"
		nakshatra.Pada = 2
	case "CH": // च
		nakshatra.Name = "Revati"
		nakshatra.Pada = 3
	case "CHI": // ची
		nakshatra.Name = "Revati"
		nakshatra.Pada = 4

	default:
		nakshatra.Name = ""
		nakshatra.Pada = 0
	}

	return nakshatra
}
