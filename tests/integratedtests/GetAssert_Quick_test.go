package integratedtests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/coretests/args"
)

func Test_GetAssert_Quick_Verification(t *testing.T) {
	for caseIndex, testCase := range quickTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.Quick

		// Act
		output := quickFunc(
			input.When(),                        // when
			input.Actual(),                      // actual
			input.Expect(),                      // expected
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

func Test_GetAssert_SortedMessage_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedMessageTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		quickFunc := asserter.SortedMessage

		// Act
		outputs := quickFunc(
			input["isPrint"].(bool),
			input["message"].(string),
			input["joiner"].(string),
		)

		actualSlice.Add(outputs)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_GetAssert_SortedArray_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedArrayTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		sortedArray := asserter.SortedArray

		// Act
		outputs := sortedArray(
			input["isPrint"].(bool),
			input["isSort"].(bool),
			input["message"].(string),
		)

		actualSlice.Adds(outputs...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
func Test_GetAssert_SortedArrayNoPrint_Verification(t *testing.T) {
	for caseIndex, testCase := range sortedArrayTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		actFunc := asserter.SortedArrayNoPrint

		// Act
		outputs := actFunc(
			input["message"].(string),
		)

		actualSlice.Adds(outputs...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}

func Test_GetAssert_StringsToSpaceString_Verification(t *testing.T) {
	for caseIndex, testCase := range stringsToSpaceStringTestCases {
		// Arrange
		input := testCase.
			ArrangeInput.(args.Map)
		actualSlice := corestr.
			New.
			SimpleSlice.
			Cap(0)
		asserter := coretests.GetAssert
		actFunc := asserter.StringsToSpaceString

		// Act
		outputs := actFunc(
			input["spaceCount"].(int),
			input["lines"].([]string)...,
		)

		actualSlice.Adds(outputs...)
		finalActLines := actualSlice.Strings()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			finalActLines...,
		)
	}
}
