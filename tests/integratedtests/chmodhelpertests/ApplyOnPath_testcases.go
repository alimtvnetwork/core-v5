package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/tests/testwrappers/chmodhelpertestwrappers"
)

var applyOnPathTestCases = []coretestcases.CaseV1{
	{
		Title:         "Apply rwx owner=*-x group=**x other=-w- on paths",
		ArrangeInput:  chmodhelpertestwrappers.RwxInstructionsApplyTestCases[0],
		ExpectedInput: "",
	},
}
