package coreapitests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// TypedRequestIn — NewTypedRequestIn
// ==========================================

var typedRequestInNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedRequestIn creates valid typed request",
		ArrangeInput: args.Map{
			"when":    "given string payload",
			"payload": "hello",
		},
		ExpectedInput: []string{
			"hello",
			"true",
		},
	},
}

// ==========================================
// TypedRequestIn — InvalidTypedRequestIn
// ==========================================

var typedRequestInInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedRequestIn creates request with zero-value payload",
		ArrangeInput: args.Map{
			"when": "given nil attribute",
		},
		ExpectedInput: []string{
			"",
			"false",
			"true",
		},
	},
}

// ==========================================
// TypedRequestIn — Clone
// ==========================================

var typedRequestInCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy of TypedRequestIn",
		ArrangeInput: args.Map{
			"when":    "given valid request",
			"payload": "cloneme",
		},
		ExpectedInput: []string{
			"cloneme",
			"true",
		},
	},
}

// ==========================================
// TypedRequestIn — nil Clone
// ==========================================

var typedRequestInCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedRequestIn returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedResponse — NewTypedResponse
// ==========================================

var typedResponseNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedResponse creates valid typed response",
		ArrangeInput: args.Map{
			"when":     "given int response",
			"response": 42,
		},
		ExpectedInput: []string{
			"42",
			"true",
		},
	},
}

// ==========================================
// TypedResponse — InvalidTypedResponse
// ==========================================

var typedResponseInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedResponse creates response with zero-value",
		ArrangeInput: args.Map{
			"when": "given nil attribute",
		},
		ExpectedInput: []string{
			"0",
			"false",
			"true",
		},
	},
}

// ==========================================
// TypedResponse — Clone
// ==========================================

var typedResponseCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy of TypedResponse",
		ArrangeInput: args.Map{
			"when":     "given valid response",
			"response": 99,
		},
		ExpectedInput: []string{
			"99",
			"true",
		},
	},
}

// ==========================================
// TypedResponseResult — NewTypedResponseResult
// ==========================================

var typedResponseResultNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedResponseResult creates valid result",
		ArrangeInput: args.Map{
			"when":     "given string response",
			"response": "ok",
		},
		ExpectedInput: []string{
			"ok",
			"true",
			"true",
		},
	},
}

// ==========================================
// TypedResponseResult — IsValid / IsInvalid
// ==========================================

var typedResponseResultInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedResponseResult reports IsInvalid correctly",
		ArrangeInput: args.Map{
			"when": "given invalid result",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

// ==========================================
// TypedResponseResult — Clone / ClonePtr
// ==========================================

var typedResponseResultCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr creates independent copy of TypedResponseResult",
		ArrangeInput: args.Map{
			"when":     "given valid result",
			"response": "cloneable",
		},
		ExpectedInput: []string{
			"cloneable",
			"true",
			"true",
		},
	},
}

var typedResponseResultCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}
