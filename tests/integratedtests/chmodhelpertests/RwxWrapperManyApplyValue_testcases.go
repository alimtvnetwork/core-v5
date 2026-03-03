package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper"
	"gitlab.com/auk-go/core/chmodhelper/chmodclasstype"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type rwxWrapperManyApplyTestCase struct {
	Case     coretestcases.CaseV1
	SingleRwx chmodhelper.SingleRwx
}

var rwxWrapperManyApplyTestCases = []rwxWrapperManyApplyTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Apply r-x on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "r-x",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply --- on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "---",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply --x on Other class",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "--x",
			ClassType: chmodclasstype.Other,
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Apply r-x on Other class (duplicate verify)",
		},
		SingleRwx: chmodhelper.SingleRwx{
			Rwx:       "r-x",
			ClassType: chmodclasstype.Other,
		},
	},
}
