package baseLib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToDMS(T *testing.T) {
	tests := []struct {
		input    float64
		expected string
	}{
		{123.456, "123°27'21.60\""},
		{-123.456, "-123°27'21.60\""},
		{1.654, "1°39'14.40\""},
		{0, "0°0'0.00\""},
	}

	for _, test := range tests {
		result := NewDMS(test.input)
		assert.Equal(T, test.expected, result.String(), "Expected %s, got %s", test.expected, result)
	}
}

func Test_ToDegrees(T *testing.T) {
	tests := []struct {
		input    DMS
		expected float64
	}{
		{DMS{false, 123, 27, 21.60}, 123.456},
		{DMS{true, 123, 27, 21.60}, -123.456},
		{DMS{false, 1, 39, 14.40}, 1.654},
		{DMS{false, 0, 0, 0.00}, 0},
	}

	for _, test := range tests {
		result := test.input.ToDegree()
		assert.InDelta(T, test.expected, result, 0.001, "Expected %f, got %f", test.expected, result)
	}
}
