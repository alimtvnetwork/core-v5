package chmodins

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
)

func ParseBaseRwxInstructionsUsingJsonResultMust(
	result *corejson.Result,
) *BaseRwxInstructions {
	baseRwxInstructions, err := ParseBaseRwxInstructionsUsingJsonResult(
		result)

	if err != nil {
		panic(err)
	}

	return baseRwxInstructions
}
