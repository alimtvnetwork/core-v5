package corecsv

import "gitlab.com/auk-go/core/constants"

func AnyItemsToStringDefault(
	references ...any,
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
