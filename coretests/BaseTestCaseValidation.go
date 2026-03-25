package coretests

import (
	"reflect"
	"testing"

	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/reflectinternal"
)

// =============================================================================
// Type Validation
// =============================================================================

func (it *BaseTestCase) TypesValidationMustPasses(t *testing.T) {
	err := it.TypeValidationError()

	if err != nil {
		t.Error(
			"any one of the type validation failed",
			err.Error(),
		)
	}
}

// TypeValidationError
//
// must use SetActual to set the actual,
// what received from the act method,
// set it using SetActual
func (it *BaseTestCase) TypeValidationError() error {
	if it.IsTypeInvalidOrSkipVerify() {
		return nil
	}

	var sliceErr []string
	arrangeInputActualType := reflect.TypeOf(it.ArrangeInput)
	actualInputActualType := reflect.TypeOf(it.ActualInput)
	expectedInputActualType := reflect.TypeOf(it.ExpectedInput)
	verifyOf := it.VerifyTypeOf

	if reflectinternal.Is.Defined(it.ArrangeInput) && arrangeInputActualType != verifyOf.ArrangeInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Arrange Type Mismatch",
				verifyOf.ArrangeInput,
				arrangeInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ActualInput) && actualInputActualType != verifyOf.ActualInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Actual Type Mismatch",
				verifyOf.ActualInput,
				actualInputActualType,
			),
		)
	}

	if reflectinternal.Is.Defined(it.ExpectedInput) && expectedInputActualType != verifyOf.ExpectedInput {
		sliceErr = append(
			sliceErr,
			errcore.ExpectingSimpleNoType(
				"Expected Type Mismatch",
				verifyOf.ExpectedInput,
				expectedInputActualType,
			),
		)
	}

	if len(sliceErr) > 0 {
		var newSlice []string

		newSlice = append(
			newSlice,
			it.Title,
		)
		sliceErr = append(
			newSlice,
			sliceErr...,
		)
	}

	return errcore.SliceToError(sliceErr)
}
