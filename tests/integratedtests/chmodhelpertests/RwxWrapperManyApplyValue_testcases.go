package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/chmodhelper"
	"github.com/alimtvnetwork/core/chmodhelper/chmodclasstype"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

type rwxWrapperManyApplyTestCase struct {
	Case      coretestcases.CaseV1
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
