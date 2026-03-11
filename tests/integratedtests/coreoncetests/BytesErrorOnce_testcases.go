package coreoncetests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// BytesErrorOnce -- Wrapper
// =============================================================================

type bytesErrorOnceTestCase struct {
	Case          coretestcases.CaseV1
	InitBytes     []byte
	InitErr       error
	IsNilReceiver bool
}

// =============================================================================
// BytesErrorOnce -- Core (Value, Length, IsEmpty, IsNull, IsDefined, HasAnyItem)
// =============================================================================

var bytesErrorOnceCoreTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns length 3 and isDefined true -- 'abc' input",
			ExpectedInput: args.Map{
				"stringValue":  "abc",
				"noError":      true,
				"length":       3,
				"hasAnyItem":   true,
				"isEmpty":      false,
				"isEmptyBytes": false,
				"isBytesEmpty": false,
				"isNull":       false,
				"isDefined":    true,
			},
		},
		InitBytes: []byte("abc"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns isEmpty true and isDefined false -- nil bytes nil error",
			ExpectedInput: args.Map{
				"stringValue":  "",
				"noError":      true,
				"length":       0,
				"hasAnyItem":   false,
				"isEmpty":      true,
				"isEmptyBytes": true,
				"isBytesEmpty": true,
				"isNull":       true,
				"isDefined":    false,
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns isNull false and isDefined true -- empty bytes nil error",
			ExpectedInput: args.Map{
				"stringValue":  "",
				"noError":      true,
				"length":       0,
				"hasAnyItem":   false,
				"isEmpty":      false,
				"isEmptyBytes": true,
				"isBytesEmpty": true,
				"isNull":       false,
				"isDefined":    true,
			},
		},
		InitBytes: []byte{},
	},
}

// =============================================================================
// BytesErrorOnce -- Caching
// =============================================================================

var bytesErrorOnceCachingTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Value caches -- initializer runs exactly once",
			ExpectedInput: args.Map{
				"value1":         "cached",
				"value2":         "cached",
				"value1EqValue2": true,
				"executeEqValue": true,
				"callCount":      1,
			},
		},
		InitBytes: []byte("cached"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Value returns cached error -- error 'test error'",
			ExpectedInput: args.Map{
				"emptyValue":   "",
				"hasError":     true,
				"errorMessage": "test error",
			},
		},
		InitErr: errors.New("test error"),
	},
}

// =============================================================================
// BytesErrorOnce -- Execute / ValueOnly / ValueWithError
// =============================================================================

var bytesErrorOnceExecuteTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.Execute returns same result as Value -- 'exec' input",
		ExpectedInput: args.Map{
			"executeEqValue": true,
		},
	},
	InitBytes: []byte("exec"),
}

var bytesErrorOnceValueOnlyTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.ValueOnly returns bytes without error -- 'only' input",
		ExpectedInput: args.Map{
			"valueOnlyResult": "only",
		},
	},
	InitBytes: []byte("only"),
}

var bytesErrorOnceValueWithErrorTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.ValueWithError returns same as Value -- 'vwe' input",
		ExpectedInput: args.Map{
			"valueWithErrorResult": "vwe",
			"noError":              true,
		},
	},
	InitBytes: []byte("vwe"),
}

// =============================================================================
// BytesErrorOnce -- Error State
// =============================================================================

var bytesErrorOnceErrorStateTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns hasError true and isFailed true -- error 'fail'",
			ExpectedInput: args.Map{
				"hasError":  true,
				"isValid":   false,
				"isSuccess": false,
				"isEmpty":   false,
				"isInvalid": true,
				"isFailed":  true,
			},
		},
		InitErr: errors.New("fail"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns isValid true and isSuccess true -- 'ok' bytes no error",
			ExpectedInput: args.Map{
				"hasError":  false,
				"isValid":   true,
				"isSuccess": true,
				"isEmpty":   true,
				"isInvalid": false,
				"isFailed":  false,
			},
		},
		InitBytes: []byte("ok"),
	},
}

// =============================================================================
// BytesErrorOnce -- HasIssuesOrEmpty / HasSafeItems
// =============================================================================

var bytesErrorOnceHasIssuesTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns hasIssuesOrEmpty true -- 'data' bytes with error",
			ExpectedInput: args.Map{
				"hasIssuesOrEmpty": true,
				"hasSafeItems":     false,
			},
		},
		InitBytes: []byte("data"),
		InitErr:   errors.New("err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce returns hasIssuesOrEmpty true -- empty bytes no error",
			ExpectedInput: args.Map{
				"hasIssuesOrEmpty": true,
				"hasSafeItems":     false,
			},
		},
		InitBytes: []byte{},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce data no error -- HasIssuesOrEmpty false, HasSafeItems true",
			ExpectedInput: args.Map{
				"hasIssuesOrEmpty": false,
				"hasSafeItems":     true,
			},
		},
		InitBytes: []byte("ok"),
	},
	// Note: Nil receiver case migrated to BytesErrorOnce_NilReceiver_testcases.go
}

