package bal

import (
	"math"
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

// Invalid speeds previously slipped past the pl_speed >= 0 check: -Inf produced a
// full-strength 100 and NaN propagated into the result.
func TestVakraBal_InvalidSpeeds(t *testing.T) {
	invalid := []struct {
		name  string
		speed float64
	}{
		{"NaN", math.NaN()},
		{"positive infinity", math.Inf(1)},
		{"negative infinity", math.Inf(-1)},
	}

	for _, tt := range invalid {
		t.Run(tt.name+" on a computed planet returns 0", func(t *testing.T) {
			for _, planet := range []string{baselib.MARS, baselib.MERCURY, baselib.JUPITER, baselib.VENUS, baselib.SATURN} {
				assert.Equal(t, 0.0, VakraBal(tt.speed, planet), "planet %s", planet)
			}
		})

		// Rahu and Ketu are always retrograde by definition and never read the speed,
		// so they must keep returning full strength even for an unusable speed value.
		t.Run(tt.name+" keeps Rahu/Ketu definitional", func(t *testing.T) {
			assert.Equal(t, 100.0, VakraBal(tt.speed, baselib.RAHU))
			assert.Equal(t, 100.0, VakraBal(tt.speed, baselib.KETU))
		})

		t.Run(tt.name+" keeps Sun/Moon at 0", func(t *testing.T) {
			assert.Equal(t, 0.0, VakraBal(tt.speed, baselib.SUN))
			assert.Equal(t, 0.0, VakraBal(tt.speed, baselib.MOON))
		})
	}
}
