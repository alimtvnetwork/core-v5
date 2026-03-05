package coreinstructiontests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coreinstruction"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/enums/stringcompareas"
)

func newStringSearchFromMap(input args.Map) *coreinstruction.StringSearch {
	method, _ := input.GetAsString("method")
	search, _ := input.GetAsString("search")

	var compareMethod stringcompareas.Variant
	switch method {
	case "equal":
		compareMethod = stringcompareas.Equal
	case "contains":
		compareMethod = stringcompareas.Contains
	case "startsWith":
		compareMethod = stringcompareas.StartsWith
	case "endsWith":
		compareMethod = stringcompareas.EndsWith
	case "regex":
		compareMethod = stringcompareas.Regex
	default:
		compareMethod = stringcompareas.Equal
	}

	return &coreinstruction.StringSearch{
		CompareMethod: compareMethod,
		Search:        search,
	}
}

func Test_StringSearch_IsMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchIsMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ss := newStringSearchFromMap(input)
		content, _ := input.GetAsString("content")

		// Act
		isMatch := ss.IsMatch(content)
		isMatchFailed := ss.IsMatchFailed(content)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isMatch),
			fmt.Sprintf("%v", isMatchFailed),
		)
	}
}

func Test_StringSearch_IsAllMatch_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchIsAllMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		ss := newStringSearchFromMap(input)
		contents, _ := input.GetAsStrings("contents")

		// Act
		isAllMatch := ss.IsAllMatch(contents...)
		isAnyFailed := ss.IsAnyMatchFailed(contents...)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isAllMatch),
			fmt.Sprintf("%v", isAnyFailed),
		)
	}
}

func Test_StringSearch_State_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil, _ := isNilVal.(bool)

		var ss *coreinstruction.StringSearch
		if !isNil {
			ss = newStringSearchFromMap(input)
		}

		// Act
		isEmpty := ss.IsEmpty()
		isExist := ss.IsExist()
		has := ss.Has()

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isEmpty),
			fmt.Sprintf("%v", isExist),
			fmt.Sprintf("%v", has),
		)
	}
}

func Test_StringSearch_VerifyError_Verification(t *testing.T) {
	for caseIndex, testCase := range stringSearchVerifyErrorTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNilVal, _ := input.Get("isNil")
		isNil, _ := isNilVal.(bool)
		content, _ := input.GetAsString("content")

		var ss *coreinstruction.StringSearch
		if !isNil {
			ss = newStringSearchFromMap(input)
		}

		// Act
		err := ss.VerifyError(content)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", err != nil),
		)
	}
}
