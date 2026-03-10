package keymktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var keyCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "Default key compiles with hyphen joiner",
		ArrangeInput: args.Map{
			"when":   "given main and chain items",
			"main":   "root",
			"chains": []string{"sub", "item"},
		},
		ExpectedInput: "root-sub-item",
	},
	{
		Title: "Default key compiles with main only",
		ArrangeInput: args.Map{
			"when":   "given main only",
			"main":   "solo",
			"chains": []string{},
		},
		ExpectedInput: "solo",
	},
	{
		Title: "Default key compiles with single chain",
		ArrangeInput: args.Map{
			"when":   "given main and single chain",
			"main":   "app",
			"chains": []string{"config"},
		},
		ExpectedInput: "app-config",
	},
	{
		Title: "Default key compiles with multiple chains",
		ArrangeInput: args.Map{
			"when":   "given main and four chain items",
			"main":   "a",
			"chains": []string{"b", "c", "d", "e"},
		},
		ExpectedInput: "a-b-c-d-e",
	},
}

var keyAppendChainTestCases = []coretestcases.CaseV1{
	{
		Title: "AppendChain adds items and compiles",
		ArrangeInput: args.Map{
			"when":    "given main then append items",
			"main":    "root",
			"initial": []string{"a"},
			"append":  []string{"b", "c"},
		},
		ExpectedInput: "root-a-b-c",
	},
}

var keyFinalizedTestCases = []coretestcases.CaseV1{
	{
		Title: "Finalized locks key and returns compiled",
		ArrangeInput: args.Map{
			"when":   "given key is finalized",
			"main":   "base",
			"chains": []string{"path"},
		},
		ExpectedInput: []string{"base-path", "true"},
	},
}

var keyHasInChainsTestCases = []coretestcases.CaseV1{
	{
		Title: "HasInChains true when item exists",
		ArrangeInput: args.Map{
			"when":   "given existing chain item",
			"main":   "root",
			"chains": []string{"sub", "item"},
			"search": "sub",
		},
		ExpectedInput: "true",
	},
	{
		Title: "HasInChains false when item missing",
		ArrangeInput: args.Map{
			"when":   "given non-existing chain item",
			"main":   "root",
			"chains": []string{"sub", "item"},
			"search": "missing",
		},
		ExpectedInput: "false",
	},
}

var keyClonePtrTestCases = []coretestcases.CaseV1{
	{
		Title: "ClonePtr creates independent copy",
		ArrangeInput: args.Map{
			"when":   "given key with chains",
			"main":   "original",
			"chains": []string{"a", "b"},
		},
		ExpectedInput: []string{"original-a-b", "original-a-b"},
	},
}

var keyLengthTestCases = []coretestcases.CaseV1{
	{
		Title: "Length returns chain count",
		ArrangeInput: args.Map{
			"when":   "given key with 3 chains",
			"main":   "root",
			"chains": []string{"a", "b", "c"},
		},
		ExpectedInput: "3",
	},
	{
		Title: "Length returns 0 for no chains",
		ArrangeInput: args.Map{
			"when":   "given key with no chains",
			"main":   "root",
			"chains": []string{},
		},
		ExpectedInput: "0",
	},
}

var keyIsEmptyTestCases = []coretestcases.CaseV1{
	{
		Title: "IsEmpty true for empty main and no chains",
		ArrangeInput: args.Map{
			"when":   "given empty main",
			"main":   "",
			"chains": []string{},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmpty false for non-empty main",
		ArrangeInput: args.Map{
			"when":   "given non-empty main",
			"main":   "root",
			"chains": []string{},
		},
		ExpectedInput: "false",
	},
}

var keyCompileWithAdditionalTestCases = []coretestcases.CaseV1{
	{
		Title: "Compile with additional items appends at the end",
		ArrangeInput: args.Map{
			"when":       "given key with chains and compile extras",
			"main":       "root",
			"chains":     []string{"a"},
			"additional": "extra",
		},
		ExpectedInput: "root-a-extra",
	},
}
