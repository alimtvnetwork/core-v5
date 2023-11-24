package integratedtests

import (
	"reflect"

	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	arrangeTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	twoArgsTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.Two{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	oneFuncTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([]args.OneFunc{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	interfaceArrayTypeVerification = &coretests.VerifyTypeOf{
		ArrangeInput:  reflect.TypeOf([][]interface{}{}),
		ActualInput:   reflect.TypeOf([]string{}),
		ExpectedInput: reflect.TypeOf([]string{}),
	}

	quickTestCases = []coretestcases.CaseV1{
		{
			Title: "Quick output as gherkins format",
			ArrangeInput: args.Map{
				"when":    "some title, or case when",
				"actual":  "actual rec",
				"expect":  "expected item",
				"counter": 3,
			},
			ExpectedInput: []string{
				"----------------------",
				"3 )  When: some title, or case when,",
				"   Actual: actual rec,",
				" Expected: expected item",
			},
			VerifyTypeOf: nil,
		},
	}

	sortedArrayTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedArray output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"isSort":  true,
				"message": "some message alim, knows, who, do you know --- #alim",
			},
			ExpectedInput: []string{
				"#alim",
				"---",
				"alim,",
				"do",
				"know",
				"knows,",
				"message",
				"some",
				"who,",
				"you",
			},
			VerifyTypeOf: nil,
		},
	}

	sortedMessageTestCases = []coretestcases.CaseV1{
		{
			Title: "SortedMessage output verification",
			ArrangeInput: args.Map{
				"isPrint": false,
				"message": "some message alim, knows, who, do you know --- #alim",
				"joiner":  " | ",
			},
			ExpectedInput: []string{
				"#alim | --- | alim, | do | know | knows, | message | some | who, | you",
			},
			VerifyTypeOf: nil,
		},
	}

	stringsToSpaceStringTestCases = []coretestcases.CaseV1{
		{
			Title: "StringsToSpaceString output verification",
			ArrangeInput: args.Map{
				"spaceCount": 4,
				"lines": []string{
					"#alim",
					"---",
					"alim,",
					"do",
					"",
					"know",
					"when",
					"you",
					"type,",
					"lines",
				},
			},
			ExpectedInput: []string{
				"    #alim",
				"    ---",
				"    alim,",
				"    do",
				"    ",
				"    know",
				"    when",
				"    you",
				"    type,",
				"    lines",
			},
			VerifyTypeOf: nil,
		},
	}
)
