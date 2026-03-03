package coredynamictests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coredata/coredynamic"
	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/errcore"
)

// ==========================================
// Diff helper — prints map state on failure
// ==========================================

func getBool(input args.Map, key string) bool {
	v, ok := input[key]
	if !ok {
		return false
	}

	b, isBool := v.(bool)
	if !isBool {
		return false
	}

	return b
}

// ==========================================
// IsEqual — table-driven
// ==========================================

func Test_MapAnyItems_IsEqual(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsIsEqualTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := getBool(input, "leftNil")
		rightNil := getBool(input, "rightNil")

		var left *coredynamic.MapAnyItems
		var right *coredynamic.MapAnyItems

		if !leftNil {
			leftMap := input["leftMap"].(map[string]any)
			left = coredynamic.NewMapAnyItemsUsingItems(leftMap)
		}
		if !rightNil {
			rightMap := input["rightMap"].(map[string]any)
			right = coredynamic.NewMapAnyItemsUsingItems(rightMap)
		}

		// Act
		result := left.IsEqual(right)

		// Print diff on failure
		resultStr := fmt.Sprintf("%v", result)
		expected := testCase.ExpectedInput.([]string)

		diag := MapDiffDiagnostics{
			CaseIndex: caseIndex,
			Title:     testCase.Title,
			Left:      left,
			Right:     right,
		}
		diag.PrintIfResultMismatch(resultStr, expected)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

// ==========================================
// IsEqualRaw — table-driven
// ==========================================

func Test_MapAnyItems_IsEqualRaw(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsIsEqualRawTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := getBool(input, "leftNil")

		var m *coredynamic.MapAnyItems
		if !leftNil {
			leftMap := input["leftMap"].(map[string]any)
			m = coredynamic.NewMapAnyItemsUsingItems(leftMap)
		}

		var rawMap map[string]any
		if rm, ok := input["rightMap"]; ok {
			rawMap = rm.(map[string]any)
		}

		// Act
		result := m.IsEqualRaw(rawMap)

		// Print diff on failure
		resultStr := fmt.Sprintf("%v", result)
		expected := testCase.ExpectedInput.([]string)

		diag := MapDiffDiagnostics{
			CaseIndex: caseIndex,
			Title:     testCase.Title,
			Left:      m,
			RawMap:    rawMap,
		}
		diag.PrintIfResultMismatch(resultStr, expected)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, resultStr)
	}
}

// ==========================================
// ClonePtr — table-driven
// ==========================================

func Test_MapAnyItems_ClonePtr(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsClonePtrTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := getBool(input, "leftNil")
		addAfterClone := getBool(input, "addAfterClone")

		var m *coredynamic.MapAnyItems
		if !leftNil {
			leftMap := input["leftMap"].(map[string]any)
			m = coredynamic.NewMapAnyItemsUsingItems(leftMap)
		}

		// Act
		clone, err := m.ClonePtr()

		var actLines []string

		hasError := err != nil
		cloneIsNil := clone == nil
		actLines = append(actLines, fmt.Sprintf("%v", hasError))
		actLines = append(actLines, fmt.Sprintf("%v", cloneIsNil))

		if !cloneIsNil && !hasError {
			actLines = append(actLines, fmt.Sprintf("%d", clone.Length()))

			if addAfterClone {
				clone.Add("new_key", "new_val")
				actLines = append(actLines, fmt.Sprintf("%v", m.HasKey("new_key")))
				actLines = append(actLines, fmt.Sprintf("%v", clone.HasKey("new_key")))
			} else {
				if _, ok := input["leftMap"]; ok {
					leftMap := input["leftMap"].(map[string]any)
					if _, has := leftMap["name"]; has {
						actLines = append(actLines, fmt.Sprintf("%v", clone.HasKey("name")))
						actLines = append(actLines, fmt.Sprintf("%v", clone.HasKey("age")))
					}
				}
			}
		}

		// Print diff on failure
		expected := testCase.ExpectedInput.([]string)

		diag := MapDiffDiagnostics{
			CaseIndex: caseIndex,
			Title:     testCase.Title,
			Left:      m,
			Clone:     clone,
			Error:     err,
		}
		diag.PrintIfMismatch(actLines, expected)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}

// ==========================================
// Edge cases — table-driven
// ==========================================

func Test_MapAnyItems_EdgeCases(t *testing.T) {
	for caseIndex, testCase := range mapAnyItemsEdgeCaseTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		leftNil := getBool(input, "leftNil")

		var m *coredynamic.MapAnyItems
		if !leftNil {
			leftMap := input["leftMap"].(map[string]any)
			m = coredynamic.NewMapAnyItemsUsingItems(leftMap)
		}

		var actLines []string

		// Act — branch by what the test case exercises
		if key, ok := input["addKey"]; ok {
			addKey := key.(string)
			addValue := input["addValue"]
			isNew := m.Add(addKey, addValue)
			actLines = append(actLines, fmt.Sprintf("%v", isNew))

			expected := testCase.ExpectedInput.([]string)
			if len(expected) > 1 {
				if expected[1] == "new" || expected[1] == "old" {
					val := m.GetValue(addKey)
					actLines = append(actLines, fmt.Sprintf("%v", val))
				} else {
					actLines = append(actLines, fmt.Sprintf("%d", m.Length()))
				}
			}
		} else if key, ok := input["key"]; ok {
			actLines = append(actLines, fmt.Sprintf("%v", m.HasKey(key.(string))))
		} else {
			actLines = append(actLines, fmt.Sprintf("%d", m.Length()))
			actLines = append(actLines, fmt.Sprintf("%v", m.IsEmpty()))
			actLines = append(actLines, fmt.Sprintf("%v", m.HasAnyItem()))
		}

		// Print diff on failure
		expected := testCase.ExpectedInput.([]string)

		diag := MapDiffDiagnostics{
			CaseIndex: caseIndex,
			Title:     testCase.Title,
			Left:      m,
		}
		diag.PrintIfMismatch(actLines, expected)

		// Assert
		testCase.ShouldBeEqual(t, caseIndex, actLines...)
	}
}
