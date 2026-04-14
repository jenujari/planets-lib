package baselib

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSignFrmDegree_BasicsAndBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		deg      float64
		expected string
	}{
		{
			name:     "degree 0 -> Aries",
			deg:      0.0,
			expected: SIGN_ARIES,
		},
		{
			name:     "degree 29.999 -> Aries (upper boundary exclusive)",
			deg:      29.999,
			expected: SIGN_ARIES,
		},
		{
			name:     "degree 30 -> Taurus (start of second sign)",
			deg:      30.0,
			expected: SIGN_TAURUS,
		},
		{
			name:     "degree 59.999 -> Taurus (just below 60)",
			deg:      59.999,
			expected: SIGN_TAURUS,
		},
		{
			name:     "degree 359.999 -> Pisces (near wrap-around)",
			deg:      359.999,
			expected: SIGN_PISCES,
		},
		{
			name:     "degree 360 -> Aries (exact multiple -> normalized to 0)",
			deg:      360.0,
			expected: SIGN_ARIES,
		},
		{
			name:     "degree 30.0*3 -> Cancer",
			deg:      30.0 * 3.0,
			expected: SIGN_CANCER,
		},
		{
			name:     "degree 30.0*11 -> Pisces (start of last sign)",
			deg:      30.0 * 11.0,
			expected: SIGN_PISCES,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSignFrmDegree(tt.deg)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestGetSignFrmDegree_Normalization(t *testing.T) {
	tests := []struct {
		name     string
		deg      float64
		expected string
	}{
		{
			name:     "large positive angle normalizes (721.5 -> 1.5) -> Aries",
			deg:      721.5,
			expected: SIGN_ARIES, // 721.5 % 360 = 1.5 -> Aries
		},
		{
			name:     "large negative angle normalizes (-358.5 -> 1.5) -> Aries",
			deg:      -358.5,
			expected: SIGN_ARIES, // -358.5 normalized -> 1.5 -> Aries
		},
		{
			name:     "small negative just below 0 -> maps to Pisces",
			deg:      -0.0001,
			expected: SIGN_PISCES,
		},
		{
			name:     "exact multiple of 360 * -2 -> Aries",
			deg:      -720.0,
			expected: SIGN_ARIES,
		},
		{
			name:     "just above 360 wraps -> small positive -> Aries",
			deg:      360.0001,
			expected: SIGN_ARIES,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSignFrmDegree(tt.deg)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestGetSignFrmDegree_InvalidInputs(t *testing.T) {
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
			got := GetSignFrmDegree(c.deg)
			assert.Equal(t, "", got, "invalid floating inputs should return empty string")
		})
	}
}

func TestGetSignFrmDegree_ExactBoundaries(t *testing.T) {
	// Verify exact boundaries around multiples of 30 behave as expected.
	tests := []struct {
		name     string
		deg      float64
		expected string
	}{
		{
			name:     "just below 30 -> Aries",
			deg:      30.0 - 1e-6,
			expected: SIGN_ARIES,
		},
		{
			name:     "exactly 30 -> Taurus",
			deg:      30.0,
			expected: SIGN_TAURUS,
		},
		{
			name:     "just below 360 -> Pisces",
			deg:      360.0 - 1e-6,
			expected: SIGN_PISCES,
		},
		{
			name:     "just above 330 (11*30) -> Pisces (start of last sign)",
			deg:      330.0 + 1e-6,
			expected: SIGN_PISCES,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSignFrmDegree(tt.deg)
			assert.Equal(t, tt.expected, got)
		})
	}
}
