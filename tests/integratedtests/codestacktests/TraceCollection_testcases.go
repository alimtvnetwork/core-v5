package codestacktests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var traceCollectionBasicTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection empty -- Length 0 IsEmpty true",
		ArrangeInput: args.Map{
			"when":  "empty collection",
			"count": 0,
		},
		ExpectedInput: args.Map{
			"length":     0,
			"isEmpty":    true,
			"hasAnyItem": false,
		},
	},
	{
		Title: "TraceCollection with items -- Length > 0 IsEmpty false",
		ArrangeInput: args.Map{
			"when":  "collection with 3 items",
			"count": 3,
		},
		ExpectedInput: args.Map{
			"length":     3,
			"isEmpty":    false,
			"hasAnyItem": true,
		},
	},
}

var traceCollectionFirstLastTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection First/Last -- returns correct items",
		ArrangeInput: args.Map{
			"when": "collection with 3 items",
		},
		ExpectedInput: args.Map{
			"firstPkg": "pkg1",
			"lastPkg":  "pkg3",
			"lastIdx":  2,
		},
	},
}

var traceCollectionSkipTakeTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Skip 1 -- returns remaining items",
		ArrangeInput: args.Map{
			"when": "skip first item",
			"skip": 1,
		},
		ExpectedInput: args.Map{
			"length":   2,
			"firstPkg": "pkg2",
		},
	},
	{
		Title: "TraceCollection Take 2 -- returns first 2 items",
		ArrangeInput: args.Map{
			"when": "take first 2",
			"take": 2,
		},
		ExpectedInput: args.Map{
			"length":  2,
			"lastPkg": "pkg2",
		},
	},
}

var traceCollectionReverseTestCases = []coretestcases.CaseV1{
	{
		Title: "TraceCollection Reverse 3 items -- reverses order",
		ArrangeInput: args.Map{
			"when": "reverse 3 item collection",
		},
		ExpectedInput: args.Map{
			"firstPkg": "pkg3",
			"lastPkg":  "pkg1",
		},
	},
}

var traceCollectionFilterTestCases = []coretestcases.CaseV1{
	{
		Title: "FilterPackageNameTraceCollection -- filters by package",
		ArrangeInput: args.Map{
			"when":    "filter by pkg2",
			"package": "pkg2",
		},
		ExpectedInput: args.Map{
			"length":   1,
			"firstPkg": "pkg2",
		},
	},
}
