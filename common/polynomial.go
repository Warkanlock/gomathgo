package common

import "fmt"

// Define a 2d-point in a graph
type Point struct {
	X float64
	Y float64
}

type Coefficient float64

// Structure of a polynomial
type Polynomial struct {
	coefficients []Coefficient
}

// Create a polynomial based on a list of coefficients
func (Polynomial) New(coefficients []Coefficient) *Polynomial {
	return &Polynomial{
		coefficients,
	}
}

// Retrieve the degree of the polynomial
func (p Polynomial) Degree() int {
	return len(p.coefficients)
}

// Show a polynomial based on the points provided
func (p Polynomial) Show() {
	fmt.Printf("f(x) = ")

	for idx, coefficient := range p.coefficients {
		if idx == 0 {
			fmt.Printf("%f", coefficient)
		}

		if coefficient == 0 {
			// skip if the coefficient is 0
			continue
		}

		fmt.Printf("%fx^%d", coefficient, idx)
	}

	fmt.Print("\n")
}
