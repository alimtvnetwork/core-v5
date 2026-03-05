package coreoncetests

import (
	"errors"

	"gitlab.com/auk-go/core/coretests/args"
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

// Note: 9 fields — exceeds args.Six, kept as []string for accuracy.
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
			ExpectedInput: args.Five[string, string, string, string, string]{
				First:  "cached", // value1
				Second: "cached", // value2
				Third:  "true",   // value1EqValue2
				Fourth: "true",   // executeEqValue
				Fifth:  "1",      // callCount
			},
		},
		InitBytes: []byte("cached"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Value caches error result",
			ExpectedInput: args.Three[string, string, string]{
				First:  "",           // emptyValue
				Second: "true",       // hasError
				Third:  "test error", // errorMessage
			},
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
		ExpectedInput: "true", // executeEqValue
	},
	InitBytes: []byte("exec"),
}

var bytesErrorOnceValueOnlyTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title:         "BytesErrorOnce.ValueOnly returns bytes without error",
		ExpectedInput: "only", // valueOnlyResult
	},
	InitBytes: []byte("only"),
}

var bytesErrorOnceValueWithErrorTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.ValueWithError aliases Value",
		ExpectedInput: args.Two[string, string]{
			First:  "vwe",  // valueWithErrorResult
			Second: "true", // noError
		},
	},
	InitBytes: []byte("vwe"),
}

// =============================================================================
// BytesErrorOnce — Error State
// =============================================================================

var bytesErrorOnceErrorStateTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce with error — HasError, IsInvalid, IsFailed true",
			ExpectedInput: args.Six[string, string, string, string, string, string]{
				First:  "true",  // hasError
				Second: "false", // isValid
				Third:  "false", // isSuccess
				Fourth: "false", // isEmpty
				Fifth:  "true",  // isInvalid
				Sixth:  "true",  // isFailed
			},
		},
		InitErr: errors.New("fail"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce without error — HasError false, IsValid, IsSuccess true",
			ExpectedInput: args.Six[string, string, string, string, string, string]{
				First:  "false", // hasError
				Second: "true",  // isValid
				Third:  "true",  // isSuccess
				Fourth: "true",  // isEmpty
				Fifth:  "false", // isInvalid
				Sixth:  "false", // isFailed
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
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // hasIssuesOrEmpty
				Second: "false", // hasSafeItems
			},
		},
		InitBytes: []byte("data"),
		InitErr:   errors.New("err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce empty bytes — HasIssuesOrEmpty true",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // hasIssuesOrEmpty
				Second: "false", // hasSafeItems
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce data no error — HasIssuesOrEmpty false, HasSafeItems true",
			ExpectedInput: args.Two[string, string]{
				First:  "false", // hasIssuesOrEmpty
				Second: "true",  // hasSafeItems
			},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil receiver — HasIssuesOrEmpty true",
			ExpectedInput: args.Two[string, string]{
				First:  "true",  // hasIssuesOrEmpty
				Second: "false", // hasSafeItems
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
			ExpectedInput: args.Three[string, string, string]{
				First:  "str-val", // stringValue
				Second: "false",   // isStringEmpty
				Third:  "false",   // isStringEmptyOrWhitespace
			},
		},
		InitBytes: []byte("str-val"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil — String empty, IsStringEmpty true",
			ExpectedInput: args.Three[string, string, string]{
				First:  "",     // stringValue
				Second: "true", // isStringEmpty
				Third:  "true", // isStringEmptyOrWhitespace
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce whitespace — IsStringEmptyOrWhitespace true",
			ExpectedInput: args.Three[string, string, string]{
				First:  "   ",   // stringValue
				Second: "false", // isStringEmpty
				Third:  "true",  // isStringEmptyOrWhitespace
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
			ExpectedInput: args.Two[string, string]{
				First:  "true", // noError
				Second: "test", // deserializedName
			},
		},
		InitJson: `{"name":"test"}`,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize returns error when source has error",
			ExpectedInput: args.Three[string, string, string]{
				First:  "true", // hasSourceError
				Second: "true", // hasDeserializeError
				Third:  "true", // errorsMatch
			},
		},
		InitErr: errors.New("source error"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize with invalid JSON returns error",
			ExpectedInput: args.Two[string, string]{
				First:  "true", // hasError
				Second: "true", // isJsonError
			},
		},
		InitJson: "not-json",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.DeserializeMust succeeds without panic",
			ExpectedInput: args.Two[string, string]{
				First:  "false", // didPanic
				Second: "val",   // deserializedKey
			},
		},
		InitJson: `{"key":"val"}`,
		IsMust:   true,
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.DeserializeMust panics on error",
			ExpectedInput: "true", // didPanic
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
		Title: "BytesErrorOnce.MarshalJSON returns bytes",
		ExpectedInput: args.Two[string, string]{
			First:  "true",    // noError
			Second: `{"a":1}`, // marshaledValue
		},
	},
	InitBytes: []byte(`{"a":1}`),
}

var bytesErrorOnceSerializeTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.Serialize returns bytes",
		ExpectedInput: args.Two[string, string]{
			First:  "true", // noError
			Second: "ser",  // serializedValue
		},
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
			Title: "BytesErrorOnce — HandleError no panic, MustBeEmptyError no panic on success",
			ExpectedInput: args.Three[string, string, string]{
				First:  "false", // handleErrorPanicked
				Second: "false", // mustBeEmptyErrorPanicked
				Third:  "false", // mustHaveSafeItemsPanicked
			},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce — HandleError panics on error",
			ExpectedInput: args.Three[string, string, string]{
				First:  "true",  // handleErrorPanicked
				Second: "false", // mustBeEmptyErrorPanicked
				Third:  "false", // mustHaveSafeItemsPanicked
			},
		},
		InitErr: errors.New("handle-err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce — MustHaveSafeItems panics when empty",
			ExpectedInput: args.Three[string, string, string]{
				First:  "false", // handleErrorPanicked
				Second: "false", // mustBeEmptyErrorPanicked
				Third:  "true",  // mustHaveSafeItemsPanicked
			},
		},
		InitBytes: []byte{},
	},
}

var bytesErrorOnceInitializedTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.IsInitialized false before, true after Value call",
			ExpectedInput: args.Two[string, string]{
				First:  "false", // isInitializedBefore
				Second: "true",  // isInitializedAfter
			},
		},
		InitBytes: []byte("x"),
	},
}

var bytesErrorOnceSerializeMustTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.SerializeMust returns bytes without panic",
			ExpectedInput: args.Two[string, string]{
				First:  "false",    // didPanic
				Second: "must-ser", // serializedValue
			},
		},
		InitBytes: []byte("must-ser"),
	},
	{
		Case: coretestcases.CaseV1{
			Title:         "BytesErrorOnce.SerializeMust panics on error",
			ExpectedInput: "true", // didPanic
		},
		InitErr: errors.New("ser-fail"),
	},
}

var bytesErrorOnceConstructorTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewBytesErrorOnce (value) works correctly",
			ExpectedInput: args.Two[string, string]{
				First:  "val",  // value
				Second: "true", // isCorrect
			},
		},
		InitBytes: []byte("val"),
	},
}
