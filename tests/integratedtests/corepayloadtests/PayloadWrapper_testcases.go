package corepayloadtests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		ExpectedInput: []string{
			"test-payload",
			"pay-1",
			"true",
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
		ExpectedInput: []string{
			"roundtrip-payload",
			"rt-1",
			"true",
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
		ExpectedInput: []string{
			"original-pay",
			"mutated-pay",
			"true",
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
		ExpectedInput: []string{
			"3",
		},
	},
}
