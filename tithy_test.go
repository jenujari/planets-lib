package baselib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcTithy(t *testing.T) {
	tests := []struct {
		name     string
		moonLong float64
		sunLong  float64
		expected int
	}{
		{
			name:     "new moon boundary gives tithy 1",
			moonLong: 100,
			sunLong:  100,
			expected: 1,
		},
		{
			name:     "end of first tithy",
			moonLong: 111.999,
			sunLong:  100,
			expected: 1,
		},
		{
			name:     "start of second tithy",
			moonLong: 112,
			sunLong:  100,
			expected: 2,
		},
		{
			name:     "purnima boundary 180 gives tithy 16",
			moonLong: 280,
			sunLong:  100,
			expected: 16,
		},
		{
			name:     "just before 180 gives tithy 15",
			moonLong: 279.999,
			sunLong:  100,
			expected: 15,
		},
		{
			name:     "just before 360 gives tithy 30",
			moonLong: 99.999,
			sunLong:  100,
			expected: 30,
		},
		{
			name:     "worked example from description gives tithy 19",
			moonLong: 84 + 12.0/60.0,
			sunLong:  227 + 46.0/60.0,
			expected: 19,
		},
		{
			name:     "normalizes longitudes outside 0 to 360",
			moonLong: 721.5,
			sunLong:  -358.5,
			expected: 1,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			moon := PlanetCord{Longitude: tt.moonLong}
			sun := PlanetCord{Longitude: tt.sunLong}
			result := CalcTithy(moon, sun)
			assert.Equal(t, tt.expected, result)
		})
	}
}
