package corevalidator

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/errcore"
)

func (it *SliceValidator) ActualInputWithExpectingMessage(
	caseIndex int,
	header string,
) string {
	actualInputMessage := it.ActualInputMessage(
		caseIndex,
		header,
	)
	userExpectingMessage := it.UserExpectingMessage(
		caseIndex,
		header,
	)
	finalMessage := actualInputMessage + constants.NewLineUnix + userExpectingMessage

	return finalMessage
}

func (it *SliceValidator) ActualInputMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		actualUserInputsV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ActualLinesString(),
	)
}

func (it *SliceValidator) UserExpectingMessage(
	caseIndex int,
	header string,
) string {
	finalHeader := fmt.Sprintf(
		expectingLinesV2MessageFormat,
		caseIndex,
		header,
	)

	return errcore.MsgHeaderPlusEnding(
		finalHeader,
		it.ExpectingLinesString(),
	)
}

// UserInputsMergeWithError
//
//   - Returns a combine error of actual and expecting inputs.
//   - If all validation successful then no error.
func (it *SliceValidator) UserInputsMergeWithError(
	parameter *Parameter,
	err error,
) error {
	if !parameter.IsAttachUserInputs {
		return err
	}

	toStr := it.ActualInputWithExpectingMessage(
		parameter.CaseIndex,
		parameter.Header,
	)

	if err == nil && len(toStr) == 0 {
		return nil
	}

	if err == nil && len(toStr) >= 0 {
		return errors.New(toStr)
	}

	msg := err.Error() + toStr

	return errors.New(msg)
}
