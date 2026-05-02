package bal

import (
	"testing"

	baselib "github.com/jenujari/planets-lib"
	"github.com/stretchr/testify/assert"
)

func TestNavanshBal(t *testing.T) {
	tests := []struct {
		name     string
		plLong   float64
		plName   string
		expected float64
	}{
		// Aries 1st Navamsha (Aries). Lord: Mars. Mars to Mars = SELF (Weight: 1.0).
		// Navamsha duration is 3°20' = 200'. Midpoint is 1°40' = 1.6666666666666667 degrees.
		// distanceFactor = 1.0, weight = 1.0 -> 100%
		{"Mars in Aries Navamsha Midpoint (Self)", 1.6666666666666667, baselib.MARS, 100.0},

		// Aries 5th Navamsha (Leo). Lord: Sun. Mars to Sun = FRIEND (Weight: 0.75).
		// Midpoint is 15°00' = 15.0 degrees.
		// distanceFactor = 1.0, weight = 0.75 -> 75%
		{"Mars in Leo Navamsha Midpoint (Friend)", 15.0, baselib.MARS, 75.0},

		// Taurus 1st Navamsha (Capricorn). Lord: Saturn. Mars to Saturn = NEUTRAL (Weight: 0.50).
		// Start of Navamsha is 30°00' = 30.0 degrees. Position in block is 0'.
		// distanceFactor = 0.0, weight = 0.50 -> 0%
		{"Mars in Capricorn Navamsha Start (Neutral)", 30.0, baselib.MARS, 0.0},

		// Aries 3rd Navamsha (Gemini). Lord: Mercury. Mars to Mercury = ENEMY (Weight: 0.25).
		// 50 minutes into Navamsha block (450 total minutes) = 7°30' = 7.5 degrees.
		// Position is 50'. distanceFactor = (100 - |100 - 50|) / 100 = 0.5.
		// navanshBalPercentage = (0.5 * 0.25) * 100 = 12.5%
		{"Mars in Gemini Navamsha 1/4th (Enemy)", 7.5, baselib.MARS, 12.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := NavanshBal(tt.plLong, tt.plName)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, actual, 1e-9)
		})
	}
}

func BenchmarkNavanshBal(b *testing.B) {
	var r float64
	for b.Loop() {
		r, _ = NavanshBal(15.0, baselib.MARS)
	}
	test_bal = r
}
