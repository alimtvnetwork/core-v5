package stringslice

import "gitlab.com/auk-go/core/constants"

func AppendStringsWithAnyItems(
	isClone,
	isSkipOnEmpty bool,
	mainSlice []any,
	appendingItems ...string,
) []any {
	slice := AnyItemsCloneIf(
		isClone,
		len(appendingItems)+constants.Capacity2,
		mainSlice)

	if len(appendingItems) == 0 {
		return slice
	}

	for _, item := range appendingItems {
		if isSkipOnEmpty && item == "" {
			continue
		}

		slice = append(
			slice,
			item)
	}

	return slice
}
