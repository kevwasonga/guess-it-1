package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Function to calculate the mean of two numbers
func Mean(num1, num2 float64) float64 {
	return (num1 + num2) / 2
}

// Function to calculate the standard deviation of two numbers
func StandardDev(num1, num2 float64) float64 {
	mean := Mean(num1, num2)
	variance := (math.Pow(num1-mean, 2) + math.Pow(num2-mean, 2)) / 2
	return math.Sqrt(variance)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var prevInput float64
	var firstInput bool = true

	fmt.Println("Enter floating-point numbers (one per line, type 'done' to finish):")

	for scanner.Scan() {
		input := scanner.Text()
		if input == "done" {
			break
		}

		// Convert input string to float64
		currentInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}

		if firstInput {
			// If it's the first input, just store it and wait for the next
			prevInput = currentInput
			firstInput = false
		} else {
			// Calculate the mean and standard deviation for the range
			mean := Mean(prevInput, currentInput)
			stdDev := StandardDev(prevInput, currentInput)
			
			// Calculate the range (mean ± standard deviation)
			lowerBound := mean - stdDev
			upperBound := mean + stdDev

			// Print the current input and the range for the next input
			fmt.Printf("%d\n", int(currentInput))
			fmt.Printf("%d %d\n", int(lowerBound), int(upperBound))

			// Update previous input
			prevInput = currentInput
		}
	}
	fmt.Println("Program finished.")
}
