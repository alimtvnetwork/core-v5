package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coretests/args"
)

// ==========================================
// Test: NewTypedRequestIn
// ==========================================

func Test_NewTypedRequestIn_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}

		// Act
		req := coreapi.NewTypedRequestIn[string](attr, payload)
		actLines := []string{
			req.Request,
			fmt.Sprintf("%v", req.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidTypedRequestIn
// ==========================================

func Test_InvalidTypedRequestIn_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInInvalidTestCases {
		// Arrange — nil attribute

		// Act
		req := coreapi.InvalidTypedRequestIn[string](nil)
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
// Test: TypedRequestIn Clone
// ==========================================

func Test_TypedRequestIn_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}
		req := coreapi.NewTypedRequestIn[string](attr, payload)

		// Act
		cloned := req.Clone()
		actLines := []string{
			cloned.Request,
			fmt.Sprintf("%v", cloned.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedRequestIn Clone nil
// ==========================================

func Test_TypedRequestIn_Clone_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedRequestInCloneNilTestCases {
		// Arrange
		var req *coreapi.TypedRequestIn[string]

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
// Test: NewTypedResponse
// ==========================================

func Test_NewTypedResponse_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response := input["response"].(int)

		attr := &coreapi.ResponseAttribute{IsValid: true}

		// Act
		resp := coreapi.NewTypedResponse[int](attr, response)
		actLines := []string{
			fmt.Sprintf("%d", resp.Response),
			fmt.Sprintf("%v", resp.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidTypedResponse
// ==========================================

func Test_InvalidTypedResponse_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseInvalidTestCases {
		// Arrange — nil attribute

		// Act
		resp := coreapi.InvalidTypedResponse[int](nil)
		actLines := []string{
			fmt.Sprintf("%d", resp.Response),
			fmt.Sprintf("%v", resp.Attribute.IsValid),
			fmt.Sprintf("%v", resp.Attribute != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponse Clone
// ==========================================

func Test_TypedResponse_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response := input["response"].(int)

		attr := &coreapi.ResponseAttribute{IsValid: true}
		resp := coreapi.NewTypedResponse[int](attr, response)

		// Act
		cloned := resp.Clone()
		actLines := []string{
			fmt.Sprintf("%d", cloned.Response),
			fmt.Sprintf("%v", cloned.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: NewTypedResponseResult
// ==========================================

func Test_NewTypedResponseResult_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultNewTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")

		attr := &coreapi.ResponseAttribute{IsValid: true, Message: "ok"}

		// Act
		result := coreapi.NewTypedResponseResult[string](attr, response)
		actLines := []string{
			result.Response,
			fmt.Sprintf("%v", result.IsValid()),
			fmt.Sprintf("%v", !result.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidTypedResponseResult
// ==========================================

func Test_InvalidTypedResponseResult_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultInvalidTestCases {
		// Arrange — nil attribute

		// Act
		result := coreapi.InvalidTypedResponseResult[string](nil)
		actLines := []string{
			fmt.Sprintf("%v", result.IsValid()),
			fmt.Sprintf("%v", result.IsInvalid()),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: GenericResponse alias
// ==========================================

func Test_GenericResponse_Alias_Verification(t *testing.T) {
	for caseIndex, testCase := range genericResponseAliasTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")

		attr := &coreapi.ResponseAttribute{IsValid: true}

		// Act — using GenericResponse which is TypedResponse[any]
		resp := coreapi.NewTypedResponse[any](attr, response)
		var gr *coreapi.GenericResponse = resp // type alias assignment
		actLines := []string{
			fmt.Sprintf("%v", gr.Response),
			fmt.Sprintf("%v", gr.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: GenericRequestIn alias
// ==========================================

func Test_GenericRequestIn_Alias_Verification(t *testing.T) {
	for caseIndex, testCase := range genericRequestInAliasTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}

		// Act — using GenericRequestIn which is TypedRequestIn[any]
		req := coreapi.NewTypedRequestIn[any](attr, payload)
		var gr *coreapi.GenericRequestIn = req // type alias assignment
		actLines := []string{
			fmt.Sprintf("%v", gr.Request),
			fmt.Sprintf("%v", gr.Attribute.IsValid),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidGenericResponse
// ==========================================

func Test_InvalidGenericResponse_Verification(t *testing.T) {
	for caseIndex, testCase := range invalidGenericResponseTestCases {
		// Arrange — nil attribute

		// Act
		resp := coreapi.InvalidGenericResponse(nil)
		actLines := []string{
			fmt.Sprintf("%v", resp.Attribute.IsValid),
			fmt.Sprintf("%v", resp.Attribute != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: InvalidGenericRequestIn
// ==========================================

func Test_InvalidGenericRequestIn_Verification(t *testing.T) {
	for caseIndex, testCase := range invalidGenericRequestInTestCases {
		// Arrange — nil attribute

		// Act
		req := coreapi.InvalidGenericRequestIn(nil)
		actLines := []string{
			fmt.Sprintf("%v", req.Attribute.IsValid),
			fmt.Sprintf("%v", req.Attribute != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: GenericResponseResult conversion
// ==========================================

func Test_GenericResponseResult_Conversion_Verification(t *testing.T) {
	for caseIndex, testCase := range genericResponseResultConversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")
		message, _ := input.GetAsString("message")

		attr := &coreapi.ResponseAttribute{IsValid: true, Message: message}

		// Act
		resp := coreapi.NewTypedResponse[string](attr, response)
		grr := resp.GenericResponseResult()
		actLines := []string{
			fmt.Sprintf("%v", grr.Attribute.IsValid),
			fmt.Sprintf("%v", grr.Response != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: SimpleGenericRequest conversion
// ==========================================

func Test_SimpleGenericRequest_Conversion_Verification(t *testing.T) {
	for caseIndex, testCase := range simpleGenericRequestConversionTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		payload, _ := input.GetAsString("payload")

		attr := &coreapi.RequestAttribute{IsValid: true}

		// Act
		req := coreapi.NewTypedRequestIn[string](attr, payload)
		sgr := req.SimpleGenericRequest(true, "")
		actLines := []string{
			fmt.Sprintf("%v", sgr.Attribute.IsValid),
			fmt.Sprintf("%v", sgr.Request != nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponseResult ClonePtr
// ==========================================

func Test_TypedResponseResult_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		response, _ := input.GetAsString("response")

		attr := &coreapi.ResponseAttribute{IsValid: true}
		result := coreapi.NewTypedResponseResult[string](attr, response)

		// Act
		cloned := result.ClonePtr()
		actLines := []string{
			cloned.Response,
			fmt.Sprintf("%v", cloned.IsValid()),
			fmt.Sprintf("%v", cloned != result),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Test: TypedResponseResult ClonePtr nil
// ==========================================

func Test_TypedResponseResult_ClonePtr_Nil_Verification(t *testing.T) {
	for caseIndex, testCase := range typedResponseResultCloneNilTestCases {
		// Arrange
		var result *coreapi.TypedResponseResult[string]

		// Act
		cloned := result.ClonePtr()
		actLines := []string{
			fmt.Sprintf("%v", cloned == nil),
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
