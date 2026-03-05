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
		ExpectedInput: args.Two[string, string]{
			First:  "hello", // payload
			Second: "true",  // isValid
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "",      // payload
			Second: "false", // isValid
			Third:  "true",  // isInvalid
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
		ExpectedInput: args.Two[string, string]{
			First:  "cloneme", // payload
			Second: "true",    // isValid
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
		ExpectedInput: "true",
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
		ExpectedInput: args.Two[string, string]{
			First:  "42",   // response
			Second: "true", // isValid
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "0",     // response
			Second: "false", // isValid
			Third:  "true",  // isInvalid
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
		ExpectedInput: args.Two[string, string]{
			First:  "99",   // response
			Second: "true", // isValid
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "ok",   // response
			Second: "true", // isValid
			Third:  "true", // hasResponse
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
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "cloneable", // response
			Second: "true",      // isValid
			Third:  "true",      // isIndependent
		},
	},
}

var typedResponseResultCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: "true",
	},
}
