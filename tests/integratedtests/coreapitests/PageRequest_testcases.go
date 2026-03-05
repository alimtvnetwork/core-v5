package coreapitests

import (
	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IsPageSizeEmpty test cases
// =============================================================================

var pageRequestIsPageSizeEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPageSizeEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - zero returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 0, PageIndex: 1},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - negative returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: -1},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageSizeEmpty - positive returns false",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 10},
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// IsPageIndexEmpty test cases
// =============================================================================

var pageRequestIsPageIndexEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsPageIndexEmpty - nil receiver returns true",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageIndexEmpty - zero returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 0, PageSize: 10},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsPageIndexEmpty - positive returns false",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 2},
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// HasPageSize test cases
// =============================================================================

var pageRequestHasPageSizeTestCases = []coretestcases.CaseV1{
	{
		Title: "HasPageSize - nil receiver returns false",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "false",
	},
	{
		Title: "HasPageSize - positive returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 25},
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// HasPageIndex test cases
// =============================================================================

var pageRequestHasPageIndexTestCases = []coretestcases.CaseV1{
	{
		Title: "HasPageIndex - nil receiver returns false",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "false",
	},
	{
		Title: "HasPageIndex - positive returns true",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageIndex: 3},
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// Clone test cases
// =============================================================================

var pageRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - nil receiver returns nil",
		ArrangeInput: args.Map{
			"req": (*coreapi.PageRequest)(nil),
		},
		ExpectedInput: "true",
	},
}

var pageRequestCloneFieldsTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - copies all fields",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5},
		},
		ExpectedInput: args.Two[string, string]{
			First:  "20", // pageSize
			Second: "5",  // pageIndex
		},
	},
}

var pageRequestCloneIndependenceTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone - independence from original",
		ArrangeInput: args.Map{
			"req": &coreapi.PageRequest{PageSize: 20, PageIndex: 5},
		},
		ExpectedInput: args.Two[string, string]{
			First:  "20", // pageSize
			Second: "5",  // pageIndex
		},
	},
}
