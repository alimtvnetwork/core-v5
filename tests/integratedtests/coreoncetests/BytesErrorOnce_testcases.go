package coreoncetests

import (
	"errors"

	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// =============================================================================
// BytesErrorOnce — Wrapper
// =============================================================================

type bytesErrorOnceTestCase struct {
	Case          coretestcases.CaseV1
	InitBytes     []byte
	InitErr       error
	IsNilReceiver bool
}

// =============================================================================
// BytesErrorOnce — Core (Value, Length, IsEmpty, IsNull, IsDefined, HasAnyItem)
// =============================================================================

var bytesErrorOnceCoreTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce 'abc' — Length 3, not empty, not null, isDefined",
			ExpectedInput: []string{
				"abc", "true", "3", "true", "false", "false", "false", "false", "true",
			},
		},
		InitBytes: []byte("abc"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil/nil — empty, null, not defined",
			ExpectedInput: []string{
				"", "true", "0", "false", "true", "true", "true", "true", "false",
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce empty bytes/nil — empty, not null",
			ExpectedInput: []string{
				"", "true", "0", "false", "true", "true", "true", "false", "false",
			},
		},
		InitBytes: []byte{},
	},
}

// =============================================================================
// BytesErrorOnce — Caching
// =============================================================================

var bytesErrorOnceCachingTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Value caches — initializer runs exactly once",
			ExpectedInput: []string{"cached", "cached", "true", "true", "1"},
		},
		InitBytes: []byte("cached"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.Value caches error result",
			ExpectedInput: []string{"", "true", "test error"},
		},
		InitErr: errors.New("test error"),
	},
}

// =============================================================================
// BytesErrorOnce — Execute / ValueOnly / ValueWithError
// =============================================================================

var bytesErrorOnceExecuteTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.Execute returns same as Value",
		ExpectedInput: []string{"true"},
	},
	InitBytes: []byte("exec"),
}

var bytesErrorOnceValueOnlyTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.ValueOnly returns bytes without error",
		ExpectedInput: []string{"only"},
	},
	InitBytes: []byte("only"),
}

var bytesErrorOnceValueWithErrorTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.ValueWithError aliases Value",
		ExpectedInput: []string{"vwe", "true"},
	},
	InitBytes: []byte("vwe"),
}

// =============================================================================
// BytesErrorOnce — Error State
// =============================================================================

var bytesErrorOnceErrorStateTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce with error — HasError, IsInvalid, IsFailed true",
			ExpectedInput: []string{"true", "false", "false", "false", "true", "true"},
		},
		InitErr: errors.New("fail"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce without error — HasError false, IsValid, IsSuccess true",
			ExpectedInput: []string{"false", "true", "true", "true", "false", "false"},
		},
		InitBytes: []byte("ok"),
	},
}

// =============================================================================
// BytesErrorOnce — HasIssuesOrEmpty / HasSafeItems
// =============================================================================

var bytesErrorOnceHasIssuesTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce with error — HasIssuesOrEmpty true",
			ExpectedInput: []string{"true", "false"},
		},
		InitBytes: []byte("data"),
		InitErr:   errors.New("err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce empty bytes — HasIssuesOrEmpty true",
			ExpectedInput: []string{"true", "false"},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce data no error — HasIssuesOrEmpty false, HasSafeItems true",
			ExpectedInput: []string{"false", "true"},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce nil receiver — HasIssuesOrEmpty true",
			ExpectedInput: []string{"true", "false"},
		},
		IsNilReceiver: true,
	},
}

// =============================================================================
// BytesErrorOnce — String
// =============================================================================

var bytesErrorOnceStringTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce 'str-val' — String returns correct value",
			ExpectedInput: []string{"str-val", "false", "false"},
		},
		InitBytes: []byte("str-val"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce nil — String empty, IsStringEmpty true",
			ExpectedInput: []string{"", "true", "true"},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce whitespace — IsStringEmptyOrWhitespace true",
			ExpectedInput: []string{"   ", "false", "true"},
		},
		InitBytes: []byte("   "),
	},
}

// =============================================================================
// BytesErrorOnce — Deserialize
// =============================================================================

type bytesErrorOnceDeserializeTestCase struct {
	Case     coretestcases.CaseV1
	InitJson string
	InitErr  error
	IsMust   bool
}

var bytesErrorOnceDeserializeTestCases = []bytesErrorOnceDeserializeTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.Deserialize succeeds with valid JSON",
			ExpectedInput: []string{"true", "test"},
		},
		InitJson: `{"name":"test"}`,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.Deserialize returns error when source has error",
			ExpectedInput: []string{"true", "true", "true"},
		},
		InitErr: errors.New("source error"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.Deserialize with invalid JSON returns error",
			ExpectedInput: []string{"true", "true"},
		},
		InitJson: "not-json",
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.DeserializeMust succeeds without panic",
			ExpectedInput: []string{"false", "val"},
		},
		InitJson: `{"key":"val"}`,
		IsMust:   true,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.DeserializeMust panics on error",
			ExpectedInput: []string{"true"},
		},
		InitErr: errors.New("must-fail"),
		IsMust:  true,
	},
}

// =============================================================================
// BytesErrorOnce — Serialization (MarshalJSON, Serialize, SerializeMust)
// =============================================================================

var bytesErrorOnceMarshalJSONTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.MarshalJSON returns bytes",
		ExpectedInput: []string{"true", `{"a":1}`},
	},
	InitBytes: []byte(`{"a":1}`),
}

var bytesErrorOnceSerializeTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.Serialize returns bytes",
		ExpectedInput: []string{"true", "ser"},
	},
	InitBytes: []byte("ser"),
}

// =============================================================================
// BytesErrorOnce — Lifecycle (panic checks, IsInitialized, constructor)
// =============================================================================

type bytesErrorOnceLifecycleTestCase struct {
	Case      coretestcases.CaseV1
	InitBytes []byte
	InitErr   error
}

var bytesErrorOnceLifecycleTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce — HandleError no panic, MustBeEmptyError no panic on success",
			ExpectedInput: []string{"false", "false", "false"},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce — HandleError panics on error",
			ExpectedInput: []string{"true", "false", "false"},
		},
		InitErr: errors.New("handle-err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce — MustHaveSafeItems panics when empty",
			ExpectedInput: []string{"false", "false", "true"},
		},
		InitBytes: []byte{},
	},
}

var bytesErrorOnceInitializedTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.IsInitialized false before, true after Value call",
			ExpectedInput: []string{"false", "true"},
		},
		InitBytes: []byte("x"),
	},
}

var bytesErrorOnceSerializeMustTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.SerializeMust returns bytes without panic",
			ExpectedInput: []string{"false", "must-ser"},
		},
		InitBytes: []byte("must-ser"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.SerializeMust panics on error",
			ExpectedInput: []string{"true"},
		},
		InitErr: errors.New("ser-fail"),
	},
}

var bytesErrorOnceConstructorTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title:         "NewBytesErrorOnce (value) works correctly",
			ExpectedInput: []string{"val", "true"},
		},
		InitBytes: []byte("val"),
	},
}
