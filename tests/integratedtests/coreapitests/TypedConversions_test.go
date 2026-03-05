package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: NewTypedSimpleGenericRequest
// ==========================================

func Test_NewTypedSimpleGenericRequest(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)

		// Act
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		actLines := []string{
			fmt.Sprintf("%v", req.IsValid()),
			fmt.Sprintf("%v", req.Attribute.IsValid),
			req.Data(),
			fmt.Sprintf("%v", req.Request != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidTypedSimpleGenericRequest
// ==========================================

func Test_InvalidTypedSimpleGenericRequest(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestInvalidTestCases {
		// Act
		req := coreapi.InvalidTypedSimpleGenericRequest[string](nil)
		actLines := []string{
			fmt.Sprintf("%v", req.IsValid()),
			fmt.Sprintf("%v", req.Attribute != nil),
			fmt.Sprintf("%v", req.Request == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest IsValid / IsInvalid
// ==========================================

func Test_TypedSimpleGenericRequest_Validity(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestValidityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequestVal, _ := input.Get("nilRequest")
		nilRequest, _ := nilRequestVal.(bool)
		invalidAttributeVal, _ := input.Get("invalidAttribute")
		invalidAttribute, _ := invalidAttributeVal.(bool)
		nilAttributeVal, _ := input.Get("nilAttribute")
		nilAttribute, _ := nilAttributeVal.(bool)

		var req *coreapi.TypedSimpleGenericRequest[string]

		var attr *coreapi.RequestAttribute

		if nilAttribute {
			attr = nil
		} else if invalidAttribute {
			attr = &coreapi.RequestAttribute{IsValid: false}
		} else {
			attr = &coreapi.RequestAttribute{IsValid: true}
		}

		if nilRequest {
			req = &coreapi.TypedSimpleGenericRequest[string]{
				Attribute: attr,
				Request:   nil,
			}
		} else {
			payload, _ := input.GetAsString("payload")
			simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
			req = coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		}

		// Act
		actLines := []string{
			fmt.Sprintf("%v", req.IsValid()),
			fmt.Sprintf("%v", req.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Message / InvalidError
// ==========================================

func Test_TypedSimpleGenericRequest_Message(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequestVal, _ := input.Get("nilRequest")
		nilRequest, _ := nilRequestVal.(bool)

		var req *coreapi.TypedSimpleGenericRequest[string]

		if nilRequest {
			req = coreapi.InvalidTypedSimpleGenericRequest[string](nil)
		} else {
			payload, _ := input.GetAsString("payload")
			message, _ := input.GetAsString("message")
			attr := &coreapi.RequestAttribute{IsValid: true}
			simpleReq := coredynamic.NewTypedSimpleRequest[string](payload, false, message)
			req = coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)
		}

		// Act
		actLines := []string{
			req.Message(),
			fmt.Sprintf("%v", req.InvalidError() == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestNilReceiverTestCases {
		// Arrange
		var req *coreapi.TypedSimpleGenericRequest[string]

		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")

		// Act
		var actLines []string

		switch method {
		case "IsValid":
			actLines = []string{fmt.Sprintf("%v", req.IsValid())}
		case "IsInvalid":
			actLines = []string{fmt.Sprintf("%v", req.IsInvalid())}
		case "Message":
			actLines = []string{req.Message()}
		case "InvalidError":
			actLines = []string{fmt.Sprintf("%v", req.InvalidError() == nil)}
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestInvalidUnderlyingTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")
		message, _ := input.GetAsString("message")
		check, _ := input.GetAsString("check")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequest[string](payload, false, message)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		var actLines []string

		switch check {
		case "validity":
			actLines = []string{
				fmt.Sprintf("%v", req.IsValid()),
				fmt.Sprintf("%v", req.IsInvalid()),
			}
		case "message":
			actLines = []string{req.Message()}
		case "invalidError":
			actLines = []string{
				fmt.Sprintf("%v", req.InvalidError() == nil),
				req.InvalidError().Error(),
			}
		case "invalidErrorNil":
			actLines = []string{
				fmt.Sprintf("%v", req.InvalidError() == nil),
				req.Message(),
			}
		}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
// ==========================================
// Test: TypedSimpleGenericRequest Clone
// ==========================================

func Test_TypedSimpleGenericRequest_Clone(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		cloned := req.Clone()
		actLines := []string{
			cloned.Data(),
			fmt.Sprintf("%v", cloned.IsValid()),
			fmt.Sprintf("%v", cloned != req),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedSimpleGenericRequest_Clone_Nil(t *testing.T) {
	for caseIndex, tc := range typedSimpleGenericRequestCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedSimpleGenericRequest[string]

		// Act
		cloned := req.Clone()
		actLines := []string{fmt.Sprintf("%v", cloned == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedRequestIn.TypedSimpleGenericRequest conversion
// ==========================================

func Test_TypedRequestIn_TypedSimpleGenericRequest(t *testing.T) {
	for caseIndex, testCase := range typedRequestInToTypedSimpleGenericTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")
		isValid := true
		message := ""

		if vRaw, ok := input.Get("isValid"); ok {
			v, _ := vRaw.(bool)
		if true {
			_ = v
			isValid = v
		}

		if v, ok := input.GetAsString("message"); ok {
			message = v
		}

		attr := &coreapi.RequestAttribute{IsValid: true}
		reqIn := coreapi.NewTypedRequestIn[string](attr, payload)

		// Act
		tsgr := reqIn.TypedSimpleGenericRequest(isValid, message)
		actLines := []string{
			fmt.Sprintf("%v", tsgr.Request.IsValid()),
			tsgr.Data(),
			fmt.Sprintf("%v", tsgr.Attribute.IsValid),
			tsgr.Message(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedRequestIn_TypedSimpleGenericRequest_Nil(t *testing.T) {
	for caseIndex, tc := range typedRequestInToTypedSimpleGenericNilTestCases {
		// Arrange
		var reqIn *coreapi.TypedRequestIn[string]

		// Act
		result := reqIn.TypedSimpleGenericRequest(true, "")
		actLines := []string{fmt.Sprintf("%v", result == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponse.TypedResponseResult conversion
// ==========================================

func Test_TypedResponse_TypedResponseResult(t *testing.T) {
	for caseIndex, testCase := range typedResponseToTypedResponseResultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")
		message, _ := input.GetAsString("message")
		isValid := true

		if v, ok := input.GetAsBool("isValid"); ok {
			isValid = v
		}

		attr := &coreapi.ResponseAttribute{IsValid: isValid, Message: message}
		resp := coreapi.NewTypedResponse[string](attr, response)

		// Act
		result := resp.TypedResponseResult()
		actLines := []string{
			result.Response,
			fmt.Sprintf("%v", result.IsValid()),
			result.Message(),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedResponse_TypedResponseResult_Nil(t *testing.T) {
	for caseIndex, tc := range typedResponseToTypedResponseResultNilTestCases {
		// Arrange
		var resp *coreapi.TypedResponse[string]

		// Act
		result := resp.TypedResponseResult()
		actLines := []string{fmt.Sprintf("%v", result == nil)}

		// Assert
		tc.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
