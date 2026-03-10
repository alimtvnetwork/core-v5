package errcoretests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// SliceToError test cases
// =============================================================================

var sliceToErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "SliceToError returns nil -- nil slice",
		ArrangeInput: args.Map{
			"when":  "given nil slice",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToError returns nil -- empty slice",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToError returns error with message -- single item ['error one']",
		ArrangeInput: args.Map{
			"when":    "given single error string",
			"input":   []string{"error one"},
			"contain": "error one",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "SliceToError returns joined error -- three items ['err1','err2','err3']",
		ArrangeInput: args.Map{
			"when":    "given three error strings",
			"input":   []string{"err1", "err2", "err3"},
			"contain": "err1",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
}

// =============================================================================
// SliceToErrorPtr test cases
// =============================================================================

var sliceToErrorPtrTestCases = []coretestcases.CaseV1{
	{
		Title: "SliceToErrorPtr returns nil -- nil slice",
		ArrangeInput: args.Map{
			"when":  "given nil slice",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToErrorPtr returns nil -- empty slice",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToErrorPtr returns error -- single item ['one']",
		ArrangeInput: args.Map{
			"when":    "given single error string",
			"input":   []string{"one"},
			"contain": "one",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
}

// =============================================================================
// MergeErrors test cases
// =============================================================================

var mergeErrorsTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors returns nil -- no arguments",
		ArrangeInput: args.Map{
			"when":   "given no arguments",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "false",
	},
	{
		Title: "MergeErrors returns nil -- three nil errors",
		ArrangeInput: args.Map{
			"when":   "given three nil errors",
			"errors": []string{},
			"nils":   3,
		},
		ExpectedInput: "false",
	},
	{
		Title: "MergeErrors returns single error -- one error 'fail'",
		ArrangeInput: args.Map{
			"when":    "given single error",
			"errors":  []string{"fail"},
			"nils":    0,
			"contain": "fail",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "MergeErrors returns joined error -- three errors ['a','b','c']",
		ArrangeInput: args.Map{
			"when":    "given three errors",
			"errors":  []string{"a", "b", "c"},
			"nils":    0,
			"contain": "a",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "MergeErrors returns error skipping nils -- 2 errors interleaved with 3 nils",
		ArrangeInput: args.Map{
			"when":    "given errors interleaved with nils",
			"errors":  []string{"real", "also real"},
			"nils":    3,
			"contain": "real",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "MergeErrors returns nil -- single nil",
		ArrangeInput: args.Map{
			"when":   "given single nil",
			"errors": []string{},
			"nils":   1,
		},
		ExpectedInput: "false",
	},
}

// =============================================================================
// SliceErrorsToStrings test cases
// =============================================================================

var sliceErrorsToStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "SliceErrorsToStrings returns empty -- no arguments",
		ArrangeInput: args.Map{
			"when":   "given no arguments",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "0",
	},
	{
		Title: "SliceErrorsToStrings returns empty -- two nil errors",
		ArrangeInput: args.Map{
			"when":   "given two nil errors",
			"errors": []string{},
			"nils":   2,
		},
		ExpectedInput: "0",
	},
	{
		Title: "SliceErrorsToStrings returns non-nil only -- 2 errors mixed with 1 nil",
		ArrangeInput: args.Map{
			"when":   "given errors mixed with nil",
			"errors": []string{"a", "b"},
			"nils":   1,
		},
		ExpectedInput: args.Map{
			"count":  "2",
			"first":  "a",
			"second": "b",
		},
	},
}

// =============================================================================
// MergeErrorsToString test cases
// =============================================================================

var mergeErrorsToStringTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrorsToString returns empty -- no errors",
		ArrangeInput: args.Map{
			"when":   "given no errors",
			"joiner": ", ",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "",
	},
	{
		Title: "MergeErrorsToString returns joined string -- two errors with pipe joiner",
		ArrangeInput: args.Map{
			"when":   "given two errors with pipe joiner",
			"joiner": " | ",
			"errors": []string{"x", "y"},
			"nils":   0,
		},
		ExpectedInput: "x | y",
	},
	{
		Title: "MergeErrorsToString returns non-nil only -- one error with 2 nils",
		ArrangeInput: args.Map{
			"when":   "given one error with nils",
			"joiner": ", ",
			"errors": []string{"only"},
			"nils":   2,
		},
		ExpectedInput: "only",
	},
}

// =============================================================================
// MergeErrorsToStringDefault test cases
// =============================================================================

var mergeErrorsToStringDefaultTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrorsToStringDefault returns empty -- no errors",
		ArrangeInput: args.Map{
			"when":   "given no errors",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "",
	},
	{
		Title: "MergeErrorsToStringDefault returns space-joined -- two errors ['a','b']",
		ArrangeInput: args.Map{
			"when":   "given two errors",
			"errors": []string{"a", "b"},
			"nils":   0,
		},
		ExpectedInput: "a b",
	},
}
