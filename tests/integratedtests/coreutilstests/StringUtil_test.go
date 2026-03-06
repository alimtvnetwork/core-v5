package coreutilstests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func Test_IsEmpty_Verification(t *testing.T) {
	for caseIndex, testCase := range isEmptyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsEmpty(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsBlank_Verification(t *testing.T) {
	for caseIndex, testCase := range isBlankTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsBlank(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsEmptyOrWhitespace_Verification(t *testing.T) {
	for caseIndex, testCase := range isEmptyOrWhitespaceTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.IsEmptyOrWhitespace(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_SafeSubstring_Verification(t *testing.T) {
	for caseIndex, testCase := range safeSubstringTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		start, _ := input.GetAsInt("start")
		end, _ := input.GetAsInt("end")

		// Act
		result := stringutil.SafeSubstring(content, start, end)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_SplitLeftRight_Verification(t *testing.T) {
	for caseIndex, testCase := range splitLeftRightTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")
		separator, _ := input.GetAsString("separator")

		// Act
		left, right := stringutil.SplitLeftRight(inputStr, separator)
		actual := args.Map{
			"left":  left,
			"right": right,
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_IsStartsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range isStartsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		startsWith, _ := input.GetAsString("startsWith")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true

		// Act
		result := stringutil.IsStartsWith(content, startsWith, isIgnoreCase)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_IsEndsWith_Verification(t *testing.T) {
	for caseIndex, testCase := range isEndsWithTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		endsWith, _ := input.GetAsString("endsWith")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase := isIgnoreCaseVal == true

		// Act
		result := stringutil.IsEndsWith(content, endsWith, isIgnoreCase)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, fmt.Sprintf("%v", result))
	}
}

func Test_RemoveMany_Verification(t *testing.T) {
	for caseIndex, testCase := range removeManyTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		content, _ := input.GetAsString("content")
		removesRaw, _ := input.Get("removes")
		removes := removesRaw.([]string)

		// Act
		result := stringutil.RemoveMany(content, removes...)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_ReplaceWhiteSpacesToSingle_Verification(t *testing.T) {
	for caseIndex, testCase := range replaceWhiteSpacesToSingleTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		inputStr, _ := input.GetAsString("input")

		// Act
		result := stringutil.ReplaceWhiteSpacesToSingle(inputStr)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, result)
	}
}
