package corecomparatortests

// Extended test cases migrated from cmd/main/enumTesting.go and
// cmd/main/compareEnumTesting02.go.

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

// -------------------------------------------------------------------------
// enumTesting — JSON marshal/unmarshal roundtrip for Compare enum
// -------------------------------------------------------------------------

var compareJsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Marshal returns quoted name and Unmarshal restores identity -- Equal value 0, unmarshal '3'",
		ArrangeInput: args.Map{
			"value":          0,
			"unmarshalInput": "3",
		},
		ExpectedInput: args.Map{
			"marshaledJson":    "\"Equal\"",
			"unmarshaledName":  "NotEqual",
			"unmarshaledValue": "3",
		},
	},
}

// -------------------------------------------------------------------------
// compareEnumTesting02 — OnlySupportedErr
// -------------------------------------------------------------------------

var onlySupportedErrTestCases = []coretestcases.CaseV1{
	{
		Title: "OnlySupportedErr returns error -- Equal (0) not in supported [6,3]",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "dining doesn't support more",
			"supported": []int{6, 3},
		},
		ExpectedInput: "true", // hasError
	},
	{
		Title: "OnlySupportedErr returns nil -- Equal (0) in supported [0,3]",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "some context message",
			"supported": []int{0, 3},
		},
		ExpectedInput: "false", // hasError
	},
	{
		Title: "OnlySupportedErr returns error -- Inconclusive (6) not in supported [0]",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "only equal allowed",
			"supported": []int{0},
		},
		ExpectedInput: "true", // hasError
	},
	{
		Title: "OnlySupportedErr returns nil -- Inconclusive (6) in supported [6,0,3]",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "inconclusive is fine",
			"supported": []int{6, 0, 3},
		},
		ExpectedInput: "false", // hasError
	},
}
