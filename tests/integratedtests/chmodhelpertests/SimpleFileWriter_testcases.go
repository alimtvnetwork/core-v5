package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var (
	nullBothTestCases = []coretestcases.CaseV1{
		{
			Title: "testing any file reading writing",
			ArrangeInput: []args.Two{
				{
					First:  pathInstructionsV3,
					Second: nil,
					Expect: nil,
				},
			},
			ActualInput:   nil,
			ExpectedInput: []string{},
			VerifyTypeOf:  nil,
		},
	}
)
