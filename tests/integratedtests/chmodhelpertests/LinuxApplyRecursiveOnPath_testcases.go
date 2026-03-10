package chmodhelpertests

import (
	"github.com/alimtvnetwork/core/coretests/coretestcases"
)

var linuxApplyRecursiveOnPathTestCase = coretestcases.CaseV1{
	Title:         "Apply recursive rwx owner=*-x group=**x other=-w- on paths",
	ArrangeInput:  rwxInstructionsUnixApplyRecursivelyTestCase,
	ExpectedInput: "",
}
