package main

import "fmt"

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

// SumMultiples ...
func SumMultiples(multiples []int, maxValue int) int {
	sum := 0
	current := 1
	divisible := false

	for current < maxValue {
		for _, multiple := range multiples {
			if current%multiple == 0 {
				divisible = true
				break
			}

		}
		if divisible {
			sum += current
			divisible = false
		}
		current++
	}

	return sum
}
