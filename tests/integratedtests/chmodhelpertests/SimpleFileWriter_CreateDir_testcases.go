package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/internal/pathinternal"
)

var (
	dirCreateBasePath = pathinternal.JoinTemp("core", "case-dir-create")

	createDirTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/some-dir/first.txt",
						"/some-dir-2/first.txt",
						"/some-dir-3/first.txt",
					},
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
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}
)
