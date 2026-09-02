package baselib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeAngle(t *testing.T) {
	tests := []struct {
		name     string
		in       float64
		expected float64
		isNaN    bool
	}{
		{"already in range", 45, 45, false},
		{"zero", 0, 0, false},
		{"just below 360", 359.999, 359.999, false},
		{"exact 360 wraps to 0", 360, 0, false},
		{"large positive", 721.5, 1.5, false},
		{"negative", -358.5, 1.5, false},
		{"negative multiple of 360", -720, 0, false},
		{"NaN", math.NaN(), 0, true},
		{"+Inf", math.Inf(1), 0, true},
		{"-Inf", math.Inf(-1), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeAngle(tt.in)
			if tt.isNaN {
				assert.True(t, math.IsNaN(got))
				return
			}
			assert.InDelta(t, tt.expected, got, 1e-9)
		})
	}
}

func TestValidAngle(t *testing.T) {
	t.Run("finite wraps", func(t *testing.T) {
		got, ok := ValidAngle(721.5)
		assert.True(t, ok)
		assert.InDelta(t, 1.5, got, 1e-9)
	})

	t.Run("already normalized", func(t *testing.T) {
		got, ok := ValidAngle(10)
		assert.True(t, ok)
		assert.Equal(t, 10.0, got)
	})

	invalid := []struct {
		name string
		in   float64
	}{
		{"NaN", math.NaN()},
		{"+Inf", math.Inf(1)},
		{"-Inf", math.Inf(-1)},
	}
	for _, tt := range invalid {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := ValidAngle(tt.in)
			assert.False(t, ok)
			assert.Equal(t, 0.0, got)
		})
	}
}

func TestLongitudeCallersShareValidAngle(t *testing.T) {
	// Mapping modules must treat invalid angles the same way now that
	// ValidAngle is the shared seam.
	invalid := []float64{math.NaN(), math.Inf(1), math.Inf(-1)}
	for _, d := range invalid {
		assert.Equal(t, "", GetSignFrmDegree(d))
		assert.Equal(t, NakshatraPada{}, GetNakshatraPadaFromDegree(d))
		num, name := CalcNavanshRashi(d)
		assert.Equal(t, 0, num)
		assert.Equal(t, "", name)
		assert.Equal(t, 0, CalcTithy(d, 100))
		assert.Equal(t, 0, CalcTithy(100, d))
	}
}
