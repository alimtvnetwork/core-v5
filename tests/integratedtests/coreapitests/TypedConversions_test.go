package coreapitests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coreapi"
	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest IsValid / IsInvalid
// ==========================================

func Test_TypedSimpleGenericRequest_Validity(t *testing.T) {
	for caseIndex, testCase := range typedSimpleGenericRequestValidityTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		nilRequest, _ := input.GetAsBool("nilRequest")
		invalidAttribute, _ := input.GetAsBool("invalidAttribute")
		nilAttribute, _ := input.GetAsBool("nilAttribute")

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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Message / InvalidError
// ==========================================

func Test_TypedSimpleGenericRequest_Message(t *testing.T) {
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver — IsValid
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver_IsValid(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actLines := []string{fmt.Sprintf("%v", req.IsValid())}
	expectedLines := []string{"false"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver IsValid returns false", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver — IsInvalid
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver_IsInvalid(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actLines := []string{fmt.Sprintf("%v", req.IsInvalid())}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver IsInvalid returns true", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver — Message
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver_Message(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actLines := []string{req.Message()}
	expectedLines := []string{""}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver Message returns empty string", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Nil Receiver — InvalidError
// ==========================================

func Test_TypedSimpleGenericRequest_NilReceiver_InvalidError(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	actLines := []string{fmt.Sprintf("%v", req.InvalidError() == nil)}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver InvalidError returns nil", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying — IsValid/IsInvalid
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying_Validity(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest[string]("some-data", false, "validation failed")
	req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

	// Act
	actLines := []string{
		fmt.Sprintf("%v", req.IsValid()),
		fmt.Sprintf("%v", req.IsInvalid()),
	}
	expectedLines := []string{"false", "true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Valid attribute with invalid underlying request reports IsValid false", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying — Message
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying_Message(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest[string]("some-data", false, "field is required")
	req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

	// Act
	actLines := []string{req.Message()}
	expectedLines := []string{"field is required"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Valid attribute with invalid underlying request returns message", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying — InvalidError non-nil
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying_InvalidError(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest[string]("some-data", false, "input rejected")
	req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

	// Act
	actLines := []string{
		fmt.Sprintf("%v", req.InvalidError() == nil),
		req.InvalidError().Error(),
	}
	expectedLines := []string{"false", "input rejected"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Valid attribute with invalid underlying request returns non-nil InvalidError", actLines, expectedLines)
}

// ==========================================
// Test: TypedSimpleGenericRequest Invalid Underlying — empty message nil error
// ==========================================

func Test_TypedSimpleGenericRequest_InvalidUnderlying_EmptyMessage(t *testing.T) {
	// Arrange
	attr := &coreapi.RequestAttribute{IsValid: true}
	simpleReq := coredynamic.NewTypedSimpleRequest[string]("some-data", false, "")
	req := coreapi.NewTypedSimpleGenericRequest[string](attr, simpleReq)

	// Act
	actLines := []string{
		fmt.Sprintf("%v", req.InvalidError() == nil),
		req.Message(),
	}
	expectedLines := []string{"true", ""}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Valid attribute with invalid underlying request and empty message returns nil InvalidError", actLines, expectedLines)
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

func Test_TypedSimpleGenericRequest_Clone_Nil(t *testing.T) {
	// Arrange
	var req *coreapi.TypedSimpleGenericRequest[string]

	// Act
	cloned := req.Clone()
	actLines := []string{fmt.Sprintf("%v", cloned == nil)}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Clone on nil TypedSimpleGenericRequest returns nil", actLines, expectedLines)
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

func Test_TypedRequestIn_TypedSimpleGenericRequest_Nil(t *testing.T) {
	// Arrange
	var reqIn *coreapi.TypedRequestIn[string]

	// Act
	result := reqIn.TypedSimpleGenericRequest(true, "")
	actLines := []string{fmt.Sprintf("%v", result == nil)}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "TypedRequestIn.TypedSimpleGenericRequest on nil returns nil", actLines, expectedLines)
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
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

func Test_TypedResponse_TypedResponseResult_Nil(t *testing.T) {
	// Arrange
	var resp *coreapi.TypedResponse[string]

	// Act
	result := resp.TypedResponseResult()
	actLines := []string{fmt.Sprintf("%v", result == nil)}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "TypedResponse.TypedResponseResult on nil returns nil", actLines, expectedLines)
}
