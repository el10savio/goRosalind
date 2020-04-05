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

	return 0, nil
}
