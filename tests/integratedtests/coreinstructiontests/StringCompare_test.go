package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

func newStringCompareFromMap(input args.Map) *coreinstruction.StringCompare {
	method, _ := input.GetAsString("method")
	search, _ := input.GetAsString("search")
	content, _ := input.GetAsString("content")
	isIgnoreCase, _ := input.GetAsBool("isIgnoreCase")

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
		actLines := []string{
			fmt.Sprintf("%v", sc.IsMatch()),
			fmt.Sprintf("%v", sc.IsMatchFailed()),
		}
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

func Test_StringCompare_VerifyError(t *testing.T) {
	for caseIndex, testCase := range stringCompareVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		err := sc.VerifyError()
		actLines := []string{fmt.Sprintf("%v", err != nil)}
		expectedLines := testCase.ExpectedInput.([]string)

		// Assert
		errcore.AssertDiffOnMismatch(t, caseIndex, testCase.Title, actLines, expectedLines)
	}
}

// ==========================================
// Nil Receiver — per-method tests
// ==========================================

func Test_StringCompare_NilReceiver_IsMatch(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actLines := []string{fmt.Sprintf("%v", sc.IsMatch())}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver - IsMatch returns true (vacuous truth)", actLines, expectedLines)
}

func Test_StringCompare_NilReceiver_IsMatchFailed(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actLines := []string{fmt.Sprintf("%v", sc.IsMatchFailed())}
	expectedLines := []string{"false"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver - IsMatchFailed returns false", actLines, expectedLines)
}

func Test_StringCompare_NilReceiver_IsInvalid(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actLines := []string{fmt.Sprintf("%v", sc.IsInvalid())}
	expectedLines := []string{"true"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver - IsInvalid returns true", actLines, expectedLines)
}

func Test_StringCompare_NilReceiver_IsDefined(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actLines := []string{fmt.Sprintf("%v", sc.IsDefined())}
	expectedLines := []string{"false"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver - IsDefined returns false", actLines, expectedLines)
}

func Test_StringCompare_NilReceiver_VerifyError(t *testing.T) {
	// Arrange
	var sc *coreinstruction.StringCompare

	// Act
	actLines := []string{fmt.Sprintf("%v", sc.VerifyError() != nil)}
	expectedLines := []string{"false"}

	// Assert
	errcore.AssertDiffOnMismatch(t, 0, "Nil receiver - VerifyError returns nil", actLines, expectedLines)
}
