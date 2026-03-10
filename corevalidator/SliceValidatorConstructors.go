package corevalidator

import (
	"strings"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/enums/stringcompareas"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/strutilinternal"
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
