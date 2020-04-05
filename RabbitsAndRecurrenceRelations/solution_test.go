package rabbitsandrecurrencerelations

import (
	"testing"
)

type testCase struct {
	Description string
	Error       error
}

var testCases = []testCase{
	testCase{},
}

func TestTranscribe(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			// RNAActual, ErrorActual := Transcribe(tC.DNA)
			// assert.Equal(t, tC.Error, ErrorActual)
			// assert.Equal(t, tC.RNA, RNAActual)
		})
	}
}
