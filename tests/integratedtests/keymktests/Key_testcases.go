package keymktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var keyCompileTestCases = []coretestcases.CaseV1{
	{
		Title: "Key.Compile returns 'root-sub-item' -- main 'root' with chains ['sub','item']",
		ArrangeInput: args.Map{
			"when":   "given main and chain items",
			"main":   "root",
			"chains": []string{"sub", "item"},
		},
		ExpectedInput: "root-sub-item",
	},
	{
		Title: "Key.Compile returns 'solo' -- main only, no chains",
		ArrangeInput: args.Map{
			"when":   "given main only",
			"main":   "solo",
			"chains": []string{},
		},
		ExpectedInput: "solo",
	},
	{
		Title: "Key.Compile returns 'app-config' -- main 'app' with single chain",
		ArrangeInput: args.Map{
			"when":   "given main and single chain",
			"main":   "app",
			"chains": []string{"config"},
		},
		ExpectedInput: "app-config",
	},
	{
		Title: "Key.Compile returns 'a-b-c-d-e' -- main 'a' with four chains",
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
		Title: "AppendChain returns 'root-a-b-c' -- main 'root', initial ['a'], appended ['b','c']",
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
		Title: "Finalized returns compiled 'base-path' and locks key -- main 'base' chain ['path']",
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
		Title: "HasInChains returns true -- search 'sub' in chains ['sub','item']",
		ArrangeInput: args.Map{
			"when":   "given existing chain item",
			"main":   "root",
			"chains": []string{"sub", "item"},
			"search": "sub",
		},
		ExpectedInput: "true",
	},
	{
		Title: "HasInChains returns false -- search 'missing' in chains ['sub','item']",
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
		Title: "ClonePtr returns independent copy -- main 'original' chains ['a','b']",
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
		Title: "Length returns 3 -- key with chains ['a','b','c']",
		ArrangeInput: args.Map{
			"when":   "given key with 3 chains",
			"main":   "root",
			"chains": []string{"a", "b", "c"},
		},
		ExpectedInput: "3",
	},
	{
		Title: "Length returns 0 -- key with no chains",
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
		Title: "IsEmpty returns true -- empty main and no chains",
		ArrangeInput: args.Map{
			"when":   "given empty main",
			"main":   "",
			"chains": []string{},
		},
		ExpectedInput: "true",
	},
	{
		Title: "IsEmpty returns false -- non-empty main 'root'",
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
		Title: "Compile returns 'root-a-extra' -- main 'root', chain ['a'], additional 'extra'",
		ArrangeInput: args.Map{
			"when":       "given key with chains and compile extras",
			"main":       "root",
			"chains":     []string{"a"},
			"additional": "extra",
		},
		ExpectedInput: "root-a-extra",
	},
}
