package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

func newStringCompareFromMap(input args.Map) *coreinstruction.StringCompare {
	method, _ := input.GetAsString("method")
	search, _ := input.GetAsString("search")
	content, _ := input.GetAsString("content")

	isIgnoreCaseRaw, _ := input.Get("isIgnoreCase")
	isIgnoreCase, _ := isIgnoreCaseRaw.(bool)

	switch method {
	case "equal":
		return coreinstruction.NewStringCompareEqual(search, content)
	case "contains":
		return coreinstruction.NewStringCompareContains(isIgnoreCase, search, content)
	case "startsWith":
		return coreinstruction.NewStringCompareStartsWith(isIgnoreCase, search, content)
	case "endsWith":
		return coreinstruction.NewStringCompareEndsWith(isIgnoreCase, search, content)
	case "regex":
		return coreinstruction.NewStringCompareRegex(search, content)
	default:

		return coreinstruction.NewStringCompare(
			stringcompareas.Equal,
			isIgnoreCase,
			search,
			content,
		)
	}
}

func Test_StringCompare_IsMatch(t *testing.T) {
	for caseIndex, testCase := range stringCompareIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		actual := args.Map{
			"isMatch":       sc.IsMatch(),
			"isMatchFailed": sc.IsMatchFailed(),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
		)
	}
}

func Test_StringCompare_VerifyError(t *testing.T) {
	for caseIndex, testCase := range stringCompareVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		hasErr := fmt.Sprintf("%v", sc.VerifyError() != nil)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, hasErr)
	}
}

// ==========================================
// Nil Receiver — CaseNilSafe pattern
// ==========================================

func Test_StringCompare_NilReceiver(t *testing.T) {
	for caseIndex, tc := range stringCompareNilSafeTestCases {
		// Arrange (implicit — nil receiver)

		// Act & Assert
		tc.ShouldBeSafe(t, caseIndex)
	}
}
