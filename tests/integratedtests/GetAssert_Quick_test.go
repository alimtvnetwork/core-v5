package integratedtests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_QuickGherkins_Verification(t *testing.T) {
	for caseIndex, testCase := range quickTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Dynamic)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.QuickGherkins
		params := input.Params

		// Act
		output := quickFunc(
			params["when"],                      // when
			params["actual"],                    // actual
			params["expect"],                    // expected
			input.GetAsIntDefault("counter", 0), // counter
		)

		actualSlice.Adds(strings.Split(output, "\n")...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
