package baselib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNakshatraVowelMap_ConsistencyAndNormalization(t *testing.T) {
	// Ensure that GetNakshatraFromVowel returns the same values as the internal map
	// for exact keys, and also for case/whitespace-normalized variations.
	for key, expected := range nakshatraFromVowelMap {
		t.Run("exact_"+key, func(t *testing.T) {
			got := GetNakshatraFromVowel(key)
			assert.Equal(t, expected.Name, got.Name, "name mismatch for key %s", key)
			assert.Equal(t, expected.Pada, got.Pada, "pada mismatch for key %s", key)
		})

		t.Run("lower_"+key, func(t *testing.T) {
			l := strings.ToLower(key)
			got := GetNakshatraFromVowel(l)
			assert.Equal(t, expected.Name, got.Name, "name mismatch for lower-case key %s", l)
			assert.Equal(t, expected.Pada, got.Pada, "pada mismatch for lower-case key %s", l)
		})

		t.Run("trimmed_"+key, func(t *testing.T) {
			withSpaces := "  " + key + "  "
			got := GetNakshatraFromVowel(withSpaces)
			assert.Equal(t, expected.Name, got.Name, "name mismatch for trimmed key %q", withSpaces)
			assert.Equal(t, expected.Pada, got.Pada, "pada mismatch for trimmed key %q", withSpaces)
		})
	}

	t.Run("unknown_key_returns_zero_value", func(t *testing.T) {
		got := GetNakshatraFromVowel("NOT_A_KEY")
		assert.Equal(t, "", got.Name, "expected empty name for unknown key")
		assert.Equal(t, 0, got.Pada, "expected pada 0 for unknown key")
	})
}

func TestSignLordMap_Consistency(t *testing.T) {
	// Ensure GetSignLord returns the same mapping as the data-driven map.
	for sign, expectedLord := range signLordMap {
		t.Run("sign_"+sign, func(t *testing.T) {
			got := GetSignLord(sign)
			assert.Equal(t, expectedLord, got, "lord mismatch for sign %s", sign)
		})
	}

	t.Run("known_constant_returns_expected", func(t *testing.T) {
		// Sanity checks for a few canonical constants.
		assert.Equal(t, MARS, GetSignLord(SIGN_ARIES))
		assert.Equal(t, VENUS, GetSignLord(SIGN_TAURUS))
		assert.Equal(t, JUPITER, GetSignLord(SIGN_PISCES))
	})

	t.Run("unknown_sign_returns_empty", func(t *testing.T) {
		assert.Equal(t, "", GetSignLord("NotASign"))
	})
}
