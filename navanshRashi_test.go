package baselib

import (
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
