package countingpointmutations

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// HammingDistance ...
func HammingDistance(first string, second string) (int, error) {
	// Check if input values are empty
	if len(first) == 0 || len(second) == 0 {
		return 0, definitions.ErrEmptyString
	}

	// Check if input values are unequal length
	if len(first) != len(second) {
		return 0, definitions.ErrUnequalLength
	}

	first = strings.ToUpper(first)
	second = strings.ToUpper(second)

	return hammingDistance(first, second), nil
}

func hammingDistance(first string, second string) int {
	distance := 0
	runesFirst := []rune(first)
	runesSecond := []rune(second)

	for index := 0; index < len(runesFirst); index++ {
		if runesFirst[index] == runesSecond[index] {
			distance++
		}
	}

	return distance
}
