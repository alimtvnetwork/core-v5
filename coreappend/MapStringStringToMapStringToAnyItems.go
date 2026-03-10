package coreappend

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func MapStringStringAppendMapStringToAnyItems(
	isSkipEmpty bool,
	mainMap map[string]string,
	appendMapItems map[string]any,
) map[string]string {
	if len(appendMapItems) == 0 {
		return mainMap
	}

	for key, valInf := range appendMapItems {
		val := fmt.Sprintf(
			constants.SprintValueFormat,
			valInf)

		if isSkipEmpty && val == "" {
			continue
		}

		mainMap[key] = val
	}

	return mainMap
}
