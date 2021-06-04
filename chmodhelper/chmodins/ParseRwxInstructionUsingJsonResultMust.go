package chmodins

import (
	"gitlab.com/evatix-go/core/coredata/corejson"
)

func ParseRwxInstructionUsingJsonResultMust(
	result *corejson.Result,
) *RwxInstruction {
	rwxInstruction, err := ParseRwxInstructionUsingJsonResult(
		result)

	if err != nil {
		panic(err)
	}

	return rwxInstruction
}
