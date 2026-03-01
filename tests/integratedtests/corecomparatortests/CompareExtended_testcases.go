package corecomparatortests

// Extended test cases migrated from cmd/main/enumTesting.go and
// cmd/main/compareEnumTesting02.go.

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

// -------------------------------------------------------------------------
// enumTesting — JSON marshal/unmarshal roundtrip for Compare enum
// -------------------------------------------------------------------------

var compareJsonRoundtripTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal marshal produces quoted name, unmarshal restores identity",
		ArrangeInput: args.Map{
			"value":          0,
			"unmarshalInput": "3",
		},
		ExpectedInput: []string{
			"\"Equal\"",
			"NotEqual",
			"3",
		},
	},
}

// -------------------------------------------------------------------------
// compareEnumTesting02 — OnlySupportedErr
// -------------------------------------------------------------------------

var onlySupportedErrTestCases = []coretestcases.CaseV1{
	{
		Title: "Equal is NOT in the supported list — should return error",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "dining doesn't support more",
			"supported": []int{6, 3},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Equal IS in the supported list — should return nil error",
		ArrangeInput: args.Map{
			"value":     0,
			"message":   "some context message",
			"supported": []int{0, 3},
		},
		ExpectedInput: []string{
			"false",
		},
	},
	{
		Title: "Inconclusive is NOT in the supported list — should return error",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "only equal allowed",
			"supported": []int{0},
		},
		ExpectedInput: []string{
			"true",
		},
	},
	{
		Title: "Inconclusive IS in the supported list — should return nil error",
		ArrangeInput: args.Map{
			"value":     6,
			"message":   "inconclusive is fine",
			"supported": []int{6, 0, 3},
		},
		ExpectedInput: []string{
			"false",
		},
	},
}
