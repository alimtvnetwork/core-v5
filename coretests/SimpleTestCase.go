package coretests

import (
	"fmt"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
)

// SimpleTestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
type SimpleTestCase struct {
	Title         string      // consider as header
	ArrangeInput  interface{} // preparing input, initial input
	ActualInput   interface{} // (dynamically set) : must be set after running Act, using SetActual
	ExpectedInput interface{} // expectation set from the test
}

func (it *SimpleTestCase) CaseTitle() string {
	return it.Title
}

// ArrangeString
//
//	returns ArrangeInput in string
//	format using constants.SprintValueFormat
func (it *SimpleTestCase) ArrangeString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ArrangeInput)
}

// Input returns ArrangeInput
func (it *SimpleTestCase) Input() interface{} {
	return it.ArrangeInput
}

func (it *SimpleTestCase) Expected() interface{} {
	return it.ExpectedInput
}

func (it *SimpleTestCase) ExpectedString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ExpectedInput)
}

func (it *SimpleTestCase) Actual() interface{} {
	return it.ActualInput
}

func (it *SimpleTestCase) ActualString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ActualInput)
}

func (it *SimpleTestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

// String
//
//	returns a string format using GetAssertMessageUsingSimpleTestCaseWrapper
//	- https://prnt.sc/lxUV0eYk_qlg
func (it *SimpleTestCase) String(caseIndex int) string {
	return GetAssertMessageUsingSimpleTestCaseWrapper(
		caseIndex, it)
}

// ShouldBe
//
// Disabled testcases will not be executed.
func (it *SimpleTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	it.SetActual(actual)

	convey.Convey(it.Title, t, func() {
		convey.SoMsg(
			it.String(caseIndex),
			actual,
			assert,
			it.ExpectedInput)
	})
}

func (it *SimpleTestCase) AsSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}
