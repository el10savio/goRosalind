package rabbitsandrecurrencerelations

import (
	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// Total
func Total(months int, litter int) (int, error) {
	if months <= 0 || litter <= 0 {
		return 0, definitions.ErrNotPositiveValue
	}

	return fibonacci(months, litter), nil
}

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
