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

func Test_StringCompare_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range stringCompareIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		isMatch := sc.IsMatch()
		isMatchFailed := sc.IsMatchFailed()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isMatch),
			fmt.Sprintf("%v", isMatchFailed),
		)
	}
}

func Test_StringCompare_VerifyError_Verification(t *testing.T) {
	for caseIndex, testCase := range stringCompareVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		sc := newStringCompareFromMap(input)

		// Act
		err := sc.VerifyError()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", err != nil),
		)
	}
}

func Test_StringCompare_NilReceiver_Verification(t *testing.T) {
	for caseIndex, testCase := range stringCompareNilReceiverTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		methodName, err := input.GetAsString("method")
		errcore.HandleErrMessage("method required", err)
		var sc *coreinstruction.StringCompare

		// Act
		var result string
		switch methodName {
		case "IsMatch":
			result = fmt.Sprintf("%v", sc.IsMatch())
		case "IsMatchFailed":
			result = fmt.Sprintf("%v", sc.IsMatchFailed())
		case "IsInvalid":
			result = fmt.Sprintf("%v", sc.IsInvalid())
		case "IsDefined":
			result = fmt.Sprintf("%v", sc.IsDefined())
		case "VerifyError":
			result = fmt.Sprintf("%v", sc.VerifyError() != nil)
		}

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			result,
		)
	}
}
