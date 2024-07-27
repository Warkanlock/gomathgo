package common

import (
	"testing"
)

func TestNewPolynomial(t *testing.T) {
	coefficients := []Coefficient{1, 2, 3}

	polynomial := Polynomial{}.New(coefficients)

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

	polynomial := Polynomial{}.New(coefficients)

	degree := polynomial.Degree()

	if degree != 3 {
		t.Errorf("Degree() = %d; want %d", degree, 3)
	}
}
