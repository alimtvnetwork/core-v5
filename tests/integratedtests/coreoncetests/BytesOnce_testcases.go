package coreoncetests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// BytesOnce -- Wrapper
// =============================================================================

type bytesOnceTestCase struct {
	Case       coretestcases.CaseV1
	InitBytes  []byte
	UseNilInit bool
}

// =============================================================================
// BytesOnce -- Core (Value, String, IsEmpty, Length, isNil)
// =============================================================================

var bytesOnceCoreTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce returns length 5 and isEmpty false -- 'hello' input",
			ExpectedInput: args.Map{
				"stringOfValue": "hello",
				"stringMethod":  "hello",
				"isEmpty":       false,
				"length":        5,
				"isNil":         false,
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce returns isEmpty true and isNil true -- nil input",
			ExpectedInput: args.Map{
				"stringOfValue": "",
				"stringMethod":  "",
				"isEmpty":       true,
				"length":        0,
				"isNil":         true,
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce returns isEmpty true and isNil false -- empty bytes input",
			ExpectedInput: args.Map{
				"stringOfValue": "",
				"stringMethod":  "",
				"isEmpty":       true,
				"length":        0,
				"isNil":         false,
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce returns isEmpty true -- nil initializer",
			ExpectedInput: args.Map{
				"stringOfValue": "",
				"stringMethod":  "",
				"isEmpty":       true,
				"length":        0,
				"isNil":         true,
			},
		},
		UseNilInit: true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.String returns 'test-string' -- 'test-string' input",
			ExpectedInput: args.Map{
				"stringOfValue": "test-string",
				"stringMethod":  "test-string",
				"isEmpty":       false,
				"length":        11,
				"isNil":         false,
			},
		},
		InitBytes: []byte("test-string"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce returns isEmpty false and length 1 -- 'x' input",
			ExpectedInput: args.Map{
				"stringOfValue": "x",
				"stringMethod":  "x",
				"isEmpty":       false,
				"length":        1,
				"isNil":         false,
			},
		},
		InitBytes: []byte("x"),
	},
}

// =============================================================================
// BytesOnce -- Caching (Value caches, Execute same as Value)
// =============================================================================

var bytesOnceCachingTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"r1":             "hello",
				"r2":             "hello",
				"r3":             "hello",
				"callCount":      1,
				"executeEqValue": true,
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.Execute returns same result as Value -- 'data' input",
			ExpectedInput: args.Map{
				"r1":             "data",
				"r2":             "data",
				"r3":             "data",
				"callCount":      1,
				"executeEqValue": true,
			},
		},
		InitBytes: []byte("data"),
	},
}

// =============================================================================
// BytesOnce -- JSON (MarshalJSON, UnmarshalJSON, Serialize)
// =============================================================================

type bytesOnceJsonTestCase struct {
	Case         coretestcases.CaseV1
	InitBytes    []byte
	ReplaceBytes []byte // non-nil triggers UnmarshalJSON test
}

var bytesOnceJsonTestCases = []bytesOnceJsonTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.MarshalJSON returns data without error -- 'hello' input",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitBytes: []byte("hello"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce.UnmarshalJSON overrides value to 'replaced'",
			ExpectedInput: args.Map{
				"noError":  true,
				"newValue": "replaced",
			},
		},
		InitBytes:    []byte("original"),
		ReplaceBytes: []byte("replaced"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesOnce 'serialize-me' -- Serialize returns JSON bytes",
			ExpectedInput: args.Map{
				"noError":             true,
				"dataLengthAboveZero": true,
			},
		},
		InitBytes: []byte("serialize-me"),
	},
}

// =============================================================================
// BytesOnce -- Constructor
// =============================================================================

var bytesOnceConstructorTestCases = []bytesOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewBytesOnce (value) works correctly",
			ExpectedInput: args.Map{
				"constructedValue": "val",
			},
		},
		InitBytes: []byte("val"),
	},
}
