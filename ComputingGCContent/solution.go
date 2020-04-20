package computinggccontent

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// FASTA ...
type FASTA struct {
	Label string
	DNA   string
}

// Compute ...
func Compute(FASTAValues []FASTA) (string, float64, error) {
	if len(FASTAValues) <= 0 {
		return "", 0.0, definitions.ErrEmptyList
	}

	GCContents := make([]float64, len(FASTAValues))

	for index, fasta := range FASTAValues {
		GCContent, err := compute(fasta.Label, fasta.DNA)
		if err != nil {
			return "", 0.0, err
		}
		GCContents[index] = GCContent
	}

	// find max & return index
	index, value := Max(GCContents)
	value = 100 * value

	return FASTAValues[int(index)].Label, value, nil
}

// compute ...
func compute(label string, DNA string) (float64, error) {
	if len(DNA) <= 0 {
		return 0.0, definitions.ErrEmptyString
	}

	if len(label) <= 0 {
		return 0.0, definitions.ErrEmptyLabel
	}

	DNA = strings.ToLower(DNA)
	GCContent := gcContent(DNA)

	return GCContent, nil
}

// gcContent ...
func gcContent(DNA string) float64 {
	gcCount := 0
	runes := []rune(DNA)

	for _, runeValue := range runes {
		if runeValue == rune('G') || runeValue == rune('C') {
			gcCount++
		}
	}

	return float64(gcCount / len(runes))
}

// Max ...
func Max(slice []float64) (index int, max float64) {
	for index, element := range slice {
		if index == 0 || max < element {
			max = element
		}
	}
}
