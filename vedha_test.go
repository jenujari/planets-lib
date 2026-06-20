package baselib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVedhaTarget(t *testing.T) {
	t.Run("No Vedha direction returns empty string", func(t *testing.T) {
		assert.Equal(t, "", VedhaTarget(NAKSHATRA_ASHWINI, NO_VEDHA))
		assert.Equal(t, "", VedhaTarget("", NO_VEDHA))
	})

	t.Run("Invalid direction returns empty string", func(t *testing.T) {
		assert.Equal(t, "", VedhaTarget(NAKSHATRA_ASHWINI, "invalid-direction"))
		assert.Equal(t, "", VedhaTarget(NAKSHATRA_ASHWINI, ""))
	})

	t.Run("Left Vedha mappings", func(t *testing.T) {
		// Test specific sample entries from LeftVedhaMap
		tests := []struct {
			nakshatra string
			expected  string
		}{
			{NAKSHATRA_ASHWINI, NAKSHATRA_ROHINI},
			{NAKSHATRA_BHARANI, NAKSHATRA_KRITTIKA},
			{NAKSHATRA_REVATI, NAKSHATRA_MRIGASHIRSHA},
		}
		for _, tt := range tests {
			t.Run(tt.nakshatra, func(t *testing.T) {
				assert.Equal(t, tt.expected, VedhaTarget(tt.nakshatra, LEFT_VEDHA))
			})
		}

		// Test invalid/empty nakshatra
		assert.Equal(t, "", VedhaTarget("NonExistentNakshatra", LEFT_VEDHA))
		assert.Equal(t, "", VedhaTarget("", LEFT_VEDHA))
	})

	t.Run("Right Vedha mappings", func(t *testing.T) {
		// Test specific sample entries from RightVedhaMap
		tests := []struct {
			nakshatra string
			expected  string
		}{
			{NAKSHATRA_ASHWINI, NAKSHATRA_JYESTHA},
			{NAKSHATRA_BHARANI, NAKSHATRA_ANURADHA},
			{NAKSHATRA_REVATI, NAKSHATRA_MOOLA},
		}
		for _, tt := range tests {
			t.Run(tt.nakshatra, func(t *testing.T) {
				assert.Equal(t, tt.expected, VedhaTarget(tt.nakshatra, RIGHT_VEDHA))
			})
		}

		// Test invalid/empty nakshatra
		assert.Equal(t, "", VedhaTarget("NonExistentNakshatra", RIGHT_VEDHA))
		assert.Equal(t, "", VedhaTarget("", RIGHT_VEDHA))
	})

	t.Run("Front Vedha mappings", func(t *testing.T) {
		// Test specific sample entries from FrontVedhaMap
		tests := []struct {
			nakshatra string
			expected  string
		}{
			{NAKSHATRA_ASHWINI, NAKSHATRA_PURVA_PHALGUNI},
			{NAKSHATRA_BHARANI, NAKSHATRA_MAGHA},
			{NAKSHATRA_REVATI, NAKSHATRA_UTTARA_PHALGUNI},
		}
		for _, tt := range tests {
			t.Run(tt.nakshatra, func(t *testing.T) {
				assert.Equal(t, tt.expected, VedhaTarget(tt.nakshatra, FRONT_VEDHA))
			})
		}

		// Test invalid/empty nakshatra
		assert.Equal(t, "", VedhaTarget("NonExistentNakshatra", FRONT_VEDHA))
		assert.Equal(t, "", VedhaTarget("", FRONT_VEDHA))
	})

	t.Run("Complete Map Validation", func(t *testing.T) {
		// Verify that VedhaTarget yields correct results for ALL mapped nakshatras
		// across Left, Right, and Front maps.
		for k, v := range LeftVedhaMap {
			assert.Equal(t, v, VedhaTarget(k, LEFT_VEDHA), "LeftVedhaMap key %s", k)
		}
		for k, v := range RightVedhaMap {
			assert.Equal(t, v, VedhaTarget(k, RIGHT_VEDHA), "RightVedhaMap key %s", k)
		}
		for k, v := range FrontVedhaMap {
			assert.Equal(t, v, VedhaTarget(k, FRONT_VEDHA), "FrontVedhaMap key %s", k)
		}
	})
}
