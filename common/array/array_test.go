package array

import "testing"

func TestFindDuplicates(t *testing.T) {
	testCases := []struct {
		points   []float64
		expected bool
	}{
		{[]float64{1, 2, 3, 4, 5}, false},
		{[]float64{1, 2, 3, 4, 5, 1}, true},
		{[]float64{1, 2, 3, 4, 5, 1, 2}, true},
		{[]float64{1, 2, 3, 4, 5, 1, 2, 3}, true},
	}

	for _, tc := range testCases {
		result := FindDuplicates(tc.points)
		if result != tc.expected {
			t.Errorf("FindDuplicates(%v) = %t; want %t", tc.points, result, tc.expected)
		}
	}
}

func TestFindDuplicatesWithStruct(t *testing.T) {
	type testStruct struct {
		x float64
		y float64
	}

	testCases := []struct {
		points   []testStruct
		expected bool
	}{
		{[]testStruct{{1, 2}, {3, 4}, {5, 6}}, false},
		{[]testStruct{{1, 2}, {3, 4}, {5, 6}, {1, 2}}, true},
		{[]testStruct{{1, 2}, {3, 4}, {5, 6}, {1, 2}, {3, 4}}, true},
	}

	for _, tc := range testCases {
		result := FindDuplicates(tc.points)
		if result != tc.expected {
			t.Errorf("FindDuplicates(%v) = %t; want %t", tc.points, result, tc.expected)
		}
	}
}
