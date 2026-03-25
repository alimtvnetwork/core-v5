package corejsontests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// =============================================================================
// serializerLogic — Apply
// =============================================================================

// Branch: valid input → serialize success
var serializeApplyTestCases = []coretestcases.CaseV1{
	{
		Title: "Serialize Apply returns no error -- valid string input",
		ArrangeInput: args.Map{
			"input": "hello",
		},
		ExpectedInput: args.Map{
			"hasError": false,
			"hasBytes": true,
		},
	},
	{
		Title: "Serialize Apply returns no error -- valid int input",
		ArrangeInput: args.Map{
			"input": 42,
		},
		ExpectedInput: args.Map{
			"hasError": false,
			"hasBytes": true,
		},
	},
}

// Branch: unmarshalable input → error
var serializeApplyUnmarshalableTestCase = coretestcases.CaseV1{
	Title:        "Serialize Apply returns error -- unmarshalable channel input",
	ArrangeInput: args.Map{},
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

// =============================================================================
// serializerLogic — FromBytes / FromStrings / FromString / FromInteger
// =============================================================================

var serializeFromBytesTestCase = coretestcases.CaseV1{
	Title: "Serialize FromBytes returns no error -- valid byte slice",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromStringsTestCase = coretestcases.CaseV1{
	Title: "Serialize FromStrings returns no error -- valid string slice",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromStringsSpreadTestCase = coretestcases.CaseV1{
	Title: "Serialize FromStringsSpread returns no error -- two strings",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromStringTestCase = coretestcases.CaseV1{
	Title: "Serialize FromString returns no error -- hello string",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromIntegerTestCase = coretestcases.CaseV1{
	Title: "Serialize FromInteger returns no error -- int 42",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromInteger64TestCase = coretestcases.CaseV1{
	Title: "Serialize FromInteger64 returns no error -- int64 99",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromBoolTestCase = coretestcases.CaseV1{
	Title: "Serialize FromBool returns no error -- true",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeFromIntegersTestCase = coretestcases.CaseV1{
	Title: "Serialize FromIntegers returns no error -- int slice",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// serializerLogic — UsingAnyPtr / UsingAny
// =============================================================================

var serializeUsingAnyPtrValidTestCase = coretestcases.CaseV1{
	Title: "Serialize UsingAnyPtr returns no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeUsingAnyPtrUnmarshalableTestCase = coretestcases.CaseV1{
	Title: "Serialize UsingAnyPtr returns error -- unmarshalable channel",
	ExpectedInput: args.Map{
		"hasError": true,
	},
}

var serializeUsingAnyTestCase = coretestcases.CaseV1{
	Title: "Serialize UsingAny returns no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

// =============================================================================
// serializerLogic — Raw / Marshal / ApplyMust / ToBytesMust / etc.
// =============================================================================

var serializeRawTestCase = coretestcases.CaseV1{
	Title: "Serialize Raw returns bytes and no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasBytes": true,
	},
}

var serializeMarshalTestCase = coretestcases.CaseV1{
	Title: "Serialize Marshal returns bytes and no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasBytes": true,
	},
}

var serializeApplyMustTestCase = coretestcases.CaseV1{
	Title: "Serialize ApplyMust returns no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
	},
}

var serializeToBytesMustTestCase = coretestcases.CaseV1{
	Title: "Serialize ToBytesMust returns bytes -- valid string",
	ExpectedInput: args.Map{
		"hasBytes": true,
	},
}

var serializeToSafeBytesMustTestCase = coretestcases.CaseV1{
	Title: "Serialize ToSafeBytesMust returns bytes -- valid string",
	ExpectedInput: args.Map{
		"hasBytes": true,
	},
}

// =============================================================================
// serializerLogic — Swallow / ToString / ToPretty
// =============================================================================

var serializeToSafeBytesSwallowErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToSafeBytesSwallowErr returns bytes -- valid string",
	ExpectedInput: args.Map{
		"hasBytes": true,
	},
}

var serializeToBytesSwallowErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToBytesSwallowErr returns bytes -- valid string",
	ExpectedInput: args.Map{
		"hasBytes": true,
	},
}

var serializeToBytesErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToBytesErr returns bytes and no error -- valid string",
	ExpectedInput: args.Map{
		"hasError": false,
		"hasBytes": true,
	},
}

var serializeToStringTestCase = coretestcases.CaseV1{
	Title: "Serialize ToString returns non-empty string -- valid input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var serializeToStringMustTestCase = coretestcases.CaseV1{
	Title: "Serialize ToStringMust returns non-empty string -- valid input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var serializeToStringErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToStringErr returns string and no error -- valid input",
	ExpectedInput: args.Map{
		"hasError":   false,
		"hasContent": true,
	},
}

var serializeToPrettyStringErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToPrettyStringErr returns pretty string -- map input",
	ExpectedInput: args.Map{
		"hasError":   false,
		"hasContent": true,
	},
}

var serializeToPrettyStringIncludingErrTestCase = coretestcases.CaseV1{
	Title: "Serialize ToPrettyStringIncludingErr returns string -- valid input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}

var serializePrettyTestCase = coretestcases.CaseV1{
	Title: "Serialize Pretty returns non-empty string -- valid input",
	ExpectedInput: args.Map{
		"hasContent": true,
	},
}
