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
				"abc",   // string(Value)
				"true",  // err == nil
				"3",     // Length
				"true",  // HasAnyItem
				"false", // IsEmpty
				"false", // IsEmptyBytes
				"false", // IsBytesEmpty
				"false", // IsNull
				"true",  // IsDefined
			},
		},
		InitBytes: []byte("abc"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil/nil — empty, null, not defined",
			ExpectedInput: []string{
				"",      // string(Value)
				"true",  // err == nil
				"0",     // Length
				"false", // HasAnyItem
				"true",  // IsEmpty
				"true",  // IsEmptyBytes
				"true",  // IsBytesEmpty
				"true",  // IsNull
				"false", // IsDefined
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce empty bytes/nil — empty, not null",
			ExpectedInput: []string{
				"",      // string(Value)
				"true",  // err == nil
				"0",     // Length
				"false", // HasAnyItem
				"true",  // IsEmpty
				"true",  // IsEmptyBytes
				"true",  // IsBytesEmpty
				"false", // IsNull
				"false", // IsDefined
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
			ExpectedInput: []string{
				"cached", // r1
				"cached", // r2
				"true",   // e1 == nil
				"true",   // e2 == nil
				"1",      // callCount
			},
		},
		InitBytes: []byte("cached"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Value caches error result",
			ExpectedInput: []string{
				"",          // string(Value)
				"true",      // val == nil
				"test error", // err.Error()
			},
		},
		InitErr: errors.New("test error"),
	},
}

// =============================================================================
// BytesErrorOnce — Execute / ValueOnly / ValueWithError
// =============================================================================

var bytesErrorOnceAccessTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Execute returns same as Value",
			ExpectedInput: []string{
				"true", // Execute == Value
			},
		},
		InitBytes: []byte("exec"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.ValueOnly returns bytes without error",
			ExpectedInput: []string{
				"only",
			},
		},
		InitBytes: []byte("only"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.ValueWithError aliases Value",
			ExpectedInput: []string{
				"vwe",  // value
				"true", // err == nil
			},
		},
		InitBytes: []byte("vwe"),
	},
}

// =============================================================================
// BytesErrorOnce — Error State
// =============================================================================

var bytesErrorOnceErrorStateTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce with error — HasError, IsInvalid, IsFailed true",
			ExpectedInput: []string{
				"true",  // HasError
				"false", // IsEmptyError
				"false", // IsValid
				"false", // IsSuccess
				"true",  // IsInvalid
				"true",  // IsFailed
			},
		},
		InitErr: errors.New("fail"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce without error — HasError false, IsValid, IsSuccess true",
			ExpectedInput: []string{
				"false", // HasError
				"true",  // IsEmptyError
				"true",  // IsValid
				"true",  // IsSuccess
				"false", // IsInvalid
				"false", // IsFailed
			},
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
			Title: "BytesErrorOnce with error — HasIssuesOrEmpty true",
			ExpectedInput: []string{
				"true",  // HasIssuesOrEmpty
				"false", // HasSafeItems
			},
		},
		InitBytes: []byte("data"),
		InitErr:   errors.New("err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce empty bytes — HasIssuesOrEmpty true",
			ExpectedInput: []string{
				"true",  // HasIssuesOrEmpty
				"false", // HasSafeItems
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce data no error — HasIssuesOrEmpty false, HasSafeItems true",
			ExpectedInput: []string{
				"false", // HasIssuesOrEmpty
				"true",  // HasSafeItems
			},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil receiver — HasIssuesOrEmpty true",
			ExpectedInput: []string{
				"true",  // HasIssuesOrEmpty
				"false", // HasSafeItems
			},
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
			Title: "BytesErrorOnce 'str-val' — String returns correct value",
			ExpectedInput: []string{
				"str-val", // String
				"false",   // IsStringEmpty
				"false",   // IsStringEmptyOrWhitespace
			},
		},
		InitBytes: []byte("str-val"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil — String empty, IsStringEmpty true",
			ExpectedInput: []string{
				"",     // String
				"true", // IsStringEmpty
				"true", // IsStringEmptyOrWhitespace
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce whitespace — IsStringEmptyOrWhitespace true",
			ExpectedInput: []string{
				"   ",   // String
				"false", // IsStringEmpty
				"true",  // IsStringEmptyOrWhitespace
			},
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
			Title: "BytesErrorOnce.Deserialize succeeds with valid JSON",
			ExpectedInput: []string{
				"true", // no error
				"test", // result["name"]
			},
		},
		InitJson: `{"name":"test"}`,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize returns error when source has error",
			ExpectedInput: []string{
				"true", // has error
				"true", // contains 'existing error cannot deserialize'
				"true", // contains 'source error'
			},
		},
		InitErr: errors.New("source error"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize with invalid JSON returns error",
			ExpectedInput: []string{
				"true", // has error
				"true", // contains 'deserialize failed'
			},
		},
		InitJson: "not-json",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.DeserializeMust succeeds without panic",
			ExpectedInput: []string{
				"false", // no panic
				"val",   // result["key"]
			},
		},
		InitJson: `{"key":"val"}`,
		IsMust:   true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.DeserializeMust panics on error",
			ExpectedInput: []string{
				"true", // panics
			},
		},
		InitErr: errors.New("must-fail"),
		IsMust:  true,
	},
}

// =============================================================================
// BytesErrorOnce — Serialization (MarshalJSON, Serialize, SerializeMust)
// =============================================================================

var bytesErrorOnceSerializationTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.MarshalJSON returns bytes",
			ExpectedInput: []string{
				"true",      // no error
				`{"a":1}`,   // data string
			},
		},
		InitBytes: []byte(`{"a":1}`),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Serialize returns bytes",
			ExpectedInput: []string{
				"true", // no error
				"ser",  // data string
			},
		},
		InitBytes: []byte("ser"),
	},
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
			Title: "BytesErrorOnce — HandleError no panic, MustBeEmptyError no panic on success",
			ExpectedInput: []string{
				"false", // HandleError panics
				"false", // MustBeEmptyError panics
				"false", // MustHaveSafeItems panics
			},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce — HandleError panics on error",
			ExpectedInput: []string{
				"true",  // HandleError panics
				"false", // (not tested, early concept)
				"false", // (not tested)
			},
		},
		InitErr: errors.New("handle-err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce — MustHaveSafeItems panics when empty",
			ExpectedInput: []string{
				"false", // HandleError panics
				"false", // MustBeEmptyError panics
				"true",  // MustHaveSafeItems panics
			},
		},
		InitBytes: []byte{},
	},
}

var bytesErrorOnceInitializedTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.IsInitialized false before, true after Value call",
			ExpectedInput: []string{
				"false", // before
				"true",  // after
			},
		},
		InitBytes: []byte("x"),
	},
}

var bytesErrorOnceSerializeMustTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.SerializeMust returns bytes without panic",
			ExpectedInput: []string{
				"false",    // panics
				"must-ser", // result
			},
		},
		InitBytes: []byte("must-ser"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.SerializeMust panics on error",
			ExpectedInput: []string{
				"true", // panics
			},
		},
		InitErr: errors.New("ser-fail"),
	},
}

var bytesErrorOnceConstructorTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewBytesErrorOnce (value) works correctly",
			ExpectedInput: []string{
				"val",  // string(Value)
				"true", // err == nil
			},
		},
		InitBytes: []byte("val"),
	},
}
