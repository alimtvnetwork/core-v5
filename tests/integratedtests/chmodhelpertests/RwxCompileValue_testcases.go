package chmodhelpertests

import (
	"gitlab.com/auk-go/core/chmodhelper/chmodins"
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

type rwxCompileValueTestCase struct {
	Case                         coretestcases.CaseV1
	Existing, Input, Expected    chmodins.RwxOwnerGroupOther
}

var rwxCompileValueTestCases = []rwxCompileValueTestCase{
	{
		Case: coretestcases.CaseV1{
			Title: "Existing [rwx,r-x,r--] Applied by [*-x,**x,-w-] should result [r-x,r-x,-w-]",
		},
		Existing: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "r--",
		},
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "*-x",
			Group: "**x",
			Other: "-w-",
		},
		Expected: chmodins.RwxOwnerGroupOther{
			Owner: "r-x",
			Group: "r-x",
			Other: "-w-",
		},
	},
	{
		Case: coretestcases.CaseV1{
			Title: "Existing [rwx,r--,--x] Applied by [***,**x,-w*] should result [rwx,r-x,-wx]",
		},
		Existing: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r--",
			Other: "--x",
		},
		Input: chmodins.RwxOwnerGroupOther{
			Owner: "***",
			Group: "**x",
			Other: "-w*",
		},
		Expected: chmodins.RwxOwnerGroupOther{
			Owner: "rwx",
			Group: "r-x",
			Other: "-wx",
		},
	},
}
