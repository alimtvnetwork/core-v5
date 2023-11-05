package coretests

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/reflectinternal"
	"gitlab.com/auk-go/core/issetter"
)

// BaseTestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
//   - Will verify type at ArrangeExpectedType
//   - Set by Developer expectation or assumptions
//   - ArrangeExpectedType : Verify type for the ActualInput
//   - ActualExpectedType : Verify type for the ExpectedInput
//   - ExpectedTypeOfExpected : Verify type for the actual method return type, ExpectedInput type == ActualExpectedType == ExpectedTypeOfExpected (all 3 should match)
type BaseTestCase struct {
	Title           string         // consider as header
	ArrangeInput    interface{}    // preparing input, initial input
	ActualInput     interface{}    // (dynamically set) : must be set after running Act, using SetActual
	ExpectedInput   interface{}    // expectation set from the test
	VerifyTypeOf    *VerifyTypeOf  // Verify the type of ArrangeInput, ActualInput, ExpectedInput
	IsEnable        issetter.Value // Only false makes it disabled.
	HasError        bool
	IsValidateError bool
}

func (it *BaseTestCase) CaseTitle() string {
	return it.Title
}

func (it *BaseTestCase) TypesValidationMustPasses(t *testing.T) {
	err := it.TypeValidationError()

	if err != nil {
		t.Error(
			"any one of the type validation failed",
			err.Error())
	}
}

// TypeValidationError
//
// must use SetActual to set the actual,
// what received from the act method,
// set it using SetActual
func (it *BaseTestCase) TypeValidationError() error {
	if it.VerifyTypeOf.IsInvalid() {
		return nil
	}

	var sliceErr []string
	arrangeInputActualType := reflect.TypeOf(it.ArrangeInput)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)
	verifyOf := it.VerifyTypeOf

	if reflectinternal.IsNotNull(it.ArrangeInput) && arrangeInputActualType != verifyOf.ArrangeInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Arrange Type Mismatch",
				verifyOf.ArrangeInput,
				arrangeInputActualType))
	}

	if reflectinternal.IsNotNull(it.ActualInput) && actualInputActualType != verifyOf.ActualInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Actual Type Mismatch",
				verifyOf.ActualInput,
				actualInputActualType))
	}

	if reflectinternal.IsNotNull(it.ExpectedInput) && expectedInputActualType != verifyOf.ExpectedInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Expected Type Mismatch",
				verifyOf.ExpectedInput,
				expectedInputActualType))
	}

	return errcore.SliceToError(sliceErr)
}

// ArrangeString
//
//	returns ArrangeInput in string
//	format using constants.SprintValueFormat
func (it *BaseTestCase) ArrangeString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ArrangeInput)
}

// Input returns ArrangeInput
func (it *BaseTestCase) Input() interface{} {
	return it.ArrangeInput
}

func (it *BaseTestCase) Expected() interface{} {
	return it.ExpectedInput
}

func (it *BaseTestCase) ExpectedString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ExpectedInput)
}

func (it *BaseTestCase) Actual() interface{} {
	return it.ActualInput
}

func (it *BaseTestCase) ActualString() string {
	return fmt.Sprintf(
		constants.SprintValueFormat,
		it.ActualInput)
}

func (it *BaseTestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

// String
//
//	returns a string format using GetAssertMessageUsingSimpleTestCaseWrapper
//	- https://prnt.sc/lxUV0eYk_qlg
func (it *BaseTestCase) String(caseIndex int) string {
	return GetAssertMessageUsingSimpleTestCaseWrapper(
		caseIndex, it)
}

func (it *BaseTestCase) IsDisabled() bool {
	return it.IsEnable.IsFalse()
}

func (it *BaseTestCase) IsSkipWithLog(caseIndex int) bool {
	if it.IsDisabled() {
		fmt.Printf(
			"Header : %s (%d), skipped: Disabled.",
			it.Title,
			caseIndex)

		return true
	}

	return false
}

// ShouldBe
//
// Disabled testcases will not be executed.
func (it *BaseTestCase) ShouldBe(
	caseIndex int,
	t *testing.T,
	assert convey.Assertion,
	actual interface{},
) {
	if it.IsEnable.IsFalse() {
		fmt.Printf(
			skippedMsgFormat,
			it.Title,
			caseIndex)

		return
	}

	it.ShouldBeExplicit(
		true,
		caseIndex,
		t,
		it.Title,
		actual,
		assert,
		it.Expected())
}

func (it *BaseTestCase) ShouldBeExplicit(
	isValidateType bool,
	caseIndex int,
	t *testing.T,
	title string,
	actual interface{},
	assert convey.Assertion,
	expected interface{},
) {
	if it.IsEnable.IsFalse() {
		t.Skipf(
			skippedMsgFormat,
			it.Title,
			caseIndex)
	}

	it.SetActual(actual)

	convey.Convey(title, t, func() {
		convey.SoMsg(it.String(caseIndex), actual, assert, expected)
	})

	if !isValidateType {
		return
	}

	err := it.TypeValidationError()
	errHeader := fmt.Sprintf(
		"case %d : test case type validation must passes - ",
		caseIndex)

	if err != nil {
		err = errors.New(errHeader + err.Error() + ", case title : " + title)
	}

	convey.Convey(errHeader, t, func() {
		convey.So(err, convey.ShouldBeNil)
	})
}

func (it *BaseTestCase) asSimpleTestCaseWrapper() SimpleTestCaseWrapper {
	return it
}

func (it *BaseTestCase) asBaseTestCaseWrapper() BaseTestCaseWrapper {
	return it
}
