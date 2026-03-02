package converterstests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// IntegersWithDefaults
// =============================================================================

var integersWithDefaultsTestCases = []coretestcases.CaseV1{
	{
		Title: "IntegersWithDefaults converts all valid integers",
		ArrangeInput: args.Map{
			"when":       "given all valid integer strings",
			"input":      []string{"1", "2", "3"},
			"defaultInt": -1,
		},
		ExpectedInput: []string{"3", "1", "2", "3", "false"},
	},
	{
		Title: "IntegersWithDefaults uses default for invalid entries",
		ArrangeInput: args.Map{
			"when":       "given mix of valid and invalid strings",
			"input":      []string{"10", "abc", "20"},
			"defaultInt": -1,
		},
		ExpectedInput: []string{"3", "10", "-1", "20", "true"},
	},
	{
		Title: "IntegersWithDefaults returns empty on empty input",
		ArrangeInput: args.Map{
			"when":       "given empty input",
			"input":      []string{},
			"defaultInt": 0,
		},
		ExpectedInput: []string{"0", "false"},
	},
	{
		Title: "IntegersWithDefaults all invalid uses default everywhere",
		ArrangeInput: args.Map{
			"when":       "given all non-numeric strings",
			"input":      []string{"x", "y", "z"},
			"defaultInt": 99,
		},
		ExpectedInput: []string{"3", "99", "99", "99", "true"},
	},
	{
		Title: "IntegersWithDefaults handles negative numbers",
		ArrangeInput: args.Map{
			"when":       "given negative number strings",
			"input":      []string{"-5", "0", "5"},
			"defaultInt": 0,
		},
		ExpectedInput: []string{"3", "-5", "0", "5", "false"},
	},
}

// =============================================================================
// BytesWithDefaults
// =============================================================================

var bytesWithDefaultsTestCases = []coretestcases.CaseV1{
	{
		Title: "BytesWithDefaults converts valid byte values",
		ArrangeInput: args.Map{
			"when":        "given valid byte strings",
			"input":       []string{"0", "127", "255"},
			"defaultByte": byte(0),
		},
		ExpectedInput: []string{"3", "0", "127", "255", "false"},
	},
	{
		Title: "BytesWithDefaults uses default for out-of-range value",
		ArrangeInput: args.Map{
			"when":        "given value > 255",
			"input":       []string{"100", "256", "50"},
			"defaultByte": byte(42),
		},
		ExpectedInput: []string{"3", "100", "42", "50", "true"},
	},
	{
		Title: "BytesWithDefaults uses default for negative value",
		ArrangeInput: args.Map{
			"when":        "given negative value",
			"input":       []string{"-1", "10"},
			"defaultByte": byte(0),
		},
		ExpectedInput: []string{"2", "0", "10", "true"},
	},
	{
		Title: "BytesWithDefaults uses default for non-numeric",
		ArrangeInput: args.Map{
			"when":        "given non-numeric string",
			"input":       []string{"abc"},
			"defaultByte": byte(99),
		},
		ExpectedInput: []string{"1", "99", "true"},
	},
	{
		Title: "BytesWithDefaults empty input",
		ArrangeInput: args.Map{
			"when":        "given empty input",
			"input":       []string{},
			"defaultByte": byte(0),
		},
		ExpectedInput: []string{"0", "false"},
	},
}

// =============================================================================
// CloneIf
// =============================================================================

var cloneIfTestCases = []coretestcases.CaseV1{
	{
		Title: "CloneIf clones when isClone is true",
		ArrangeInput: args.Map{
			"when":    "given isClone true",
			"input":   []string{"a", "b", "c"},
			"isClone": true,
		},
		ExpectedInput: []string{"3", "a", "b", "c", "true"},
	},
	{
		Title: "CloneIf returns same slice when isClone is false",
		ArrangeInput: args.Map{
			"when":    "given isClone false",
			"input":   []string{"x", "y"},
			"isClone": false,
		},
		ExpectedInput: []string{"2", "x", "y", "false"},
	},
	{
		Title: "CloneIf returns empty on empty input regardless of isClone",
		ArrangeInput: args.Map{
			"when":    "given empty input with isClone true",
			"input":   []string{},
			"isClone": true,
		},
		ExpectedInput: []string{"0", "false"},
	},
}

// =============================================================================
// PtrOfPtrToPtrStrings
// =============================================================================

var ptrOfPtrToPtrStringsTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOfPtrToPtrStrings converts valid pointer strings",
		ArrangeInput: args.Map{
			"when":  "given valid pointer string array",
			"input": []string{"hello", "world"},
		},
		ExpectedInput: []string{"2", "hello", "world"},
	},
	{
		Title: "PtrOfPtrToPtrStrings handles nil entries",
		ArrangeInput: args.Map{
			"when":    "given array with nil entry",
			"input":   []string{"hello"},
			"hasNil":  true,
			"nilIdx":  1,
		},
		ExpectedInput: []string{"2", "hello", ""},
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty for nil outer pointer",
		ArrangeInput: args.Map{
			"when":    "given nil outer pointer",
			"isNil":   true,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty for nil inner pointer",
		ArrangeInput: args.Map{
			"when":      "given nil inner pointer",
			"isNilInner": true,
		},
		ExpectedInput: []string{"0"},
	},
}

// =============================================================================
// PtrOfPtrToMapStringBool
// =============================================================================

var ptrOfPtrToMapStringBoolTestCases = []coretestcases.CaseV1{
	{
		Title: "PtrOfPtrToMapStringBool converts valid entries",
		ArrangeInput: args.Map{
			"when":  "given valid pointer string array",
			"input": []string{"key1", "key2"},
		},
		ExpectedInput: []string{"2", "true", "true"},
	},
	{
		Title: "PtrOfPtrToMapStringBool skips nil entries",
		ArrangeInput: args.Map{
			"when":   "given array with nil entry",
			"input":  []string{"key1"},
			"hasNil": true,
		},
		ExpectedInput: []string{"1", "true"},
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty for nil input",
		ArrangeInput: args.Map{
			"when":  "given nil outer pointer",
			"isNil": true,
		},
		ExpectedInput: []string{"0"},
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty for empty array",
		ArrangeInput: args.Map{
			"when":  "given empty array",
			"input": []string{},
		},
		ExpectedInput: []string{"0"},
	},
}
