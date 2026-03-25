package chmodhelper

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

func ParseRwxInstructionToVarWrapper(
	rwxInstruction *chmodins.RwxInstruction,
) (
	*RwxVariableWrapper, error,
) {
	if rwxInstruction == nil {
		return nil, rwxInstructionNilErr
	}

	return ParseRwxOwnerGroupOtherToRwxVariableWrapper(
		&rwxInstruction.RwxOwnerGroupOther)
}
