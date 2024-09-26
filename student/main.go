package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"guess/utils" // Ensure this path is correct based on your project structure
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputs []float64

	for scanner.Scan() {
		// Read input from the user

		input := scanner.Text()

		// Exit condition
		if input == "exit" {
			break
		}

		// Convert input string to float64
		currentInput, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error parsing number:", err)
			continue
		}

		// Add the current input to the slice
		inputs = append(inputs, currentInput)

		// If more than 10 inputs, remove the oldest one
		// if len(inputs) > 10 {
		// 	inputs = inputs[1:]
		// }
		if len(inputs) > 1 {
			// Calculate the mean using the last 10 or fewer inputs
			mean := utils.Mean(inputs)

			// Calculate the standard deviation using the current inputs
			stdDev := utils.StandardDev(inputs)

			// Calculate the lower and upper bounds
			lowerBound := mean - (2 * stdDev)
			upperBound := mean + (2 * stdDev)

			// // Print the mean, standard deviation, and range
			// fmt.Printf("Mean of last %d inputs: %.2f\n", len(inputs), mean)
			// fmt.Printf("Standard Deviation: %.2f\n", stdDev)

			fmt.Printf("%.0f %.0f\n", lowerBound, upperBound)
		}
	}
}
