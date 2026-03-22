package corejsontests

import (
	"errors"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// deserializerLogic — Apply / UsingResult / ApplyMust
// =============================================================================

// Branch: valid result → unmarshal success
var deserializeApplyTestCase = coretestcases.CaseV1{
	Title: "Deserialize Apply returns no error -- valid result with string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingResultTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingResult returns no error -- valid result with string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// Branch: result has error → ApplyMust panics
var deserializeApplyMustPanicsTestCase = coretestcases.CaseV1{
	Title: "Deserialize ApplyMust panics -- result with error",
	ExpectedInput: args.Map{
		"panicked": true,
	},
}

// =============================================================================
// deserializerLogic — UsingString / FromString / FromStringMust
// =============================================================================

var deserializeUsingStringTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingString returns no error -- valid json string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeFromStringTestCase = coretestcases.CaseV1{
	Title: "Deserialize FromString returns no error -- valid json string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeFromStringMustPanicsTestCase = coretestcases.CaseV1{
	Title: "Deserialize FromStringMust panics -- invalid json string",
	ExpectedInput: args.Map{
		"panicked": true,
	},
}

// =============================================================================
// deserializerLogic — UsingStringPtr / UsingStringOption / UsingStringIgnoreEmpty
// =============================================================================

var deserializeUsingStringPtrNilTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingStringPtr returns error -- nil pointer",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingStringPtrValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingStringPtr returns no error -- valid pointer",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingStringOptionSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingStringOption returns nil -- skip flag true",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeUsingStringOptionProcessTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingStringOption returns no error -- skip false valid json",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingStringIgnoreEmptyTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingStringIgnoreEmpty returns nil -- empty string",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// deserializerLogic — UsingError / UsingBytes / UsingBytesMust
// =============================================================================

var deserializeUsingErrorNilTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingError returns nil -- nil error input",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeUsingBytesValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytes returns no error -- valid json bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingBytesInvalidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytes returns error -- invalid json bytes",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingBytesMustPanicsTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytesMust panics -- invalid json bytes",
	ExpectedInput: args.Map{
		"panicked": true,
	},
}

// =============================================================================
// deserializerLogic — UsingBytesPointer / UsingBytesIf / UsingSafeBytesMust
// =============================================================================

var deserializeUsingBytesPointerNilTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytesPointer returns error -- nil bytes",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingBytesPointerValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytesPointer returns no error -- valid json bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingBytesIfSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytesIf returns nil -- condition false skip",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeUsingBytesIfProcessTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingBytesIf returns no error -- condition true valid bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var deserializeUsingSafeBytesMustEmptyTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingSafeBytesMust returns without panic -- empty bytes",
	ExpectedInput: args.Map{
		"result": "",
	},
}

var deserializeUsingSafeBytesMustValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingSafeBytesMust returns value -- valid json bytes",
	ExpectedInput: args.Map{
		"result": "hello",
	},
}

// =============================================================================
// deserializerLogic — MapAnyToPointer / FromTo
// =============================================================================

type deserializeTestStruct struct {
	Name string
}

var deserializeMapAnyToPointerSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize MapAnyToPointer returns nil -- skip empty map",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeMapAnyToPointerValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize MapAnyToPointer returns no error -- valid map with Name",
	ExpectedInput: args.Map{
		"hasError": false,
		"name":     "test",
	},
}

