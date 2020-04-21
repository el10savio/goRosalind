package computinggccontent

import (
	"testing"

	definitions "github.com/el10savio/goRosalind/Definitions"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Description       string
	FASTAValues       []FASTA
	ExpectedLabel     string
	ExpectedGCContent float64
	ExpectedError     error
}

var testCases = []testCase{
	testCase{
		Description:       "Empty List",
		FASTAValues:       []FASTA{},
		ExpectedLabel:     "",
		ExpectedGCContent: 0.0,
		ExpectedError:     definitions.ErrEmptyList,
	},
}

func TestComplement(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			LabelActual, GCContentActual, ErrorActual := Compute(tC.FASTAValues)
			assert.Equal(t, tC.ExpectedError, ErrorActual)
			assert.Equal(t, tC.ExpectedLabel, LabelActual)
			assert.Equal(t, tC.ExpectedGCContent, GCContentActual)
		})
	}
}
