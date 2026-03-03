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
		isMatch := fmt.Sprintf("%v", sc.IsMatch())
		isMatchFailed := fmt.Sprintf("%v", sc.IsMatchFailed())

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, isMatch, isMatchFailed)
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
// Nil Receiver — per-method tests
// ==========================================

func Test_StringCompare_NilReceiver_IsMatch(t *testing.T) {
	for caseIndex, tc := range stringCompareNilReceiverTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		if method != "IsMatch" {
			continue
		}

		var sc *coreinstruction.StringCompare

		// Act
		result := fmt.Sprintf("%v", sc.IsMatch())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringCompare_NilReceiver_IsMatchFailed(t *testing.T) {
	for caseIndex, tc := range stringCompareNilReceiverTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		if method != "IsMatchFailed" {
			continue
		}

		var sc *coreinstruction.StringCompare

		// Act
		result := fmt.Sprintf("%v", sc.IsMatchFailed())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringCompare_NilReceiver_IsInvalid(t *testing.T) {
	for caseIndex, tc := range stringCompareNilReceiverTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		if method != "IsInvalid" {
			continue
		}

		var sc *coreinstruction.StringCompare

		// Act
		result := fmt.Sprintf("%v", sc.IsInvalid())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringCompare_NilReceiver_IsDefined(t *testing.T) {
	for caseIndex, tc := range stringCompareNilReceiverTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		if method != "IsDefined" {
			continue
		}

		var sc *coreinstruction.StringCompare

		// Act
		result := fmt.Sprintf("%v", sc.IsDefined())

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}

func Test_StringCompare_NilReceiver_VerifyError(t *testing.T) {
	for caseIndex, tc := range stringCompareNilReceiverTestCases {
		// Arrange
		input := tc.ArrangeInput.(args.Map)
		method, _ := input.GetAsString("method")
		if method != "VerifyError" {
			continue
		}

		var sc *coreinstruction.StringCompare

		// Act
		result := fmt.Sprintf("%v", sc.VerifyError() != nil)

		// Assert
		tc.ShouldBeEqual(t, caseIndex, result)
	}
}
