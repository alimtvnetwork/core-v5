package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/tests/testwrappers/chmodhelpertestwrappers"
)

var applyOnPathTestCases = []coretestcases.CaseV1{
	{
		Title:         "Apply rwx owner=*-x group=**x other=-w- on paths",
		ArrangeInput:  chmodhelpertestwrappers.RwxInstructionsApplyTestCases[0],
		ExpectedInput: "",
	},
}
