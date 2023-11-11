package testsinternal

import (
	"fmt"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"
	"gitlab.com/auk-go/core/corevalidator"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
)

// TestCase
//
//   - Title : Test case header
//   - ArrangeInput : Preparing input
//   - ActualInput : Input for the act method
//   - ExpectedInput : Set expectations for the unit test (what we are going receive from invoking something)
//   - Will verify type using VerifyTypeOf
type TestCase coretests.BaseTestCase

func (it *TestCase) SetActual(actual interface{}) {
	it.ActualInput = actual
}

func (it *TestCase) SetExpected(expected interface{}) {
	it.ExpectedInput = expected
}

func (it *TestCase) VerifyAllEqual(
	caseIndex int,
	actualElement ...string,
) error {
	return it.VerifyAll(
		caseIndex,
		stringcompareas.Equal,
		actualElement,
	)
}

func (it *TestCase) VerifyAll(
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElements []string,
) error {
	it.SetActual(actualElements)

	sliceValidator := corevalidator.SliceValidator{
		ValidatorCoreCondition: corevalidator.DefaultTrimCoreCondition,
		CompareAs:              compareAs,
		ActualLines:            actualElements,
		ExpectedLines:          it.ExpectedInput.([]string),
	}

	return it.VerifyAllSliceValidator(
		caseIndex,
		sliceValidator)
}

func (it *TestCase) VerifyAllSliceValidator(
	caseIndex int,
	validator corevalidator.SliceValidator,
) error {
	baseParameter := corevalidator.ValidatorParamsBase{
		CaseIndex:          caseIndex,
		Header:             it.Title,
		IsAttachUserInputs: true,
		IsCaseSensitive:    true,
	}

	return validator.AllVerifyError(&baseParameter)
}

func (it *TestCase) VerifyError(
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElement ...string,
) error {
	toBaseTestCase := it.AsBaseTestCase()
	validationFinalError := it.VerifyAll(
		caseIndex,
		compareAs,
		actualElement)

	if toBaseTestCase.IsTypeInvalidOrSkipVerify() {
		return validationFinalError
	}

	typeVerifyErr := toBaseTestCase.TypeValidationError()

	return errcore.MergeErrors(
		validationFinalError,
		typeVerifyErr)
}

func (it *TestCase) Assert(
	t *testing.T,
	caseIndex int,
	compareAs stringcompareas.Variant,
	actualElement ...string,
) error {
	toBaseTestCase := it.AsBaseTestCase()
	validationFinalError := it.VerifyAll(
		caseIndex,
		compareAs,
		actualElement)

	convey.Convey(toBaseTestCase.Title, t, func() {
		errcore.PrintErrorWithTestIndex(
			caseIndex,
			toBaseTestCase.Title,
			validationFinalError)

		convey.So(
			validationFinalError,
			should.BeNil)
	})

	if toBaseTestCase.IsTypeInvalidOrSkipVerify() {
		return validationFinalError
	}

	typeVerifyErr := toBaseTestCase.TypeValidationError()
	typeVerifyTitle := fmt.Sprintf(
		typeVerifyTitleFormat,
		it.Title)

	convey.Convey(typeVerifyTitle, t, func() {
		convey.So(
			typeVerifyErr,
			should.BeNil)
	})

	return errcore.MergeErrors(
		validationFinalError,
		typeVerifyErr)
}

func (it *TestCase) AssertEqual(
	t *testing.T,
	caseIndex int,
	actualElements ...string,
) {
	_ = it.Assert(t,
		caseIndex,
		stringcompareas.Equal,
		actualElements...)
}

func (it TestCase) AsBaseTestCase() coretests.BaseTestCase {
	return coretests.BaseTestCase(it)
}
