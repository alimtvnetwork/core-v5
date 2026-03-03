package coreapitests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// ==========================================
// TypedSimpleGenericRequest — New
// ==========================================

var typedSimpleGenericRequestNewTestCases = []coretestcases.CaseV1{
	{
		Title: "NewTypedSimpleGenericRequest creates valid request with typed data",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and typed simple request",
			"payload": "hello-world",
		},
		ExpectedInput: []string{
			"true",
			"true",
			"hello-world",
			"true",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — Invalid
// ==========================================

var typedSimpleGenericRequestInvalidTestCases = []coretestcases.CaseV1{
	{
		Title: "InvalidTypedSimpleGenericRequest creates invalid with nil request",
		ArrangeInput: args.Map{
			"when": "given nil attribute",
		},
		ExpectedInput: []string{
			"false",
			"true",
			"true",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — IsValid / IsInvalid
// ==========================================

var typedSimpleGenericRequestValidityTestCases = []coretestcases.CaseV1{
	{
		Title: "Valid TypedSimpleGenericRequest reports IsValid true",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and valid request",
			"payload": "data",
		},
		ExpectedInput: []string{
			"true",
			"false",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":       "given valid attribute but nil request",
			"nilRequest": true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and valid request",
			"payload":          "data",
			"invalidAttribute": true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":          "given nil attribute and valid request",
			"payload":       "data",
			"nilAttribute":  true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute and nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and nil request",
			"nilRequest":       true,
			"invalidAttribute": true,
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — Message / InvalidError
// ==========================================

var typedSimpleGenericRequestMessageTestCases = []coretestcases.CaseV1{
	{
		Title: "Message returns request message from underlying TypedSimpleRequest",
		ArrangeInput: args.Map{
			"when":    "given request with message",
			"payload": "data",
			"message": "validation failed",
		},
		ExpectedInput: []string{
			"validation failed",
			"false",
		},
	},
	{
		Title: "Message returns empty string when request is nil",
		ArrangeInput: args.Map{
			"when":       "given nil request",
			"nilRequest": true,
		},
		ExpectedInput: []string{
			"",
			"true",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — Nil Receiver Edge Cases
// ==========================================

var typedSimpleGenericRequestNilReceiverTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil receiver IsValid returns false",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "IsValid",
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Nil receiver IsInvalid returns true",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "IsInvalid",
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Nil receiver Message returns empty string",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "Message",
		},
		ExpectedInput: []string{
			"",
		},
	},
	{
		Title: "Nil receiver InvalidError returns nil",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "InvalidError",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — Invalid Underlying Request Edge Cases
// ==========================================

var typedSimpleGenericRequestInvalidUnderlyingTestCases = []coretestcases.CaseV1{
	{
		Title: "Valid attribute with invalid underlying request reports IsValid false",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and invalid underlying TypedSimpleRequest",
			"payload": "some-data",
			"message": "validation failed",
			"check":   "validity",
		},
		ExpectedInput: []string{
			"false",
			"true",
		},
	},
	{
		Title: "Valid attribute with invalid underlying request returns message",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and invalid underlying TypedSimpleRequest with message",
			"payload": "some-data",
			"message": "field is required",
			"check":   "message",
		},
		ExpectedInput: []string{
			"field is required",
		},
	},
	{
		Title: "Valid attribute with invalid underlying request returns non-nil InvalidError",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and invalid underlying TypedSimpleRequest with error message",
			"payload": "some-data",
			"message": "input rejected",
			"check":   "invalidError",
		},
		ExpectedInput: []string{
			"false",
			"input rejected",
		},
	},
	{
		Title: "Valid attribute with invalid underlying request and empty message returns nil InvalidError",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and invalid underlying TypedSimpleRequest with empty message",
			"payload": "some-data",
			"message": "",
			"check":   "invalidErrorNil",
		},
		ExpectedInput: []string{
			"true",
			"",
		},
	},
}

// ==========================================
// TypedSimpleGenericRequest — Clone
// ==========================================

var typedSimpleGenericRequestCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone creates independent copy of TypedSimpleGenericRequest",
		ArrangeInput: args.Map{
			"when":    "given valid request",
			"payload": "clone-me",
		},
		ExpectedInput: []string{
			"clone-me",
			"true",
			"true",
		},
	},
}

var typedSimpleGenericRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedSimpleGenericRequest returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedRequestIn — TypedSimpleGenericRequest conversion
// ==========================================

var typedRequestInToTypedSimpleGenericTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedRequestIn.TypedSimpleGenericRequest creates valid conversion",
		ArrangeInput: args.Map{
			"when":    "given valid typed request in",
			"payload": "wrapped-data",
		},
		ExpectedInput: []string{
			"true",
			"wrapped-data",
			"true",
			"",
		},
	},
	{
		Title: "TypedRequestIn.TypedSimpleGenericRequest with invalid message",
		ArrangeInput: args.Map{
			"when":    "given request with invalid flag",
			"payload": "bad-data",
			"isValid": false,
			"message": "input rejected",
		},
		ExpectedInput: []string{
			"false",
			"bad-data",
			"false",
			"input rejected",
		},
	},
}

var typedRequestInToTypedSimpleGenericNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedRequestIn.TypedSimpleGenericRequest on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

// ==========================================
// TypedResponse — TypedResponseResult conversion
// ==========================================

var typedResponseToTypedResponseResultTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedResponse.TypedResponseResult creates valid result",
		ArrangeInput: args.Map{
			"when":     "given valid typed response",
			"response": "result-data",
			"message":  "success",
		},
		ExpectedInput: []string{
			"result-data",
			"true",
			"success",
		},
	},
	{
		Title: "TypedResponse.TypedResponseResult preserves invalid state",
		ArrangeInput: args.Map{
			"when":     "given invalid typed response",
			"response": "error-data",
			"isValid":  false,
			"message":  "failed",
		},
		ExpectedInput: []string{
			"error-data",
			"false",
			"failed",
		},
	},
}

var typedResponseToTypedResponseResultNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedResponse.TypedResponseResult on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil response",
		},
		ExpectedInput: []string{
			"true",
		},
	},
}

