package countingpointmutations

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given: Two DNA strings s and t of equal length.
//
// Return: The Hamming distance dH(s,t).

// HammingDistance is a wrapper to validate the
// input values and preprocess them before
// calculating their hamming distance
func HammingDistance(first string, second string) (int, error) {
	// Check if input values are empty
	if len(first) == 0 || len(second) == 0 {
		return 0, definitions.ErrEmptyString
	}

	// Check if input values are unequal length
	if len(first) != len(second) {
		return 0, definitions.ErrUnequalLength
	}

	// Convert both strings to uppercase
	first = strings.ToUpper(first)
	second = strings.ToUpper(second)

	// Check if they're equal 
	if first == second {
		return 0, nil
	}

	return hammingDistance(first, second), nil
}

// hammingDistance compares two strings and
// returns how many times they are
// different with each other
func hammingDistance(first string, second string) int {
	distance := 0
	runesFirst := []rune(first)
	runesSecond := []rune(second)

	for index := 0; index < len(runesFirst); index++ {
		if runesFirst[index] != runesSecond[index] {
			distance++
		}
	}

	return distance
}
