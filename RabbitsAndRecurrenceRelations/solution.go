package rabbitsandrecurrencerelations

import (
	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given: Positive integers n≤40 and k≤5.
//
// Return: The total number of rabbit pairs that will 
// be present after n months, if we begin with 1 pair 
// and in each generation, every pair of reproduction-age 
// rabbits produces a litter of k rabbit pairs
// (instead of only 1 pair).

// Total returns the total rabbit pairs based on 
// the number of months and pairs after 
// each reproductive stage
func Total(months int, litter int) (int, error) {
	if months <= 0 || litter <= 0 {
		return 0, definitions.ErrNotPositiveValue
	}

	return fibonacci(months, litter), nil
}

// fibonacci returns the nth fibonacci number 
// given that the 2nd fibonacci number is 
// a parameter supplied
func fibonacci(n int, init int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return init
	}

	if n <= 4 {
		return fibonacci(n-1, init) + fibonacci(n-2, init)
	}

	return fibonacci(n-1, init) + fibonacci(n-2, init)*init
}
