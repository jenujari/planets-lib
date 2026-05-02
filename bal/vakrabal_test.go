package bal

import (
	"testing"

	baselib "github.com/jenujari/planets-lib"
	"github.com/stretchr/testify/assert"
)

func TestVakraBal(t *testing.T) {
	tests := []struct {
		name     string
		plSpeed  float64
		plName   string
		expected float64
	}{
		{"Sun should return 0", -1.0, baselib.SUN, 0},
		{"Moon should return 0", -1.0, baselib.MOON, 0},
		{"Rahu should return 100", 0.1, baselib.RAHU, 100},
		{"Ketu should return 100", 0.1, baselib.KETU, 100},
		{"Direct Mars should return 0", 0.5, baselib.MARS, 0},
		{"Retro Mars max speed", -0.436666, baselib.MARS, 100},
		{"Retro Mars half speed", -0.218333, baselib.MARS, 50},
		{"Retro Mercury max speed", -1.5, baselib.MERCURY, 100},
		{"Retro Jupiter max speed", -0.136666, baselib.JUPITER, 100},
		{"Retro Venus max speed", -0.686666, baselib.VENUS, 100},
		{"Retro Saturn max speed", -0.0833333, baselib.SATURN, 100},
		{"Retro Saturn exceeding max", -0.1, baselib.SATURN, 100},
		{"Unknown planet", -1.0, "Uranus", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := VakraBal(tt.plSpeed, tt.plName)
			assert.InDelta(t, tt.expected, actual, 1e-9)
		})
	}
}

func BenchmarkVakraBal(b *testing.B) {
	var r float64
	for b.Loop() {
		r = VakraBal(-0.4, baselib.MARS)
	}
	test_bal = r
}
