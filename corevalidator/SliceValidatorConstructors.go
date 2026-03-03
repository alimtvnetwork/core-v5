package corevalidator

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/enums/stringcompareas"
	"gitlab.com/auk-go/core/errcore"
	"gitlab.com/auk-go/core/internal/strutilinternal"
)

func NewSliceValidatorUsingErr(
	errActual error,
	compareLinesContentAsExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	inputLines := errcore.ErrorToSplitLines(errActual)
	compareLines := strings.Split(
		compareLinesContentAsExpected,
		constants.NewLineUnix,
	)

	return &SliceValidator{
		ActualLines:   inputLines,
		ExpectedLines: compareLines,
		Condition: Condition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}

func NewSliceValidatorUsingAny(
	anyValActual any,
	compareLinesContentExpected string,
	isTrimLineCompare,
	isNonEmptyWhitespace,
	isSortStringsBySpace bool,
	compareAs stringcompareas.Variant,
) *SliceValidator {
	anyToString := strutilinternal.AnyToString(anyValActual)
	splitLines := strings.Split(anyToString, constants.NewLineUnix)
	compareLines := strings.Split(
		compareLinesContentExpected,
		constants.NewLineUnix,
	)

	return &SliceValidator{
		ActualLines:   splitLines,
		ExpectedLines: compareLines,
		Condition: Condition{
			IsTrimCompare:        isTrimLineCompare,
			IsNonEmptyWhitespace: isNonEmptyWhitespace,
			IsSortStringsBySpace: isSortStringsBySpace,
		},
		CompareAs:           compareAs,
		comparingValidators: nil,
	}
}
