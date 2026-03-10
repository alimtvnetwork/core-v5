package chmodhelper

import "github.com/alimtvnetwork/core/chmodhelper/chmodins"

func ParseBaseRwxInstructionsToExecutors(
	baseRwxInstructions *chmodins.BaseRwxInstructions,
) (
	*RwxInstructionExecutors, error,
) {
	if baseRwxInstructions == nil || baseRwxInstructions.RwxInstructions == nil {
		return NewRwxInstructionExecutors(0), rwxInstructionNilErr
	}

	return ParseRwxInstructionsToExecutors(
		baseRwxInstructions.RwxInstructions)
}
