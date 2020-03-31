package countingdnanucleotides

import (
	"strings"
)

func Count(DNA string) map[string]int {
	if len(DNA) <= 0 {
		return map[string]int{}
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

	return symbols
}
