package coretests

import (
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/internal/msgformats"
)

type getAssert struct {
	SimpleTestCaseWrapper getAssertSimpleTestCaseWrapper
}

// Quick
//
// Gives generic and consistent
// test message using msgformats.QuickIndexInputActualExpectedMessageFormat
func (it getAssert) Quick(
	when,
	actual,
	expected interface{},
	counter int,
) string {
	return fmt.Sprintf(
		msgformats.QuickIndexInputActualExpectedMessageFormat,
		counter,
		when,
		actual,
		expected,
	)
}

func (it getAssert) SortedMessage(
	isPrint bool,
	message,
	joiner string,
) string {
	whitespaceRemovedSplits := it.SortedArray(
		isPrint,
		true,
		message,
	)

	return strings.Join(whitespaceRemovedSplits, joiner)
}

func (it getAssert) SortedArray(
	isPrint bool,
	isSort bool,
	message string,
) []string {
	if isPrint {
		fmt.Println(message)
	}

	return SplitByEachWordTrimmedNoSpace(
		message,
		isSort,
	)
}

// SortedArrayNoPrint
//
// isPrint: false, isSort: true
func (it getAssert) SortedArrayNoPrint(
	message string,
) []string {
	return it.SortedArray(
		false,
		true, message,
	)
}
