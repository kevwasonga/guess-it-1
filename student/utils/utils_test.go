package utils

import (
	"testing"
)

func TestMean(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{10, 20, 30, 40, 50}, 30},
		{[]float64{10.5, 20.5, 30.5}, 21},
	}

	for _, test := range tests {
		result := Mean(test.numbers)
		if result != test.expected {
			t.Errorf("Mean(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}

func TestVariance(t *testing.T) {
	tests := []struct {
		numbers  []float64
		mean     float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3, 2},
		{[]float64{10, 20, 30, 40, 50}, 30, 200},
		{[]float64{10.5, 20.5, 30.5}, 20, 67},
	}

	for _, test := range tests {
		result := Variance(test.numbers, test.mean)
		if result != test.expected {
			t.Errorf("Variance(%v, %v) = %v; expected %v", test.numbers, test.mean, result, test.expected)
		}
	}
}

func TestStandardDev(t *testing.T) {
	tests := []struct {
		numbers  []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 1},
		{[]float64{10, 20, 30, 40, 50}, 14},
		{[]float64{10.5, 20.5, 30.5}, 8},
	}

	for _, test := range tests {
		result := StandardDev(test.numbers)
		if result != test.expected {
			t.Errorf("StandardDev(%v) = %v; expected %v", test.numbers, result, test.expected)
		}
	}
}
