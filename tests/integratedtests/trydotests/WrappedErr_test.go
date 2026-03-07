package trydotests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/internal/trydo"
)

func Test_WrappedErr_State(t *testing.T) {
	for caseIndex, testCase := range wrappedErrStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		we := input["wrappedErr"].(*trydo.WrappedErr)

		// Act
		actual := args.Map{
			"isDefined":          we.IsDefined(),
			"isInvalid":          we.IsInvalid(),
			"isInvalidException": we.IsInvalidException(),
			"hasErrorOrExc":      we.HasErrorOrException(),
			"isBothPresent":      we.IsBothPresent(),
			"hasException":       we.HasException(),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_WrappedErr_Strings(t *testing.T) {
	for caseIndex, testCase := range wrappedErrStringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		we := input["wrappedErr"].(*trydo.WrappedErr)

		// Act
		errStr := we.ErrorString()
		excStr := we.ExceptionString()
		str := we.String()

		// Assert
		expected := testCase.ExpectedInput.(args.Map)

		actual := args.Map{}

		if _, has := expected["errorString"]; has {
			actual["errorString"] = errStr
		}

		if _, has := expected["exceptionString"]; has {
			actual["exceptionString"] = excStr
		}

		if _, has := expected["string"]; has {
			actual["string"] = str
		}

		if _, has := expected["hasExceptionValue"]; has {
			actual["hasExceptionValue"] = excStr != ""
		}

		if _, has := expected["hasStringValue"]; has {
			actual["hasStringValue"] = str != ""
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_WrappedErr_ExceptionType(t *testing.T) {
	for caseIndex, testCase := range wrappedErrExceptionTypeTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		we := input["wrappedErr"].(*trydo.WrappedErr)

		// Act
		excType := we.ExceptionType()

		// Assert
		actual := args.Map{
			"isNilType": excType == nil,
		}

		if excType != nil {
			actual["typeName"] = excType.Name()
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_ErrorFuncWrapPanic(t *testing.T) {
	for caseIndex, testCase := range errorFuncWrapPanicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		fn := input["func"].(func() error)

		// Act
		result := trydo.ErrorFuncWrapPanic(fn)

		// Assert
		expected := testCase.ExpectedInput.(args.Map)

		actual := args.Map{
			"hasError":  result.HasError,
			"hasThrown": result.HasThrown,
		}

		if _, has := expected["errorString"]; has {
			actual["errorString"] = result.ErrorString()
		}

		if _, has := expected["hasException"]; has {
			actual["hasException"] = result.HasException()
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_WrapPanic(t *testing.T) {
	for caseIndex, testCase := range wrapPanicTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		fn := input["func"].(func())

		// Act
		exception := trydo.WrapPanic(fn)

		// Assert
		actual := args.Map{
			"isNil": exception == nil,
		}

		if exception != nil {
			actual["value"] = fmt.Sprintf("%v", exception)
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Block_Do(t *testing.T) {
	for caseIndex, testCase := range blockDoTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		hasCatch := input["hasCatch"].(bool)
		hasFinally := input["hasFinally"].(bool)
		panics := input["panics"].(bool)

		tryRan := false
		catchRan := false
		finallyRan := false

		block := trydo.Block{
			Try: func() {
				tryRan = true

				if panics {
					panic("test-panic")
				}
			},
		}

		if hasCatch {
			block.Catch = func(e trydo.Exception) {
				catchRan = true
			}
		}

		if hasFinally {
			block.Finally = func() {
				finallyRan = true
			}
		}

		// Act
		block.Do()

		// Assert
		actual := args.Map{
			"tryRan":     tryRan,
			"catchRan":   catchRan,
			"finallyRan": finallyRan,
		}

		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}
