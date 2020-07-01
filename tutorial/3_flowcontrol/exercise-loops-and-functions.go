package main

import (
	"fmt"
	"math"
)

// TODO: a function to implement Newtonian
// input is x and accuracy
// z := float64(1)
// returns z and number of loops

// This is a function to get the squreroot of a number
// until the given accuracy
func Sqrt(x, accuracy float64) (float64, int) {
	z := float64(1)
	count := 1
	// Keep looping if accuracy is not achieved
	for math.Abs(z*z-x) > accuracy {
		z -= (z*z - x) / (2 * z)
		count += 1
	}

	return z, count
}

// compares Newtonian and sqrt results
func main() {
	number := 64.0
	acc := 0.00001
	result, count := Sqrt(number, acc)
	fmt.Printf("The square root of %v is %v after %v loops\n", number, result, count)
}
