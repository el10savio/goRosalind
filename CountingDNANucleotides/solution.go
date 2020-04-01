package countingdnanucleotides

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given: A DNA string s of length at most 1000 nt.
//
// Return: Four integers (separated by spaces)
// counting the respective number of times
// that the symbols 'A', 'C', 'G',
// and 'T' occur in s.

// Count return the letter map
// from the given DNA string
func Count(DNA string) (map[string]int, error) {
	if len(DNA) <= 0 {
		return map[string]int{}, definitions.ErrEmptyString
	}

	DNA = strings.ToUpper(DNA)

	symbols := map[string]int{
		"A": 0,
		"C": 0,
		"G": 0,
		"T": 0,
	}

	for _, alphabet := range DNA {
		alphabetString := string(alphabet)
		if alphabetString == "A" {
			symbols["A"]++
		}
		if alphabetString == "C" {
			symbols["C"]++
		}
		if alphabetString == "G" {
			symbols["G"]++
		}
		if alphabetString == "T" {
			symbols["T"]++
		}
	}

	return symbols, nil
}
