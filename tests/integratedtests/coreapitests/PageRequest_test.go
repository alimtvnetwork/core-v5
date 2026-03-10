package coreapitests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coredata/coreapi"
	"github.com/alimtvnetwork/core/coretests/args"
)

func Test_PageRequest_IsPageSizeEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageSizeEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.IsPageSizeEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_IsPageIndexEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageIndexEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.IsPageIndexEmpty())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_HasPageSize(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageSizeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.HasPageSize())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_HasPageIndex(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.HasPageIndex())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_Clone_Nil(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneNilTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		result := fmt.Sprintf("%v", req.Clone() == nil)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_PageRequest_Clone_Fields(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneFieldsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		clone := req.Clone()
		actual := args.Map{
			"pageSize":  fmt.Sprintf("%v", clone.PageSize),
			"pageIndex": fmt.Sprintf("%v", clone.PageIndex),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_PageRequest_Clone_Independence(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneIndependenceTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		clone := req.Clone()
		clone.PageSize = 99
		clone.PageIndex = 99
		actual := args.Map{
			"pageSize":  fmt.Sprintf("%v", req.PageSize),
			"pageIndex": fmt.Sprintf("%v", req.PageIndex),
		}

		// Assert
		tc.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
