package chmodhelpertests

import (
	"gitlab.com/auk-go/core/coretests/coretestcases"
)

var linuxApplyRecursiveOnPathTestCase = coretestcases.CaseV1{
	Title:         "Apply recursive rwx owner=*-x group=**x other=-w- on paths",
	ArrangeInput:  rwxInstructionsUnixApplyRecursivelyTestCase,
	ExpectedInput: "",
}
