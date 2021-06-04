package chmodins

import (
	"encoding/json"

	"gitlab.com/evatix-go/core/coredata/corejson"
	"gitlab.com/evatix-go/core/msgtype"
)

func ParseRwxInstructionUsingJsonResult(
	result *corejson.Result,
) (*RwxInstruction, error) {
	if result == nil {
		return nil,
			msgtype.JsonResultBytesAreNilOrEmpty.Error(
				"ParseRwxInstructionUsingJsonResult", nil)
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		return nil, result.MeaningfulError()
	}

	var rwxInstruction RwxInstruction
	err := json.Unmarshal(*result.Bytes, &rwxInstruction)

	if err != nil {
		return nil, msgtype.MeaningFulError(
			msgtype.FailedToParse,
			"ParseRwxInstructionUsingJsonResult",
			err)
	}

	return &rwxInstruction, err
}
