package main

import "fmt"

func main() {
	mult := []int{3, 5}
	max := 1000

	result := SumMultiples(mult, max)

	fmt.Println(result)
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
