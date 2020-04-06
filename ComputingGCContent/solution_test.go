package computinggccontent

import (
	"testing"
)

type testCase struct {
	Description       string
	DNA               string
	ReverseComplement string
	Error             error
}

var testCases = []testCase{
	testCase{},
}

func TestComplement(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			// ReverseComplementActual, ErrorActual := Complement(tC.DNA)
			// assert.Equal(t, tC.Error, ErrorActual)
			// assert.Equal(t, tC.ReverseComplement, ReverseComplementActual)
		})
	}
}
