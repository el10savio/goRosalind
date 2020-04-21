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
func Compute(FASTAValues []FASTA) (string, float32, error) {
	if len(FASTAValues) <= 0 {
		return "", 0.0, definitions.ErrEmptyList
	}

	GCContents := make([]float32, len(FASTAValues))

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

	// Remove '>' from label
	label := FASTAValues[int(index)].Label
	label = strings.ReplaceAll(label, ">", "")

	return label, value, nil
}

// compute ...
func compute(label string, DNA string) (float32, error) {
	if len(DNA) <= 0 {
		return 0.0, definitions.ErrEmptyString
	}

	if len(label) <= 0 {
		return 0.0, definitions.ErrEmptyLabel
	}

	DNA = strings.ToUpper(DNA)
	GCContent := gcContent(DNA)

	return GCContent, nil
}

// gcContent ...
func gcContent(DNA string) (ratio float32) {
	gcCount := 0
	runes := []rune(DNA)

	for _, runeValue := range runes {
		if runeValue == rune('G') || runeValue == rune('C') {
			gcCount++
		}
	}

	if gcCount != 0 {
		ratio = float32(float32(gcCount)/float32(len(DNA))*100) / 100
	}

	return
}

// Max ...
func Max(slice []float32) (maxIndex int, max float32) {
	for index, element := range slice {
		if index == 0 || max < element {
			max = element
			maxIndex = index
		}
	}
	return
}
