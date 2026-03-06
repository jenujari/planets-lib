package baselib

import "math"

// CalcTithy calculates the tithy based on the longitudes of the moon and the sun.
// 1 to 15 - Sukla Paksha
// 16 to 30 - Krishna Paksha
func CalcTithy(moon, sun float64) int {
	moonLongitude := math.Mod(moon, 360)
	if moonLongitude < 0 {
		moonLongitude += 360
	}

	sunLongitude := math.Mod(sun, 360)
	if sunLongitude < 0 {
		sunLongitude += 360
	}

	delta := moonLongitude - sunLongitude
	if delta < 0 {
		delta += 360
	}

	tithy := int(math.Floor(delta/12.0)) + 1

	if tithy > 30 {
		return 30
	}

	return tithy
}
