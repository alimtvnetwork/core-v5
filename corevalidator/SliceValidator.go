package corevalidator

import (
	"errors"
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/enums/stringcompareas"
	"gitlab.com/evatix-go/core/internal/utilstringinternal"
	"gitlab.com/evatix-go/core/msgtype"
)

type SliceValidator struct {
	ValidatorCoreCondition
	// InputLines considered to be actual
	// ComparingLines considered to be expected
	InputLines, ComparingLines []string
	CompareAs                  stringcompareas.Variant
	comparingValidators        *TextValidators
}

func NewSliceValidatorUsingErr(
	errActual error,
	compareLinesContentAsExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	inputLines := msgtype.ErrorToSplitLines(errActual)
	compareLines := strings.Split(
		compareLinesContentAsExpected,
		constants.NewLineUnix)

	return &SliceValidator{
		InputLines:     inputLines,
		ComparingLines: compareLines,
		ValidatorCoreCondition: ValidatorCoreCondition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}

func NewSliceValidatorUsingAny(
	anyValActual interface{},
	compareLinesContentExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	anyToString := utilstringinternal.AnyToString(anyValActual)
	splitLines := strings.Split(anyToString, constants.NewLineUnix)
	compareLines := strings.Split(
		compareLinesContentExpected,
		constants.NewLineUnix)

	return &SliceValidator{
		InputLines:     splitLines,
		ComparingLines: compareLines,
		ValidatorCoreCondition: ValidatorCoreCondition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}

func (it *SliceValidator) InputLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.InputLines)
}

func (it *SliceValidator) MethodName() string {
	return it.CompareAs.Name()
}

func (it *SliceValidator) IsValidOtherLines(
	isCaseSensitive bool,
	otherActualLines []string,
) bool {
	return it.
		isValidLines(
			isCaseSensitive,
			otherActualLines)
}

func (it *SliceValidator) InputLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return msgtype.StringLinesToQuoteLinesToSingle(
		it.InputLines)
}

func (it *SliceValidator) ComparingLinesString() string {
	if it == nil {
		return constants.EmptyString
	}

	return msgtype.StringLinesToQuoteLinesToSingle(
		it.ComparingLines)
}

func (it *SliceValidator) ComparingLinesLength() int {
	if it == nil {
		return 0
	}

	return len(it.ComparingLines)
}

func (it *SliceValidator) ComparingValidators() *TextValidators {
	if it.comparingValidators != nil {
		return it.comparingValidators
	}

	validators := NewTextValidators(it.ComparingLinesLength())

	for _, line := range it.ComparingLines {
		validators.Add(TextValidator{
			Search:                 line,
			ValidatorCoreCondition: it.ValidatorCoreCondition,
			SearchAs:               it.CompareAs,
		})
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
		it.InputLines)
}

func (it *SliceValidator) isValidLines(
	isCaseSensitive bool,
	lines []string,
) bool {
	if it == nil && lines == nil {
		return true
	}

	if lines == nil && it.ComparingLines == nil {
		return true
	}

	if lines == nil || it.ComparingLines == nil {
		return false
	}

	inputLength := len(lines)
	comparingLength := len(it.ComparingLines)

	if inputLength != comparingLength {
		return false
	}

	validators := it.ComparingValidators()

	for i, validator := range validators.Items {
		isNotMatch := !validator.IsMatch(
			lines[i],
			isCaseSensitive)

		if isNotMatch {
			return false
		}
	}

	return true
}

func (it *SliceValidator) VerifyFirstError(
	paramsBase *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.VerifyFirstLengthUptoError(
		paramsBase,
		it.ComparingLinesLength())
}

func (it *SliceValidator) VerifyFirstLengthUptoError(
	params *ValidatorParamsBase,
	lengthUpTo int,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		true,
		params,
		lengthUpTo)
}

func (it *SliceValidator) AllVerifyError(
	params *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ComparingLinesLength())
}

func (it *SliceValidator) AllVerifyErrorExceptLast(
	params *ValidatorParamsBase,
) error {
	if it == nil {
		return nil
	}

	return it.AllVerifyErrorUptoLength(
		false,
		params,
		it.ComparingLinesLength()-1)
}

