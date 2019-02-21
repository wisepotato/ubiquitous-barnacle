package main

import (
	"fmt"
	"math"
)

func main() {
	mult := []int{3, 5}
	max := 1000

	result := SumMultiples(mult, max)

	fmt.Printf(`
Challenge 1: Multiples of 3 and 5

Problem: Sum all multiples of 3 and 5 up to 1000

Solution: %d

Approach:
- Loop through every natural number up to 999
- Sum those that are divisible by 3 or 5

Notes:
- The challenge has a trap: multiples of 3 and 5 can overlap, but should be
  counted only once. Therefore summing all multiples under 1000 is not a
  good strategy.
`, result)
}

// Need to know if the slice tonains my int
func contains(slice []int, item int) bool {
	set := make(map[int]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// SumMultiples ...
func SumMultiples(multiples []int, maxValue int) int {
	sum := 0
	//current := 1
	var combinations []int
	var maxIterations int
	// Get all combinations
	for _, multiple := range multiples {
		for _, multipleNested := range multiples {
			if multiple != multipleNested && !contains(combinations, multiple*multipleNested) {
				combinations = append(combinations, multiple*multipleNested)
				fmt.Printf("added %d", (multiple * multipleNested))
			}
		}
		// Get max amount of iterations
		maxIterations = int(math.Floor(float64(maxValue-1) / float64(multiple)))
		// sum up to that amount
		if maxIterations != 0 {
			sum += multiple * (maxIterations * (1 + maxIterations) / 2) // 5*( sum_i^maxIterations i)
		}
	}

	// take out the combinations
	for _, combination := range combinations {
		maxIterations = int(math.Floor(float64(maxValue-1) / float64(combination)))
		if maxIterations != 0 {
			sum -= combination * (maxIterations * (1 + maxIterations) / 2)
		}

	}

	return sum
}
