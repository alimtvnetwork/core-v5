package coretestcasestests

import (
	"fmt"
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
	"gitlab.com/auk-go/core/errcore"
)

func Test_GenericGherkins_IsFailedToMatch_WhenMatching(t *testing.T) {
	tc := isFailedToMatchWhenMatchingTestCase

	// Arrange — IsMatching is true

	// Act
	result := tc.IsFailedToMatch()

	// Assert
	actLines := []string{fmt.Sprintf("%v", result)}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_IsFailedToMatch_WhenNotMatching(t *testing.T) {
	tc := isFailedToMatchWhenNotMatchingTestCase

	// Arrange — IsMatching is false

	// Act
	result := tc.IsFailedToMatch()

	// Assert
	actLines := []string{fmt.Sprintf("%v", result)}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_CompareWith_Equal(t *testing.T) {
	tc := compareWithEqualTestCase

	// Arrange
	a := &coretestcases.StringBoolGherkins{
		Title: "same",
		Input: "hello",
	}
	b := &coretestcases.StringBoolGherkins{
		Title: "same",
		Input: "hello",
	}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actLines := []string{
		fmt.Sprintf("%v", isEqual),
		diff,
	}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_CompareWith_DiffTitle(t *testing.T) {
	tc := compareWithDiffTitleTestCase

	// Arrange
	a := &coretestcases.StringBoolGherkins{
		Title: "A",
		Input: "hello",
	}
	b := &coretestcases.StringBoolGherkins{
		Title: "B",
		Input: "hello",
	}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actLines := []string{
		fmt.Sprintf("%v", isEqual),
		diff,
	}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_CompareWith_DiffInput(t *testing.T) {
	tc := compareWithDiffInputTestCase

	// Arrange
	a := &coretestcases.StringBoolGherkins{
		Title: "same",
		Input: "alpha",
	}
	b := &coretestcases.StringBoolGherkins{
		Title: "same",
		Input: "beta",
	}

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actLines := []string{
		fmt.Sprintf("%v", isEqual),
		diff,
	}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_CompareWith_BothNil(t *testing.T) {
	tc := compareWithBothNilTestCase

	// Arrange
	var a *coretestcases.StringBoolGherkins
	var b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actLines := []string{
		fmt.Sprintf("%v", isEqual),
		diff,
	}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_CompareWith_OneNil(t *testing.T) {
	tc := compareWithOneNilTestCase

	// Arrange
	a := &coretestcases.StringBoolGherkins{Title: "exists"}
	var b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert
	actLines := []string{
		fmt.Sprintf("%v", isEqual),
		diff,
	}
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_FullString_Basic(t *testing.T) {
	tc := fullStringBasicTestCase

	// Arrange
	g := &coretestcases.StringBoolGherkins{
		Title:      "FullString includes all fields",
		Feature:    "regex",
		Given:      "a valid pattern",
		When:       "struct has all fields populated",
		Then:       "output is formatted",
		Input:      "test-pattern",
		Expected:   true,
		Actual:     false,
		IsMatching: true,
	}

	// Act
	result := g.FullString()
	actLines := strings.Split(strings.TrimRight(result, "\n"), "\n")

	// Assert
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_FullString_Nil(t *testing.T) {
	tc := fullStringNilTestCase

	// Arrange
	var g *coretestcases.StringBoolGherkins

	// Act
	result := g.FullString()
	actLines := []string{result}

	// Assert
	tc.ShouldBeEqualUsingExpected(t, 0, actLines)
}

func Test_GenericGherkins_ShouldBeEqualArgs_Passing(t *testing.T) {
	tc := shouldBeEqualPassingTestCase

	// Arrange — expected lines defined in test case

	// Act + Assert
	tc.ShouldBeEqualArgs(
		t,
		0,
		"line-a",
		"line-b",
	)
}

func Test_GenericGherkins_CaseTitle_UsesTitle(t *testing.T) {
	tc := caseTitleUseTitleTestCase

	// Arrange — Title is set

	// Act
	result := tc.CaseTitle()

	// Assert
	actLines := []string{result}
	errcore.AssertDiffOnMismatch(
		t,
		0,
		"CaseTitle uses Title",
		actLines,
		tc.ExpectedLines,
	)
}

func Test_GenericGherkins_CaseTitle_FallsBackToWhen(t *testing.T) {
	tc := caseTitleFallbackToWhenTestCase

	// Arrange — Title is empty

	// Act
	result := tc.CaseTitle()

	// Assert
	actLines := []string{result}
	errcore.AssertDiffOnMismatch(
		t,
		0,
		"CaseTitle falls back to When",
		actLines,
		tc.ExpectedLines,
	)
}
