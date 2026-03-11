package coreversiontests

import (
	"github.com/alimtvnetwork/core/corecomparator"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var comparisonValueIndexesTestCases = []coretestcases.CaseV1{
	{
		Title: "ComparisonValueIndexes returns Equal for identical versions using all indexes",
		ArrangeInput: args.Map{
			"when":  "given identical versions compared by all indexes",
			"left":  "v0.0.1",
			"right": "v0.0.1",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
	{
		Title: "ComparisonValueIndexes returns LeftGreater when left major is greater",
		ArrangeInput: args.Map{
			"when":  "given left major version greater compared by all indexes",
			"left":  "v3.0",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftGreater",
		},
	},
	{
		Title: "ComparisonValueIndexes returns LeftLess when left minor is less",
		ArrangeInput: args.Map{
			"when":  "given left minor version less compared by all indexes",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftLess",
		},
	},
	{
		Title: "ComparisonValueIndexes returns Equal for zero-padded equivalents",
		ArrangeInput: args.Map{
			"when":  "given v4 vs v4.0 compared by all indexes",
			"left":  "v4",
			"right": "v4.0",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
}

var versionSliceIntegerTestCases = []coretestcases.CaseV1{
	{
		Title: "VersionSliceInteger returns Equal for identical version value slices",
		ArrangeInput: args.Map{
			"when":  "given identical version value slices",
			"left":  "v1.2.3",
			"right": "v1.2.3",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
	{
		Title: "VersionSliceInteger returns LeftGreater when left has higher major",
		ArrangeInput: args.Map{
			"when":  "given left has higher major",
			"left":  "v3.0",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftGreater",
		},
	},
	{
		Title: "VersionSliceInteger returns LeftLess when left has lower minor",
		ArrangeInput: args.Map{
			"when":  "given left has lower minor",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftLess",
		},
	},
	{
		Title: "VersionSliceInteger returns Equal for v4 vs v4.0",
		ArrangeInput: args.Map{
			"when":  "given v4 vs v4.0 same effective version",
			"left":  "v4",
			"right": "v4.0",
		},
		ExpectedInput: args.Map{
			"result": "Equal",
		},
	},
	{
		Title: "VersionSliceInteger returns LeftLess for build difference",
		ArrangeInput: args.Map{
			"when":  "given equal major.minor.patch but left build less",
			"left":  "v2.0.0.1",
			"right": "v2.0.0.5",
		},
		ExpectedInput: args.Map{
			"result": "LeftLess",
		},
	},
	{
		Title: "VersionSliceInteger returns LeftGreater for build difference",
		ArrangeInput: args.Map{
			"when":  "given equal major.minor.patch but left build greater",
			"left":  "v2.0.0.5",
			"right": "v2.0.0.1",
		},
		ExpectedInput: args.Map{
			"result": "LeftGreater",
		},
	},
}

var isAtLeastTestCases = []coretestcases.CaseV1{
	{
		Title: "IsAtLeast returns true when left is greater",
		ArrangeInput: args.Map{
			"when":  "given left version greater than right",
			"left":  "3.0",
			"right": "0.2.1",
		},
		ExpectedInput: args.Map{
			"isAtLeast": true,
		},
	},
	{
		Title: "IsAtLeast returns true when versions are equal",
		ArrangeInput: args.Map{
			"when":  "given equal versions",
			"left":  "v0.0.1",
			"right": "v0.0.1",
		},
		ExpectedInput: args.Map{
			"isAtLeast": true,
		},
	},
	{
		Title: "IsAtLeast returns false when left is less",
		ArrangeInput: args.Map{
			"when":  "given left version less than right",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"isAtLeast": false,
		},
	},
}

var isLowerTestCases = []coretestcases.CaseV1{
	{
		Title: "IsLower returns true when left is less",
		ArrangeInput: args.Map{
			"when":  "given left version less than right",
			"left":  "v0.0.2",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"isLower": true,
		},
	},
	{
		Title: "IsLower returns false when versions are equal",
		ArrangeInput: args.Map{
			"when":  "given equal versions",
			"left":  "v0.0.1",
			"right": "v0.0.1",
		},
		ExpectedInput: args.Map{
			"isLower": false,
		},
	},
	{
		Title: "IsLower returns false when left is greater",
		ArrangeInput: args.Map{
			"when":  "given left version greater than right",
			"left":  "v3.0",
			"right": "v0.2.1",
		},
		ExpectedInput: args.Map{
			"isLower": false,
		},
	},
}

var isExpectedVersionTestCases = []coretestcases.CaseV1{
	{
		Title: "IsExpectedVersion returns true -- equal versions v0.0.1 with Equal expectation",
		ArrangeInput: args.Map{
			"when":     "given equal versions with Equal expectation",
			"left":     "v0.0.1",
			"right":    "v0.0.1",
			"expected": corecomparator.Equal,
		},
		ExpectedInput: args.Map{
			"isExpected": true,
		},
	},
	{
		Title: "IsExpectedVersion returns true -- v3.0 vs v0.2.1 with LeftGreater expectation",
		ArrangeInput: args.Map{
			"when":     "given left greater with LeftGreater expectation",
			"left":     "v3.0",
			"right":    "v0.2.1",
			"expected": corecomparator.LeftGreater,
		},
		ExpectedInput: args.Map{
			"isExpected": true,
		},
	},
	{
		Title: "IsExpectedVersion returns true -- equal versions v4 vs v4.0 with LeftGreater expectation",
		ArrangeInput: args.Map{
			"when":     "given equal versions with LeftGreater expectation",
			"left":     "v4",
			"right":    "v4.0",
			"expected": corecomparator.LeftGreater,
		},
		ExpectedInput: args.Map{
			"isExpected": true,
		},
	},
}
