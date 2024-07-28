package polynomial

import (
	"testing"
)

func TestNewPolynomial(t *testing.T) {
	coefficients := []Coefficient{1, 2, 3}

	polynomial := NewPolynomial(coefficients)

	if len(polynomial.coefficients) != 3 {
		t.Errorf("New(%v) = %v; want %v", coefficients, polynomial.coefficients, coefficients)
	}

	for idx, coefficient := range coefficients {
		if polynomial.coefficients[idx] != coefficient {
			t.Errorf("New(%v) = %v; want %v", coefficients, polynomial.coefficients, coefficients)
		}
	}
}

func TestGetDegree(t *testing.T) {
	coefficients := []Coefficient{1, 2, 3}

	polynomial := NewPolynomial(coefficients)

	degree := polynomial.Degree()

	if degree != 3 {
		t.Errorf("Degree() = %d; want %d", degree, 3)
	}
}

func TestInterpolate(t *testing.T) {
	points := []Point{
		{1, 2},
		{2, 3},
	}

	polynomial, err := Interpolate(points)

	if err != nil {
		t.Errorf("Interpolate(%v) = %v; want %v", points, err, nil)
	}

	if polynomial.Degree() != 1 {
		t.Errorf("Interpolate(%v) = %d; want %d", points, polynomial.Degree(), 1)
	}

	structure := polynomial.String()

	if structure != "f(x) = 1.00 + 1.00x^1" {
		t.Errorf("Interpolate(%v) = %s; want %s", points, structure, "f(x) = 1.00 + 1.00x^1")
	}
}
