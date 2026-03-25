package coredatatests

import (
	"sort"
	"testing"

	"github.com/alimtvnetwork/core/coredata"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_IntegersDsc_Len(t *testing.T) {
	for caseIndex, tc := range integersDscLenTestCases {
		// Arrange
		var integers coredata.IntegersDsc
		if tc.ArrangeInput != nil {
			input := tc.ArrangeInput.(args.Map)
			if vals, ok := input["values"]; ok {
				integers = coredata.IntegersDsc(vals.([]int))
			}
		}

		// Act
		actual := args.Map{
			"length": integers.Len(),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IntegersDsc_Sort(t *testing.T) {
	for caseIndex, tc := range integersDscSortTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		src := input["values"].([]int)
		integers := make(coredata.IntegersDsc, len(src))
		copy(integers, src)

		// Act
		sort.Sort(integers)

		actual := args.Map{
			"first": integers[0],
			"last":  integers[len(integers)-1],
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
