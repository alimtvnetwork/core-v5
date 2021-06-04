package chmodhelper

import "gitlab.com/evatix-go/core/chmodhelper/chmodins"

func ParseRwxInstructionToVarWrapper(
	rwxInstruction *chmodins.RwxInstruction,
) (
	*VarWrapper, error,
) {
	if rwxInstruction == nil {
		return nil, rwxInstructionNilErr
	}

	return ParseRwxOwnerGroupOtherInstructionToVarWrapper(
		&rwxInstruction.RwxOwnerGroupOther)
}
