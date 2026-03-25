package isanytests

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/isany"
)

func Test_IsAny_Defined_Null_Verification(t *testing.T) {
	for caseIndex, testCase := range isAnyDefinedNullTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		var item any

		useError, _ := input.Get("useError")
		if useError == true {
			item = errors.New("test error")
		} else {
			item = input.GetDirectLower("input")
		}

		// Act
		actual := args.Map{
			"isDefined": fmt.Sprintf("%v", isany.Defined(item)),
			"isNull":    fmt.Sprintf("%v", isany.Null(item)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsAny_Both_Verification(t *testing.T) {
	for caseIndex, testCase := range isAnyBothTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		first := input.GetDirectLower("first")
		second := input.GetDirectLower("second")

		// Act
		actual := args.Map{
			"definedBoth": fmt.Sprintf("%v", isany.DefinedBoth(first, second)),
			"nullBoth":    fmt.Sprintf("%v", isany.NullBoth(first, second)),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
