package coreoncetests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// BytesOnce — Wrapper
// =============================================================================

type bytesOnceTestCase struct {
	Case       coretestcases.CaseV1
	InitBytes  []byte
	UseNilInit bool
}

// =============================================================================
// BytesOnce — Core (Value, String, IsEmpty, Length, isNil)
// =============================================================================

var bytesOnceCoreTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'hello' — Value, String, IsEmpty false, Length 5",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "hello", // stringOfValue
				Second: "hello", // stringMethod
				Third:  "false", // isEmpty
				Fourth: "5",     // length
				Fifth:  "false", // isNil
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce nil — IsEmpty true, Length 0, isNil true",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "",     // stringOfValue
				Second: "",     // stringMethod
				Third:  "true", // isEmpty
				Fourth: "0",    // length
				Fifth:  "true", // isNil
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce empty — IsEmpty true, Length 0, not nil",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "",      // stringOfValue
				Second: "",      // stringMethod
				Third:  "true",  // isEmpty
				Fourth: "0",     // length
				Fifth:  "false", // isNil (empty slice is not nil)
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce nil initializer — Length 0, IsEmpty true",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "",     // stringOfValue
				Second: "",     // stringMethod
				Third:  "true", // isEmpty
				Fourth: "0",    // length
				Fifth:  "true", // isNil
			},
		},
		UseNilInit: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'test-string' — String returns correct conversion",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "test-string", // stringOfValue
				Second: "test-string", // stringMethod
				Third:  "false",       // isEmpty
				Fourth: "11",          // length
				Fifth:  "false",       // isNil
			},
		},
		InitBytes: []byte("test-string"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'x' — IsEmpty false",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "x",     // stringOfValue
				Second: "x",     // stringMethod
				Third:  "false", // isEmpty
				Fourth: "1",     // length
				Fifth:  "false", // isNil
			},
		},
		InitBytes: []byte("x"),
	},
}

// =============================================================================
// BytesOnce — Caching (Value caches, Execute same as Value)
// =============================================================================

var bytesOnceCachingTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.Value caches — initializer runs exactly once",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "hello", // r1
				Second: "hello", // r2
				Third:  "hello", // r3
				Fourth: "1",     // callCount
				Fifth:  "true",  // executeEqValue
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.Execute returns same result as Value for 'data'",
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "data", // r1
				Second: "data", // r2
				Third:  "data", // r3
				Fourth: "1",    // callCount
				Fifth:  "true", // executeEqValue
			},
		},
		InitBytes: []byte("data"),
	},
}

// =============================================================================
// BytesOnce — JSON (MarshalJSON, UnmarshalJSON, Serialize)
// =============================================================================

type bytesOnceJsonTestCase struct {
	Case         coretestcases.CaseV1
	InitBytes    []byte
	ReplaceBytes []byte // non-nil triggers UnmarshalJSON test
}

var bytesOnceJsonTestCases = []bytesOnceJsonTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'hello' — MarshalJSON succeeds with data",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "true", // dataLengthAboveZero
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.UnmarshalJSON overrides value to 'replaced'",
			ExpectedInput: args.Two[string, string]{
				First:  "true",     // noError
				Second: "replaced", // newValue
			},
		},
		InitBytes:    []byte("original"),
		ReplaceBytes: []byte("replaced"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'serialize-me' — Serialize returns JSON bytes",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "true", // dataLengthAboveZero
			},
		},
		InitBytes: []byte("serialize-me"),
	},
}

// =============================================================================
// BytesOnce — Constructor
// =============================================================================

var bytesOnceConstructorTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "NewBytesOnce (value) works correctly",
			ExpectedInput: "val", // constructedValue
		},
		InitBytes: []byte("val"),
	},
}
