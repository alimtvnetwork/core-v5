package chmodins

import (
	"github.com/alimtvnetwork/core/coredata/corejson"
	"github.com/alimtvnetwork/core/errcore"
)

func ParseBaseRwxInstructionsUsingJsonResult(
	result *corejson.Result,
) (*BaseRwxInstructions, error) {
	if result == nil {
		return nil,
			errcore.BytesAreNilOrEmptyType.Error(
				"ParseBaseRwxInstructionsUsingJsonResult", nil)
	}

	if result.IsEmptyJsonBytes() || result.HasError() {
		return nil, result.MeaningfulError()
	}

	var baseRwxInstructions BaseRwxInstructions
	err := result.Unmarshal(&baseRwxInstructions)

	if err != nil {
		return nil, errcore.MeaningfulError(
			errcore.FailedToParseType,
			"ParseBaseRwxInstructionsUsingJsonResult",
			err)
	}

	return &baseRwxInstructions, nil
}
