package corecsv

import "gitlab.com/auk-go/core/constants"

func DefaultAnyCsv(
	references ...any,
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
