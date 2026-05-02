package baselib

// CalcNavanshRashi calculates the Navansh Rashi for a given longitude.
// It returns the 1-indexed rashi number (1..12) and its name.
func CalcNavanshRashi(pl_long float64) (int, string) {
	// Normalize longitude to [0, 360)
	lon := NormalizeAngle(pl_long)

	// Determine the base sign (0..11)
	rashiIdx := int(lon / 30.0)
	// Degrees within the sign [0, 30)
	degInSign := lon - float64(rashiIdx*30.0)

	// Navansh size is exactly 3°20' (3.333... degrees)
	const nvmDegSize = 30.0 / 9.0
	// Determine which navansh of the sign (0..8)
	navanshIdx := int(degInSign / nvmDegSize)

	// Each sign group starts from a specific sign index:
	// - Fire signs (Aries, Leo, Sagittarius): Start from Aries (0)
	// - Earth signs (Taurus, Virgo, Capricorn): Start from Capricorn (9)
	// - Air signs (Gemini, Libra, Aquarius): Start from Libra (6)
	// - Water signs (Cancer, Scorpio, Pisces): Start from Cancer (3)
	var startRashiIdx int
	switch rashiIdx % 4 {
	case 0: // Fire: 0, 4, 8
		startRashiIdx = 0
	case 1: // Earth: 1, 5, 9
		startRashiIdx = 9
	case 2: // Air: 2, 6, 10
		startRashiIdx = 6
	case 3: // Water: 3, 7, 11
		startRashiIdx = 3
	}

	// Calculate resulting rashi index using modulo 12 for cyclicity
	resultRashiIdx := (startRashiIdx + navanshIdx) % 12

	// Return 1-indexed Rashi number and its name from the SIGNS array
	return resultRashiIdx + 1, SIGNS[resultRashiIdx]
}
