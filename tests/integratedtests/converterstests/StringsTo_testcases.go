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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",     // count
			Second: "1",     // val0
			Third:  "2",     // val1
			Fourth: "3",     // val2
			Fifth:  "false", // hadDefaultUsed
		},
	},
	{
		Title: "IntegersWithDefaults uses default for invalid entries",
		ArrangeInput: args.Map{
			"when":       "given mix of valid and invalid strings",
			"input":      []string{"10", "abc", "20"},
			"defaultInt": -1,
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",    // count
			Second: "10",   // val0
			Third:  "-1",   // val1 (default)
			Fourth: "20",   // val2
			Fifth:  "true", // hadDefaultUsed
		},
	},
	{
		Title: "IntegersWithDefaults returns empty on empty input",
		ArrangeInput: args.Map{
			"when":       "given empty input",
			"input":      []string{},
			"defaultInt": 0,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",     // count
			Second: "false", // hadDefaultUsed
		},
	},
	{
		Title: "IntegersWithDefaults all invalid uses default everywhere",
		ArrangeInput: args.Map{
			"when":       "given all non-numeric strings",
			"input":      []string{"x", "y", "z"},
			"defaultInt": 99,
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",    // count
			Second: "99",   // val0 (default)
			Third:  "99",   // val1 (default)
			Fourth: "99",   // val2 (default)
			Fifth:  "true", // hadDefaultUsed
		},
	},
	{
		Title: "IntegersWithDefaults handles negative numbers",
		ArrangeInput: args.Map{
			"when":       "given negative number strings",
			"input":      []string{"-5", "0", "5"},
			"defaultInt": 0,
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",     // count
			Second: "-5",    // val0
			Third:  "0",     // val1
			Fourth: "5",     // val2
			Fifth:  "false", // hadDefaultUsed
		},
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",     // count
			Second: "0",     // val0
			Third:  "127",   // val1
			Fourth: "255",   // val2
			Fifth:  "false", // hadDefaultUsed
		},
	},
	{
		Title: "BytesWithDefaults uses default for out-of-range value",
		ArrangeInput: args.Map{
			"when":        "given value > 255",
			"input":       []string{"100", "256", "50"},
			"defaultByte": byte(42),
		},
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",    // count
			Second: "100",  // val0
			Third:  "42",   // val1 (default)
			Fourth: "50",   // val2
			Fifth:  "true", // hadDefaultUsed
		},
	},
	{
		Title: "BytesWithDefaults uses default for negative value",
		ArrangeInput: args.Map{
			"when":        "given negative value",
			"input":       []string{"-1", "10"},
			"defaultByte": byte(0),
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "2",    // count
			Second: "0",    // val0 (default)
			Third:  "10",   // val1
			Fourth: "true", // hadDefaultUsed
		},
	},
	{
		Title: "BytesWithDefaults uses default for non-numeric",
		ArrangeInput: args.Map{
			"when":        "given non-numeric string",
			"input":       []string{"abc"},
			"defaultByte": byte(99),
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "1",    // count
			Second: "99",   // val0 (default)
			Third:  "true", // hadDefaultUsed
		},
	},
	{
		Title: "BytesWithDefaults empty input",
		ArrangeInput: args.Map{
			"when":        "given empty input",
			"input":       []string{},
			"defaultByte": byte(0),
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",     // count
			Second: "false", // hadDefaultUsed
		},
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
		ExpectedInput: args.Five[string, string, string, string, string]{
			First:  "3",    // count
			Second: "a",    // item0
			Third:  "b",    // item1
			Fourth: "c",    // item2
			Fifth:  "true", // isIndependent
		},
	},
	{
		Title: "CloneIf returns same slice when isClone is false",
		ArrangeInput: args.Map{
			"when":    "given isClone false",
			"input":   []string{"x", "y"},
			"isClone": false,
		},
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "2",     // count
			Second: "x",     // item0
			Third:  "y",     // item1
			Fourth: "false", // isIndependent
		},
	},
	{
		Title: "CloneIf returns empty on empty input regardless of isClone",
		ArrangeInput: args.Map{
			"when":    "given empty input with isClone true",
			"input":   []string{},
			"isClone": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "0",     // count
			Second: "false", // isIndependent
		},
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",     // count
			Second: "hello", // item0
			Third:  "world", // item1
		},
	},
	{
		Title: "PtrOfPtrToPtrStrings handles nil entries",
		ArrangeInput: args.Map{
			"when":   "given array with nil entry",
			"input":  []string{"hello"},
			"hasNil": true,
			"nilIdx": 1,
		},
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",     // count
			Second: "hello", // item0
			Third:  "",      // item1 (nil)
		},
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty for nil outer pointer",
		ArrangeInput: args.Map{
			"when":  "given nil outer pointer",
			"isNil": true,
		},
		ExpectedInput: "0", // count
	},
	{
		Title: "PtrOfPtrToPtrStrings returns empty for nil inner pointer",
		ArrangeInput: args.Map{
			"when":       "given nil inner pointer",
			"isNilInner": true,
		},
		ExpectedInput: "0", // count
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "2",    // count
			Second: "true", // hasKey1
			Third:  "true", // hasKey2
		},
	},
	{
		Title: "PtrOfPtrToMapStringBool skips nil entries",
		ArrangeInput: args.Map{
			"when":   "given array with nil entry",
			"input":  []string{"key1"},
			"hasNil": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "1",    // count
			Second: "true", // hasKey1
		},
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty for nil input",
		ArrangeInput: args.Map{
			"when":  "given nil outer pointer",
			"isNil": true,
		},
		ExpectedInput: "0", // count
	},
	{
		Title: "PtrOfPtrToMapStringBool returns empty for empty array",
		ArrangeInput: args.Map{
			"when":  "given empty array",
			"input": []string{},
		},
		ExpectedInput: "0", // count
	},
}
