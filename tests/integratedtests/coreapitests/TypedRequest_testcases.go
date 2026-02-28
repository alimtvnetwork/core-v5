package coreapitests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		ExpectedInput: []string{
			"my-request",
			"true",
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
		ExpectedInput: []string{
			"",
			"false",
			"true",
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
		ExpectedInput: []string{
			"clone-payload",
			"true",
			"true",
		},
	},
}

var typedRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedRequest returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedRequest — ToGenericRequestIn
// ==========================================

var typedRequestToGenericRequestInTestCases = []coretestcases.CaseV1{
	{
		Title: "ToGenericRequestIn converts TypedRequest to GenericRequestIn",
		ArrangeInput: args.Map{
			"when":    "given valid typed request",
			"payload": "legacy-convert",
		},
		ExpectedInput: []string{
			"true",
			"legacy-convert",
		},
	},
}

var typedRequestToGenericRequestInNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToGenericRequestIn on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedRequest — ToSimpleGenericRequest
// ==========================================

var typedRequestToSimpleGenericRequestTestCases = []coretestcases.CaseV1{
	{
		Title: "ToSimpleGenericRequest wraps request in SimpleRequest",
		ArrangeInput: args.Map{
			"when":    "given valid typed request",
			"payload": "wrap-data",
		},
		ExpectedInput: []string{
			"true",
			"true",
			"true",
		},
	},
}

var typedRequestToSimpleGenericRequestNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToSimpleGenericRequest on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedResponseResult — ToGenericResponseResult
// ==========================================

var typedResponseResultToGenericTestCases = []coretestcases.CaseV1{
	{
		Title: "ToGenericResponseResult converts typed result to legacy",
		ArrangeInput: args.Map{
			"when":     "given valid typed response result",
			"response": "data",
		},
		ExpectedInput: []string{
			"true",
			"true",
		},
	},
}

var typedResponseResultToGenericNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToGenericResponseResult on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: []string{
			"true",
		},
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
		ExpectedInput: []string{
			"back-convert",
			"true",
		},
	},
}

var typedResponseResultToTypedResponseNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ToTypedResponse on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: []string{
			"true",
		},
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
		ExpectedInput: []string{
			"operation completed",
		},
	},
	{
		Title: "Message returns empty string on nil result",
		ArrangeInput: args.Map{
			"when": "given nil result",
		},
		ExpectedInput: []string{
			"",
		},
	},
}
