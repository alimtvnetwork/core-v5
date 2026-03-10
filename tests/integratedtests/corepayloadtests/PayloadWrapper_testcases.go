package corepayloadtests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// ==========================================
// PayloadWrapper — Create and Serialize
// ==========================================

var payloadWrapperCreateTestCases = []coretestcases.CaseV1{
	{
		Title: "PayloadWrapper creation produces valid JSON",
		ArrangeInput: args.Map{
			"when": "given payload with name and id",
			"name": "test-payload",
			"id":   "pay-1",
		},
		ExpectedInput: args.Map{
			"name":       "test-payload",
			"identifier": "pay-1",
			"hasJson":    true,
		},
	},
}

// ==========================================
// PayloadWrapper — Deserialization roundtrip
// ==========================================

var payloadWrapperDeserializeRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "PayloadWrapper serialization then deserialization preserves data",
		ArrangeInput: args.Map{
			"when": "given payload serialized then deserialized",
			"name": "roundtrip-payload",
			"id":   "rt-1",
		},
		ExpectedInput: args.Map{
			"restoredName":       "roundtrip-payload",
			"restoredIdentifier": "rt-1",
			"jsonIsEqual":        true,
		},
	},
}

// ==========================================
// PayloadWrapper — Deep Clone
// ==========================================

var payloadWrapperCloneTestCases = []coretestcases.CaseV1{
	{
		Title: "PayloadWrapper ClonePtr creates independent copy",
		ArrangeInput: args.Map{
			"when":     "given payload cloned and mutated",
			"name":     "original-pay",
			"id":       "clone-1",
			"new_name": "mutated-pay",
		},
		ExpectedInput: args.Map{
			"originalName":  "original-pay",
			"clonedName":    "mutated-pay",
			"isIndependent": true,
		},
	},
}

// ==========================================
// PayloadWrapper — DeserializeToMany
// ==========================================

var payloadWrapperDeserializeToManyTestCases = []coretestcases.CaseV1{
	{
		Title: "DeserializeToMany parses array of payloads",
		ArrangeInput: args.Map{
			"when":  "given 3 payloads serialized as array",
			"count": 3,
		},
		ExpectedInput: args.Map{
			"deserializedCount": 3,
		},
	},
}
