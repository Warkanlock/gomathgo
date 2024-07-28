package mathops

import "fmt"

// Add adds two integers and returns the result.
func SafeAdd(a, b float64) float64 {
	return a + b
}

// Subtract subtracts the second integer from the first and returns the result.
func SafeSub(a, b float64) float64 {
	return a - b
}

// Multiply multiplies two integers and returns the result.
func SafeMul(a, b float64) float64 {
	return a * b
}

// Divide divides the first integer by the second and returns the result.
// Returns an error if the second integer is zero.
func SafeDiv(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return a / b, nil
}
