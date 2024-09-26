package utils

import (
	"math"
)

// Calculate the mean of a slice of numbers and round the result
func Mean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return math.Round(sum / float64(len(numbers)))
}

// Calculate the variance of a slice of numbers and round the result
func Variance(numbers []float64, mean float64) float64 {
	var varianceSum float64
	for _, num := range numbers {
		varianceSum += (num - mean) * (num - mean)
	}
	return math.Round(varianceSum / float64(len(numbers)))
}

// Calculate the standard deviation of a slice of numbers and round the result
func StandardDev(numbers []float64) float64 {
	mean := Mean(numbers)
	variance := Variance(numbers, mean)
	return math.Round(math.Sqrt(variance))
}
