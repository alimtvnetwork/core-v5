package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var (
	dirFilesWithContentCreateReadTestCases = []coretestcases.CaseV1{
		{
			Title: "DirFilesWithContent - testing any file reading - writing",
			ArrangeInput: []args.OneAny{
				{
					First: pathInstructionsV3,
				},
			},
			ExpectedInput: []string{
				"0 : file-1.txt",
				"         0. some lines",
				"         1. alim",
				"0 : file-2.txt",
				"         0. some lines file - 2",
				"         1. alim",
				"0 : file-3.txt",
				"         0. some lines file - 3",
				"         1. alim",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]args.OneAny{}),
		},
	}
)