var deserializeFromToTestCase = coretestcases.CaseV1{
	Title: "Deserialize FromTo returns no error -- string to string",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// =============================================================================
// deserializerLogic — UsingDeserializerFuncDefined
// =============================================================================

var deserializeUsingDeserializerFuncDefinedNilTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingDeserializerFuncDefined returns error -- nil func",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingDeserializerFuncDefinedValidTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingDeserializerFuncDefined returns no error -- valid func",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

// =============================================================================
// deserializerLogic — UsingJsonerToAny / UsingDeserializerToOption / UsingDeserializerDefined
// =============================================================================

var deserializeUsingJsonerToAnySkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingJsonerToAny returns nil -- skip nil true",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeUsingJsonerToAnyNotSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingJsonerToAny returns error -- nil jsoner skip false",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingDeserializerToOptionSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingDeserializerToOption returns nil -- skip nil true",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var deserializeUsingDeserializerToOptionNotSkipTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingDeserializerToOption returns error -- nil deserializer skip false",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var deserializeUsingDeserializerDefinedNilTestCase = coretestcases.CaseV1{
	Title: "Deserialize UsingDeserializerDefined returns nil -- nil deserializer",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// deserializerLogic — Result / ResultPtr
// =============================================================================

var deserializeResultPtrInvalidTestCase = coretestcases.CaseV1{
	Title: "Deserialize ResultPtr returns error -- invalid json bytes",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// =============================================================================
// deserializeFromBytesTo — String / Strings / Integer / Bool / Map
// =============================================================================

// Branch: valid json bytes → typed result
var bytesToStringTestCases = []coretestcases.CaseV1{
	{
		Title:        "BytesTo String returns hello -- valid json string bytes",
		ArrangeInput: args.Map{"bytes": []byte(`"hello"`)},
		ExpectedInput: args.Map{
			"hasError": false,
			"result":   "hello",
		},
	},
}

var bytesToStringMustTestCase = coretestcases.CaseV1{
	Title: "BytesTo StringMust returns hello -- valid json string bytes",
	ExpectedInput: args.Map{
		"result": "hello",
	},
}

var bytesToStringMustPanicsTestCase = coretestcases.CaseV1{
	Title: "BytesTo StringMust panics -- invalid json bytes",
	ExpectedInput: args.Map{
		"panicked": true,
	},
}

var bytesToStringsTestCase = coretestcases.CaseV1{
	Title: "BytesTo Strings returns 2 items -- valid json array",
	ExpectedInput: args.Map{
		"hasError": false,
		"length":   2,
	},
}

var bytesToIntegerTestCase = coretestcases.CaseV1{
	Title: "BytesTo Integer returns 42 -- valid json int bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   42,
	},
}

var bytesToInteger64TestCase = coretestcases.CaseV1{
	Title: "BytesTo Integer64 returns 99 -- valid json int64 bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   int64(99),
	},
}

var bytesToBoolTestCase = coretestcases.CaseV1{
	Title: "BytesTo Bool returns true -- valid json bool bytes",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   true,
	},
}

var bytesToMapAnyItemTestCase = coretestcases.CaseV1{
	Title: "BytesTo MapAnyItem returns map with key a -- valid json object",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasKeyA":  true,
	},
}

var bytesToMapStringStringTestCase = coretestcases.CaseV1{
	Title: "BytesTo MapStringString returns a=b -- valid json string map",
	ExpectedInput: args.Map{
		"hasError": false,
		"valueA":   "b",
	},
}

// =============================================================================
// deserializeFromResultTo — String / Bool / Byte / MapAnyItem / MapStringString
// =============================================================================

var resultToStringTestCase = coretestcases.CaseV1{
	Title: "ResultTo String returns hello -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   "hello",
	},
}

var resultToBoolTestCase = coretestcases.CaseV1{
	Title: "ResultTo Bool returns true -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   true,
	},
}

var resultToByteTestCase = coretestcases.CaseV1{
	Title: "ResultTo Byte returns 65 -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"result":   byte(65),
	},
}

var resultToMapAnyItemTestCase = coretestcases.CaseV1{
	Title: "ResultTo MapAnyItem returns map with key a -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasKeyA":  true,
	},
}

var resultToMapStringStringTestCase = coretestcases.CaseV1{
	Title: "ResultTo MapStringString returns a=b -- valid result",
	ExpectedInput: args.Map{
		"hasError": false,
		"valueA":   "b",
	},
}

// Branch: invalid bytes → error
var resultToResultCollectionInvalidTestCase = coretestcases.CaseV1{
	Title: "ResultTo ResultCollection returns error -- invalid bytes in result",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var resultToResultsPtrCollectionInvalidTestCase = coretestcases.CaseV1{
	Title: "ResultTo ResultsPtrCollection returns error -- invalid bytes in result",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var resultToMapResultsInvalidTestCase = coretestcases.CaseV1{
	Title: "ResultTo MapResults returns error -- invalid bytes in result",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// =============================================================================
// Panic helper
// =============================================================================

var invalidResultForPanic = errors.New("fail")
