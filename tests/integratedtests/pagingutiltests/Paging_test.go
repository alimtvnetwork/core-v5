package pagingutiltests

import (
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
		actual := args.Map{
			"pagesSize": result,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
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
		actual := args.Map{
			"pageIndex":        info.PageIndex,
			"skipItems":        info.SkipItems,
			"endingLength":     info.EndingLength,
			"isPagingPossible": info.IsPagingPossible,
		}
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
