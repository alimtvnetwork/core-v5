package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: NewTypedRequest
// ==========================================

func Test_NewTypedRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}

		// Act
		req := coreapi.NewTypedRequest[string](attr, payload)
		actLines := []string{
			req.Request,
			fmt.Sprintf("%v", req.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidTypedRequest
// ==========================================

func Test_InvalidTypedRequest_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInvalidTestCases {
		// Arrange — nil attribute

		// Act
		req := coreapi.InvalidTypedRequest[string](nil)
		actLines := []string{
			req.Request,
			fmt.Sprintf("%v", req.Attribute.IsValid),
			fmt.Sprintf("%v", req.Attribute != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedRequest Clone
// ==========================================

func Test_TypedRequest_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		req := coreapi.NewTypedRequest[string](attr, payload)

		// Act
		cloned := req.Clone()
		actLines := []string{
			cloned.Request,
			fmt.Sprintf("%v", cloned.Attribute.IsValid),
			fmt.Sprintf("%v", cloned != req),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedRequest_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedRequest[string]

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
// Test: TypedResponseResult ToTypedResponse
// ==========================================

func Test_TypedResponseResult_ToTypedResponse_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultToTypedResponseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")

		attr := &coreapi.ResponseAttribute{IsValid: true}
		result := coreapi.NewTypedResponseResult[string](attr, response)

		// Act
		resp := result.ToTypedResponse()
		actLines := []string{
			resp.Response,
			fmt.Sprintf("%v", resp.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

func Test_TypedResponseResult_ToTypedResponse_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultToTypedResponseNilTestCases {
		// Arrange
		var result *coreapi.TypedResponseResult[string]

		// Act
		resp := result.ToTypedResponse()
		actLines := []string{
			fmt.Sprintf("%v", resp == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponseResult Message
// ==========================================

func Test_TypedResponseResult_Message_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultMessageTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)

		var result *coreapi.TypedResponseResult[string]

		if response, ok := input.GetAsString("response"); ok {
			message, _ := input.GetAsString("message")
			attr := &coreapi.ResponseAttribute{IsValid: true, Message: message}
			result = coreapi.NewTypedResponseResult[string](attr, response)
		}

		// Act
		var msg string
		if result != nil {
			msg = result.Message()
		}
		actLines := []string{
			msg,
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
