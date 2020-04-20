package computinggccontent

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// Compute
func Compute(label string, DNA string) (string, error) {
	if len(DNA) <= 0 {
		return "", definitions.ErrEmptyString
	}

	if len(label) <= 0 {
		return "", definitions.ErrEmptyLabel
	}

	DNA = strings.ToLower(DNA)
	gcContent := GCContent(DNA)

	return "", nil
}

func GCContent(DNA string) float64 {
	gcCount := 0
	runes := []rune(DNA)

	for _, runeValue := range runes {
		if runeValue == rune('G') || runeValue == rune('C') {
			gcCount++
		}
	}

	return float64(gcCount / len(runes))
}
