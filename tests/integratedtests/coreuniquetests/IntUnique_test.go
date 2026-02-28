package coreuniquetests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreunique/intunique"
)

func Test_IntUnique_Get_Verification(t *testing.T) {
	for caseIndex, testCase := range intUniqueGetTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil := isNilVal == true

		// Act
		if isNil {
			result := intunique.Get(nil)
			isResultNil := fmt.Sprintf("%v", result == nil)
			testCase.ShouldBeEqual(t, caseIndex, isResultNil)
		} else {
			inputVal, _ := input.Get("input")
			slice := inputVal.([]int)
			clone := make([]int, len(slice))
			copy(clone, slice)

			result := intunique.Get(&clone)
			length := fmt.Sprintf("%v", len(*result))
			testCase.ShouldBeEqual(t, caseIndex, length)
		}
	}
}
