package corevalidator

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func (it *SliceValidator) VerifyFirstError(
	parameter *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.VerifyFirstLengthUptoError(
		parameter,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) VerifyFirstLengthUptoError(
	params *Parameter,
	lengthUpTo int,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		true,
		params,
		lengthUpTo,
	)
}

func (it *SliceValidator) AllVerifyErrorQuick(
	caseIndex int,
	header string,
	actualElements ...string,
) error {
	if it == nil {
		return nil
	}

	var params = Parameter{
		CaseIndex:                  caseIndex,
		Header:                     header,
		IsSkipCompareOnActualEmpty: true,
		IsAttachUserInputs:         true,
		IsCaseSensitive:            true,
	}

	it.SetActual(actualElements)

	return it.AllVerifyErrorUptoLength(
		false,
		&params,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) AllVerifyError(
	params *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength(),
	)
}

func (it *SliceValidator) AllVerifyErrorTestCase(
	caseIndex int,
	header string,
	isCaseSensitive bool,
) error {
	if it == nil {
		return nil
	}

	params := Parameter{
		CaseIndex:                  caseIndex,
		Header:                     header,
		IsSkipCompareOnActualEmpty: false,
		IsAttachUserInputs:         true,
		IsCaseSensitive:            isCaseSensitive,
	}

	err := it.AllVerifyErrorUptoLength(
		false,
		&params,
		it.ExpectingLinesLength(),
	)

	errcore.PrintErrorWithTestIndex(caseIndex, header, err)

	return err
}

// AllVerifyErrorExceptLast
//
// Verify up to the second last item.
func (it *SliceValidator) AllVerifyErrorExceptLast(
	params *Parameter,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ExpectingLinesLength()-1,
	)
}

func (it *SliceValidator) AllVerifyErrorUptoLength(
	isFirstOnly bool,
	params *Parameter,
	lengthUpto int,
) error {
	if it == nil {
		return nil
	}

	if it.isEmptyIgnoreCase(params) {
		return nil
	}

	initialVerifyErr := it.initialVerifyErrorWithMerged(
		params,
		lengthUpto,
	)

	if initialVerifyErr != nil {
		return initialVerifyErr
	}

	lengthErr := it.lengthVerifyError(params, lengthUpto)
	if lengthErr != nil {
		return lengthErr
	}

	validators := it.ComparingValidators()
	var sliceErr []string

	for i, validator := range validators.Items[:lengthUpto] {
		err := validator.VerifySimpleError(
			i,
			params,
			it.ActualLines[i],
		)

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error(),
			)
		}

		if isFirstOnly && err != nil {
			break
		}
	}

	hasErrors := len(sliceErr) > constants.Zero

	if hasErrors {
		diffMsg := errcore.LineDiffToString(
			params.CaseIndex,
			params.Header,
			it.ActualLines,
			it.ExpectedLines,
		)

		if len(diffMsg) > 0 {
			sliceErr = append(sliceErr, diffMsg)
		}
	}

	if params.IsAttachUserInputs && hasErrors {
		sliceErr = append(
			sliceErr,
			it.ActualInputWithExpectingMessage(
				params.CaseIndex,
				params.Header,
			),
		)
	}

	return errcore.SliceToError(sliceErr)
}

func (it *SliceValidator) lengthVerifyError(
	params *Parameter,
	lengthUpto int,
) error {
	hasLengthUpto := lengthUpto > constants.InvalidValue
	comparingLength := it.ExpectingLinesLength()

	var comparingLengthError error
	if hasLengthUpto && lengthUpto > comparingLength {
		comparingLengthError = errcore.OutOfRangeLengthType.Error(
			"Asked comparingLength is out of range!",
			comparingLength,
		)
	}

	if comparingLengthError != nil {
		return it.UserInputsMergeWithError(
			params,
			comparingLengthError,
		)
	}

	var inputLengthErr error
	if it.ActualLinesLength() > 0 && comparingLength == 0 {
		inputLengthErr = errcore.LengthIssueType.Error(
			"Input comparison has some text but comparing length is 0! Must set comparing text!",
			comparingLength,
		)
	}

	if inputLengthErr != nil {
		return it.UserInputsMergeWithError(
			params,
			inputLengthErr,
		)
	}

	return nil
}

// initialVerifyError, verifyLengthUpto less than 0 will check actual length
func (it *SliceValidator) initialVerifyError(
	lengthUpto int,
) error {
	if it.ActualLines == nil && it.ExpectedLines == nil {
		return nil
	}

	isAnyNilCase := it.ActualLines == nil ||
		it.ExpectedLines == nil

	if isAnyNilCase {
		return errcore.ExpectingErrorSimpleNoTypeNewLineEnds(
			"ActualLines, ExpectedLines any is nil and other is not.",
			it.ActualLines,
			it.ExpectedLines,
		)
	}

	if !it.isLengthOkay(lengthUpto) {
		return errcore.ExpectingErrorSimpleNoTypeNewLineEnds(
			"ActualLines, ExpectedLines Length is not equal.",
			len(it.ActualLines),
			len(it.ExpectedLines),
		)
	}

	return nil
}

func (it *SliceValidator) initialVerifyErrorWithMerged(
	params *Parameter,
	lengthUpto int,
) error {
	initialVerifyErr := it.initialVerifyError(
		lengthUpto,
	)

	if initialVerifyErr != nil {
		return it.UserInputsMergeWithError(
			params,
			initialVerifyErr,
		)
	}

	return nil
}
