package coretests

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

// =============================================================================
// Input / Output Getters and Formatting
// =============================================================================

// Input returns ArrangeInput
func (it *BaseTestCase) Input() any {
	return it.ArrangeInput
}

func (it *BaseTestCase) Expected() any {
	return it.ExpectedInput
}

func (it *BaseTestCase) ExpectedString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ExpectedInput,
	)
}

func (it *BaseTestCase) Actual() any {
	return it.ActualInput
}

func (it *BaseTestCase) ActualLines() []string {
	return GetAssert.ToStrings(it.ActualInput)
}

func (it *BaseTestCase) ExpectedLines() []string {
	return GetAssert.ToStrings(it.ExpectedInput)
}

func (it *BaseTestCase) ActualString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ActualInput,
	)
}

func (it *BaseTestCase) SetActual(actual any) {
	it.ActualInput = actual
}

// String
//
//	returns a string format using GetAssertMessageUsingSimpleTestCaseWrapper
//	- https://prnt.sc/lxUV0eYk_qlg
func (it *BaseTestCase) String(caseIndex int) string {
	return GetAssert.SimpleTestCaseWrapper.String(
		caseIndex, it,
	)
}

func (it *BaseTestCase) LinesString(caseIndex int) string {
	return GetAssert.SimpleTestCaseWrapper.CaseLinesUsingDoubleQuoteLinesToString(
		caseIndex, it,
	)
}

func (it *BaseTestCase) IsDisabled() bool {
	return it.IsEnable.IsFalse()
}

func (it *BaseTestCase) IsSkipWithLog(caseIndex int) bool {
	if it.IsDisabled() {
		fmt.Printf(
			"Header : %s (%d), skipped: Disabled.",
			it.Title,
			caseIndex,
		)

		return true
	}

	return false
}

func (it *BaseTestCase) FormTitle(caseIndex int) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		it.Title,
	)
}

func (it *BaseTestCase) CustomTitle(caseIndex int, title string) string {
	return fmt.Sprintf(
		skippedMsgFormat,
		caseIndex,
		title,
	)
}
