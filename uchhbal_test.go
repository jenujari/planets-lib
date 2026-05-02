package baselib

import (
	"math"
	"testing"
)

func TestCalcUchhBal(t *testing.T) {
	tests := []struct {
		name     string
		pl_name  string
		pl_long  float64
		expected float64
	}{
		{"Sun at exaltation", SUN, 10, 100.0},
		{"Sun at 90 deg from exaltation", SUN, 100, 75.0},
		{"Sun at 180 deg from exaltation", SUN, 190, 50.0},
		{"Sun at 270 deg from exaltation", SUN, 280, 75.0},
		{"Moon at exaltation", MOON, 33, 100.0},
		{"Jupiter at exaltation", JUPITER, 95, 100.0},
		{"Unknown planet", "Unknown", 100, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalcUchhBal(tt.pl_long, tt.pl_name)

			// Using a small epsilon for float comparison
			if math.Abs(got-tt.expected) > 1e-9 {
				t.Errorf("CalcUchhBal(%v, %v) = %v, want %v", tt.pl_long, tt.pl_name, got, tt.expected)
			}
		})
	}
}

var test_bal float64

func BenchmarkCalcUchhBal(b *testing.B) {
	var result float64
	for b.Loop() {
		result = CalcUchhBal(10, SUN)
	}
	test_bal = result
}
