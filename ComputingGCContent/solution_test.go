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
	ExpectedGCContent float32
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
	testCase{
		Description: "Empty DNA",
		FASTAValues: []FASTA{
			FASTA{
				Label: ">Rosalind_6404",
				DNA:   "",
			},
		},
		ExpectedLabel:     "",
		ExpectedGCContent: 0.0,
		ExpectedError:     definitions.ErrEmptyString,
	},
	testCase{
		Description: "Empty Label",
		FASTAValues: []FASTA{
			FASTA{
				Label: "",
				DNA:   "CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCCTCCCACTAATAATTCTGAGG",
			},
		},
		ExpectedLabel:     "",
		ExpectedGCContent: 0.0,
		ExpectedError:     definitions.ErrEmptyLabel,
	},
	testCase{
		Description: "Basic List",
		FASTAValues: []FASTA{
			FASTA{
				Label: ">Rosalind_6404",
				DNA:   "CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCCTCCCACTAATAATTCTGAGG",
			},
			FASTA{
				Label: ">Rosalind_5959",
				DNA:   "CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCTATATCCATTTGTCAGCAGACACGC",
			},
			FASTA{
				Label: ">Rosalind_5951",
				DNA:   "AGCTATAG",
			},
			FASTA{
				Label: ">Rosalind_0808",
				DNA:   "CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGACTGGGAACCTGCGGGCAGTAGGTGGAAT",
			},
		},
		ExpectedLabel:     "Rosalind_0808",
		ExpectedGCContent: 60.919540,
		ExpectedError:     nil,
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
