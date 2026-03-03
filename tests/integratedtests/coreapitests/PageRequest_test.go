package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

func Test_PageRequest_IsPageSizeEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageSizeEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		actLines := []string{fmt.Sprintf("%v", req.IsPageSizeEmpty())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_PageRequest_IsPageIndexEmpty(t *testing.T) {
	for caseIndex, tc := range pageRequestIsPageIndexEmptyTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		actLines := []string{fmt.Sprintf("%v", req.IsPageIndexEmpty())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_PageRequest_HasPageSize(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageSizeTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		actLines := []string{fmt.Sprintf("%v", req.HasPageSize())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_PageRequest_HasPageIndex(t *testing.T) {
	for caseIndex, tc := range pageRequestHasPageIndexTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		actLines := []string{fmt.Sprintf("%v", req.HasPageIndex())}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_PageRequest_Clone_Nil(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneNilTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		actLines := []string{fmt.Sprintf("%v", req.Clone() == nil)}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}

func Test_PageRequest_Clone_Fields(t *testing.T) {
	for caseIndex, tc := range pageRequestCloneFieldsTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		req := input["req"].(*coreapi.PageRequest)

		// Act
		clone := req.Clone()
		actLines := []string{
			fmt.Sprintf("%v", clone.PageSize),
			fmt.Sprintf("%v", clone.PageIndex),
		}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
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

		actLines := []string{
			fmt.Sprintf("%v", req.PageSize),
			fmt.Sprintf("%v", req.PageIndex),
		}
		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