// =============================================================================
// BytesErrorOnce -- String
// =============================================================================

var bytesErrorOnceStringTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce 'str-val' -- String returns correct value",
			ExpectedInput: args.Map{
				"stringValue":               "str-val",
				"isStringEmpty":             false,
				"isStringEmptyOrWhitespace": false,
			},
		},
		InitBytes: []byte("str-val"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce nil -- String empty, IsStringEmpty true",
			ExpectedInput: args.Map{
				"stringValue":               "",
				"isStringEmpty":             true,
				"isStringEmptyOrWhitespace": true,
			},
		},
		InitBytes: nil,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce whitespace -- IsStringEmptyOrWhitespace true",
			ExpectedInput: args.Map{
				"stringValue":               "   ",
				"isStringEmpty":             false,
				"isStringEmptyOrWhitespace": true,
			},
		},
		InitBytes: []byte("   "),
	},
}

// =============================================================================
// BytesErrorOnce -- Deserialize
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
			ExpectedInput: args.Map{
				"noError":          true,
				"deserializedName": "test",
			},
		},
		InitJson: `{"name":"test"}`,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize returns error when source has error",
			ExpectedInput: args.Map{
				"hasSourceError":      true,
				"hasDeserializeError": true,
				"errorsMatch":         true,
			},
		},
		InitErr: errors.New("source error"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.Deserialize with invalid JSON returns error",
			ExpectedInput: args.Map{
				"hasError":    true,
				"isJsonError": true,
			},
		},
		InitJson: "not-json",
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.DeserializeMust succeeds without panic",
			ExpectedInput: args.Map{
				"didPanic":        false,
				"deserializedKey": "val",
			},
		},
		InitJson: `{"key":"val"}`,
		IsMust:   true,
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.DeserializeMust panics on error",
			ExpectedInput: args.Map{
				"didPanic": true,
			},
		},
		InitErr: errors.New("must-fail"),
		IsMust:  true,
	},
}

// =============================================================================
// BytesErrorOnce -- Serialization (MarshalJSON, Serialize, SerializeMust)
// =============================================================================

var bytesErrorOnceMarshalJSONTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.MarshalJSON returns bytes",
		ExpectedInput: args.Map{
			"noError":        true,
			"marshaledValue": `{"a":1}`,
		},
	},
	InitBytes: []byte(`{"a":1}`),
}

var bytesErrorOnceSerializeTestCase = bytesErrorOnceTestCase{
	Case: coretestcases.CaseV1{
		Title: "BytesErrorOnce.Serialize returns bytes",
		ExpectedInput: args.Map{
			"noError":         true,
			"serializedValue": "ser",
		},
	},
	InitBytes: []byte("ser"),
}

// =============================================================================
// BytesErrorOnce -- Lifecycle (panic checks, IsInitialized, constructor)
// =============================================================================

type bytesErrorOnceLifecycleTestCase struct {
	Case      coretestcases.CaseV1
	InitBytes []byte
	InitErr   error
}

var bytesErrorOnceLifecycleTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce -- HandleError no panic, MustBeEmptyError no panic on success",
			ExpectedInput: args.Map{
				"handleErrorPanicked":       false,
				"mustBeEmptyErrorPanicked":  false,
				"mustHaveSafeItemsPanicked": false,
			},
		},
		InitBytes: []byte("ok"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce -- HandleError panics on error",
			ExpectedInput: args.Map{
				"handleErrorPanicked":       true,
				"mustBeEmptyErrorPanicked":  true,
				"mustHaveSafeItemsPanicked": true,
			},
		},
		InitErr: errors.New("handle-err"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce -- MustHaveSafeItems panics when empty",
			ExpectedInput: args.Map{
				"handleErrorPanicked":       false,
				"mustBeEmptyErrorPanicked":  false,
				"mustHaveSafeItemsPanicked": true,
			},
		},
		InitBytes: []byte{},
	},
}

var bytesErrorOnceInitializedTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.IsInitialized false before, true after Value call",
			ExpectedInput: args.Map{
				"isInitializedBefore": false,
				"isInitializedAfter":  true,
			},
		},
		InitBytes: []byte("x"),
	},
}

var bytesErrorOnceSerializeMustTestCases = []bytesErrorOnceLifecycleTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.SerializeMust returns bytes without panic",
			ExpectedInput: args.Map{
				"didPanic":        false,
				"serializedValue": "must-ser",
			},
		},
		InitBytes: []byte("must-ser"),
	},
	{
		Case: coretestcases.CaseV1{
			Title: "BytesErrorOnce.SerializeMust panics on error",
			ExpectedInput: args.Map{
				"didPanic": true,
			},
		},
		InitErr: errors.New("ser-fail"),
	},
}

var bytesErrorOnceConstructorTestCases = []bytesErrorOnceTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "NewBytesErrorOnce (value) works correctly",
			ExpectedInput: args.Map{
				"value":     "val",
				"isCorrect": true,
			},
		},
		InitBytes: []byte("val"),
	},
}
