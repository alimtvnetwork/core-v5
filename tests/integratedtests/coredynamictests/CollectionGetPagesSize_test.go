package coredynamictests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coredynamic"
	"github.com/alimtvnetwork/core/coretests/args"
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
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}
