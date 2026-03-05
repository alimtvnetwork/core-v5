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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "true",        // isValid
			Second: "true",        // hasAttribute
			Third:  "hello-world", // payload
			Fourth: "true",        // hasRequest
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
			Third:  "true",  // isNilRequest
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
		ExpectedInput: args.Two[string, string]{
			First:  "true",  // isValid
			Second: "false", // isInvalid
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":       "given valid attribute but nil request",
			"nilRequest": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and valid request",
			"payload":          "data",
			"invalidAttribute": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
		},
	},
	{
		Title: "TypedSimpleGenericRequest with nil attribute but valid request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":         "given nil attribute and valid request",
			"payload":      "data",
			"nilAttribute": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
		},
	},
	{
		Title: "TypedSimpleGenericRequest with invalid attribute and nil request reports IsInvalid",
		ArrangeInput: args.Map{
			"when":             "given invalid attribute and nil request",
			"nilRequest":       true,
			"invalidAttribute": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
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
		ExpectedInput: args.Two[string, string]{
			First:  "validation failed", // message
			Second: "false",             // isNilError
		},
	},
	{
		Title: "Message returns empty string when request is nil",
		ArrangeInput: args.Map{
			"when":       "given nil request",
			"nilRequest": true,
		},
		ExpectedInput: args.Two[string, string]{
			First:  "",     // message
			Second: "true", // isNilError
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
		ExpectedInput: "false",
	},
	{
		Title: "Nil receiver IsInvalid returns true",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "IsInvalid",
		},
		ExpectedInput: "true",
	},
	{
		Title: "Nil receiver Message returns empty string",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "Message",
		},
		ExpectedInput: "",
	},
	{
		Title: "Nil receiver InvalidError returns nil",
		ArrangeInput: args.Map{
			"when":   "given nil receiver",
			"method": "InvalidError",
		},
		ExpectedInput: "true",
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
		ExpectedInput: args.Two[string, string]{
			First:  "false", // isValid
			Second: "true",  // isInvalid
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
		ExpectedInput: args.Two[string, string]{
			First:  "false",          // isNilError
			Second: "input rejected", // errorMessage
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
		ExpectedInput: args.Two[string, string]{
			First:  "true", // isNilError
			Second: "",     // errorMessage
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "clone-me", // payload
			Second: "true",     // isValid
			Third:  "true",     // isIndependent
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "true",         // isValid
			Second: "wrapped-data", // payload
			Third:  "true",         // hasRequest
			Fourth: "",             // message
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
		ExpectedInput: args.Four[string, string, string, string]{
			First:  "false",          // isValid
			Second: "bad-data",       // payload
			Third:  "false",          // hasValidRequest
			Fourth: "input rejected", // message
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "result-data", // response
			Second: "true",        // isValid
			Third:  "success",     // message
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
		ExpectedInput: args.Three[string, string, string]{
			First:  "error-data", // response
			Second: "false",      // isValid
			Third:  "failed",     // message
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
