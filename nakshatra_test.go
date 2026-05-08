package baselib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNakshatraPadaFromDegree_NormalizationAndBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		deg      float64
		expected NakshatraPada
	}{
		{
			name: "small positive inside first pada -> Ashwini p1",
			deg:  1.5,
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
		{
			name: "just below 3.333333 -> Ashwini p1",
			deg:  3.333333 - 1e-6,
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
		{
			name: "exact 3.333333 -> Ashwini p2",
			deg:  3.333333,
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 2,
			},
		},
		{
			name: "near top of circle -> Revati p4",
			deg:  359.999,
			expected: NakshatraPada{
				Name: NAKSHATRA_REVATI,
				Pada: 4,
			},
		},
		{
			name: "exact 360 normalizes to 0 -> Ashwini p1",
			deg:  360.0,
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
		{
			name: "large positive normalizes -> Ashwini p1",
			deg:  721.5, // 721.5 % 360 = 1.5
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
		{
			name: "large negative normalizes -> Ashwini p1",
			deg:  -358.5, // -358.5 normalized = 1.5
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNakshatraPadaFromDegree(tt.deg)
			assert.Equal(t, tt.expected.Name, got.Name, "nakshatra name mismatch")
			assert.Equal(t, tt.expected.Pada, got.Pada, "nakshatra pada mismatch")
		})
	}
}

func TestGetNakshatraPadaFromDegree_InvalidInputs(t *testing.T) {
	cases := []struct {
		name string
		deg  float64
	}{
		{"NaN input", math.NaN()},
		{"+Inf input", math.Inf(1)},
		{"-Inf input", math.Inf(-1)},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetNakshatraPadaFromDegree(c.deg)
			assert.Equal(t, "", got.Name, "expected empty name for invalid input")
			assert.Equal(t, 0, got.Pada, "expected pada 0 for invalid input")
		})
	}
}

func TestGetNakshatraFromVowel_BasicMappings(t *testing.T) {
	tests := []struct {
		code     string
		expected NakshatraPada
	}{
		{
			code: "CHU",
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 1,
			},
		},
		{
			code: "CHE",
			expected: NakshatraPada{
				Name: NAKSHATRA_ASHWINI,
				Pada: 2,
			},
		},
		{
			code: "AA",
			expected: NakshatraPada{
				Name: NAKSHATRA_KRITTICA,
				Pada: 1,
			},
		},
		{
			// some code mapped in original file
			code: "O",
			expected: NakshatraPada{
				Name: NAKSHATRA_ROHINI,
				Pada: 1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			got := GetNakshatraFromVowel(tt.code)
			assert.Equal(t, tt.expected.Name, got.Name)
			assert.Equal(t, tt.expected.Pada, got.Pada)
		})
	}
}

