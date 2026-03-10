package defaultcapacitytests

import (
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var predictiveTestCases = []coretestcases.CaseV1{
	{
		Title: "Zero length returns additionalCap only",
		ArrangeInput: args.Map{
			"possibleLen":   0,
			"multiplier":    1.5,
			"additionalCap": 10,
		},
		ExpectedInput: "10",
	},
	{
		Title: "Negative length returns additionalCap only",
		ArrangeInput: args.Map{
			"possibleLen":   -5,
			"multiplier":    1.5,
			"additionalCap": 10,
		},
		ExpectedInput: "10",
	},
	{
		Title: "10 items with 1.5x multiplier and 5 additional = 20",
		ArrangeInput: args.Map{
			"possibleLen":   10,
			"multiplier":    1.5,
			"additionalCap": 5,
		},
		ExpectedInput: "20",
	},
	{
		Title: "100 items with 2.0x multiplier and 0 additional = 200",
		ArrangeInput: args.Map{
			"possibleLen":   100,
			"multiplier":    2.0,
			"additionalCap": 0,
		},
		ExpectedInput: "200",
	},
	{
		Title: "1 item with 1.0x multiplier and 0 additional = 1",
		ArrangeInput: args.Map{
			"possibleLen":   1,
			"multiplier":    1.0,
			"additionalCap": 0,
		},
		ExpectedInput: "1",
	},
	{
		Title: "3 items with 1.2x multiplier and 4 additional = 8",
		ArrangeInput: args.Map{
			"possibleLen":   3,
			"multiplier":    1.2,
			"additionalCap": 4,
		},
		ExpectedInput: "8",
	},
}

var maxLimitTestCases = []coretestcases.CaseV1{
	{
		Title: "Limit -1 (no limit) returns wholeLength",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       -1,
		},
		ExpectedInput: "100",
	},
	{
		Title: "Limit >= wholeLength returns wholeLength",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       100,
		},
		ExpectedInput: "50",
	},
	{
		Title: "Limit == wholeLength returns wholeLength",
		ArrangeInput: args.Map{
			"wholeLength": 50,
			"limit":       50,
		},
		ExpectedInput: "50",
	},
	{
		Title: "Limit < wholeLength returns limit",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       25,
		},
		ExpectedInput: "25",
	},
	{
		Title: "Limit 0 with items returns 0",
		ArrangeInput: args.Map{
			"wholeLength": 100,
			"limit":       0,
		},
		ExpectedInput: "0",
	},
}

var ofSearchTestCases = []coretestcases.CaseV1{
	{
		Title:         "Length 1 returns 1",
		ArrangeInput:  1,
		ExpectedInput: "1",
	},
	{
		Title:         "Length 3 returns 3",
		ArrangeInput:  3,
		ExpectedInput: "3",
	},
	{
		Title:         "Length 15 returns 15/3 = 5",
		ArrangeInput:  15,
		ExpectedInput: "5",
	},
	{
		Title:         "Length 1000 returns 100",
		ArrangeInput:  1000,
		ExpectedInput: "100",
	},
	{
		Title:         "Length 500 returns 500/20 = 25",
		ArrangeInput:  500,
		ExpectedInput: "25",
	},
	{
		Title:         "Length 100 returns 100/10 = 10",
		ArrangeInput:  100,
		ExpectedInput: "10",
	},
	{
		Title:         "Length 50 returns 50/5 = 10",
		ArrangeInput:  50,
		ExpectedInput: "10",
	},
}

var predictiveDefaultTestCases = []coretestcases.CaseV1{
	{
		Title:         "Positive length yields positive result",
		ArrangeInput:  100,
		ExpectedInput: "true",
	},
	{
		Title:         "Zero length yields positive result (from additionalCap)",
		ArrangeInput:  0,
		ExpectedInput: "true",
	},
}
