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
		ExpectedInput: args.Map{
			"isValid":      "true",
			"hasAttribute": "true",
			"payload":      "hello-world",
			"hasRequest":   "true",
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
		ExpectedInput: args.Map{
			"isValid":      "false",
			"isInvalid":    "true",
			"isNilRequest": "true",
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
		ExpectedInput: args.Map{
			"isValid":   "true",
			"isInvalid": "false",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":       "given valid attribute but nil request",
			"nilRequest": true,
		},
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and valid request",
			"payload":          "data",
			"invalidAttribute": true,
		},
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":         "given nil attribute and valid request",
			"payload":      "data",
			"nilAttribute": true,
		},
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute and nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and nil request",
			"nilRequest":       true,
			"invalidAttribute": true,
		},
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
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
		ExpectedInput: args.Map{
			"message":    "validation failed",
			"isNilError": "false",
		},
	},
	{
		Title: "Message returns empty string when request is nil",
		ArrangeInput: args.Map{
			"when":       "given nil request",
			"nilRequest": true,
		},
		ExpectedInput: args.Map{
			"message":    "",
			"isNilError": "true",
		},
	},
}
// Note: Nil receiver test cases migrated to TypedConversions_NilReceiver_testcases.go
// using CaseNilSafe pattern with function literal wrappers for generic types.

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
		ExpectedInput: args.Map{
			"isValid":   "false",
			"isInvalid": "true",
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
		ExpectedInput: "field is required",
	},
	{
		Title: "Valid attribute with invalid underlying request returns non-nil InvalidError",
		ArrangeInput: args.Map{
			"when":    "given valid attribute and invalid underlying TypedSimpleRequest with error message",
			"payload": "some-data",
			"message": "input rejected",
			"check":   "invalidError",
		},
		ExpectedInput: args.Map{
			"isNilError":   "false",
			"errorMessage": "input rejected",
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
		ExpectedInput: args.Map{
			"isNilError":   "true",
			"errorMessage": "",
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
		ExpectedInput: args.Map{
			"payload":       "clone-me",
			"isValid":       "true",
			"isIndependent": "true",
		},
	},
}

var typedSimpleGenericRequestCloneNilTestCases = []coretestcases.CaseV1{
	{
		Title: "Clone on nil TypedSimpleGenericRequest returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: "true",
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
		ExpectedInput: args.Map{
			"isValid":    "true",
			"payload":    "wrapped-data",
			"hasRequest": "true",
			"message":    "",
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
		ExpectedInput: args.Map{
			"isValid":         "false",
			"payload":         "bad-data",
			"hasValidRequest": "false",
			"message":         "input rejected",
		},
	},
}

var typedRequestInToTypedSimpleGenericNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedRequestIn.TypedSimpleGenericRequest on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil request",
		},
		ExpectedInput: "true",
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
		ExpectedInput: args.Map{
			"response": "result-data",
			"isValid":  "true",
			"message":  "success",
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
		ExpectedInput: args.Map{
			"response": "error-data",
			"isValid":  "false",
			"message":  "failed",
		},
	},
}

var typedResponseToTypedResponseResultNilTestCases = []coretestcases.CaseV1{
	{
		Title: "TypedResponse.TypedResponseResult on nil returns nil",
		ArrangeInput: args.Map{
			"when": "given nil response",
		},
		ExpectedInput: "true",
	},
}
