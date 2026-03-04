package isanytests

// Extended test cases migrated from cmd/main/nullTesting01.go and nullTesting02.go.
// These cover typed-nil scenarios with error and *int that the original
// nullTestCases and definedTestCases do not exercise.

import (
	"errors"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/issetter"
)

var (
	nilError    error
	nilIntPtr   *int
	nilIntPtr2  *int
	liveError   = errors.New("")
	liveErrorX  = errors.New("x")

	// -------------------------------------------------------------------------
	// nullTesting01 — Defined and Null on typed nils (error, *int)
	// -------------------------------------------------------------------------

	extendedDefinedTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Defined on typed-nil error and *int — " +
					"nil literal is not defined, typed-nil error is not defined, " +
					"typed-nil *int is not defined, non-nil error IS defined.",
				ArrangeInput: []any{
					nil,
					liveError,
					nilError,
					nilIntPtr,
				},
				ExpectedInput: []string{
					"0 : false (value: <nil>, type: <nil>)",
					"1 : true (value: , type: *errors.errorString)",
					"2 : false (value: <nil>, type: <nil>)",
					"3 : false (value: <nil>, type: *int)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	extendedNullTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "Null on typed-nil error and *int — " +
					"nil literal is null, typed-nil error is null, " +
					"typed-nil *int is null, non-nil error is NOT null.",
				ArrangeInput: []any{
					nil,
					liveError,
					nilError,
					nilIntPtr,
				},
				ExpectedInput: []string{
					"0 : true (value: <nil>, type: <nil>)",
					"1 : false (value: , type: *errors.errorString)",
					"2 : true (value: <nil>, type: <nil>)",
					"3 : true (value: <nil>, type: *int)",
				},
				VerifyTypeOf: arrangeTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	// -------------------------------------------------------------------------
	// nullTesting02 — DefinedBoth and NullBoth with error and *int typed nils
	// -------------------------------------------------------------------------

	extendedDefinedBothTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "DefinedBoth with typed-nil error, *int, and live error — " +
					"migrated from nullTesting02.",
				ArrangeInput: []args.TwoAny{
					{
						First:  nil,
						Second: liveErrorX,
					},
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  nilIntPtr,
						Second: nilIntPtr2,
					},
					{
						First:  liveErrorX,
						Second: liveErrorX,
					},
					{
						First:  liveErrorX,
						Second: nilIntPtr,
					},
				},
				ExpectedInput: []string{
					"0 : false (<nil>, *errors.errorString)",
					"1 : false (<nil>, <nil>)",
					"2 : false (*int, *int)",
					"3 : true (*errors.errorString, *errors.errorString)",
					"4 : false (*errors.errorString, *int)",
				},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}

	extendedNullBothTestCases = []testWrapper{
		{
			BaseTestCase: coretests.BaseTestCase{
				Title: "NullBoth with typed-nil error, *int, and live error — " +
					"migrated from nullTesting02.",
				ArrangeInput: []args.TwoAny{
					{
						First:  nil,
						Second: liveErrorX,
					},
					{
						First:  nil,
						Second: nil,
					},
					{
						First:  nilIntPtr,
						Second: nilIntPtr2,
					},
					{
						First:  liveErrorX,
						Second: liveErrorX,
					},
					{
						First:  liveErrorX,
						Second: nilIntPtr,
					},
				},
				ExpectedInput: []string{
					"0 : false (<nil>, *errors.errorString)",
					"1 : true (<nil>, <nil>)",
					"2 : true (*int, *int)",
					"3 : false (*errors.errorString, *errors.errorString)",
					"4 : false (*errors.errorString, *int)",
				},
				VerifyTypeOf: twoArgsTypeVerification,
				IsEnable:     issetter.True,
			},
		},
	}
)
