package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

var pageRequestTestCases = []coretestcases.CaseV1{
	// IsPageSizeEmpty
	{Title: "IsPageSizeEmpty - nil receiver returns true", ArrangeInput: args.Map{"method": "IsPageSizeEmpty", "req": (*coreapi.PageRequest)(nil)}, ExpectedInput: []string{"true"}},
	{Title: "IsPageSizeEmpty - zero returns true", ArrangeInput: args.Map{"method": "IsPageSizeEmpty", "req": &coreapi.PageRequest{PageSize: 0, PageIndex: 1}}, ExpectedInput: []string{"true"}},
	{Title: "IsPageSizeEmpty - negative returns true", ArrangeInput: args.Map{"method": "IsPageSizeEmpty", "req": &coreapi.PageRequest{PageSize: -1}}, ExpectedInput: []string{"true"}},
	{Title: "IsPageSizeEmpty - positive returns false", ArrangeInput: args.Map{"method": "IsPageSizeEmpty", "req": &coreapi.PageRequest{PageSize: 10}}, ExpectedInput: []string{"false"}},
	// IsPageIndexEmpty
	{Title: "IsPageIndexEmpty - nil receiver returns true", ArrangeInput: args.Map{"method": "IsPageIndexEmpty", "req": (*coreapi.PageRequest)(nil)}, ExpectedInput: []string{"true"}},
	{Title: "IsPageIndexEmpty - zero returns true", ArrangeInput: args.Map{"method": "IsPageIndexEmpty", "req": &coreapi.PageRequest{PageIndex: 0, PageSize: 10}}, ExpectedInput: []string{"true"}},
	{Title: "IsPageIndexEmpty - positive returns false", ArrangeInput: args.Map{"method": "IsPageIndexEmpty", "req": &coreapi.PageRequest{PageIndex: 2}}, ExpectedInput: []string{"false"}},
	// HasPageSize
	{Title: "HasPageSize - nil receiver returns false", ArrangeInput: args.Map{"method": "HasPageSize", "req": (*coreapi.PageRequest)(nil)}, ExpectedInput: []string{"false"}},
	{Title: "HasPageSize - positive returns true", ArrangeInput: args.Map{"method": "HasPageSize", "req": &coreapi.PageRequest{PageSize: 25}}, ExpectedInput: []string{"true"}},
	// HasPageIndex
	{Title: "HasPageIndex - nil receiver returns false", ArrangeInput: args.Map{"method": "HasPageIndex", "req": (*coreapi.PageRequest)(nil)}, ExpectedInput: []string{"false"}},
	{Title: "HasPageIndex - positive returns true", ArrangeInput: args.Map{"method": "HasPageIndex", "req": &coreapi.PageRequest{PageIndex: 3}}, ExpectedInput: []string{"true"}},
	// Clone
	{Title: "Clone - nil receiver returns nil", ArrangeInput: args.Map{"method": "Clone", "req": (*coreapi.PageRequest)(nil)}, ExpectedInput: []string{"true"}},
	{Title: "Clone - copies all fields", ArrangeInput: args.Map{"method": "CloneFields", "req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5}}, ExpectedInput: []string{"20", "5"}},
	{Title: "Clone - independence from original", ArrangeInput: args.Map{"method": "CloneIndependence", "req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5}}, ExpectedInput: []string{"20", "5"}},
}

func Test_PageRequest_Verification(t *testing.T) {
	for caseIndex, tc := range pageRequestTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method := input["method"].(string)
		req := input["req"].(*coreapi.PageRequest)

		var actLines []string

		// Act
		switch method {
		case "IsPageSizeEmpty":
			actLines = []string{fmt.Sprintf("%v", req.IsPageSizeEmpty())}
		case "IsPageIndexEmpty":
			actLines = []string{fmt.Sprintf("%v", req.IsPageIndexEmpty())}
		case "HasPageSize":
			actLines = []string{fmt.Sprintf("%v", req.HasPageSize())}
		case "HasPageIndex":
			actLines = []string{fmt.Sprintf("%v", req.HasPageIndex())}
		case "Clone":
			actLines = []string{fmt.Sprintf("%v", req.Clone() == nil)}
		case "CloneFields":
			clone := req.Clone()
			actLines = []string{fmt.Sprintf("%v", clone.PageSize), fmt.Sprintf("%v", clone.PageIndex)}
		case "CloneIndependence":
			clone := req.Clone()
			clone.PageSize = 99
			clone.PageIndex = 99
			actLines = []string{fmt.Sprintf("%v", req.PageSize), fmt.Sprintf("%v", req.PageIndex)}
		}

		expectedLines := tc.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, tc.Title, actLines, expectedLines)
	}
}
