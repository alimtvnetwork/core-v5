package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/errcore"
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

// StringByLines
//
// Actual lines and then expected lines.
func (it getAssertSimpleTestCaseWrapper) StringByLines(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) string {
	toStringsFunc := GetAssert.ToStrings
	actualLines := toStringsFunc(testCaseWrapper.Actual())
	expectedLines := toStringsFunc(testCaseWrapper.Expected())

	actual := errcore.StringLinesToQuoteLinesToSingle(actualLines)
	expected := errcore.StringLinesToQuoteLinesToSingle(expectedLines)

	return fmt.Sprintf(
		msgformats.QuickIndexTitleInputActualExpectedMessageFormat,
		testCaseIndex,
		testCaseWrapper.CaseTitle(),
		testCaseWrapper.Input(),
		actual,
		expected,
	)
}
