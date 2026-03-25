package stringcompareastests

import (
	"fmt"
	"testing"

	"github.com/alimtvnetwork/core/coretests/args"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
)

func Test_Glob_Match_Verification(t *testing.T) {
	for caseIndex, testCase := range globMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, _ := input.GetAsString("pattern")
		content, _ := input.GetAsString("content")
		isIgnoreCaseVal, _ := input.Get("isIgnoreCase")
		isIgnoreCase, _ := isIgnoreCaseVal.(bool)

		// Act
		isGlobMatch := stringcompareas.Glob.IsCompareSuccess(isIgnoreCase, content, pattern)
		isNonGlobMatch := stringcompareas.NonGlob.IsCompareSuccess(isIgnoreCase, content, pattern)

		actual := args.Map{
			"isMatch":   fmt.Sprintf("%v", isGlobMatch),
			"isInverse": fmt.Sprintf("%v", isNonGlobMatch),
		}

		// Assert
		testCase.ShouldBeEqualMap(t, caseIndex, actual)
	}
}

func Test_Glob_IsGlob_ReturnsTrue(t *testing.T) {
	if !stringcompareas.Glob.IsGlob() {
		t.Error("Glob.IsGlob() should return true")
	}
}

func Test_Glob_IsNonGlob_ReturnsTrue(t *testing.T) {
	if !stringcompareas.NonGlob.IsNonGlob() {
		t.Error("NonGlob.IsNonGlob() should return true")
	}
}

func Test_NonGlob_IsNegativeCondition(t *testing.T) {
	if !stringcompareas.NonGlob.IsNegativeCondition() {
		t.Error("NonGlob should be a negative condition")
	}
}

func Test_Glob_IsNotNegativeCondition(t *testing.T) {
	if stringcompareas.Glob.IsNegativeCondition() {
		t.Error("Glob should not be a negative condition")
	}
}

func Test_Glob_Name(t *testing.T) {
	name := stringcompareas.Glob.Name()
	if name != "Glob" {
		t.Errorf("expected 'Glob', got '%s'", name)
	}
}

func Test_NonGlob_Name(t *testing.T) {
	name := stringcompareas.NonGlob.Name()
	if name != "NonGlob" {
		t.Errorf("expected 'NonGlob', got '%s'", name)
	}
}
