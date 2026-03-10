package coreapitests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// TypedRequest — NewTypedRequest
// ==========================================

var typedRequestNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedRequest creates valid typed request",
		ArrangeInput: args.Map{
			"when":    "given string payload",
			"payload": "my-request",
		},
		ExpectedInput: args.Map{
			"payload": "my-request",
			"isValid": "true",
		},
	},
}

// ==========================================
// TypedRequest — InvalidTypedRequest
// ==========================================

var typedRequestInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedRequest creates request with zero-value payload",
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
// TypedRequest — Clone
// ==========================================

var typedRequestCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy of TypedRequest",
		ArrangeInput: args.Map{
			"when":    "given valid request",
			"payload": "clone-payload",
		},
		ExpectedInput: args.Map{
			"payload":       "clone-payload",
			"isValid":       "true",
			"isIndependent": "true",
		},
	},
}

var typedRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedRequest returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// TypedResponseResult — ToTypedResponse
// ==========================================

var typedResponseResultToTypedResponseTestCases = []coretestcases.CaseV1{
	{
		Title: "ToTypedResponse converts result back to TypedResponse",
		ArrangeInput: args.Map{
			"when":     "given valid typed response result",
			"response": "back-convert",
		},
		ExpectedInput: args.Map{
			"response": "back-convert",
			"isValid":  "true",
		},
	},
}

var typedResponseResultToTypedResponseNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToTypedResponse on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: "true",
	},
}

// ==========================================
// TypedResponseResult — Message
// ==========================================

var typedResponseResultMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "Message returns attribute message",
		ArrangeInput: args.Map{
			"when":     "given result with message",
			"response": "data",
			"message":  "operation completed",
		},
		ExpectedInput: "operation completed",
	},
	{
		Title: "Message returns empty string on nil result",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: "",
	},
}
