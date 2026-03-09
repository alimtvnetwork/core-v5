package errcoretests

import (
	"fmt"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// ConcatMessageWithErr nil passthrough test case
//
// Note: ConcatMessageWithErr is a package function, not a method.
// CaseNilSafe is inappropriate here — using CaseV1 instead.
// =============================================================================

var concatMessageWithErrNilPassthroughTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErr nil error returns nil",
	ArrangeInput: args.Map{
		"message": "should not appear",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var concatMessageWithErrNonNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErr non-nil error wraps message",
	ArrangeInput: args.Map{
		"message": "context:",
		"error":   "original error",
	},
	ExpectedInput: args.Map{
		"isNil":    false,
		"contains": true,
	},
}

var concatMessageWithErrWithStackTraceNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErrWithStackTrace nil error returns nil",
	ArrangeInput: args.Map{
		"message": "should not appear",
	},
	ExpectedInput: args.Map{
		"isNil": true,
	},
}

var concatMessageWithErrWithStackTraceNonNilTestCase = coretestcases.CaseV1{
	Title: "ConcatMessageWithErrWithStackTrace non-nil error wraps message",
	ArrangeInput: args.Map{
		"message": "context:",
		"error":   "original error",
	},
	ExpectedInput: args.Map{
		"isNil":    false,
		"contains": true,
	},
}

// =============================================================================
// Helper for constructing error from string
// =============================================================================

func errFromString(input args.Map) error {
	errStr, hasErr := input.GetAsString("error")

	if !hasErr {
		return nil
	}

	return fmt.Errorf("%s", errStr)
}
