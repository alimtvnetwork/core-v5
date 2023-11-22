package coretests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/convertinteranl"
	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssertSimpleTestCaseWrapper struct{}

// GetAssertMessageUsingSimpleTestCaseWrapper
//
//	Gives generic and consistent test message using msgformats.QuickIndexTitleInputActualExpectedMessageFormat
func (it getAssertSimpleTestCaseWrapper) String(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexTitleInputActualExpectedMessageFormat,
		testCaseIndex,
		testCaseWrapper.CaseTitle(),
		testCaseWrapper.Input(),
		testCaseWrapper.Actual(),
		testCaseWrapper.Expected(),
	)
}

func (it getAssertSimpleTestCaseWrapper) Lines(
	testCaseWrapper SimpleTestCaseWrapper,
) (actualLines, expectedLines []string) {
	toStringsFunc := GetAssert.ToStrings
	actualLines = toStringsFunc(testCaseWrapper.Actual())
	expectedLines = toStringsFunc(testCaseWrapper.Expected())

	return actualLines, expectedLines
}

// ToQuoteLines
//
// Converts from below lines to
//
//	line 1,
//	line 2,
//	line 3,
//
// Converts a strings lines to
//
//	{spaces} "line 1",
//	{spaces} "line 2",
//	{spaces} "line 3",
func (it getAssertSimpleTestCaseWrapper) ToQuoteLines(
	tabCount int,
	lines []string,
) []string {
	return errcore.LinesToDoubleQuoteLinesWithTabs(
		tabCount,
		lines,
	)
}

func (it getAssertSimpleTestCaseWrapper) AnyToSingleQuoteLines(
	tabCount int,
	anyItem interface{},
) []string {
	lines := convertinteranl.AnyTo.Strings(anyItem)

	return it.ToQuoteLines(
		tabCount,
		lines,
	)
}

func (it getAssertSimpleTestCaseWrapper) SingleQuoteLinesToString(
	tabCount int,
	lines []string,
) string {
	finalLines := it.ToQuoteLines(
		tabCount,
		lines,
	)

	return strings.Join(finalLines, constants.NewLineUnix)
}

func (it getAssertSimpleTestCaseWrapper) AnyToStringQuoteLine(
	tabCount int,
	anyItem interface{},
) string {
	lines := convertinteranl.AnyTo.Strings(anyItem)

	return it.SingleQuoteLinesToString(tabCount, lines)
}

// StringByLines
//
// Actual lines and then expected lines.
func (it getAssertSimpleTestCaseWrapper) StringByLines(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	toStringsFunc := GetAssert.ToStrings
	prefixSpaces := 4
	actualLines := toStringsFunc(testCaseWrapper.Actual())
	expectedLines := toStringsFunc(testCaseWrapper.Expected())

	actual := it.SingleQuoteLinesToString(prefixSpaces, actualLines)
	expected := it.SingleQuoteLinesToString(prefixSpaces, expectedLines)
	title := testCaseWrapper.CaseTitle()

	return fmt.Sprintf(
		msgformats.QuickLinesFormat,
		testCaseIndex,
		title,
		testCaseIndex,
		title,
		actual,
		testCaseIndex,
		title,
		expected,
	)
}
