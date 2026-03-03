package coreoncetests

import "gitlab.com/auk-go/core/coretests/coretestcases"

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
			ExpectedInput: []string{
				"hello", // string(Value)
				"hello", // String
				"false", // IsEmpty
				"5",     // Length
				"false", // Value == nil
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce nil — IsEmpty true, Length 0, isNil true",
			ExpectedInput: []string{
				"",     // string(Value)
				"",     // String
				"true", // IsEmpty
				"0",    // Length
				"true", // Value == nil
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce empty — IsEmpty true, Length 0, not nil",
			ExpectedInput: []string{
				"",      // string(Value)
				"",      // String
				"true",  // IsEmpty
				"0",     // Length
				"false", // Value == nil (empty slice is not nil)
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce nil initializer — Length 0, IsEmpty true",
			ExpectedInput: []string{
				"",     // string(Value)
				"",     // String
				"true", // IsEmpty
				"0",    // Length
				"true", // Value == nil
			},
		},
		UseNilInit: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'test-string' — String returns correct conversion",
			ExpectedInput: []string{
				"test-string",
				"test-string",
				"false",
				"11",
				"false",
			},
		},
		InitBytes: []byte("test-string"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'x' — IsEmpty false",
			ExpectedInput: []string{
				"x",
				"x",
				"false",
				"1",
				"false",
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
			ExpectedInput: []string{
				"hello", // r1
				"hello", // r2
				"hello", // r3
				"1",     // callCount
				"true",  // Execute == Value
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.Execute returns same result as Value for 'data'",
			ExpectedInput: []string{
				"data", // r1
				"data", // r2
				"data", // r3
				"1",    // callCount
				"true", // Execute == Value
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
			ExpectedInput: []string{
				"true", // no error
				"true", // data length > 0
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.UnmarshalJSON overrides value to 'replaced'",
			ExpectedInput: []string{
				"true",     // no error
				"replaced", // new value
			},
		},
		InitBytes:    []byte("original"),
		ReplaceBytes: []byte("replaced"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'serialize-me' — Serialize returns JSON bytes",
			ExpectedInput: []string{
				"true", // no error
				"true", // data length > 0
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
			Title: "NewBytesOnce (value) works correctly",
			ExpectedInput: []string{
				"val",
			},
		},
		InitBytes: []byte("val"),
	},
}
