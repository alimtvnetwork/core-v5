package coretests

import (
	"fmt"

	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssertSimpleTestCaseWrapper struct{}

// GetAssertMessageUsingSimpleTestCaseWrapper
//
//	Gives generic and consistent test message using msgformats.QuickIndexTitleInputActualExpectedMessageFormat
func (it getAssertSimpleTestCaseWrapper) ToString(
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

func (it getAssertSimpleTestCaseWrapper) ToStrings(
	testCaseIndex int,
	testCaseWrapper SimpleTestCaseWrapper,
) []string {
	return fmt.Sprintf(
		msgformats.QuickIndexTitleInputActualExpectedMessageFormat,
		testCaseIndex,
		testCaseWrapper.CaseTitle(),
		testCaseWrapper.Input(),
		testCaseWrapper.Actual(),
		testCaseWrapper.Expected(),
	)
}
