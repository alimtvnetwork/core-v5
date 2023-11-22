package integratedtests

import (
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range quickTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Four)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.QuickGherkins

		// Act
		actualSlice.AppendFmt(
			"%s",
			quickFunc(
				input.First,        // when
				input.Second,       // actual
				input.Third,        // expected
				input.Fourth.(int), // counter
			),
			input,
			input,
		)

		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
