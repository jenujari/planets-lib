package bal

import (
	"fmt"
	"math"

	base "github.com/jenujari/planets-lib"
)

const (
	WeightSelf    = 1.00
	WeightFriend  = 0.75
	WeightNeutral = 0.50
	WeightEnemy   = 0.25
)

// NavanshBal calculates the planetary strength (Bal) of a planet based on its position in the Navamsha (D9) chart.
// It considers the relationship between the planet and the lord of the Navamsha sign it occupies.
// The strength peaks when the planet is exactly in the middle of the Navamsha segment (at 100 arc-minutes).
//
// Parameters:
//   - pl_long: The absolute longitude of the planet in degrees (0-360).
//   - pl_name: The name of the planet (e.g., SUN, MARS, JUPITER).
//
// Returns:
//   - The calculated Navansh Bal as a percentage (0.0 to 100.0).
//   - An error if the planet's relationship with the Navamsha lord cannot be determined.
func NavanshBal(pl_long float64, pl_name string) (float64, error) {

	// 1. Get the Navansh Rashi (1 to 12)
	_, navanshRashi := base.CalcNavanshRashi(pl_long)

	// 2. Get the Lord of that Navansh Rashi
	navanshLord := base.GetSignLord(navanshRashi)

	// 3. Get relationship between the Planet and the Navansh Lord
	// Expected returns: "swayam", "mitra", "sama", "shatru"
	relation, err := base.GetGrahaMaitri(pl_name, navanshLord)

	if err != nil {
		return 0, fmt.Errorf("error in getting graha maitri in NavanshBal : %w", err)
	}

	// 4. Calculate Position within the Navansh (0 to 200 minutes)
	// totalDegree * 60 gives total minutes.
	// Modulo 200 gives the current position in the 200-minute navansh block.
	totalMinutes := pl_long * 60
	positionInNavansh := math.Mod(totalMinutes, 200)

	// 5. Calculate Distance Factor (Peak at 100 mins)
	// Formula: (100 - |100 - position|) / 100
	distanceFactor := (100 - math.Abs(100-positionInNavansh)) / 100

	// 6. Assign Weightage based on relationship
	var weight float64
	switch relation {
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

	// 7. Final Percentage Calculation
	navanshBalPercentage := (distanceFactor * weight) * 100

	return navanshBalPercentage, nil
}
