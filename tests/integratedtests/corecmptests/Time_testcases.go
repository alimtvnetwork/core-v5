package corecmptests

import (
	"time"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
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
		ExpectedInput: []string{
			"Equal",
		},
	},
	{
		Title: "Time returns LeftLess when left is before right",
		ArrangeInput: args.Map{
			"when":  "given left time before right time",
			"left":  baseTime,
			"right": laterTime,
		},
		ExpectedInput: []string{
			"LeftLess",
		},
	},
	{
		Title: "Time returns LeftGreater when left is after right",
		ArrangeInput: args.Map{
			"when":  "given left time after right time",
			"left":  laterTime,
			"right": baseTime,
		},
		ExpectedInput: []string{
			"LeftGreater",
		},
	},
	{
		Title: "Time returns LeftGreater for small duration difference",
		ArrangeInput: args.Map{
			"when":  "given left time slightly after right by nanoseconds",
			"left":  baseTime.Add(time.Duration(600000)),
			"right": baseTime,
		},
		ExpectedInput: []string{
			"LeftGreater",
		},
	},
	{
		Title: "Time returns LeftLess for small duration difference reverse",
		ArrangeInput: args.Map{
			"when":  "given left time slightly before right by nanoseconds",
			"left":  baseTime,
			"right": baseTime.Add(time.Duration(600000)),
		},
		ExpectedInput: []string{
			"LeftLess",
		},
	},
}
