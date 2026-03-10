package isanytests

import (
	"testing"

	"github.com/alimtvnetwork/core/conditional"
	"github.com/alimtvnetwork/core/corecsv"
	"github.com/alimtvnetwork/core/coredata/corestr"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/coretests/coretestcases"
	"github.com/alimtvnetwork/core/isany"
)

func Test_Conclusive_Verification(t *testing.T) {
	for caseIndex, testCase := range conclusiveTestCases {
		// Arrange
		inputs := testCase.
			ArrangeInput.([]args.TwoAny)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(len(inputs))

		// Act
		for i, parameter := range inputs {
			f := parameter.First
			s := parameter.Second
			isEqual, isConclusive := isany.Conclusive(f, s)
			values := corecsv.AnyToValuesTypeString(f, s)
			conclusive := conditional.IfString(
				isConclusive,
				"Conclusive",
				"Inconclusive",
			)

			actualSlice.AppendFmt(
				conclusiveFormat,
				i,
				isEqual,
				conclusive,
				values,
			)
		}

		finalActLines := actualSlice.Strings()
		finalTestCase := coretestcases.
			CaseV1(testCase.BaseTestCase)

		// Assert
		finalTestCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
