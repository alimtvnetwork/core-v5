package coreapitests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
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
		ExpectedInput: args.Map{
			"payload": "hello",
			"isValid": "true",
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
		ExpectedInput: args.Map{
			"payload":      "",
			"isValid":      "false",
			"hasAttribute": "true",
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
		ExpectedInput: args.Map{
			"payload": "cloneme",
			"isValid": "true",
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
		ExpectedInput: args.Map{
			"response": "42",
			"isValid":  "true",
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
		ExpectedInput: args.Map{
			"response":     "0",
			"isValid":      "false",
			"hasAttribute": "true",
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
		ExpectedInput: args.Map{
			"response": "99",
			"isValid":  "true",
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
		ExpectedInput: args.Map{
			"response":    "ok",
			"isValid":     "true",
			"hasResponse": "true",
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
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
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
		ExpectedInput: args.Map{
			"response":      "cloneable",
			"isValid":       "true",
			"isIndependent": "true",
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
