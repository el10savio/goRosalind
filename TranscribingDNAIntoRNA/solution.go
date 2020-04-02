package transcribingdnaintorna

import (
	"strings"

	definitions "github.com/el10savio/goRosalind/Definitions"
)

// Given:
//
// Return:

// Transcribe ...
func Transcribe(DNA string) (string, error) {
	if len(DNA) <= 0 {
		return "", definitions.ErrEmptyString
	}

	DNA = strings.ToUpper(DNA)
	RNA := strings.ReplaceAll(DNA, "T", "U")

	return RNA, nil
}
