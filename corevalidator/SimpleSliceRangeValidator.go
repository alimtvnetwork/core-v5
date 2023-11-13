package corevalidator

import (
	"fmt"

	"gitlab.com/auk-go/core/coredata/corestr"
)

type SimpleSliceRangeValidator struct {
	actual           *corestr.SimpleSlice
	VerifierSegments []RangesSegment
}

func (it *SimpleSliceRangeValidator) LengthOfVerifierSegments() int {
	return len(it.VerifierSegments)
}

func (it *SimpleSliceRangeValidator) SetActual(
	lines []string,
) *SimpleSliceRangeValidator {
	it.actual = corestr.New.SimpleSlice.Direct(
		false,
		lines)

	return it
}

func (it *SimpleSliceRangeValidator) Validators() HeaderSliceValidators {
	validators := make([]HeaderSliceValidator, it.LengthOfVerifierSegments())

	for _, segment := range it.VerifierSegments {
		expectedSegments := segment.ExpectedLines
		start := segment.RangeInt.Start
		end := segment.RangeInt.End
		actualSegments := it.actual.Items[start:end]
		totalItems := end - start + 1
		header := fmt.Sprintf("validate for range %d to %d (total: %d lines)",
			start,
			end,
			totalItems)
		validator := HeaderSliceValidator{
			Header: header,
			SliceValidator: SliceValidator{
				CompareAs:              segment.CompareAs,
				ValidatorCoreCondition: segment.ValidatorCoreCondition,
				ActualLines:            actualSegments,
				ExpectedLines:          expectedSegments,
			},
		}

		validators = append(validators, validator)
	}

	return validators
}

func (it *SimpleSliceRangeValidator) VerifyAll(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	it.SetActual(actual)

	return it.Validators().VerifyAll(
		header,
		params,
		isPrintError)
}

func (it *SimpleSliceRangeValidator) VerifyFirst(
	header string,
	actual []string,
	params *Parameter,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyFirst(
		params,
		isPrintError)
}

func (it *SimpleSliceRangeValidator) VerifyUpto(
	header string,
	actual []string,
	params *Parameter,
	length int,
	isPrintError bool,
) error {
	params.Header = header
	it.SetActual(actual)

	return it.Validators().VerifyUpto(
		isPrintError,
		false,
		length,
		params,
	)
}
