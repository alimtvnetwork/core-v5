package bytetypetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/bytetype"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_Variant_Verification(t *testing.T) {
	for caseIndex, testCase := range newVariantTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputVal, _ := input.GetAsInt("input")

		// Act
		v := bytetype.New(byte(inputVal))
		valueStr := fmt.Sprintf("%v", v.ValueInt())
		isZero := fmt.Sprintf("%v", v.IsZero())
		isInvalid := fmt.Sprintf("%v", v.IsInvalid())
		isValid := fmt.Sprintf("%v", v.IsValid())

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			valueStr,
			isZero,
			isInvalid,
			isValid,
		)
	}
}
