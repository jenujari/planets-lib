package bal

import (
	"fmt"
	"math"

	base "github.com/jenujari/planets-lib"
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
	_, navanshRashi := base.CalcNavanshRashi(pl_long)
	navanshLord := base.GetSignLord(navanshRashi)

	// Minutes use the raw longitude, matching the historical formula.
	// Wrapping here would change results for values outside [0, 360).
	positionInNavansh := math.Mod(pl_long*60, 200)
	distanceFactor := (100 - math.Abs(100-positionInNavansh)) / 100

	score, err := weightedPositionalStrength(pl_name, navanshLord, distanceFactor)
	if err != nil {
		return 0, fmt.Errorf("error in getting graha maitri in NavanshBal : %w", err)
	}
	return score, nil
}
