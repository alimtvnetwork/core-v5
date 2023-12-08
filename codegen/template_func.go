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
		actFunc$FuncName := args.FuncDetector.GetFuncWrap(input)
		$returnArgs := actFunc$FuncName($actArgs)

		actualSlice.Add("$FuncName($valueActArgs) ->")
		actualSlice.AppendFmt(
			"$fmtJoin",
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