func (it *SliceValidator) AllVerifyErrorUptoLength(
	isFirstOnly bool,
	params *ValidatorParamsBase,
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
		lengthUpto)

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
			it.InputLines[i])

		if err != nil {
			sliceErr = append(
				sliceErr,
				err.Error())
		}

		if isFirstOnly && err != nil {
			break
		}
	}

	if params.IsAttachUserInputs && len(sliceErr) > constants.Zero {
		sliceErr = append(
			sliceErr,
			it.UserInputs())
	}

	return msgtype.SliceToError(sliceErr)
}

func (it *SliceValidator) lengthVerifyError(
	params *ValidatorParamsBase,
	lengthUpto int,
) error {
	hasLengthUpto := lengthUpto > constants.InvalidValue
	comparingLength := it.ComparingLinesLength()

	var comparingLengthError error
	if hasLengthUpto && lengthUpto > comparingLength {
		comparingLengthError = msgtype.OutOfRangeLength.Error(
			"Asked comparingLength is out of range!",
			comparingLength,
		)
	}

	if comparingLengthError != nil {
		return it.UserInputsMergeWithError(
			params,
			comparingLengthError)
	}

	var inputLengthErr error
	if it.InputLinesLength() > 0 && comparingLength == 0 {
		inputLengthErr = msgtype.LengthIssue.Error(
			"Input comparison has some text but comparing length is 0! Must set comparing text!",
			comparingLength,
		)
	}

	if inputLengthErr != nil {
		return it.UserInputsMergeWithError(
			params,
			inputLengthErr)
	}

	return nil
}

// initialVerifyError, verifyLengthUpto less than 0 will check actual length
func (it *SliceValidator) initialVerifyError(
	lengthUpto int,
) error {
	if it.InputLines == nil && it.ComparingLines == nil {
		return nil
	}

	isAnyNilCase := it.InputLines == nil ||
		it.ComparingLines == nil

	if isAnyNilCase {
		return msgtype.ExpectingErrorSimpleNoType(
			"InputLines, ComparingLines any is nil and other is not.",
			it.InputLines,
			it.ComparingLines,
		)
	}

	if !it.isLengthOkay(lengthUpto) {
		return msgtype.ExpectingErrorSimpleNoType(
			"InputLines, ComparingLines Length is not equal",
			len(it.InputLines),
			len(it.ComparingLines),
		)
	}

	return nil
}

func (it *SliceValidator) isLengthOkay(lengthUpto int) bool {
	inputLength := len(it.InputLines)
	comparingLength := len(it.ComparingLines)
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

func (it *SliceValidator) initialVerifyErrorWithMerged(
	params *ValidatorParamsBase,
	lengthUpto int,
) error {
	initialVerifyErr := it.initialVerifyError(
		lengthUpto)

	if initialVerifyErr != nil {
		return it.UserInputsMergeWithError(
			params,
			initialVerifyErr)
	}

	return nil
}

func (it *SliceValidator) UserInputs() string {
	return it.UserInputMessage() +
		it.UserComparingMessage()
}

func (it *SliceValidator) UserInputMessage() string {
	return msgtype.MsgHeaderPlusEnding(
		actualUserInputsMessage,
		it.InputLinesString())
}

func (it *SliceValidator) UserComparingMessage() string {
	return msgtype.MsgHeaderPlusEnding(
		expectingLinesMessage,
		it.ComparingLinesString())
}

func (it *SliceValidator) isEmptyIgnoreCase(
	params *ValidatorParamsBase,
) bool {
	return params.IsIgnoreCompareOnActualInputEmpty &&
		len(it.InputLines) == 0
}

func (it *SliceValidator) UserInputsMergeWithError(
	paramsBase *ValidatorParamsBase,
	err error,
) error {
	if !paramsBase.IsAttachUserInputs {
		return err
	}

	if err == nil {
		return errors.New(it.UserInputs())
	}

	msg := err.Error() +
		it.UserInputs()

	return errors.New(msg)
}
