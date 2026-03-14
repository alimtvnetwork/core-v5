package corevalidator

import (
	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/errcore"
)

// SliceValidator
//
// Use this only for one time verification only.
//
// If IsUsedAlready, don't mutate the ActualLines or ExpectedLines
// it will not work.
type SliceValidator struct {
	Condition
	CompareAs stringcompareas.Variant
	// ActualLines considered to be actual
	// ExpectedLines considered to be expected
	ActualLines, ExpectedLines []string
	comparingValidators        *TextValidators // lazy
	isUsed                     bool
}

func (it *SliceValidator) IsUsedAlready() bool {
	if it == nil {
		return false
	}

	return it.isUsed
}

func (it *SliceValidator) ActualLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.ActualLines)
}

func (it *SliceValidator) MethodName() string {
	return it.CompareAs.Name()
}

func (it *SliceValidator) SetActual(
	actual []string,
) *SliceValidator {
	it.ActualLines = actual
	it.isUsed = true

	return it
}

func (it *SliceValidator) SetActualVsExpected(
	actual, expected []string,
) *SliceValidator {
	it.ActualLines = actual
	it.ExpectedLines = expected
	it.isUsed = true

	return it
}

func (it *SliceValidator) ActualLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return errcore.StringLinesToQuoteLinesToSingle(
		it.ActualLines,
	)
}

func (it *SliceValidator) ExpectingLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return errcore.StringLinesToQuoteLinesToSingle(
		it.ExpectedLines,
	)
}

func (it *SliceValidator) ExpectingLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.ExpectedLines)
}

func (it *SliceValidator) ComparingValidators() *TextValidators {
	if it.comparingValidators != nil {
		return it.comparingValidators
	}

	validators := NewTextValidators(it.ExpectingLinesLength())

	for _, line := range it.ExpectedLines {
		validators.Add(
			TextValidator{
				Search:    line,
				Condition: it.Condition,
				SearchAs:  it.CompareAs,
			},
		)
	}

	it.comparingValidators = validators

	return it.comparingValidators
}

func (it *SliceValidator) IsValid(isCaseSensitive bool) bool {
	if it == nil {
		return true
	}

	return it.isValidLines(
		isCaseSensitive,
		it.ActualLines,
	)
}

func (it *SliceValidator) IsValidOtherLines(
	isCaseSensitive bool,
	otherActualLines []string,
) bool {
	return it.
		isValidLines(
			isCaseSensitive,
			otherActualLines,
		)
}

func (it *SliceValidator) isValidLines(
	isCaseSensitive bool,
	lines []string,
) bool {
	if it == nil && lines == nil {
		return true
	}

	if lines == nil && it.ExpectedLines == nil {
		return true
	}

	if lines == nil || it.ExpectedLines == nil {
		return false
	}

	inputLength := len(lines)
	comparingLength := len(it.ExpectedLines)

	if inputLength != comparingLength {
		return false
	}

	validators := it.ComparingValidators()

	for i, validator := range validators.Items {
		isNotMatch := !validator.IsMatch(
			lines[i],
			isCaseSensitive,
		)

		if isNotMatch {
			return false
		}
	}

	return true
}

func (it *SliceValidator) isEmptyIgnoreCase(
	params *Parameter,
) bool {
	return params.IsSkipCompareOnActualEmpty &&
		len(it.ActualLines) == 0
}

func (it *SliceValidator) isLengthOkay(lengthUpto int) bool {
	inputLength := len(it.ActualLines)
	comparingLength := len(it.ExpectedLines)
	isLengthCheckUpto := lengthUpto > constants.InvalidValue
	var isMinLengthMeet bool

	if isLengthCheckUpto {
		remainingInputLength := inputLength - lengthUpto
		remainingCompareLength := comparingLength - lengthUpto

		isMinLengthMeet = remainingInputLength == remainingCompareLength
	}

	isLengthOkay := isMinLengthMeet ||
		inputLength == comparingLength

	return isLengthOkay
}

func (it *SliceValidator) Dispose() {
	if it == nil {
		return
	}

	it.ActualLines = nil
	it.ExpectedLines = nil
	it.comparingValidators.Dispose()
	it.comparingValidators = nil
}
