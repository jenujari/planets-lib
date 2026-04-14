package baselib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanetSpeedCategory_BoundariesAndEdges(t *testing.T) {
	t.Run("Sun thresholds exact and around boundaries", func(t *testing.T) {
		// exact thresholds should map to the lower/equal category
		cat, err := PlanetSpeedCategory(SUN, 0.9639352)
		assert.NoError(t, err)
		assert.Equal(t, ATI_MAND, cat)

		cat, err = PlanetSpeedCategory(SUN, 0.9750926)
		assert.NoError(t, err)
		assert.Equal(t, MAND, cat)

		cat, err = PlanetSpeedCategory(SUN, 0.98625)
		assert.NoError(t, err)
		assert.Equal(t, MADHYAM, cat)

		cat, err = PlanetSpeedCategory(SUN, 0.9974074)
		assert.NoError(t, err)
		assert.Equal(t, SAMA, cat)

		cat, err = PlanetSpeedCategory(SUN, 1.0085648)
		assert.NoError(t, err)
		assert.Equal(t, SHEEGHRA, cat)

		cat, err = PlanetSpeedCategory(SUN, 1.0197222)
		assert.NoError(t, err)
		assert.Equal(t, ATI_SHEEGHRA, cat)

		// just above a threshold moves to the next category
		eps := 1e-9
		cat, err = PlanetSpeedCategory(SUN, 0.9750926+eps)
		assert.NoError(t, err)
		assert.Equal(t, MADHYAM, cat)

		// negative speed for Sun (not in table) falls back to ATI_MAND per implementation
		cat, err = PlanetSpeedCategory(SUN, -0.5)
		assert.NoError(t, err)
		assert.Equal(t, ATI_MAND, cat)
	})

	t.Run("Moon thresholds exact", func(t *testing.T) {
		cat, err := PlanetSpeedCategory(MOON, 12.3662037)
		assert.NoError(t, err)
		assert.Equal(t, ATI_MAND, cat)

		cat, err = PlanetSpeedCategory(MOON, 15.3930556)
		assert.NoError(t, err)
		assert.Equal(t, ATI_SHEEGHRA, cat)

		// small above boundary
		cat, err = PlanetSpeedCategory(MOON, 13.5769444+1e-9)
		assert.NoError(t, err)
		assert.Equal(t, SAMA, cat)
	})

	t.Run("Mars retro boundaries and positive", func(t *testing.T) {
		// retro KUTIL exact
		cat, err := PlanetSpeedCategory(MARS, -0.3342593)
		assert.NoError(t, err)
		assert.Equal(t, KUTIL, cat)

		// retro ATI_VAKRA exact
		cat, err = PlanetSpeedCategory(MARS, -0.2674074)
		assert.NoError(t, err)
		assert.Equal(t, ATI_VAKRA, cat)

		// retro small negative -> VAKRA
		cat, err = PlanetSpeedCategory(MARS, -0.1)
		assert.NoError(t, err)
		assert.Equal(t, VAKRA, cat)

		// positive boundaries
		cat, err = PlanetSpeedCategory(MARS, 0.1318981)
		assert.NoError(t, err)
		assert.Equal(t, ATI_MAND, cat)

		cat, err = PlanetSpeedCategory(MARS, 0.7913889)
		assert.NoError(t, err)
		assert.Equal(t, ATI_SHEEGHRA, cat)
	})

	t.Run("Mercury and Jupiter retro mapping and extremes", func(t *testing.T) {
		cat, err := PlanetSpeedCategory(MERCURY, -1.1550926)
		assert.NoError(t, err)
		assert.Equal(t, KUTIL, cat)

		cat, err = PlanetSpeedCategory(JUPITER, -0.0911111)
		assert.NoError(t, err)
		assert.Equal(t, ATI_VAKRA, cat)

		cat, err = PlanetSpeedCategory(VENUS, 1.2588889)
		assert.NoError(t, err)
		assert.Equal(t, ATI_SHEEGHRA, cat)

		cat, err = PlanetSpeedCategory(SATURN, 0.021713)
		assert.NoError(t, err)
		assert.Equal(t, ATI_MAND, cat)
	})

	t.Run("Rahu/Ketu special logic and NaN error", func(t *testing.T) {
		cat, err := PlanetSpeedCategory(RAHU, -0.2145833)
		assert.NoError(t, err)
		assert.Equal(t, KUTIL, cat)

		cat, err = PlanetSpeedCategory(RAHU, -0.1716667)
		assert.NoError(t, err)
		assert.Equal(t, ATI_VAKRA, cat)

		cat, err = PlanetSpeedCategory(RAHU, -0.17)
		assert.NoError(t, err)
		assert.Equal(t, VAKRA, cat)

		// invalid speed => error
		_, err = PlanetSpeedCategory(SUN, math.NaN())
		assert.Error(t, err)
	})
}

func TestPlanetSBCLRFVedha_MappingsAndBoundaries(t *testing.T) {
	t.Run("Sun vedha mapping", func(t *testing.T) {
		// ATI_MAND -> NO_VEDHA
		vedha, err := PlanetSBCLRFVedha(SUN, 0.9639352)
		assert.NoError(t, err)
		assert.Equal(t, NO_VEDHA, vedha)

		// MADHYAM -> FRONT_VEDHA
		vedha, err = PlanetSBCLRFVedha(SUN, 0.98625)
		assert.NoError(t, err)
		assert.Equal(t, FRONT_VEDHA, vedha)

		// SAMA -> LEFT_VEDHA
		vedha, err = PlanetSBCLRFVedha(SUN, 0.9974074)
		assert.NoError(t, err)
		assert.Equal(t, LEFT_VEDHA, vedha)
	})

	t.Run("Moon vedha mapping", func(t *testing.T) {
		// SAMA -> LEFT_VEDHA
		vedha, err := PlanetSBCLRFVedha(MOON, 14.1823148)
		assert.NoError(t, err)
		assert.Equal(t, LEFT_VEDHA, vedha)

		// MADHYAM -> FRONT_VEDHA
		vedha, err = PlanetSBCLRFVedha(MOON, 13.5769444)
		assert.NoError(t, err)
		assert.Equal(t, FRONT_VEDHA, vedha)
	})

	t.Run("Other planets vedha mapping", func(t *testing.T) {
		// Retro KUTIL -> RIGHT_VEDHA
		vedha, err := PlanetSBCLRFVedha(MARS, -0.3342593)
		assert.NoError(t, err)
		assert.Equal(t, RIGHT_VEDHA, vedha)

		// Retro ATI_VAKRA -> RIGHT_VEDHA (use Mercury)
		vedha, err = PlanetSBCLRFVedha(MERCURY, -0.9240741)
		assert.NoError(t, err)
		assert.Equal(t, RIGHT_VEDHA, vedha)

		// Positive ATI_SHEEGHRA -> LEFT_VEDHA (use Mercury max)
		vedha, err = PlanetSBCLRFVedha(MERCURY, 2.2025)
		assert.NoError(t, err)
		assert.Equal(t, LEFT_VEDHA, vedha)

		// Rahu always LEFT_VEDHA
		vedha, err = PlanetSBCLRFVedha(RAHU, -0.2)
		assert.NoError(t, err)
		assert.Equal(t, LEFT_VEDHA, vedha)
	})

	t.Run("Error propagation for NaN speeds", func(t *testing.T) {
		_, err := PlanetSBCLRFVedha(SUN, math.NaN())
		assert.Error(t, err)
	})
}
