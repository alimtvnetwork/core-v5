package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_Collection_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range collectionGetPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		items, _ := input.Get("items")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		slice := items.([]int)
		collection := coredynamic.New.Collection.Int.Cap(len(slice))
		for _, v := range slice {
			collection.Add(v)
		}

		// Act
		result := collection.GetPagesSize(eachPageSize)

		// Assert
		errcore.AssertDiffOnMismatch(
			t,
			caseIndex,
			testCase.Title,
			[]string{fmt.Sprintf("%v", result)},
			testCase.ExpectedInput.([]string),
		)
	}
}
