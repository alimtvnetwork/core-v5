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
		Name: "Standard key=value split",
		ArrangeInput: args.Map{
			"input": "key=value",
			"sep":   "=",
		},
		WantLines: []string{
			"key",
			"value",
			"true",
			"",
		},
	},
	{
		Name: "No separator found produces invalid pair",
		ArrangeInput: args.Map{
			"input": "noseparator",
			"sep":   "=",
		},
		WantLines: []string{
			"noseparator",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Empty input produces invalid pair",
		ArrangeInput: args.Map{
			"input": "",
			"sep":   "=",
		},
		WantLines: []string{
			"",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Multiple separators takes first two parts only",
		ArrangeInput: args.Map{
			"input": "a=b=c=d",
			"sep":   "=",
		},
		WantLines: []string{
			"a",
			"b=c=d",
			"true",
			"",
		},
	},
	{
		Name: "Separator at start produces empty left",
		ArrangeInput: args.Map{
			"input": "=value",
			"sep":   "=",
		},
		WantLines: []string{
			"",
			"value",
			"true",
			"",
		},
	},
	{
		Name: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key=",
			"sep":   "=",
		},
		WantLines: []string{
			"key",
			"",
			"true",
			"",
		},
	},
	{
		Name: "Multi-char separator",
		ArrangeInput: args.Map{
			"input": "hello::world",
			"sep":   "::",
		},
		WantLines: []string{
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
		Name: "Trims whitespace from both parts",
		ArrangeInput: args.Map{
			"input": "  key  =  value  ",
			"sep":   "=",
		},
		WantLines: []string{
			"key",
			"value",
			"true",
			"",
		},
	},
	{
		Name: "No separator trims single part",
		ArrangeInput: args.Map{
			"input": "  onlypart  ",
			"sep":   "=",
		},
		WantLines: []string{
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
		Name: "Splits at first separator only, right gets remainder",
		ArrangeInput: args.Map{
			"input": "a:b:c:d",
			"sep":   ":",
		},
		WantLines: []string{
			"a",
			"b:c:d",
			"true",
			"",
		},
	},
	{
		Name: "No separator produces invalid pair",
		ArrangeInput: args.Map{
			"input": "nosep",
			"sep":   ":",
		},
		WantLines: []string{
			"nosep",
			"",
			"false",
			"separator not found",
		},
	},
	{
		Name: "Separator at end produces empty right",
		ArrangeInput: args.Map{
			"input": "key:",
			"sep":   ":",
		},
		WantLines: []string{
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
		Name: "Splits at first separator with trimming",
		ArrangeInput: args.Map{
			"input": "  a  :  b : c : d  ",
			"sep":   ":",
		},
		WantLines: []string{
			"a",
			"b : c : d",
			"true",
			"",
		},
	},
	{
		Name: "No separator trims and marks invalid",
		ArrangeInput: args.Map{
			"input": "  nosep  ",
			"sep":   ":",
		},
		WantLines: []string{
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
		Name: "Two-element slice produces valid pair",
		ArrangeInput: args.Map{
			"parts": []string{"left", "right"},
		},
		WantLines: []string{
			"left",
			"right",
			"true",
			"",
		},
	},
	{
		Name: "Single-element slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{"only"},
		},
		WantLines: []string{
			"only",
			"",
			"false",
			"only one part found",
		},
	},
	{
		Name: "Empty slice produces invalid pair",
		ArrangeInput: args.Map{
			"parts": []string{},
		},
		WantLines: []string{
			"",
			"",
			"false",
			"empty input",
		},
	},
	{
		Name: "Three-element slice uses first and last",
		ArrangeInput: args.Map{
			"parts": []string{"first", "middle", "last"},
		},
		WantLines: []string{
			"first",
			"last",
			"true",
			"",
		},
	},
}
