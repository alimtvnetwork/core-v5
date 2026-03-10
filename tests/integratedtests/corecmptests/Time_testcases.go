package corecmptests

import (
	"time"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var baseTime = time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
var laterTime = baseTime.Add(10 * time.Minute)

var timeCompareTestCases = []coretestcases.CaseV1{
	{
		Title: "Time returns Equal for identical times",
		ArrangeInput: args.Map{
			"when":  "given identical time values",
			"left":  baseTime,
			"right": baseTime,
		},
		ExpectedInput: "Equal", // compareResult
	},
	{
		Title: "Time returns LeftLess when left is before right",
		ArrangeInput: args.Map{
			"when":  "given left time before right time",
			"left":  baseTime,
			"right": laterTime,
		},
		ExpectedInput: "LeftLess", // compareResult
	},
	{
		Title: "Time returns LeftGreater when left is after right",
		ArrangeInput: args.Map{
			"when":  "given left time after right time",
			"left":  laterTime,
			"right": baseTime,
		},
		ExpectedInput: "LeftGreater", // compareResult
	},
	{
		Title: "Time returns LeftGreater for small duration difference",
		ArrangeInput: args.Map{
			"when":  "given left time slightly after right by nanoseconds",
			"left":  baseTime.Add(time.Duration(600000)),
			"right": baseTime,
		},
		ExpectedInput: "LeftGreater", // compareResult
	},
	{
		Title: "Time returns LeftLess for small duration difference reverse",
		ArrangeInput: args.Map{
			"when":  "given left time slightly before right by nanoseconds",
			"left":  baseTime,
			"right": baseTime.Add(time.Duration(600000)),
		},
		ExpectedInput: "LeftLess", // compareResult
	},
}
