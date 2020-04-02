package transcribingdnaintorna

import (
	"testing"

	definitions "github.com/el10savio/goRosalind/Definitions"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Description string
	DNA         string
	RNA         string
	Error       error
}

var testCases = []testCase{
	testCase{
		Description: "Empty DNA string",
		DNA:         "",
		RNA:         "",
		Error:       definitions.ErrEmptyString,
	},
	testCase{
		Description: "Basic DNA string",
		DNA:         "GATGGAACTTGACTACGTAAATT",
		RNA:         "GAUGGAACUUGACUACGUAAAUU",
		Error:       nil,
	},
	testCase{
		Description: "Lowercase DNA string",
		DNA:         "gatggaacttgactacgtaaatt",
		RNA:         "GAUGGAACUUGACUACGUAAAUU",
		Error:       nil,
	},
	testCase{
		Description: "Large DNA string",
		DNA:         "TGGGTTCCAGTGACCATAGACCCCCAGACTTTGTCAGACACTTTAAACGTATCCGCACACCTTACCAGTCTGACAAACATGATGGGCCCCCTTTTTAACCGCCGAGGAGCAAACGCCGGCCTTCCATGACAGACATGTCGTTACCGTGTTTACGCTCTAACGGCTCCGTGATGGACTACCCATAGTTAGCGCGAAAGATCGAACTGCGTTATTAGAACCACATCTGGCAGCTGTGCTATTAAGATGATAAACGGATACCAAAAGAACGCGTACGCCGCGGGGCGTGTGCTCAAATAGTTCTCGAACATTCAGAAAACCAATTCCTGGCGTGCGAGTACGACCATACCCGGAGGACTGGCAGGACCTGGCGTACTCACAGAGTAGGTGCGCTACTTGCCCTCACACGTGCCTGCAGCCCCTTCGTCCTATTTCACGTACGAGGCAGGGTTGCAGGTTTGAACATTATTGTGCTCAAGCAGACCCTCGGACTGTCGGCGCCAGACGAATCCTTGAACTTTGCAAAGTTGAACGGGTGGCAGATCGTGTTTCGGAACGGGCAACGAGCGCACCGTTCTAAAGGTCCCCCTTGACCATACACGACGCTAATTGGCAGTAATCACCGACGCGAAGGTCTAGGGCCAAGAAGCTCGCGAACGATGGGAACAGTAATCGCTATCGTTCCAGCGTCCGCGGGGGGGAATGTTGAACTGGACATTGTAACGTAAGCGAGCGGCGTCTGCCTGTAGGCAGGTCTGCTCCCTAGAGACAAGAGTTTTACGGCAAATTTAAAGGTGTTCTGGACCGGTAATACCCTGAATCGGGCATTCATGGCAGCTATCTCGGCGGTGGTGGACGAGATGCACAGCAACTTGTTCGTTCGCCCGGTATCAGCGTTACCCTACAGGTGGCCTGTAGGTAAGACATGCTGGATAAAACGGTAGATCCAATTGACGCTTGT",
		RNA:         "UGGGUUCCAGUGACCAUAGACCCCCAGACUUUGUCAGACACUUUAAACGUAUCCGCACACCUUACCAGUCUGACAAACAUGAUGGGCCCCCUUUUUAACCGCCGAGGAGCAAACGCCGGCCUUCCAUGACAGACAUGUCGUUACCGUGUUUACGCUCUAACGGCUCCGUGAUGGACUACCCAUAGUUAGCGCGAAAGAUCGAACUGCGUUAUUAGAACCACAUCUGGCAGCUGUGCUAUUAAGAUGAUAAACGGAUACCAAAAGAACGCGUACGCCGCGGGGCGUGUGCUCAAAUAGUUCUCGAACAUUCAGAAAACCAAUUCCUGGCGUGCGAGUACGACCAUACCCGGAGGACUGGCAGGACCUGGCGUACUCACAGAGUAGGUGCGCUACUUGCCCUCACACGUGCCUGCAGCCCCUUCGUCCUAUUUCACGUACGAGGCAGGGUUGCAGGUUUGAACAUUAUUGUGCUCAAGCAGACCCUCGGACUGUCGGCGCCAGACGAAUCCUUGAACUUUGCAAAGUUGAACGGGUGGCAGAUCGUGUUUCGGAACGGGCAACGAGCGCACCGUUCUAAAGGUCCCCCUUGACCAUACACGACGCUAAUUGGCAGUAAUCACCGACGCGAAGGUCUAGGGCCAAGAAGCUCGCGAACGAUGGGAACAGUAAUCGCUAUCGUUCCAGCGUCCGCGGGGGGGAAUGUUGAACUGGACAUUGUAACGUAAGCGAGCGGCGUCUGCCUGUAGGCAGGUCUGCUCCCUAGAGACAAGAGUUUUACGGCAAAUUUAAAGGUGUUCUGGACCGGUAAUACCCUGAAUCGGGCAUUCAUGGCAGCUAUCUCGGCGGUGGUGGACGAGAUGCACAGCAACUUGUUCGUUCGCCCGGUAUCAGCGUUACCCUACAGGUGGCCUGUAGGUAAGACAUGCUGGAUAAAACGGUAGAUCCAAUUGACGCUUGU",
		Error:       nil,
	},
}

func TestTranscribe(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			RNAActual, ErrorActual := Transcribe(tC.DNA)
			assert.Equal(t, tC.Error, ErrorActual)
			assert.Equal(t, tC.RNA, RNAActual)
		})
	}
}
