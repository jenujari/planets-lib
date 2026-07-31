package baselib

import (
	"math"
	"testing"
)

func TestCalcNavanshRashi(t *testing.T) {
	tests := []struct {
		pl_long  float64
		expected int
		name     string
	}{
		{0.0, 1, "Aries 0.0 -> Aries Navansh"},
		{3.4, 2, "Aries 3.4 -> Taurus Navansh"}, // 30/9 = 3.333
		{29.9, 9, "Aries 29.9 -> Sagittarius Navansh"},
		{30.0, 10, "Taurus 0.0 -> Capricorn Navansh"},
		{33.4, 11, "Taurus 3.4 -> Aquarius Navansh"},
		{60.0, 7, "Gemini 0.0 -> Libra Navansh"},
		{90.0, 4, "Cancer 0.0 -> Cancer Navansh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := CalcNavanshRashi(tt.pl_long)
			if got != tt.expected {
				t.Errorf("CalcNavanshRashi(%v) = %v, want %v", tt.pl_long, got, tt.expected)
			}
		})
	}
}

// Invalid floats previously reached int(NaN / 30.0), whose result is undefined by the
// Go spec and indexed SIGNS out of range on amd64, panicking. They must now return the
// zero-value signal instead.
func TestCalcNavanshRashi_InvalidInputs(t *testing.T) {
	tests := []struct {
		name    string
		pl_long float64
	}{
		{"NaN", math.NaN()},
		{"positive infinity", math.Inf(1)},
		{"negative infinity", math.Inf(-1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("CalcNavanshRashi(%v) panicked: %v", tt.pl_long, r)
				}
			}()

			num, name := CalcNavanshRashi(tt.pl_long)
			if num != 0 || name != "" {
				t.Errorf("CalcNavanshRashi(%v) = (%v, %q), want (0, \"\")", tt.pl_long, num, name)
			}
		})
	}
}

// Valid longitudes must always land inside SIGNS, including angles needing normalization.
func TestCalcNavanshRashi_ResultAlwaysInRange(t *testing.T) {
	for d := -720.0; d <= 1080.0; d += 0.37 {
		num, name := CalcNavanshRashi(d)
		if num < 1 || num > 12 {
			t.Fatalf("CalcNavanshRashi(%v) = %v, out of 1..12", d, num)
		}
		if name != SIGNS[num-1] {
			t.Fatalf("CalcNavanshRashi(%v) name %q does not match SIGNS[%d]", d, name, num-1)
		}
	}
}
