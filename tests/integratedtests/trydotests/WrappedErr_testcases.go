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
		Title: "WrappedErr returns invalid and undefined -- nil receiver",
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
		Title: "WrappedErr returns no error no exception -- zero-value struct",
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
		Title: "WrappedErr returns hasErrorOrExc true -- error only, no exception",
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
		Title: "WrappedErr returns hasException true -- exception only, no error",
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
		Title: "WrappedErr returns isBothPresent true -- error and exception both set",
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
		Title: "WrappedErr returns isInvalidException true -- HasThrown true but nil Exception",
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
		Title: "WrappedErr returns empty strings -- nil receiver",
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
		Title: "WrappedErr returns error message -- error only, message 'something failed'",
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
		Title: "WrappedErr returns exception string -- exception only, value 'panic-data'",
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
		Title: "WrappedErr returns all empty strings -- zero-value struct",
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
		Title: "ExceptionType returns nil -- nil receiver",
		ArrangeInput: args.Map{
			"wrappedErr": (*trydo.WrappedErr)(nil),
		},
		ExpectedInput: args.Map{
			"isNilType": true,
		},
	},
	{
		Title: "ExceptionType returns nil -- invalid exception on zero-value struct",
		ArrangeInput: args.Map{
			"wrappedErr": &trydo.WrappedErr{},
		},
		ExpectedInput: args.Map{
			"isNilType": true,
		},
	},
	{
		Title: "ExceptionType returns string type -- string exception 'panic!'",
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
		Title: "ExceptionType returns int type -- int exception 42",
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
		Title: "ErrorFuncWrapPanic returns no error no exception -- func returns nil",
		ArrangeInput: args.Map{
			"func": func() error { return nil },
		},
		ExpectedInput: args.Map{
			"hasError":  false,
			"hasThrown": false,
		},
	},
	{
		Title: "ErrorFuncWrapPanic returns hasError true -- func returns error 'fail'",
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
		Title: "ErrorFuncWrapPanic returns hasThrown true -- func panics with 'kaboom'",
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
		Title: "WrapPanic returns nil exception -- no panic",
		ArrangeInput: args.Map{
			"func": func() {},
		},
		ExpectedInput: args.Map{
			"isNil": true,
		},
	},
	{
		Title: "WrapPanic returns exception value 'oops' -- func panics",
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
		Title: "Block.Do returns tryRan only -- no catch, no finally, no panic",
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
		Title: "Block.Do returns tryRan and finallyRan -- no catch, has finally, no panic",
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
		Title: "Block.Do returns all ran -- has catch, has finally, panics",
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
		Title: "Block.Do returns tryRan and catchRan -- has catch, no finally, panics",
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
