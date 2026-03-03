package corepayloadtests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/corepayload"
	"gitlab.com/auk-go/core/coretests/args"
)

// pagingInfoDiff returns a diff-style string comparing two PagingInfo pointers.
// Prints all fields side-by-side so failures show exactly what differs.
func pagingInfoDiff(label string, left, right *corepayload.PagingInfo) string {
	leftStr := "<nil>"
	rightStr := "<nil>"

	if left != nil {
		leftStr = fmt.Sprintf(
			"{TotalPages:%d, CurrentPageIndex:%d, PerPageItems:%d, TotalItems:%d}",
			left.TotalPages, left.CurrentPageIndex, left.PerPageItems, left.TotalItems,
		)
	}

	if right != nil {
		rightStr = fmt.Sprintf(
			"{TotalPages:%d, CurrentPageIndex:%d, PerPageItems:%d, TotalItems:%d}",
			right.TotalPages, right.CurrentPageIndex, right.PerPageItems, right.TotalItems,
		)
	}

	return fmt.Sprintf(
		"\n[%s]\n  Left:  %s\n  Right: %s",
		label, leftStr, rightStr,
	)
}

// pagingInfoStateString returns all state check results as a formatted string for diff output.
func pagingInfoStateString(label string, info *corepayload.PagingInfo) string {
	infoStr := "<nil>"
	if info != nil {
		infoStr = fmt.Sprintf(
			"{TotalPages:%d, CurrentPageIndex:%d, PerPageItems:%d, TotalItems:%d}",
			info.TotalPages, info.CurrentPageIndex, info.PerPageItems, info.TotalItems,
		)
	}

	return fmt.Sprintf(
		"\n[%s] Input: %s\n"+
			"  IsEmpty:%v | HasTotalPages:%v, HasCurrentPageIndex:%v, HasPerPageItems:%v, HasTotalItems:%v\n"+
			"  IsInvalidTotalPages:%v, IsInvalidCurrentPageIndex:%v, IsInvalidPerPageItems:%v, IsInvalidTotalItems:%v",
		label, infoStr,
		info.IsEmpty(),
		info.HasTotalPages(), info.HasCurrentPageIndex(), info.HasPerPageItems(), info.HasTotalItems(),
		info.IsInvalidTotalPages(), info.IsInvalidCurrentPageIndex(), info.IsInvalidPerPageItems(), info.IsInvalidTotalItems(),
	)
}

func Test_PagingInfo_IsEqual_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isLeftNil, _ := input.GetAsBool("isLeftNil")
		isRightNil, _ := input.GetAsBool("isRightNil")

		var left, right *corepayload.PagingInfo
		if !isLeftNil {
			left = buildPagingInfoPrefixed(input, "left")
		}
		if !isRightNil {
			right = buildPagingInfoPrefixed(input, "right")
		}

		// Act
		result := left.IsEqual(right)

		// Assert
		actual := fmt.Sprintf("%v", result)
		expected := testCase.ExpectedInput.([]string)

		if actual != expected[0] {
			t.Errorf("Case %d [%s] FAILED: expected IsEqual=%s, got=%s%s",
				caseIndex,
				testCase.Title,
				expected[0],
				actual,
				pagingInfoDiff("diff", left, right),
			)
		}

		testCase.ShouldBeEqual(t, caseIndex, actual)
	}
}

