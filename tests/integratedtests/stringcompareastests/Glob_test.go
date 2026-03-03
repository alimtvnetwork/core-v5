package stringcompareastests

import (
	"fmt"
	"testing"

	"gitlab.com/auk-go/core/coretests/args"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

func Test_Glob_Match_Verification(t *testing.T) {
	for caseIndex, testCase := range globMatchTestCases {
		// Arrange
		input := testCase.ArrangeInput.(args.Map)
		pattern, err := input.GetAsString("pattern")
		errcore.HandleErrMessage("pattern required", err)
		content, err := input.GetAsString("content")
		errcore.HandleErrMessage("content required", err)
		isIgnoreCase, _ := input.GetAsBool("isIgnoreCase")

		// Act
		isGlobMatch := stringcompareas.Glob.IsCompareSuccess(isIgnoreCase, content, pattern)
		isNotGlobMatch := stringcompareas.NotGlob.IsCompareSuccess(isIgnoreCase, content, pattern)

		// Assert
		testCase.ShouldBeEqual(
			t,
			caseIndex,
			fmt.Sprintf("%v", isGlobMatch),
			fmt.Sprintf("%v", isNotGlobMatch),
		)
	}
}

func Test_Glob_IsGlob_ReturnsTrue(t *testing.T) {
	if !stringcompareas.Glob.IsGlob() {
		t.Error("Glob.IsGlob() should return true")
	}
}

func Test_Glob_IsNotGlob_ReturnsTrue(t *testing.T) {
	if !stringcompareas.NotGlob.IsNotGlob() {
		t.Error("NotGlob.IsNotGlob() should return true")
	}
}

func Test_NotGlob_IsNegativeCondition(t *testing.T) {
	if !stringcompareas.NotGlob.IsNegativeCondition() {
		t.Error("NotGlob should be a negative condition")
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

func Test_NotGlob_Name(t *testing.T) {
	name := stringcompareas.NotGlob.Name()
	if name != "NotGlob" {
		t.Errorf("expected 'NotGlob', got '%s'", name)
	}
}
