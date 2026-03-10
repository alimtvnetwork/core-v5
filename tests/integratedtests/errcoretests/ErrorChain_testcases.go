package errcoretests

import (
	"errors"
	"io"
	"os"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// Sentinel errors for chain verification
var (
	errSentinelA = errors.New("sentinel-a")
	errSentinelB = errors.New("sentinel-b")
	errSentinelC = errors.New("sentinel-c")
)

// =============================================================================
// MergeErrors + errors.Is
// =============================================================================

var mergeErrorsIsTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors preserves sentinel via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given sentinel among multiple errors",
			"sentinel": errSentinelA,
			"errors":   []error{errors.New("other"), errSentinelA, errors.New("another")},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "true",
		},
	},
	{
		Title: "MergeErrors single error matches itself via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given single sentinel",
			"sentinel": errSentinelB,
			"errors":   []error{errSentinelB},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "true",
		},
	},
	{
		Title: "MergeErrors does not match absent sentinel",
		ArrangeInput: args.Map{
			"when":     "given sentinel not in the list",
			"sentinel": errSentinelC,
			"errors":   []error{errSentinelA, errSentinelB},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsIsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr + errors.Is
// =============================================================================

var concatMessageErrorsIsTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr preserves sentinel via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given sentinel wrapped with message",
			"sentinel": errSentinelA,
			"message":  "context:",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"errorsIsOk":      "true",
			"containsMessage": "true",
		},
	},
	{
		Title: "ConcatMessageWithErr preserves io.EOF via errors.Is",
		ArrangeInput: args.Map{
			"when":     "given io.EOF wrapped",
			"sentinel": io.EOF,
			"message":  "read failed:",
		},
		ExpectedInput: args.Map{
			"hasError":        "true",
			"errorsIsOk":      "true",
			"containsMessage": "true",
		},
	},
}

// =============================================================================
// MergeErrors + errors.As
// =============================================================================

var mergeErrorsAsTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors preserves *os.PathError via errors.As",
		ArrangeInput: args.Map{
			"when": "given PathError among plain errors",
			"errors": []error{
				errors.New("plain"),
				&os.PathError{Op: "open", Path: "/tmp/test", Err: errors.New("not found")},
			},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "true",
		},
	},
	{
		Title: "MergeErrors errors.As returns false when type absent",
		ArrangeInput: args.Map{
			"when":   "given only plain errors",
			"errors": []error{errors.New("a"), errors.New("b")},
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr + errors.As
// =============================================================================

var concatMessageErrorsAsTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr preserves *os.PathError via errors.As",
		ArrangeInput: args.Map{
			"when":    "given PathError wrapped with message",
			"error":   &os.PathError{Op: "read", Path: "/etc/cfg", Err: errors.New("denied")},
			"message": "config load:",
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "true",
		},
	},
	{
		Title: "ConcatMessageWithErr errors.As false for plain error",
		ArrangeInput: args.Map{
			"when":    "given plain error",
			"error":   errors.New("simple"),
			"message": "wrap:",
		},
		ExpectedInput: args.Map{
			"hasError":   "true",
			"errorsAsOk": "false",
		},
	},
}

// =============================================================================
// ConcatMessageWithErr nil passthrough
// =============================================================================

var concatMessageNilTestCases = []coretestcases.CaseV1{
	{
		Title: "ConcatMessageWithErr nil returns nil",
		ArrangeInput: args.Map{
			"when":    "given nil error",
			"message": "should not appear",
		},
		ExpectedInput: "true",
	},
}

// =============================================================================
// MergeErrors multiple sentinels
// =============================================================================

var mergeErrorsMultiSentinelTestCases = []coretestcases.CaseV1{
	{
		Title: "MergeErrors all three sentinels found via errors.Is",
		ArrangeInput: args.Map{
			"when":      "given three distinct sentinels",
			"errors":    []error{errSentinelA, errSentinelB, errSentinelC},
			"sentinels": []error{errSentinelA, errSentinelB, errSentinelC},
		},
		ExpectedInput: args.Map{
			"hasError":       "true",
			"allSentinelsOk": "true",
		},
	},
	{
		Title: "MergeErrors partial sentinels — missing one fails",
		ArrangeInput: args.Map{
			"when":      "given two of three sentinels",
			"errors":    []error{errSentinelA, errSentinelB},
			"sentinels": []error{errSentinelA, errSentinelB, errSentinelC},
		},
		ExpectedInput: args.Map{
			"hasError":       "true",
			"allSentinelsOk": "false",
		},
	},
}
