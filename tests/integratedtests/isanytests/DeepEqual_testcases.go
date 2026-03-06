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
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
		},
	},
	{
		Title: "DeepEqual false for different primitives",
		ArrangeInput: args.Map{
			"when": "given different int values",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "false",
			"isNotDeepEqual": "true",
		},
	},
	{
		Title: "DeepEqual true for identical slices",
		ArrangeInput: args.Map{
			"when": "given same string slices",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
		},
	},
	{
		Title: "DeepEqual false for different slices",
		ArrangeInput: args.Map{
			"when": "given different string slices",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "false",
			"isNotDeepEqual": "true",
		},
	},
	{
		Title: "DeepEqual true for both nil",
		ArrangeInput: args.Map{
			"when": "given both nil",
		},
		ExpectedInput: args.Map{
			"isDeepEqual":    "true",
			"isNotDeepEqual": "false",
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
		ExpectedInput: args.Map{
			"intZero":   "true",
			"int42":     "false",
			"emptyStr":  "true",
			"helloStr":  "false",
			"boolFalse": "true",
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
		ExpectedInput: args.Map{
			"nilPtr":    "true",
			"nonNilPtr": "false",
			"nilSlice":  "true",
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
		ExpectedInput: args.Map{
			"notNullNil":    "false",
			"notNull42":     "true",
			"inverseEquals": "false",
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
		ExpectedInput: args.Map{
			"sameStrings": "true",
			"diffStrings": "false",
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
		ExpectedInput: args.Map{
			"intPtr":    "true",
			"intVal":    "false",
			"stringPtr": "true",
		},
	},
}
