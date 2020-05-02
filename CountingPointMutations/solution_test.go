package countingpointmutations

import (
	"testing"

	definitions "github.com/el10savio/goRosalind/Definitions"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Description      string
	First            string
	Second           string
	ExpectedDistance int
	ExpectedError    error
}

var testCases = []testCase{
	testCase{
		Description:      "Empty Strings",
		First:            "",
		Second:           "",
		ExpectedDistance: 0,
		ExpectedError:    definitions.ErrEmptyList,
	},
	testCase{
		Description:      "Unequal Strings",
		First:            "GAGCCTACTAACGGGAT",
		Second:           "CATCGTAATGACGGCCTGGC",
		ExpectedDistance: 0,
		ExpectedError:    definitions.ErrUnequalLength,
	},
	testCase{
		Description:      "Equal Strings",
		First:            "GAGCCTACTAACGGGAT",
		Second:           "GAGCCTACTAACGGGAT",
		ExpectedDistance: 0,
		ExpectedError:    nil,
	},
	testCase{
		Description:      "Basic Test",
		First:            "GAGCCTACTAACGGGAT",
		Second:           "CATCGTAATGACGGCCT",
		ExpectedDistance: 7,
		ExpectedError:    nil,
	},
	testCase{
		Description:      "Case Insensitive Test",
		First:            "GAGCCTACTAACGGGAT",
		Second:           "catcgtaatgacggcct",
		ExpectedDistance: 7,
		ExpectedError:    nil,
	},
}

func TestHammingDistance(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			DistanceActual, ErrorActual := Compute(tC.First, tC.Second)
			assert.Equal(t, tC.ExpectedError, ErrorActual)
			assert.Equal(t, tC.ExpectedDistance, DistanceActual)
		})
	}
}
