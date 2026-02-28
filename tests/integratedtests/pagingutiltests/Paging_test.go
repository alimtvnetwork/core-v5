package pagingutiltests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/pagingutil"
)

func Test_GetPagesSize_Verification(t *testing.T) {
	for caseIndex, testCase := range getPagesSizeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		eachPageSize, _ := input.GetAsInt("eachPageSize")
		totalLength, _ := input.GetAsInt("totalLength")

		// Act
		result := pagingutil.GetPagesSize(eachPageSize, totalLength)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_GetPagingInfo_Verification(t *testing.T) {
	for caseIndex, testCase := range getPagingInfoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		length, _ := input.GetAsInt("length")
		pageIndex, _ := input.GetAsInt("pageIndex")
		eachPageSize, _ := input.GetAsInt("eachPageSize")

		// Act
		info := pagingutil.GetPagingInfo(pagingutil.PagingRequest{
			Length:       length,
			PageIndex:    pageIndex,
			EachPageSize: eachPageSize,
		})

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", info.PageIndex),
			fmt.Sprintf("%v", info.SkipItems),
			fmt.Sprintf("%v", info.EndingLength),
			fmt.Sprintf("%v", info.IsPagingPossible),
		)
	}
}
