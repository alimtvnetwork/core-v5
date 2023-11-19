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
						"/if/some-dir/first.txt",
						"/if/some-dir-2/first.txt",
						"/if/some-dir-3/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core\\case-dir-create\\if\\some-dir - isCreated : true",
				"0 - 1 : core\\case-dir-create\\if\\some-dir-2 - isCreated : true",
				"0 - 2 : core\\case-dir-create\\if\\some-dir-3 - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}
)
