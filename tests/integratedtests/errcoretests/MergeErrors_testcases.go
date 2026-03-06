package errcoretests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// SliceToError test cases
// =============================================================================

var sliceToErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "SliceToError - nil slice returns nil",
		ArrangeInput: args.Map{
			"when":  "given nil slice",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToError - empty slice returns nil",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToError - single item returns error with message",
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
		Title: "SliceToError - multiple items joins all",
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
		Title: "SliceToErrorPtr - nil slice returns nil",
		ArrangeInput: args.Map{
			"when":  "given nil slice",
			"isNil": true,
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToErrorPtr - empty slice returns nil",
		ArrangeInput: args.Map{
			"when":  "given empty slice",
			"input": []string{},
		},
		ExpectedInput: "false",
	},
	{
		Title: "SliceToErrorPtr - single item returns error",
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
		Title: "MergeErrors - no args returns nil",
		ArrangeInput: args.Map{
			"when":   "given no arguments",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "false",
	},
	{
		Title: "MergeErrors - all nil returns nil",
		ArrangeInput: args.Map{
			"when":   "given three nil errors",
			"errors": []string{},
			"nils":   3,
		},
		ExpectedInput: "false",
	},
	{
		Title: "MergeErrors - single error returns it",
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
		Title: "MergeErrors - multiple errors joins all",
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
		Title: "MergeErrors - mixed nil and errors skips nil",
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
		Title: "MergeErrors - single nil returns nil",
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
		Title: "SliceErrorsToStrings - no args returns empty",
		ArrangeInput: args.Map{
			"when":   "given no arguments",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "0",
	},
	{
		Title: "SliceErrorsToStrings - all nil returns empty",
		ArrangeInput: args.Map{
			"when":   "given two nil errors",
			"errors": []string{},
			"nils":   2,
		},
		ExpectedInput: "0",
	},
	{
		Title: "SliceErrorsToStrings - mixed returns non-nil only",
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
		Title: "MergeErrorsToString - no errors returns empty",
		ArrangeInput: args.Map{
			"when":   "given no errors",
			"joiner": ", ",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "",
	},
	{
		Title: "MergeErrorsToString - joins with custom joiner",
		ArrangeInput: args.Map{
			"when":   "given two errors with pipe joiner",
			"joiner": " | ",
			"errors": []string{"x", "y"},
			"nils":   0,
		},
		ExpectedInput: "x | y",
	},
	{
		Title: "MergeErrorsToString - skips nil errors",
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
		Title: "MergeErrorsToStringDefault - no errors returns empty",
		ArrangeInput: args.Map{
			"when":   "given no errors",
			"errors": []string{},
			"nils":   0,
		},
		ExpectedInput: "",
	},
	{
		Title: "MergeErrorsToStringDefault - joins with space",
		ArrangeInput: args.Map{
			"when":   "given two errors",
			"errors": []string{"a", "b"},
			"nils":   0,
		},
		ExpectedInput: "a b",
	},
}
