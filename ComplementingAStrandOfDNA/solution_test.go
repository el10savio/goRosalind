package complementingastrandofdna

import (
	"testing"

	definitions "github.com/el10savio/goRosalind/Definitions"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Description       string
	DNA               string
	ReverseComplement string
	Error             error
}

var testCases = []testCase{
	testCase{
		Description:       "Empty DNA string",
		DNA:               "",
		ReverseComplement: "",
		Error:             definitions.ErrEmptyString,
	},
	testCase{
		Description:       "Basic DNA string",
		DNA:               "AAAACCCGGT",
		ReverseComplement: "ACCGGGTTTT",
		Error:             nil,
	},
	testCase{
		Description:       "Lowercase DNA string",
		DNA:               "aaaacccggt",
		ReverseComplement: "ACCGGGTTTT",
		Error:             nil,
	},
}

func TestComplement(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			ReverseComplementActual, ErrorActual := Complement(tC.DNA)
			assert.Equal(t, tC.Error, ErrorActual)
			assert.Equal(t, tC.ReverseComplement, ReverseComplementActual)
		})
	}
}
