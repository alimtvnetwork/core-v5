package chmodhelpertests

import (
	"testing"

	"github.com/alimtvnetwork/core/chmodhelper/chmodins"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/errcore"
)

type rwxCompileValueTestCase struct {
	Case                      coretestcases.CaseV1
	Existing, Input, Expected chmodins.RwxOwnerGroupOther
}

// ShouldBeEqual asserts actLines match expectedLines using
// the embedded Case.Title, with optional context lines for diagnostics.
func (it rwxCompileValueTestCase) ShouldBeEqual(
	t *testing.T,
	caseIndex int,
	actLines []string,
	expectedLines []string,
	contextLines ...string,
) {
	t.Helper()

	errcore.AssertDiffOnMismatch(
		t,
		caseIndex,
		it.Case.Title,
		actLines,
		expectedLines,
		contextLines...,
	)
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
