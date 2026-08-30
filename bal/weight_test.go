package bal

import (
	"math"
	"testing"

	baselib "github.com/jenujari/planets-lib"
	"github.com/stretchr/testify/assert"
)

func TestRelationWeight(t *testing.T) {
	tests := []struct {
		rel      string
		expected float64
	}{
		{baselib.SELF, WeightSelf},
		{baselib.FRIEND, WeightFriend},
		{baselib.NEUTRAL, WeightNeutral},
		{baselib.ENEMY, WeightEnemy},
		{"unknown", WeightNeutral},
	}
	for _, tt := range tests {
		t.Run(tt.rel, func(t *testing.T) {
			assert.Equal(t, tt.expected, relationWeight(tt.rel))
		})
	}
}

func TestWeightedPositionalStrength(t *testing.T) {
	// Mars in Aries (lord Mars) at the midpoint distance factor 1.0 -> SELF * 100
	got, err := weightedPositionalStrength(baselib.MARS, baselib.MARS, 1.0)
	assert.NoError(t, err)
	assert.InDelta(t, 100.0, got, 1e-9)

	// Mars to Sun is FRIEND
	got, err = weightedPositionalStrength(baselib.MARS, baselib.SUN, 1.0)
	assert.NoError(t, err)
	assert.InDelta(t, 75.0, got, 1e-9)

	_, err = weightedPositionalStrength(baselib.URANUS, baselib.SUN, 1.0)
	assert.Error(t, err)
}

func TestNavanshBal_NormalizesWrappedLongitude(t *testing.T) {
	// 1°40' is the Aries navamsha midpoint for Mars (SELF -> 100).
	mid := 1.6666666666666667
	baseScore, err := NavanshBal(mid, baselib.MARS)
	assert.NoError(t, err)

	wrapped, err := NavanshBal(mid+360, baselib.MARS)
	assert.NoError(t, err)
	assert.InDelta(t, baseScore, wrapped, 1e-9)

	// Negative wrap used to skip NormalizeAngle on the minutes path.
	neg, err := NavanshBal(mid-360, baselib.MARS)
	assert.NoError(t, err)
	assert.InDelta(t, baseScore, neg, 1e-9)
}

func TestKshetraBal_InvalidLongitude(t *testing.T) {
	got, err := KshetraBal(math.NaN(), baselib.SUN)
	assert.Error(t, err)
	assert.Equal(t, 0.0, got)
}
