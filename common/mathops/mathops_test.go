package mathops

import (
	"testing"
)

func TestSafeAdd(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{1, 2, 3},
		{1, 5, 6},
	}

	for _, tc := range testCases {
		result := SafeAdd(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSafeSub(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{1, 2, -1},
		{2, 3, -1},
		{10, -10, 20},
	}

	for _, tc := range testCases {
		result := SafeSub(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Subtract(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSafeMul(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{1, 2, 2},
		{2, 3, 6},
		{10, -10, -100},
	}

	for _, tc := range testCases {
		result := SafeMul(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Multiply(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestSafeDiv(t *testing.T) {
	testCases := []struct {
		a, b, expected float64
	}{
		{1, 2, 0.5},
		{2, 3, 0.6666666666666666},
		{10, -10, -1},
	}

	for _, tc := range testCases {
		result, err := SafeDiv(tc.a, tc.b)

		if err != nil {
			t.Errorf("Error: %s", err)
		}

		if result != tc.expected {
			t.Errorf("Divide(%f, %f) = %f; want %f", tc.a, tc.b, result, tc.expected)
		}
	}
}
