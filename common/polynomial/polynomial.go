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
func NewPolynomial(coefficients []Coefficient) *Polynomial {
	return &Polynomial{
		coefficients,
	}
}

// Add two polynomials together
func (p Polynomial) Add(p2 Polynomial) *Polynomial {
	// pick the maximum degree of the two polynomials
	newCoefficients := make([]Coefficient, max(len(p.coefficients), len(p2.coefficients)))

	// coeffients of the first polynomial
	for i, c := range p2.coefficients {
		newCoefficients[i] += c

		// if there's a coefficient in the base polynomial
		if len(p.coefficients) > 0 && i < len(p.coefficients) {
			newCoefficients[i] += p.coefficients[i]
		}
	}

	return &Polynomial{
		coefficients: newCoefficients,
	}
}

// Multiply two polynomials together
func (p Polynomial) Multiply(p2 Polynomial) *Polynomial {
	coefficients := make([]Coefficient, len(p.coefficients)+len(p2.coefficients)-1)

	for i, c1 := range p.coefficients {
		for j, c2 := range p2.coefficients {
			coefficients[i+j] += c1 * c2
		}
	}

	return &Polynomial{
		coefficients,
	}
}

// Interpolate a polynomial based on a list of points
func Interpolate(points []Point) (*Polynomial, error) {
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

	var terms []Polynomial

	// iterate over the points
	for i, pointA := range points {
		var term Polynomial
		// evaluate polynomial at (x,y) on point
		x, y := pointA.X, pointA.Y

		for j, pointB := range points {
			if i == j {
				continue
			}

			xj := pointB.X

			a0 := Coefficient(-xj / (x - xj))
			a1 := Coefficient(1 / (x - xj))

			// create a polynomial based on the x values
			p := NewPolynomial([]Coefficient{a0, a1})

			// multiply the polynomial with the term
			term = *p
		}

		// multiply the term with the y value
		term = *term.Multiply(*NewPolynomial([]Coefficient{Coefficient(y)}))

		// append term to a list of terms
		terms = append(terms, term)
	}

	if terms == nil {
		return nil, fmt.Errorf("cannot perform an interpolation without any valid terms")
	}

	// add the terms together
	base := NewPolynomial([]Coefficient{0})

	for _, term := range terms {
		base = base.Add(term)
	}

	return base, nil
}

// Retrieve the degree of the polynomial
func (p Polynomial) Degree() int {
	return len(p.coefficients) - 1
}

// Show a polynomial based on the points provided
func (p Polynomial) String() string {
	var completePolynomial string

	completePolynomial = "f(x) = "

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

	completePolynomial = fmt.Sprintf("%s", completePolynomial)

	return fmt.Sprintf(completePolynomial)
}
