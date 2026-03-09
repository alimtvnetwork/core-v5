package codefuncstests

import (
	"encoding/json"
	"strings"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// InOutErrFuncWrapperOf — Exec
// =============================================================================

var inOutErrOfExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns typed output on success",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": false,
			"name":         "strlen",
		},
		ExpectedInput: args.Map{
			"output":   5,
			"hasError": false,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": true,
			"name":         "strlen",
		},
		ExpectedInput: args.Map{
			"output":   0,
			"hasError": true,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

var inOutErrOfAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": true,
			"name":         "parse",
		},
		ExpectedInput: args.Map{
			"hasError":      true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": false,
			"name":         "parse",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// InOutErrFuncWrapperOf — ToLegacy
// =============================================================================

var inOutErrOfToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title: "ToLegacy preserves behavior on success",
		ArrangeInput: args.Map{
			"input":        "hello",
			"hasActionErr": false,
			"name":         "legacy-test",
		},
		ExpectedInput: args.Map{
			"output":   5,
			"hasError": false,
		},
	},
}

// =============================================================================
// InOutFuncWrapperOf — Exec
// =============================================================================

var inOutFuncOfExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns typed output",
		ArrangeInput: args.Map{
			"input": "hello",
			"name":  "upper",
		},
		ExpectedInput: args.Map{
			"output": "HELLO",
		},
	},
}

// =============================================================================
// InOutFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

var inOutFuncOfAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc always returns nil",
		ArrangeInput: args.Map{
			"input": "data",
			"name":  "process",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — Exec
// =============================================================================

var inActionReturnsErrOfExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns nil on success",
		ArrangeInput: args.Map{
			"input":        "valid@example.com",
			"hasActionErr": false,
			"name":         "validate-email",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"input":        "invalid",
			"hasActionErr": true,
			"name":         "validate-email",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

var inActionReturnsErrOfAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"input":        "bad-data",
			"hasActionErr": true,
			"name":         "validate",
		},
		ExpectedInput: args.Map{
			"hasError":      true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"input":        "good-data",
			"hasActionErr": false,
			"name":         "validate",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// InActionReturnsErrFuncWrapperOf — ToLegacy
// =============================================================================

var inActionReturnsErrOfToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title: "ToLegacy preserves error behavior",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": true,
			"name":         "legacy-validate",
		},
		ExpectedInput: args.Map{
			"hasError": true,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — Exec
// =============================================================================

var resultDelegatingOfExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns nil on success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "unmarshal-user",
		},
		ExpectedInput: args.Map{
			"hasError": false,
			"filled":   true,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "unmarshal-user",
		},
		ExpectedInput: args.Map{
			"hasError": true,
			"filled":   false,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

var resultDelegatingOfAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"hasActionErr": true,
			"name":         "decode",
		},
		ExpectedInput: args.Map{
			"hasError":      true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "decode",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// ResultDelegatingFuncWrapperOf — ToLegacy
// =============================================================================

var resultDelegatingOfToLegacyTestCases = []coretestcases.CaseV1{
	{
		Title: "ToLegacy preserves success behavior",
		ArrangeInput: args.Map{
			"hasActionErr": false,
			"name":         "legacy-decode",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// SerializeOutputFuncWrapperOf — Exec
// =============================================================================

var serializeOutputOfExecTestCases = []coretestcases.CaseV1{
	{
		Title: "Exec returns serialized bytes on success",
		ArrangeInput: args.Map{
			"input":        "test-value",
			"hasActionErr": false,
			"name":         "json-marshal",
		},
		ExpectedInput: args.Map{
			"hasError":  false,
			"hasOutput": true,
		},
	},
	{
		Title: "Exec returns error on failure",
		ArrangeInput: args.Map{
			"input":        "test-value",
			"hasActionErr": true,
			"name":         "json-marshal",
		},
		ExpectedInput: args.Map{
			"hasError":  true,
			"hasOutput": false,
		},
	},
}

// =============================================================================
// SerializeOutputFuncWrapperOf — AsActionReturnsErrorFunc
// =============================================================================

var serializeOutputOfAsActionReturnsErrorTestCases = []coretestcases.CaseV1{
	{
		Title: "AsActionReturnsErrorFunc wraps error with name",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": true,
			"name":         "marshal",
		},
		ExpectedInput: args.Map{
			"hasError":      true,
			"containsName": true,
		},
	},
	{
		Title: "AsActionReturnsErrorFunc returns nil on success",
		ArrangeInput: args.Map{
			"input":        "data",
			"hasActionErr": false,
			"name":         "marshal",
		},
		ExpectedInput: args.Map{
			"hasError": false,
		},
	},
}

// =============================================================================
// Helper factory functions for generic wrappers
// =============================================================================

func makeStrToIntErrFunc(hasErr bool) func(string) (int, error) {
	return func(s string) (int, error) {
		if hasErr {
			return 0, errTest
		}

		return len(s), nil
	}
}

func makeStrToStrFunc() func(string) string {
	return func(s string) string {
		return strings.ToUpper(s)
	}
}

func makeStrErrFunc(hasErr bool) func(string) error {
	return func(_ string) error {
		if hasErr {
			return errTest
		}

		return nil
	}
}

type fillTarget struct {
	Filled bool
}

func makeResultDelegatingFunc(hasErr bool) func(*fillTarget) error {
	return func(target *fillTarget) error {
		if hasErr {
			return errTest
		}

		target.Filled = true

		return nil
	}
}

func makeSerializeFunc(hasErr bool) func(string) ([]byte, error) {
	return func(s string) ([]byte, error) {
		if hasErr {
			return nil, errTest
		}

		return json.Marshal(s)
	}
}
