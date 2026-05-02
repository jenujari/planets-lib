package baselib

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetGrahaMaitri(t *testing.T) {
	tests := []struct {
		base     string
		target   string
		expected string
		hasError bool
	}{
		// Sun's relationships
		{SUN, MOON, FRIEND, false},
		{SUN, MERCURY, NEUTRAL, false},
		{SUN, SATURN, ENEMY, false},
		{SUN, RAHU, ENEMY, false},

		// Moon's relationships
		{MOON, SUN, FRIEND, false},
		{MOON, MARS, NEUTRAL, false},
		{MOON, RAHU, ENEMY, false},

		// Rahu/Ketu mutual friendship
		{RAHU, KETU, FRIEND, false},
		{KETU, RAHU, FRIEND, false},

		// Rahu's enemies
		{RAHU, JUPITER, ENEMY, false},

		// Saturn missing Ketu (should error based on current implementation)
		{SATURN, KETU, "", true},

		// Self relationship
		{JUPITER, JUPITER, "Self", false},

		// Invalid planet
		{"Uranus", SUN, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.base+"_to_"+tt.target, func(t *testing.T) {
			rel, err := GetGrahaMaitri(tt.base, tt.target)
			if tt.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, rel)
			}
		})
	}
}
