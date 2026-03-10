package errcoretests

import (
	"strings"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/errcore"
)

// ==========================================
// ConcatMessageWithErr — nil passthrough (positive)
// ==========================================

func Test_ConcatMessageWithErr_NilPassthrough(t *testing.T) {
	tc := concatMessageWithErrNilPassthroughTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	msg, _ := input.GetAsString("message")

	// Act
	result := errcore.ConcatMessageWithErr(msg, nil)
	actual := args.Map{
		"isNil": result == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// ConcatMessageWithErr — non-nil wraps (negative)
// ==========================================

func Test_ConcatMessageWithErr_NonNilWraps(t *testing.T) {
	tc := concatMessageWithErrNonNilTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	msg, _ := input.GetAsString("message")
	inputErr := errFromString(input)

	// Act
	result := errcore.ConcatMessageWithErr(msg, inputErr)
	actual := args.Map{
		"isNil":    result == nil,
		"contains": result != nil && strings.Contains(result.Error(), "original error"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// ConcatMessageWithErrWithStackTrace — nil passthrough
// ==========================================

func Test_ConcatMessageWithErrWithStackTrace_NilPassthrough(t *testing.T) {
	tc := concatMessageWithErrWithStackTraceNilTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	msg, _ := input.GetAsString("message")

	// Act
	result := errcore.ConcatMessageWithErrWithStackTrace(msg, nil)
	actual := args.Map{
		"isNil": result == nil,
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}

// ==========================================
// ConcatMessageWithErrWithStackTrace — non-nil wraps
// ==========================================

func Test_ConcatMessageWithErrWithStackTrace_NonNilWraps(t *testing.T) {
	tc := concatMessageWithErrWithStackTraceNonNilTestCase

	// Arrange
	input := tc.ArrangeInput.(args.Map)
	msg, _ := input.GetAsString("message")
	inputErr := errFromString(input)

	// Act
	result := errcore.ConcatMessageWithErrWithStackTrace(msg, inputErr)
	actual := args.Map{
		"isNil":    result == nil,
		"contains": result != nil && strings.Contains(result.Error(), "original error"),
	}

	// Assert
	tc.ShouldBeEqualMapFirst(t, actual)
}
