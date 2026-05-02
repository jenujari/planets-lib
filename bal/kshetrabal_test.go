package bal

import (
	"github.com/jenujari/planets-lib"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKshetraBal(t *testing.T) {
	tests := []struct {
		name     string
		plLong   float64
		plName   string
		expected float64
	}{
		// baselib.SUN (Sun) in Leo (Lord: Sun) -> baselib.SELF relationship (inc_factor = 4)
		// Leo is 120° to 150°. 135° is the midpoint.
		// rem_deg = 135 % 30 = 15. low_bound = true.
		// final_ksBal = 15 * 4 = 60.
		// Result = (60 / 60) * 100 = 100.
		{"Sun in Leo Midpoint", 135.0, baselib.SUN, 100.0},

		// baselib.SUN in Leo 120° (start of sign)
		// rem_deg = 0. low_bound = true.
		// final_ksBal = 0 * 4 = 0.
		// Result = 0.
		{"Sun in Leo Start", 120.0, baselib.SUN, 0.0},

		// baselib.SUN in Leo 150° (end of sign)
		// rem_deg = 0 (math.Mod(150, 30) = 0). low_bound = true.
		// final_ksBal = 0.
		// Result = 0.
		{"Sun in Leo End", 150.0, baselib.SUN, 0.0},

		// baselib.MARS (Mars) in Leo (Lord: Sun) -> baselib.FRIEND relationship (inc_factor = 3)
		// Midpoint 135°. rem_deg = 15. low_bound = true.
		// final_ksBal = 15 * 3 = 45.
		// Result = (45 / 60) * 100 = 75.
		{"Mars in Leo Midpoint", 135.0, baselib.MARS, 75.0},

		// baselib.SUN in Libra (Lord: Venus) -> baselib.ENEMY relationship (inc_factor = 1)
		// Libra is 180° to 210°. Midpoint 195°.
		// rem_deg = 15. low_bound = true.
		// final_ksBal = 15 * 1 = 15.
		// Result = (15 / 60) * 100 = 25.
		{"Sun in Libra Midpoint", 195.0, baselib.SUN, 25.0},

		// Testing non-integer degrees: baselib.SUN in Leo 127.5° (7.5° into sign)
		// rem_float = 7.5. low_bound = true.
		// dms = 7°30'0".
		// deg_ksbal = 7 * 4 * 216000 = 6048000
		// min_ksbal = 30 * 4 * 3600 = 432000
		// final_ksBal = (6480000) / 216000 = 30.
		// Result = (30 / 60) * 100 = 50.
		{"Sun in Leo 7.5 deg", 127.5, baselib.SUN, 50.0},

		// Testing high boundary: baselib.SUN in Leo 142.5° (22.5° into sign)
		// rem_float = 22.5. low_bound = false.
		// dms = 22°30'0".
		// ksbal = (30 - 22) * 4 = 32.
		// deg_ksbal = 32 * 216000 = 6912000
		// min_ksbal = 30 * 4 * 3600 = 432000
		// final_ksBal = (6912000 - 432000) / 216000 = 6480000 / 216000 = 30.
		// Result = (30 / 60) * 100 = 50.
		{"Sun in Leo 22.5 deg", 142.5, baselib.SUN, 50.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := KshetraBal(tt.plLong, tt.plName)
			assert.NoError(t, err)
			assert.InDelta(t, tt.expected, actual, 1e-9)
		})
	}
}

func BenchmarkKshetraBal(b *testing.B) {
	var r float64
	for b.Loop() {
		r, _ = KshetraBal(135.0, baselib.SUN)
	}
	test_bal = r
}
