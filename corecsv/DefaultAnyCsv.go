package corecsv

import "github.com/alimtvnetwork/core/constants"

func DefaultAnyCsv(
	references ...any,
) string {
	return AnyItemsToCsvString(
		constants.CommaSpace,
		true,
		false,
		references...)
}
