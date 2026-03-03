package coregenerictests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// PairFromSplit
// ==========================================

var pairFromSplitTestCases = []coretestcases.CaseV1{
	{
		Title: "Standard key=value split",
		ArrangeInput: args.Map{
			"input": "key=value",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"key",
			"value",
			"true",
			"",
		},
	},
	{
		Title: "No separator found produces invalid pair",
		ArrangeInput: args.Map{
			"input": "noseparator",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"noseparator",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Empty input produces invalid pair",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Multiple separators takes first two parts only",
		ArrangeInput: args.Map{
			"input": "a=b=c=d",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"a",
			"b=c=d",
			"true",
			"",
		},
	},
	{
		Title: "Separator at start produces empty left",
		ArrangeInput: args.Map{
			"input": "=value",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"",
			"value",
			"true",
			"",
		},
	},
	{
		Title: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key=",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"key",
			"",
			"true",
			"",
		},
	},
	{
		Title: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "hello::world",
			"sep":   "::",
		},
		ExpectedInput: []string{
			"hello",
			"world",
			"true",
			"",
		},
	},
}

// ==========================================
// PairFromSplitTrimmed
// ==========================================

var pairFromSplitTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Trims whitespace from both parts",
		ArrangeInput: args.Map{
			"input": "  key  =  value  ",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"key",
			"value",
			"true",
			"",
		},
	},
	{
		Title: "No separator trims single part",
		ArrangeInput: args.Map{
			"input": "  onlypart  ",
			"sep":   "=",
		},
		ExpectedInput: []string{
			"onlypart",
			"",
			"false",
			"only one part found",
		},
	},
}

// ==========================================
// PairFromSplitFull
// ==========================================

var pairFromSplitFullTestCases = []coretestcases.CaseV1{
	{
		Title: "Splits at first separator only, right gets remainder",
		ArrangeInput: args.Map{
			"input": "a:b:c:d",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"a",
			"b:c:d",
			"true",
			"",
		},
	},
	{
		Title: "No separator produces invalid pair",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"nosep",
			"",
			"false",
			"separator not found",
		},
	},
	{
		Title: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key:",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"key",
			"",
			"true",
			"",
		},
	},
}

// ==========================================
// PairFromSplitFullTrimmed
// ==========================================

var pairFromSplitFullTrimmedTestCases = []coretestcases.CaseV1{
	{
		Title: "Splits at first separator with trimming",
		ArrangeInput: args.Map{
			"input": "  a  :  b : c : d  ",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"a",
			"b : c : d",
			"true",
			"",
		},
	},
	{
		Title: "No separator trims and marks invalid",
		ArrangeInput: args.Map{
			"input": "  nosep  ",
			"sep":   ":",
		},
		ExpectedInput: []string{
			"nosep",
			"",
			"false",
			"separator not found",
		},
	},
}

// ==========================================
// PairFromSlice
// ==========================================

var pairFromSliceTestCases = []coretestcases.CaseV1{
	{
		Title: "Two-element slice produces valid pair",
		ArrangeInput: args.Map{
			"parts": []string{"left", "right"},
		},
		ExpectedInput: []string{
			"left",
			"right",
			"true",
			"",
		},
	},
	{
		Title: "Single-element slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		ExpectedInput: []string{
			"only",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Title: "Empty slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		ExpectedInput: []string{
			"",
			"",
			"false",
			"empty input",
		},
	},
	{
		Title: "Three-element slice uses first and last",
		ArrangeInput: args.Map{
			"parts": []string{"first", "middle", "last"},
		},
		ExpectedInput: []string{
			"first",
			"last",
			"true",
			"",
		},
	},
}
