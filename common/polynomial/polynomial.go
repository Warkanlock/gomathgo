package polynomial

import (
	"fmt"
	"gomathgo/common/array"
)

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

func (Polynomial) Interpolate(points []Point) (*Polynomial, error) {
	if len(points) == 0 {
		return nil, fmt.Errorf("cannot perform an interpolation without any valid points")
	}

	XValues := func(points []Point) []float64 {
		xValues := make([]float64, len(points))

		for idx, point := range points {
			xValues[idx] = point.X
		}

		return xValues
	}

	isDuplicate := array.FindDuplicates(XValues(points))

	if isDuplicate {
		return nil, fmt.Errorf("cannot perform an interpolation with duplicate (x,y) pair values")
	}

	return nil, nil
}

// Retrieve the degree of the polynomial
func (p Polynomial) Degree() int {
	return len(p.coefficients)
}

// Show a polynomial based on the points provided
func (p Polynomial) String() string {
	var completePolynomial string

	completePolynomial = "\nf(x) = "

	for idx, coefficient := range p.coefficients {
		if idx == 0 {
			completePolynomial = fmt.Sprintf("%s%.2f", completePolynomial, coefficient)
			continue
		}

		if coefficient == 0 {
			// skip if the coefficient is 0
			continue
		}

		completePolynomial = fmt.Sprintf("%s + %.2fx^%d", completePolynomial, coefficient, idx)
	}

	completePolynomial = fmt.Sprintf("%s\n", completePolynomial)

	return fmt.Sprintf(completePolynomial)
}
