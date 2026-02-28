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

func Test_NewTypedSimpleGenericRequest_Verification(t *testing.T) {
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

func Test_InvalidTypedSimpleGenericRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestInvalidTestCases {
		// Arrange — nil attribute

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

func Test_TypedSimpleGenericRequest_Validity_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestValidityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequest, _ := input.GetAsBool("nilRequest")

		var req *coreapi.TypedSimpleGenericRequest[string]
		attr := &coreapi.RequestAttribute{IsValid: true}

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

func Test_TypedSimpleGenericRequest_Message_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequest, _ := input.GetAsBool("nilRequest")

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
// Test: TypedSimpleGenericRequest Clone
// ==========================================

func Test_TypedSimpleGenericRequest_Clone_Verification(t *testing.T) {
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

func Test_TypedSimpleGenericRequest_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedSimpleGenericRequest[string]

		// Act
		cloned := req.Clone()
		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest ToSimpleGenericRequest
// ==========================================

func Test_TypedSimpleGenericRequest_ToSimpleGenericRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestToLegacyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		legacy := req.ToSimpleGenericRequest()
		actLines := []string{
			fmt.Sprintf("%v", legacy != nil),
			fmt.Sprintf("%v", legacy.Attribute.IsValid),
			fmt.Sprintf("%v", legacy.Request != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedSimpleGenericRequest_ToSimpleGenericRequest_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestToLegacyNilTestCases {
		// Arrange
		var req *coreapi.TypedSimpleGenericRequest[string]

		// Act
		legacy := req.ToSimpleGenericRequest()
		actLines := []string{
			fmt.Sprintf("%v", legacy == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest ToGenericRequestIn
// ==========================================

func Test_TypedSimpleGenericRequest_ToGenericRequestIn_Verification(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestToGenericRequestInTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		simpleReq := coredynamic.NewTypedSimpleRequestValid[string](payload)
		req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

		// Act
		gri := req.ToGenericRequestIn()
		actLines := []string{
			fmt.Sprintf("%v", gri.Attribute.IsValid),
			fmt.Sprintf("%v", gri.Request),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedRequestIn.TypedSimpleGenericRequest conversion
// ==========================================

func Test_TypedRequestIn_TypedSimpleGenericRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInToTypedSimpleGenericTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")
		isValid := true
		message := ""

		if v, ok := input.GetAsBool("isValid"); ok {
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

func Test_TypedRequestIn_TypedSimpleGenericRequest_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInToTypedSimpleGenericNilTestCases {
		// Arrange
		var reqIn *coreapi.TypedRequestIn[string]

		// Act
		result := reqIn.TypedSimpleGenericRequest(true, "")
		actLines := []string{
			fmt.Sprintf("%v", result == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponse.TypedResponseResult conversion
// ==========================================

func Test_TypedResponse_TypedResponseResult_Verification(t *testing.T) {
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

func Test_TypedResponse_TypedResponseResult_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseToTypedResponseResultNilTestCases {
		// Arrange
		var resp *coreapi.TypedResponse[string]

		// Act
		result := resp.TypedResponseResult()
		actLines := []string{
			fmt.Sprintf("%v", result == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponse.GenericResponseResult conversion
// ==========================================

func Test_TypedResponse_GenericResponseResult_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseToGenericResponseResultTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")
		message, _ := input.GetAsString("message")

		attr := &coreapi.ResponseAttribute{IsValid: true, Message: message}
		resp := coreapi.NewTypedResponse[string](attr, response)

		// Act
		grr := resp.GenericResponseResult()
		actLines := []string{
			fmt.Sprintf("%v", grr != nil),
			fmt.Sprintf("%v", grr.Attribute.IsValid),
			fmt.Sprintf("%v", grr.Response != nil),
			grr.Response.Message,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedResponse_GenericResponseResult_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseToGenericResponseResultNilTestCases {
		// Arrange
		var resp *coreapi.TypedResponse[string]

		// Act
		grr := resp.GenericResponseResult()
		actLines := []string{
			fmt.Sprintf("%v", grr == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
