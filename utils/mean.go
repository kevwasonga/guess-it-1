package utils

import (
	"math"
)

// Step 1: Calculate the mean
func Mean(numbers []float64) float64 {
	var total float64
	for _, num := range numbers {
		total += num
	}
	return total / float64(len(numbers))
}

// Step 2: Calculate the variance
func Variance(numbers []float64, mean float64) float64 {
	var varianceSum float64
	for _, num := range numbers {
		varianceSum += (num - mean) * (num - mean)
	}
	return varianceSum / float64(len(numbers))
}

// Step 3: Calculate the standard deviation
func StandardDev(numbers []float64) float64 {
	mean := Mean(numbers)
	variance := Variance(numbers, mean)
	return math.Sqrt(variance)
}
