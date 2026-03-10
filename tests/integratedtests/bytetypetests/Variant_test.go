package bytetypetests

import (
	"testing"

	"github.com/alimtvnetwork/core/bytetype"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_Variant_Verification(t *testing.T) {
	for caseIndex, testCase := range newVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsInt("input")

		// Act
		v := bytetype.New(byte(inputVal))

		actual := args.Map{
			"value":     v.ValueInt(),
			"isZero":    v.IsZero(),
			"isInvalid": v.IsInvalid(),
			"isValid":   v.IsValid(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
