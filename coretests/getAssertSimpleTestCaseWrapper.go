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

func (it getAssertSimpleTestCaseWrapper) ToQuoteLines(
	tabCount int,
	lines []string,
) []string {
	return errcore.StringLinesToQuoteLinesWithTabs(
		tabCount,
		lines)
}

func (it getAssertSimpleTestCaseWrapper) AnyToQuoteLines(
	tabCount int,
	anyItem interface{},
) []string {
	lines := convertinteranl.AnyTo.Strings(anyItem)

	return it.ToQuoteLines(
		tabCount,
		lines)
}

func (it getAssertSimpleTestCaseWrapper) StringQuoteLine(
	tabCount int,
	lines []string,
) string {
	finalLines := it.ToQuoteLines(
		tabCount,
		lines)

	return strings.Join(finalLines, constants.NewLineUnix)
}

func (it getAssertSimpleTestCaseWrapper) AnyToStringQuoteLine(
	tabCount int,
	anyItem interface{},
) string {
	lines := convertinteranl.AnyTo.Strings(anyItem)

	return it.StringQuoteLine(tabCount, lines)
}

// StringByLines
//
// Actual lines and then expected lines.
func (it getAssertSimpleTestCaseWrapper) StringByLines(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	toStringsFunc := it.AnyToQuoteLines
	prefixSpaces := 4
	actualLines := toStringsFunc(prefixSpaces, testCaseWrapper.Actual())
	expectedLines := toStringsFunc(prefixSpaces, testCaseWrapper.Expected())

	actual := errcore.StringLinesToQuoteLinesToSingle(actualLines)
	expected := errcore.StringLinesToQuoteLinesToSingle(expectedLines)
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
