package baselib

import "math"

// isInvalidFloat reports whether v is NaN or ±Inf.
func isInvalidFloat(v float64) bool {
	return math.IsNaN(v) || math.IsInf(v, 0)
}

// NormalizeAngle returns an angle normalized into the range [0, 360).
// If the input is NaN or infinite, it returns NaN.
func NormalizeAngle(angle float64) float64 {
	// Fast path: already in range, so skip the math.Mod call. NaN fails both
	// comparisons and +/-Inf fail one each, so invalid inputs fall through below.
	if angle >= 0 && angle < 360.0 {
		return angle
	}

	if isInvalidFloat(angle) {
		return math.NaN()
	}

	a := math.Mod(angle, 360.0)
	if a < 0 {
		a += 360.0
	}

	// Mod can return 360.0 for exact multiples; force wrap to 0.0
	if a >= 360.0 {
		a -= 360.0
	}

	return a
}

// ValidAngle reports whether d is a finite angle and returns it normalized
// into [0, 360). Callers that previously guarded NaN/Inf themselves and then
// called NormalizeAngle should use this seam instead.
func ValidAngle(d float64) (normalized float64, ok bool) {
	if isInvalidFloat(d) {
		return 0, false
	}
	return NormalizeAngle(d), true
}
