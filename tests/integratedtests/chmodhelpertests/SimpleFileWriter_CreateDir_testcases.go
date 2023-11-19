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

	createDirIfMissingTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - if-missing",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/if-missing/some-dir/first.txt",
						"/if-missing/some-dir-2/first.txt",
						"/if-missing/some-dir-3/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core\\case-dir-create\\if-missing\\some-dir - isCreated : true",
				"0 - 1 : core\\case-dir-create\\if-missing\\some-dir-2 - isCreated : true",
				"0 - 2 : core\\case-dir-create\\if-missing\\some-dir-3 - isCreated : true",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}

	createDirDirectTestCases = []coretestcases.CaseV1{
		{
			Title: "create dir check - direct create - if exist fails",
			ArrangeInput: []chmodhelper.DirWithFiles{
				{
					Dir: dirCreateBasePath,
					Files: []string{
						"/direct-create/some-dir/first.txt",
						"/direct-create/some-dir-2/first.txt",
						"/direct-create/some-dir-3/first.txt",
					},
				},
			},
			ExpectedInput: []string{
				"0 - 0 : core\\case-dir-create\\direct-create\\some-dir\\first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 1 : core\\case-dir-create\\direct-create\\some-dir-2\\first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
				"0 - 2 : core\\case-dir-create\\direct-create\\some-dir-3\\first.txt - already exist as file, err: dir : , applyChmod :-rwxr-xr-x, path exist but it is not a dir.",
			},
			VerifyTypeOf: coretests.NewVerifyTypeOf([]chmodhelper.DirWithFiles{}),
		},
	}
)
