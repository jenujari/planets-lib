package bal

import (
	base "github.com/jenujari/planets-lib"
)

const (
	WeightSelf    = 1.00
	WeightFriend  = 0.75
	WeightNeutral = 0.50
	WeightEnemy   = 0.25
)

// relationWeight maps a Graha Maitri relationship onto the positional-strength
// scale used by KshetraBal and NavanshBal.
func relationWeight(rel string) float64 {
	switch rel {
	case base.SELF:
		return WeightSelf
	case base.FRIEND:
		return WeightFriend
	case base.NEUTRAL:
		return WeightNeutral
	case base.ENEMY:
		return WeightEnemy
	default:
		return WeightNeutral
	}
}

// weightedPositionalStrength scales a 0..1 distance factor by the planet's
// relationship to the lord of the occupied sign (or navamsha sign).
func weightedPositionalStrength(plName, lord string, distanceFactor float64) (float64, error) {
	rel, err := base.GetGrahaMaitri(plName, lord)
	if err != nil {
		return 0, err
	}
	return distanceFactor * relationWeight(rel) * 100, nil
}
