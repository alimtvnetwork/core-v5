package codegen

const testPkgHeaderTemplate = `
package $packageName

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"

	$newPackages
)
`

const funcTemplate = `
func Test_$FuncName_$Behaviour(t *testing.T) {
	for caseIndex, testCase := range $testCaseName {
		// Arrange
		input := testCase.
			ArrangeInput.($ArrangeType)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap($linesPossible)
		$actArgsSetup

		// Act
		actFunc$FuncName := $directFuncInvoke
		$outArgs := actFunc$FuncName($inArgs)

		actualSlice.AppendFmt(
			"$fmtJoin",
			caseIndex,
			$fmtOutputs,
		)

		finalActLines := actualSlice.Strings()
		actualSlice.Dispose()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
`
