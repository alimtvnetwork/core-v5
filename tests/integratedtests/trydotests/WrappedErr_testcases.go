package trydotests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/internal/trydo"
)

// ==========================================================================
// WrappedErr state inspection
// ==========================================================================

var wrappedErrStateTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil receiver is invalid and undefined",
		ArrangeInput: args.Map{
			"wrappedErr": (*trydo.WrappedErr)(nil),
		},
		ExpectedInput: args.Map{
			"isDefined":          false,
			"isInvalid":          true,
			"isInvalidException": true,
			"hasErrorOrExc":      false,
			"isBothPresent":      false,
			"hasException":       false,
		},
	},
	{
		Title: "Zero-value WrappedErr — no error, no exception",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{},
		},
		ExpectedInput: args.Map{
			"isDefined":          true,
			"isInvalid":          false,
			"isInvalidException": true,
			"hasErrorOrExc":      false,
			"isBothPresent":      false,
			"hasException":       false,
		},
	},
	{
		Title: "Error only — no exception",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Error:    errors.New("boom"),
				HasError: true,
			},
		},
		ExpectedInput: args.Map{
			"isDefined":          true,
			"isInvalid":          false,
			"isInvalidException": true,
			"hasErrorOrExc":      true,
			"isBothPresent":      false,
			"hasException":       false,
		},
	},
	{
		Title: "Exception only — no error",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Exception: "panic-value",
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"isDefined":          true,
			"isInvalid":          false,
			"isInvalidException": false,
			"hasErrorOrExc":      true,
			"isBothPresent":      false,
			"hasException":       true,
		},
	},
	{
		Title: "Both error and exception present",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Error:     errors.New("err"),
				HasError:  true,
				Exception: "exc",
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"isDefined":          true,
			"isInvalid":          false,
			"isInvalidException": false,
			"hasErrorOrExc":      true,
			"isBothPresent":      true,
			"hasException":       true,
		},
	},
	{
		Title: "HasThrown true but nil Exception is invalid exception",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				HasThrown: true,
				Exception: nil,
			},
		},
		ExpectedInput: args.Map{
			"isDefined":          true,
			"isInvalid":          false,
			"isInvalidException": true,
			"hasErrorOrExc":      true,
			"isBothPresent":      false,
			"hasException":       false,
		},
	},
}

// ==========================================================================
// WrappedErr string outputs
// ==========================================================================

var wrappedErrStringTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil receiver returns empty strings",
		ArrangeInput: args.Map{
			"wrappedErr": (*trydo.WrappedErr)(nil),
		},
		ExpectedInput: args.Map{
			"errorString":     "",
			"exceptionString": "",
			"string":          "",
		},
	},
	{
		Title: "Error only returns error string",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Error:    errors.New("something failed"),
				HasError: true,
			},
		},
		ExpectedInput: args.Map{
			"errorString":     "something failed",
			"exceptionString": "",
			"string":          "something failed",
		},
	},
	{
		Title: "Exception only returns exception string",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Exception: "panic-data",
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"errorString":       "",
			"hasExceptionValue": true,
			"hasStringValue":    true,
		},
	},
	{
		Title: "Zero-value returns all empty",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{},
		},
		ExpectedInput: args.Map{
			"errorString":     "",
			"exceptionString": "",
			"string":          "",
		},
	},
}

// ==========================================================================
// ExceptionType
// ==========================================================================

var wrappedErrExceptionTypeTestCases = []coretestcases.CaseV1{
	{
		Title: "Nil receiver returns nil type",
		ArrangeInput: args.Map{
			"wrappedErr": (*trydo.WrappedErr)(nil),
		},
		ExpectedInput: args.Map{
			"isNilType": true,
		},
	},
	{
		Title: "Invalid exception returns nil type",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{},
		},
		ExpectedInput: args.Map{
			"isNilType": true,
		},
	},
	{
		Title: "Valid string exception returns string type",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Exception: "panic!",
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"isNilType": false,
			"typeName":  "string",
		},
	},
	{
		Title: "Valid int exception returns int type",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{
				Exception: 42,
				HasThrown: true,
			},
		},
		ExpectedInput: args.Map{
			"isNilType": false,
			"typeName":  "int",
		},
	},
}

// ==========================================================================
// ErrorFuncWrapPanic — integration
// ==========================================================================

var errorFuncWrapPanicTestCases = []coretestcases.CaseV1{
	{
		Title: "Func returns nil error — no error, no exception",
		ArrangeInput: args.Map{
			"func": func() error { return nil },
		},
		ExpectedInput: args.Map{
			"hasError":  false,
			"hasThrown": false,
		},
	},
	{
		Title: "Func returns error — hasError true, no exception",
		ArrangeInput: args.Map{
			"func": func() error { return errors.New("fail") },
		},
		ExpectedInput: args.Map{
			"hasError":    true,
			"hasThrown":   false,
			"errorString": "fail",
		},
	},
	{
		Title: "Func panics — hasThrown true, exception captured",
		ArrangeInput: args.Map{
			"func": func() error { panic("kaboom") },
		},
		ExpectedInput: args.Map{
			"hasError":     false,
			"hasThrown":    true,
			"hasException": true,
		},
	},
}

// ==========================================================================
// WrapPanic
// ==========================================================================

var wrapPanicTestCases = []coretestcases.CaseV1{
	{
		Title: "No panic returns nil exception",
		ArrangeInput: args.Map{
			"func": func() {},
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "Panic returns exception value",
		ArrangeInput: args.Map{
			"func": func() { panic("oops") },
		},
		ExpectedInput: args.Map{
			"isNil": false,
			"value": "oops",
		},
	},
}

// ==========================================================================
// Block.Do
// ==========================================================================

var blockDoTestCases = []coretestcases.CaseV1{
	{
		Title: "Try executes without catch or finally",
		ArrangeInput: args.Map{
			"hasCatch":   false,
			"hasFinally": false,
			"panics":     false,
		},
		ExpectedInput: args.Map{
			"tryRan":     true,
			"catchRan":   false,
			"finallyRan": false,
		},
	},
	{
		Title: "Try + Finally both execute",
		ArrangeInput: args.Map{
			"hasCatch":   false,
			"hasFinally": true,
			"panics":     false,
		},
		ExpectedInput: args.Map{
			"tryRan":     true,
			"catchRan":   false,
			"finallyRan": true,
		},
	},
	{
		Title: "Panic triggers Catch and Finally",
		ArrangeInput: args.Map{
			"hasCatch":   true,
			"hasFinally": true,
			"panics":     true,
		},
		ExpectedInput: args.Map{
			"tryRan":     true,
			"catchRan":   true,
			"finallyRan": true,
		},
	},
	{
		Title: "Panic with Catch but no Finally",
		ArrangeInput: args.Map{
			"hasCatch":   true,
			"hasFinally": false,
			"panics":     true,
		},
		ExpectedInput: args.Map{
			"tryRan":     true,
			"catchRan":   true,
			"finallyRan": false,
		},
	},
}
