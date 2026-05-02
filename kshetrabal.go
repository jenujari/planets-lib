package baselib

import (
	"math"
)

// KshetraBal calculates the Kshetra Bal (positional strength within a sign) of a planet.
// It depends on the planet's relationship with the lord of the sign it currently occupies
// and its distance from the sign's midpoint (15°).
//
// Parameters:
//   - pl_long: The longitude of the planet in degrees.
//   - pl_name: The name of the planet (e.g., SUN, MARS).
//
// Returns:
//   - The calculated Kshetra Bal value (0 to 100).
//   - An error if the planet or its relationship is not found.
func KshetraBal(pl_long float64, pl_name string) (float64, error) {
	// 1. Determine the sign based on pl_long
	sign := GetSignFrmDegree(pl_long)

	// 2. Get the lord (swami) of that sign
	pl_swami := GetSignLord(sign)

	// 3. Determine the relationship between pl_name and pl_swami
	pl_rel, err := GetGrahaMaitri(pl_name, pl_swami)
	if err != nil {
		return 0, err
	}

	// 4. Determine inc_factor based on pl_rel
	var inc_factor float64
	switch pl_rel {
	case SELF:
		inc_factor = 4
	case FRIEND:
		inc_factor = 3
	case NEUTRAL:
		inc_factor = 2
	case ENEMY:
		inc_factor = 1
	default:
		inc_factor = 2 // Default to neutral if somehow undefined
	}

	// 5. Get remaining degrees in current sign
	normLon := normalizeAngle(pl_long)
	rem := math.Mod(normLon, 30.0)
	dms := NewDMS(rem)

	// 6. Check low_bound (midpoint is 15°)
	low_bound := rem <= 15.0

	var final_ksBal float64
	if low_bound {
		ksbal := float64(dms.D) * inc_factor
		deg_ksbal := ksbal * 216000
		min_ksbal := float64(dms.M) * inc_factor * 3600
		sec_ksbal := float64(dms.S) * inc_factor * 60
		final_ksBal = (deg_ksbal + min_ksbal + sec_ksbal) / 216000
	} else {
		ksbal := (30.0 - float64(dms.D)) * inc_factor
		deg_ksbal := ksbal * 216000
		min_ksbal := float64(dms.M) * inc_factor * 3600
		sec_ksbal := float64(dms.S) * inc_factor * 60
		final_ksBal = (deg_ksbal - min_ksbal - sec_ksbal) / 216000
	}

	return (final_ksBal / 60.0) * 100.0, nil
}
