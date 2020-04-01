package transcribingdnaintorna

import (
	"testing"
)

type testCase struct {
	description string
	DNA         string
	RNA         string
}

var testCases = []testCase{
	testCase{},
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {})
	}
}
