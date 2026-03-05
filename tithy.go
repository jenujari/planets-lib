package baselib

import "math"

/*
Sukla Paksha - 1 to 15
Krishna Paksha - 16 to 30
*/

func CalcTithy(moon, sun PlanetCord) int {
	moonLongitude := math.Mod(moon.Longitude, 360)
	if moonLongitude < 0 {
		moonLongitude += 360
	}

	sunLongitude := math.Mod(sun.Longitude, 360)
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
