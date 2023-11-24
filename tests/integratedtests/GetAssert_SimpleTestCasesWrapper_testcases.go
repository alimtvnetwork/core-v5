package integratedtests

import (
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	stringSimpleTestCasesTestCases = []coretestcases.CaseV1{
		{
			Title: "giving []string or slice string - outputs as is.",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving string - output split to lines by newlines",
				ArrangeInput: args.Map{
					"any": "some string contains\nnewline\nin between",
				},
				ActualInput: "some actual text v1! - length okay but diff text",
				ExpectedInput: []string{
					"diff text but not in length!",
				},
				VerifyTypeOf: commonType,
			},
			ExpectedInput: []string{
				"----------------------",
				"0 )  Title:\"giving string - output split to lines by newlines\"",
				"      Input:`args.Map{\"any\":\"some string contains\\nnewline\\nin between\"}` ,",
				"",
				"  Actual:",
				"  `\"some actual text v1! - length okay but diff text\"` ,",
				"",
				"Expected:",
				"  `[]string{\"diff text but not in length!\"}`",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
		{
			Title: "giving []string or slice string - outputs as is.",
			ArrangeInput: coretestcases.CaseV1{
				Title: "giving empty empty slice",
				ArrangeInput: args.Map{
					"any": []string{},
				},
				ActualInput:   "some actual text v2! - empty slice",
				ExpectedInput: []string{},
				VerifyTypeOf:  commonType,
			},
			ExpectedInput: []string{
				"----------------------",
				"1 )  Title:\"giving empty empty slice\"",
				"      Input:`args.Map{\"any\":[]string{}}` ,",
				"",
				"  Actual:",
				"  `\"some actual text v2! - empty slice\"` ,",
				"",
				"Expected:",
				"  `[]string{}`",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf(coretestcases.CaseV1{}),
		},
	}
)
