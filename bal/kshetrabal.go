package bal

import (
	base "github.com/jenujari/planets-lib"

	"math"
)

// KshetraBal calculates the Kshetra Bal (positional strength within a sign) of a planet.
// It depends on the planet's relationship with the lord of the sign it currently occupies
// and its distance from the sign's midpoint (15°).
//
// Parameters:
//   - pl_long: The longitude of the planet in degrees.
//   - pl_name: The name of the planet (e.g., baselib.SUN, baselib.MARS).
//
// Returns:
//   - The calculated Kshetra Bal value (0 to 100).
//   - An error if the planet or its relationship is not found.
func KshetraBal(pl_long float64, pl_name string) (float64, error) {
	// 1. Determine the sign based on pl_long
	sign := base.GetSignFrmDegree(pl_long)

	// 2. Get the lord (swami) of that sign
	pl_swami := base.GetSignLord(sign)

	// 3. Determine the relationship between pl_name and pl_swami
	pl_rel, err := base.GetGrahaMaitri(pl_name, pl_swami)
	if err != nil {
		return 0, err
	}

	var weight float64
	switch pl_rel {
	case base.SELF:
		weight = WeightSelf
	case base.FRIEND:
		weight = WeightFriend
	case base.NEUTRAL:
		weight = WeightNeutral
	case base.ENEMY:
		weight = WeightEnemy
	default:
		weight = WeightNeutral // Default fallback
	}

	// 5. Get remaining degrees in current sign
	normLon := base.NormalizeAngle(pl_long)
	rem := math.Mod(normLon, 30.0) * 60

	distanceFactor := (900 - math.Abs(900-rem)) / 900

	final_ksBal := weight * distanceFactor * 100
	return final_ksBal, nil
}
