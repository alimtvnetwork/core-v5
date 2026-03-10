package coreappend

import (
	"fmt"

	"github.com/alimtvnetwork/core/constants"
)

func PrependAppendAnyItemsToStringsSkipOnNil(
	prependItem, appendItem any,
	anyItems ...any,
) []string {
	slice := make([]string, 0, len(anyItems)+3)

	if prependItem != nil {
		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, prependItem))
	}

	for _, item := range anyItems {
		if item == nil {
			continue
		}

		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, item))
	}

	if appendItem != nil {
		slice = append(
			slice,
			fmt.Sprintf(constants.SprintValueFormat, appendItem))
	}

	return slice
}
