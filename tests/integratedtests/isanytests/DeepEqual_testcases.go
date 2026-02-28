package isanytests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// DeepEqual / NotDeepEqual
// ==========================================

var deepEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "DeepEqual true for identical primitives",
		ArrangeInput: args.Map{
			"when": "given same int values",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "DeepEqual false for different primitives",
		ArrangeInput: args.Map{
			"when": "given different int values",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "DeepEqual true for identical slices",
		ArrangeInput: args.Map{
			"when": "given same string slices",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "DeepEqual false for different slices",
		ArrangeInput: args.Map{
			"when": "given different string slices",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "DeepEqual true for both nil",
		ArrangeInput: args.Map{
			"when": "given both nil",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

// ==========================================
// Zero
// ==========================================

var zeroTestCases = []coretestcases.CaseV1{
	{
		Title: "Zero returns true for zero values and false for non-zero",
		ArrangeInput: args.Map{
			"when": "given various zero and non-zero values",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
			"false",
			"true",
		},
	},
}

// ==========================================
// ReflectNull vs Null comparison
// ==========================================

var reflectNullTestCases = []coretestcases.CaseV1{
	{
		Title: "ReflectNull returns true for nil pointer, false for non-nil",
		ArrangeInput: args.Map{
			"when": "given nil and non-nil pointers",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
		},
	},
}

// ==========================================
// NotNull
// ==========================================

var notNullTestCases = []coretestcases.CaseV1{
	{
		Title: "NotNull is inverse of Null",
		ArrangeInput: args.Map{
			"when": "given nil and non-nil values",
		},
		ExpectedInput: []string{
			"false",
			"true",
			"true",
		},
	},
}

// ==========================================
// StringEqual
// ==========================================

var stringEqualTestCases = []coretestcases.CaseV1{
	{
		Title: "StringEqual compares string representation of values",
		ArrangeInput: args.Map{
			"when": "given values with same/different string representation",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
}

// ==========================================
// Pointer
// ==========================================

var pointerTestCases = []coretestcases.CaseV1{
	{
		Title: "Pointer returns true for pointers, false for values",
		ArrangeInput: args.Map{
			"when": "given pointer and non-pointer values",
		},
		ExpectedInput: []string{
			"true",
			"false",
			"true",
		},
	},
}
