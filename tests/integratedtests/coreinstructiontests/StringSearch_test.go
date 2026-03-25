package coreinstructiontests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coreinstruction"
	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
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
		actual := args.Map{
			"isMatch":       ss.IsMatch(content),
			"isMatchFailed": ss.IsMatchFailed(content),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
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
		actual := args.Map{
			"isAllMatch":       ss.IsAllMatch(contents...),
			"isAnyMatchFailed": ss.IsAnyMatchFailed(contents...),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
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
		actual := args.Map{
			"isEmpty": ss.IsEmpty(),
			"isExist": ss.IsExist(),
			"has":     ss.Has(),
		}

		// Assert
		testCase.ShouldBeEqualMap(
			t,
			caseIndex,
			actual,
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
