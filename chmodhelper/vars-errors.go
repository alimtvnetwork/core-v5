package chmodhelper

import (
	"errors"

	"github.com/alimtvnetwork/core/constants"
	"github.com/alimtvnetwork/core/errcore"
	"github.com/alimtvnetwork/core/internal/messages"
)

var (
	rwxInstructionNilErr = errcore.
				CannotBeNilType.
				Error(
			"rwx (...) - parsing failed",
			"rwxInstruction / rwxOwnerGroupOther - given as nil")

	failedToCompileVarWrapperToWrapper = errcore.
						FailedToExecuteType.
						Error(
			messages.FailedToCompileChmodhelperVarWrapperToWrapper,
			constants.EmptyString)

	errHyphenedRwxLength          = errors.New("length should be " + HyphenedRwxLengthString)
	errFullRwxLengthWithoutHyphen = errors.New("length should be " + FullRwxLengthWithoutHyphenString)
)
