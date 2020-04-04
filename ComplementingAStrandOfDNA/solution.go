package complementingastrandofdna

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// Complement
func Complement(DNA string) (string, error) {
	if len(DNA) <= 0 {
		return "", definitions.ErrEmptyString
	}

	DNA = strings.ToUpper(DNA)
	reversed := reverse(DNA)

	complement, err := complement(reversed)
	if err != nil {
		return "", err
	}

	return complement, nil
}

func reverse(sentence string) string {
	runes := []rune(sentence)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func complement(sentence string) (string, error) {
	complements := map[rune]rune{
		rune('A'): rune('T'),
		rune('T'): rune('A'),
		rune('C'): rune('G'),
		rune('G'): rune('H'),
	}

	runes := []rune(sentence)

	for _, rune := range runes {
		if complement, ok := complements[rune]; !ok {
			return "", definitions.ErrInvalidDNAString
		}
		rune = complements[rune]

	}

	return string(runes), nil
}
