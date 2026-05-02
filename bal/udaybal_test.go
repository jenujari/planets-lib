package bal

import (
	"testing"

	baselib "github.com/jenujari/planets-lib"
	"github.com/stretchr/testify/assert"
)

func TestUdayBal(t *testing.T) {
	tests := []struct {
		name     string
		sunLong  float64
		plLong   float64
		plSpeed  float64
		plName   string
		expected float64
	}{
		{
			name:     "Sun always MaxRetrunValue",
			sunLong:  10,
			plLong:   20,
			plSpeed:  1,
			plName:   baselib.SUN,
			expected: 100,
		},
		{
			name:     "Rahu always MinReturnValue",
			sunLong:  10,
			plLong:   20,
			plSpeed:  -0.05,
			plName:   baselib.RAHU,
			expected: 0,
		},
		{
			name:     "Ketu always MinReturnValue",
			sunLong:  10,
			plLong:   20,
			plSpeed:  -0.05,
			plName:   baselib.KETU,
			expected: 0,
		},
		{
			name:     "Moon exactly at Astance",
			sunLong:  0,
			plLong:   12,
			plSpeed:  13,
			plName:   baselib.MOON,
			expected: 0,
		},
		{
			name:     "Moon halfway to 180",
			sunLong:  0,
			plLong:   96, // 12 + (180-12)/2 = 12 + 84 = 96
			plSpeed:  13,
			plName:   baselib.MOON,
			expected: 50,
		},
		{
			name:     "Moon at 180 degrees",
			sunLong:  0,
			plLong:   180,
			plSpeed:  13,
			plName:   baselib.MOON,
			expected: 100,
		},
		{
			name:     "Mars exactly at Astance",
			sunLong:  0,
			plLong:   17,
			plSpeed:  0.5,
			plName:   baselib.MARS,
			expected: 0,
		},
		{
			name:     "Mars at 180 degrees",
			sunLong:  0,
			plLong:   180,
			plSpeed:  0.5,
			plName:   baselib.MARS,
			expected: 100,
		},
		{
			name:     "Jupiter exactly at Astance",
			sunLong:  0,
			plLong:   11,
			plSpeed:  0.1,
			plName:   baselib.JUPITER,
			expected: 0,
		},
		{
			name:     "Jupiter at 180 degrees",
			sunLong:  0,
			plLong:   180,
			plSpeed:  0.1,
			plName:   baselib.JUPITER,
			expected: 100,
		},
		{
			name:     "Mercury Direct at Astance",
			sunLong:  0,
			plLong:   14,
			plSpeed:  1.5,
			plName:   baselib.MERCURY,
			expected: 0,
		},
		{
			name:     "Mercury Direct at MaxDiff (27)",
			sunLong:  0,
			plLong:   27,
			plSpeed:  1.5,
			plName:   baselib.MERCURY,
			expected: 100,
		},
		{
			name:     "Mercury Retro at Astance (12)",
			sunLong:  0,
			plLong:   12,
			plSpeed:  -0.5,
			plName:   baselib.MERCURY,
			expected: 0,
		},
		{
			name:     "Mercury Retro at MaxDiff (27)",
			sunLong:  0,
			plLong:   27,
			plSpeed:  -0.5,
			plName:   baselib.MERCURY,
			expected: 100,
		},
		{
			name:     "Venus Direct at Astance (10)",
			sunLong:  0,
			plLong:   10,
			plSpeed:  1.2,
			plName:   baselib.VENUS,
			expected: 0,
		},
		{
			name:     "Venus Direct at MaxDiff (47)",
			sunLong:  0,
			plLong:   47,
			plSpeed:  1.2,
			plName:   baselib.VENUS,
			expected: 100,
		},
		{
			name:     "Venus Retro at Astance (8)",
			sunLong:  0,
			plLong:   8,
			plSpeed:  -0.4,
			plName:   baselib.VENUS,
			expected: 0,
		},
		{
			name:     "Venus Retro at MaxDiff (47)",
			sunLong:  0,
			plLong:   47,
			plSpeed:  -0.4,
			plName:   baselib.VENUS,
			expected: 100,
		},
		{
			name:     "Shortest distance wrap around 360 (Mercury)",
			sunLong:  355,
			plLong:   5,
			plSpeed:  1.5,
			plName:   baselib.MERCURY,
			expected: 0, // distance is 10, which is <= 14
		},
		{
			name:     "Shortest distance wrap around 360 (Mercury 2)",
			sunLong:  355,
			plLong:   20,
			plSpeed:  1.5,
			plName:   baselib.MERCURY,
			expected: 100, // distance is 25, which is > 14 and > (27-14)? Wait.
			// distance = 25. pl_ast = 14. max_diff = 27.
			// availabe_range = 13. out_of_ast = 11.
			// uday = (100/13) * 11 = 84.615...
			// Wait, let's recalculate.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UdayBal(tt.sunLong, tt.plLong, tt.plSpeed, tt.plName)
			if tt.name == "Shortest distance wrap around 360 (Mercury 2)" {
				// (100.0 / (27 - 14)) * (25 - 14) = (100/13)*11 = 84.61538461538461
				assert.InDelta(t, 84.61538461538461, result, 0.0001)
			} else {
				assert.InDelta(t, tt.expected, result, 0.0001)
			}
		})
	}
}

func BenchmarkUdayBal(b *testing.B) {
	var result float64
	for b.Loop() {
		result = UdayBal(0, 45, 1.2, baselib.VENUS)
	}
	test_bal = result
}
