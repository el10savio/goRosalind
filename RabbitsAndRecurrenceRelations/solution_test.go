package rabbitsandrecurrencerelations

import (
	"testing"

	definitions "github.com/el10savio/goRosalind/Definitions"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Description string
	Months      int
	Litter      int
	Total       int
	Error       error
}

var testCases = []testCase{
	testCase{
		Description: "0 Total() Parameters",
		Months:      0,
		Litter:      0,
		Total:       0,
		Error:       definitions.ErrNotPositiveValue,
	},
	testCase{
		Description: "Negative Total() Parameters",
		Months:      0,
		Litter:      0,
		Total:       0,
		Error:       definitions.ErrNotPositiveValue,
	},
	testCase{
		Description: "Basic Total() Parameters",
		Months:      5,
		Litter:      3,
		Total:       19,
		Error:       nil,
	},
}

func TestTotal(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.Description, func(t *testing.T) {
			TotalActual, ErrorActual := Total(tC.Months, tC.Litter)
			assert.Equal(t, tC.Error, ErrorActual)
			assert.Equal(t, tC.Total, TotalActual)
		})
	}
}
