package main

import (
	"strconv"
	"testing"
)

// TestSumMultiples ...
func TestSumMultiples(t *testing.T) {
	testCases := []struct {
		multiples      []int
		maxValue       int
		expectedResult int
	}{
		{[]int{3}, 5, 3},
		{[]int{5}, 24, 50},
		{[]int{3, 5}, 10, 23},
		{[]int{3, 5}, 20, 78},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actualResult := SumMultiples(tc.multiples, tc.maxValue)
			if actualResult != tc.expectedResult {
				t.Logf("Multiples: %v, max value: %d", tc.multiples, tc.maxValue)
				t.Errorf("Expected %d, got %d", tc.expectedResult, actualResult)
			}
		})
	}
}
