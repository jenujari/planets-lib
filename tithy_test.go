package baselib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTithy_BasicsAndBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		moonLong float64
		sunLong  float64
		expected int
	}{
		{
			name:     "new moon (same longitude) -> tithy 1",
			moonLong: 100,
			sunLong:  100,
			expected: 1,
		},
		{
			name:     "end of first tithy (just below 112 delta) -> tithy 1",
			moonLong: 111.999,
			sunLong:  100,
			expected: 1,
		},
		{
			name:     "start of second tithy (exactly 112 delta) -> tithy 2",
			moonLong: 112,
			sunLong:  100,
			expected: 2,
		},
		{
			name:     "purnima boundary 180 delta -> tithy 16",
			moonLong: 280,
			sunLong:  100,
			expected: 16,
		},
		{
			name:     "just before 180 delta -> tithy 15",
			moonLong: 279.999,
			sunLong:  100,
			expected: 15,
		},
		{
			name:     "just before 360 delta -> tithy 30",
			moonLong: 99.999,
			sunLong:  100,
			expected: 30,
		},
		{
			name:     "worked example from description -> tithy 19",
			moonLong: 84 + 12.0/60.0,  // 84°12'
			sunLong:  227 + 46.0/60.0, // 227°46'
			expected: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcTithy(tt.moonLong, tt.sunLong)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCalcTithy_Normalization(t *testing.T) {
	tests := []struct {
		name     string
		moonLong float64
		sunLong  float64
		expected int
	}{
		{
			name:     "large positive moon and negative sun normalize to same angle -> tithy 1",
			moonLong: 721.5,  // 721.5 mod 360 = 1.5
			sunLong:  -358.5, // -358.5 mod 360 = 1.5
			expected: 1,
		},
		{
			name:     "exact multiples of 360 map to 0 -> tithy 1",
			moonLong: 360.0,
			sunLong:  360.0,
			expected: 1,
		},
		{
			name:     "negative multiples of 360 map to 0 -> tithy 1",
			moonLong: -720.0,
			sunLong:  0.0,
			expected: 1,
		},
		{
			name:     "moon slightly above 360 and sun at 0 -> small delta -> tithy 1",
			moonLong: 360.0001,
			sunLong:  0.0,
			expected: 1,
		},
		{
			name:     "moon 360 + 348 delta -> tithy 30",
			moonLong: 360.0 + 448.0, // normalized moon = 88, sun = 100 -> delta = 348? (we construct to test normalization behavior)
			sunLong:  100.0,
			expected: 30, // here we mainly assert normalization keeps behavior predictable
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcTithy(tt.moonLong, tt.sunLong)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCalcTithy_InvalidInputs(t *testing.T) {
	cases := []struct {
		name string
		moon float64
		sun  float64
	}{
		{"moon NaN", math.NaN(), 100},
		{"sun NaN", 100, math.NaN()},
		{"both NaN", math.NaN(), math.NaN()},
		{"moon +Inf", math.Inf(1), 100},
		{"sun -Inf", 100, math.Inf(-1)},
		{"both +Inf", math.Inf(1), math.Inf(1)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := CalcTithy(c.moon, c.sun)
			assert.Equal(t, 0, got, "invalid inputs should produce 0")
		})
	}
}

func TestCalcTithy_BoundaryPrecision(t *testing.T) {
	sun := 100.0

	tests := []struct {
		name     string
		delta    float64 // moon - sun (before normalization)
		expected int
	}{
		{"delta 0 -> tithy 1", 0.0, 1},
		{"delta exactly 12 -> tithy 2", 12.0, 2},
		{"delta exactly 24 -> tithy 3", 24.0, 3},
		{"delta exactly 348 -> tithy 30", 348.0, 30},            // 29*12 = 348 -> tithy 30
		{"delta 347.9999 -> tithy 29", 347.9999, 29},            // just below 348
		{"delta near 360 (359.9999) -> tithy 30", 359.9999, 30}, // near wrap-around
		{"delta tiny positive (1e-12) -> tithy 1", 1e-12, 1},    // numerical tiny delta treated as >0 -> still floor(0) +1 =1
		{"delta exactly multiple of 12 large (12*29) -> tithy 30", 12.0 * 29.0, 30},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moon := sun + tt.delta
			got := CalcTithy(moon, sun)
			assert.Equal(t, tt.expected, got)
		})
	}
}