func Test_PagingInfo_State_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoStateTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil, _ := input.GetAsBool("isNil")

		var info *corepayload.PagingInfo
		if !isNil {
			info = buildPagingInfoFromMap(input)
		}

		// Act
		results := []string{
			fmt.Sprintf("%v", info.IsEmpty()),
			fmt.Sprintf("%v", info.HasTotalPages()),
			fmt.Sprintf("%v", info.HasCurrentPageIndex()),
			fmt.Sprintf("%v", info.HasPerPageItems()),
			fmt.Sprintf("%v", info.HasTotalItems()),
			fmt.Sprintf("%v", info.IsInvalidTotalPages()),
			fmt.Sprintf("%v", info.IsInvalidCurrentPageIndex()),
			fmt.Sprintf("%v", info.IsInvalidPerPageItems()),
			fmt.Sprintf("%v", info.IsInvalidTotalItems()),
		}

		// Diff-style failure output
		expected := testCase.ExpectedInput.([]string)
		hasMismatch := false
		for i, r := range results {
			if i < len(expected) && r != expected[i] {
				hasMismatch = true
				break
			}
		}

		if hasMismatch {
			t.Errorf("Case %d [%s] FAILED%s\n  Actual:   %v\n  Expected: %v",
				caseIndex,
				testCase.Title,
				pagingInfoStateString("input", info),
				results,
				expected,
			)
		}

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_PagingInfo_Clone_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoCloneTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		info := buildPagingInfoFromMap(input)

		// Act
		clone := info.Clone()

		// Assert
		results := []string{
			fmt.Sprintf("%v", clone.TotalPages),
			fmt.Sprintf("%v", clone.CurrentPageIndex),
			fmt.Sprintf("%v", clone.PerPageItems),
			fmt.Sprintf("%v", clone.TotalItems),
		}

		expected := testCase.ExpectedInput.([]string)
		hasMismatch := false
		for i, r := range results {
			if i < len(expected) && r != expected[i] {
				hasMismatch = true
				break
			}
		}

		if hasMismatch {
			original := info
			cloned := &clone
			t.Errorf("Case %d [%s] FAILED%s",
				caseIndex,
				testCase.Title,
				pagingInfoDiff("Clone", original, cloned),
			)
		}

		testCase.ShouldBeEqual(t, caseIndex, results...)
	}
}

func Test_PagingInfo_ClonePtr_Verification(t *testing.T) {
	for caseIndex, testCase := range pagingInfoClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		isNil, _ := input.GetAsBool("isNil")

		var info *corepayload.PagingInfo
		if !isNil {
			info = buildPagingInfoFromMap(input)
		}

		// Act
		result := info.ClonePtr()

		// Assert
		if isNil {
			isResultNil := fmt.Sprintf("%v", result == nil)
			testCase.ShouldBeEqual(t, caseIndex, isResultNil)
		} else {
			results := []string{
				fmt.Sprintf("%v", result == nil),
				fmt.Sprintf("%v", result.TotalPages),
				fmt.Sprintf("%v", result.CurrentPageIndex),
				fmt.Sprintf("%v", result.PerPageItems),
				fmt.Sprintf("%v", result.TotalItems),
			}

			expected := testCase.ExpectedInput.([]string)
			hasMismatch := false
			for i, r := range results {
				if i < len(expected) && r != expected[i] {
					hasMismatch = true
					break
				}
			}

			if hasMismatch {
				t.Errorf("Case %d [%s] FAILED%s",
					caseIndex,
					testCase.Title,
					pagingInfoDiff("ClonePtr", info, result),
				)
			}

			testCase.ShouldBeEqual(t, caseIndex, results...)
		}
	}
}

// === Independence tests (cannot be table-driven — require mutation) ===

func Test_PagingInfo_ClonePtr_Independence(t *testing.T) {
	// Arrange
	info := &corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	clone := info.ClonePtr()
	clone.TotalPages = 99
	clone.CurrentPageIndex = 99

	// Assert
	if info.TotalPages != 5 || info.CurrentPageIndex != 3 {
		t.Errorf("ClonePtr Independence FAILED - mutating clone affected original%s",
			pagingInfoDiff("after mutation", info, clone),
		)
	}
}

func Test_PagingInfo_Clone_Independence(t *testing.T) {
	// Arrange
	info := corepayload.PagingInfo{TotalPages: 5, CurrentPageIndex: 3, PerPageItems: 10, TotalItems: 50}

	// Act
	clone := info.Clone()
	clone.TotalPages = 99

	// Assert
	if info.TotalPages != 5 {
		clonePtr := &clone
		infoPtr := &info
		t.Errorf("Clone Independence FAILED - mutating clone affected original%s",
			pagingInfoDiff("after mutation", infoPtr, clonePtr),
		)
	}
}
