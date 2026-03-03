package coretestcasestests

import (
	"strings"
	"testing"

	"gitlab.com/auk-go/core/coretests/coretestcases"
)

func Test_GenericGherkins_IsFailedToMatch_WhenMatching(t *testing.T) {
	tc := isFailedToMatchWhenMatchingTestCase

	// Arrange — IsMatching is true

	// Act
	result := tc.IsFailedToMatch()

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		result,
	)
}

func Test_GenericGherkins_IsFailedToMatch_WhenNotMatching(t *testing.T) {
	tc := isFailedToMatchWhenNotMatchingTestCase

	// Arrange — IsMatching is false

	// Act
	result := tc.IsFailedToMatch()

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		result,
	)
}

func Test_GenericGherkins_CompareWith_Equal(t *testing.T) {
	tc := compareWithEqualTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		isEqual,
	)
	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		[]string{diff},
	)
}

func Test_GenericGherkins_CompareWith_DiffTitle(t *testing.T) {
	tc := compareWithDiffTitleTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		isEqual,
	)
	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		[]string{diff},
	)
}

func Test_GenericGherkins_CompareWith_DiffInput(t *testing.T) {
	tc := compareWithDiffInputTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	b := tc.GetExtra("b").(*coretestcases.StringBoolGherkins)

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		isEqual,
	)
	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		[]string{diff},
	)
}

func Test_GenericGherkins_CompareWith_BothNil(t *testing.T) {
	tc := compareWithBothNilTestCase

	// Arrange
	var a *coretestcases.StringBoolGherkins
	var b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		isEqual,
	)
	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		[]string{diff},
	)
}

func Test_GenericGherkins_CompareWith_OneNil(t *testing.T) {
	tc := compareWithOneNilTestCase

	// Arrange
	a := tc.GetExtra("a").(*coretestcases.StringBoolGherkins)
	var b *coretestcases.StringBoolGherkins

	// Act
	isEqual, diff := a.CompareWith(b)

	// Assert

	tc.ShouldMatchExpectedFirst(
		t,
		isEqual,
	)
	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		[]string{diff},
	)
}

func Test_GenericGherkins_FullString_Basic(t *testing.T) {
	tc := fullStringBasicTestCase

	// Arrange
	g := tc.GetExtra("subject").(*coretestcases.StringBoolGherkins)

	// Act
	result := g.FullString()
	actLines := strings.Split(strings.TrimRight(result, "\n"), "\n")

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_GenericGherkins_FullString_Nil(t *testing.T) {
	tc := fullStringNilTestCase

	// Arrange
	var g *coretestcases.StringBoolGherkins

	// Act
	result := g.FullString()
	actLines := []string{result}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_GenericGherkins_ShouldBeEqualArgs_Passing(t *testing.T) {
	tc := shouldBeEqualPassingTestCase

	// Arrange — expected lines defined in test case

	// Act + Assert

	tc.ShouldBeEqualArgsFirst(
		t,
		"line-a",
		"line-b",
	)
}

func Test_GenericGherkins_CaseTitle_UsesTitle(t *testing.T) {
	tc := caseTitleUseTitleTestCase

	// Arrange — Title is set

	// Act
	result := tc.CaseTitle()
	actLines := []string{result}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}

func Test_GenericGherkins_CaseTitle_FallsBackToWhen(t *testing.T) {
	tc := caseTitleFallbackToWhenTestCase

	// Arrange — Title is empty

	// Act
	result := tc.CaseTitle()
	actLines := []string{result}

	// Assert

	tc.ShouldBeEqualUsingExpectedFirst(
		t,
		actLines,
	)
}