func TestPlanetCord_CalculateDerivedValues_NormalizationAndNaN(t *testing.T) {
	// Case A: finite longitude that needs normalization and negative speed -> retrograde true
	// Name is set so PlanetSpeedCategory and PlanetSBCLRFVedha can resolve.
	p := &PlanetCord{
		Name:      VENUS,
		Longitude: 721.5, // normalizes to 1.5 -> Aries, Ashwini p1
		Latitude:  -5.25,
		SpeedLong: -0.512, // retrograde Venus -> ati-vakra speed category -> right vedha
	}

	p.CalculateDerivedValues()

	// Sign based on normalized longitude
	assert.Equal(t, SIGN_ARIES, p.Sign)
	// Nakshatra should be Ashwini pada 1
	assert.Equal(t, NAKSHATRA_ASHWINI, p.Nakshatra.Name)
	assert.Equal(t, 1, p.Nakshatra.Pada)
	// Retrograde must be true for negative finite speed
	assert.True(t, p.IsRetro)
	// Speed category and vedha for retrograde Venus
	assert.Equal(t, ATI_VAKRA, p.SpeedCategory, "retrograde Venus at -0.512 should be ati-vakra")
	assert.Equal(t, RIGHT_VEDHA, p.Vedha, "ati-vakra category planet should have right vedha")

	// DMS rounding: ensure LongitudeDMS.ToDegree approximates original absolute degree
	lonDeg := p.LongitudeDMS.ToDegree()
	if math.Abs(lonDeg-math.Abs(p.Longitude)) > 0.02 {
		t.Fatalf("LongitudeDMS.ToDegree expected approx %v, got %v", math.Abs(p.Longitude), lonDeg)
	}

	// Case B: NaN longitude and NaN speed -> sign empty, nakshatra zero, isRetro false,
	// DMS zeroed, speedCategory/vedha empty (PlanetSpeedCategory returns error for NaN)
	p2 := &PlanetCord{
		Name:      SUN,
		Longitude: math.NaN(),
		Latitude:  math.NaN(),
		SpeedLong: math.NaN(),
	}

	p2.CalculateDerivedValues()

	assert.Equal(t, "", p2.Sign, "expected empty sign for NaN longitude")
	assert.Equal(t, "", p2.Nakshatra.Name, "expected empty nakshatra name for NaN longitude")
	assert.Equal(t, 0, p2.Nakshatra.Pada, "expected pada 0 for NaN longitude")
	assert.False(t, p2.IsRetro, "expected IsRetro false for NaN speed")
	assert.Equal(t, "", p2.SpeedCategory, "expected empty speed category for NaN speed")
	assert.Equal(t, "", p2.Vedha, "expected empty vedha for NaN speed")

	// DMS fields for NaN should be zeroed as per NewDMS/ParseFromDegree contract
	assert.Equal(t, 0, p2.LongitudeDMS.D)
	assert.Equal(t, 0, p2.LongitudeDMS.M)
	assert.Equal(t, float32(0), p2.LongitudeDMS.S)

	// Case C: boundary values - sign and nakshatra at exact boundaries
	p3 := &PlanetCord{
		Name:      MARS,
		Longitude: 3.333333, // should map to Ashwini pada 2
		SpeedLong: 0.5,
	}
	p3.CalculateDerivedValues()
	assert.Equal(t, SIGN_ARIES, p3.Sign)
	assert.Equal(t, NAKSHATRA_ASHWINI, p3.Nakshatra.Name)
	assert.Equal(t, 2, p3.Nakshatra.Pada)
	assert.False(t, p3.IsRetro)
	assert.Equal(t, SAMA, p3.SpeedCategory, "Mars at 0.5 should be sama")
	assert.Equal(t, FRONT_VEDHA, p3.Vedha, "sama Mars should have front vedha")

	// Case D: near 360 boundary -> sign Pisces, nakshatra Revati p4
	p4 := &PlanetCord{
		Name:      JUPITER,
		Longitude: 359.999,
		SpeedLong: 0.1,
	}
	p4.CalculateDerivedValues()
	assert.Equal(t, SIGN_PISCES, p4.Sign)
	assert.Equal(t, NAKSHATRA_REVATI, p4.Nakshatra.Name)
	assert.Equal(t, 4, p4.Nakshatra.Pada)
	assert.False(t, p4.IsRetro)
	assert.Equal(t, MADHYAM, p4.SpeedCategory, "Jupiter at 0.1 should be madhyam")
	assert.Equal(t, FRONT_VEDHA, p4.Vedha, "madhyam Jupiter should have front vedha")

	// Case E: Sun with normal positive speed -> non-retro, ati-sheeghra category, left vedha
	p5 := &PlanetCord{
		Name:      SUN,
		Longitude: 45.0, // Taurus
		SpeedLong: 1.05, // above ati-sheeghra threshold for Sun
	}
	p5.CalculateDerivedValues()
	assert.Equal(t, SIGN_TAURUS, p5.Sign)
	assert.False(t, p5.IsRetro)
	assert.Equal(t, ATI_SHEEGHRA, p5.SpeedCategory, "Sun at 1.05 should be ati-sheeghra")
	assert.Equal(t, LEFT_VEDHA, p5.Vedha, "ati-sheeghra Sun should have left vedha")

	// Case F: Rahu always gets left vedha regardless of speed
	p6 := &PlanetCord{
		Name:      RAHU,
		Longitude: 120.0, // Leo
		SpeedLong: -0.05, // negative speed
	}
	p6.CalculateDerivedValues()
	assert.Equal(t, SIGN_LEO, p6.Sign)
	assert.True(t, p6.IsRetro)
	assert.Equal(t, VAKRA, p6.SpeedCategory, "Rahu at -0.05 should be vakra")
	assert.Equal(t, LEFT_VEDHA, p6.Vedha, "Rahu should always have left vedha")
}
