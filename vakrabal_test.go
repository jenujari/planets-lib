package baselib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcVakraBal(t *testing.T) {
	tests := []struct {
		name     string
		plSpeed  float64
		plName   string
		expected float64
	}{
		{"Sun should return 0", -1.0, SUN, 0},
		{"Moon should return 0", -1.0, MOON, 0},
		{"Rahu should return 100", 0.1, RAHU, 100},
		{"Ketu should return 100", 0.1, KETU, 100},
		{"Direct Mars should return 0", 0.5, MARS, 0},
		{"Retro Mars max speed", -0.8, MARS, 100},
		{"Retro Mars half speed", -0.4, MARS, 50},
		{"Retro Mercury max speed", -1.6, MERCURY, 100},
		{"Retro Jupiter max speed", -0.35, JUPITER, 100},
		{"Retro Venus max speed", -0.8, VENUS, 100},
		{"Retro Saturn max speed", -0.25, SATURN, 100},
		{"Retro Saturn exceeding max", -0.3, SATURN, 100},
		{"Unknown planet", -1.0, "Uranus", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := CalcVakraBal(tt.plSpeed, tt.plName)
			assert.InDelta(t, tt.expected, actual, 1e-9)
		})
	}
}

var benchmarkResult float64

func BenchmarkCalcVakraBal(b *testing.B) {
	var r float64
	for b.Loop() {
		r = CalcVakraBal(-0.4, MARS)
	}
	benchmarkResult = r
}
