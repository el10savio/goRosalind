package computinggccontent

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given: At most 10 DNA strings in FASTA format
// (of length at most 1 kbp each).
//
// Return: The ID of the string having the highest GC-content,
// followed by the GC-content of that string.

// FASTA struct defines the sequence 
// defined in FASTA format
type FASTA struct {
	Label string
	DNA   string
}

// Compute processes the GCContent for each value 
// in the FASTA collection and returns the label
// of the sequence with the highest
// GCContent and it value
func Compute(FASTAValues []FASTA) (string, float32, error) {
	// Check if input values are empty
	if len(FASTAValues) <= 0 {
		return "", 0.0, definitions.ErrEmptyList
	}

	// Allocate GCContents slice to 
	// store each value's GCContent
	GCContents := make([]float32, len(FASTAValues))

	// Compute each value's GCContent
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

// compute the GCContent for a given FASTA struct
func compute(label string, DNA string) (float32, error) {
	// Check if DNA sequence is empty
	if len(DNA) <= 0 {
		return 0.0, definitions.ErrEmptyString
	}

	// Check if label is empty
	if len(label) <= 0 {
		return 0.0, definitions.ErrEmptyLabel
	}

	// Convert DNA sequence to uppercase
	DNA = strings.ToUpper(DNA)
	
	// Compute the GCContent for 
	// the DNA sequence
	GCContent := gcContent(DNA)

	return GCContent, nil
}

// gcContent returns the ratio of number of 
// 'G' & 'C' to the length of the sequence
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

// Max returns the index of the highest 
// element in the slice and its value
func Max(slice []float32) (maxIndex int, max float32) {
	for index, element := range slice {
		if index == 0 || max < element {
			max = element
			maxIndex = index
		}
	}
	return
}
