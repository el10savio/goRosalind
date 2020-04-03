package transcribingdnaintorna

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given: A DNA string t having length at 
// most 1000 nt.
//
// Return: The transcribed RNA string of t.

// Transcribe converts a DNA string to RNA
func Transcribe(DNA string) (string, error) {
	if len(DNA) <= 0 {
		return "", definitions.ErrEmptyString
	}

	DNA = strings.ToUpper(DNA)
	RNA := strings.ReplaceAll(DNA, "T", "U")

	return RNA, nil
}
